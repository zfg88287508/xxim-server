package logic

import (
	"context"
	"github.com/cherish-chat/xxim-server/app/group/groupmodel"
	"github.com/cherish-chat/xxim-server/app/notice/noticemodel"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/cherish-chat/xxim-server/common/xorm"
	"github.com/cherish-chat/xxim-server/common/xtrace"
	"go.opentelemetry.io/otel/propagation"
	"gorm.io/gorm"
	"time"

	"github.com/cherish-chat/xxim-server/app/group/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type KickGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKickGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KickGroupMemberLogic {
	return &KickGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// KickGroupMember 踢出群成员
func (l *KickGroupMemberLogic) KickGroupMember(in *pb.KickGroupMemberReq) (*pb.KickGroupMemberResp, error) {
	// 获取群里所有的管理员
	groupManagers, err := getAllGroupManager(l.ctx, l.svcCtx, in.GroupId, true)
	if err != nil {
		l.Errorf("HandleGroupApply getAllGroupManager error: %v", err)
		return &pb.KickGroupMemberResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	// 判断是否是管理员
	isManager := false
	for _, manager := range groupManagers {
		if manager.UserId == in.CommonReq.UserId {
			isManager = true
			break
		}
	}
	if !isManager {
		// 如果要踢自己，可以直接踢
		if in.MemberId != in.CommonReq.UserId {
			return &pb.KickGroupMemberResp{CommonResp: pb.NewAlertErrorResp(
				l.svcCtx.T(in.CommonReq.Language, "操作失败"),
				l.svcCtx.T(in.CommonReq.Language, "您不是群管理员"),
			)}, nil
		}
	} else {
		// 判断是不是群主
		var memberInfo *pb.GetGroupMemberInfoResp
		xtrace.StartFuncSpan(l.ctx, "GetGroupMemberInfo", func(ctx context.Context) {
			memberInfo, err = NewGetGroupMemberInfoLogic(l.ctx, l.svcCtx).GetGroupMemberInfo(&pb.GetGroupMemberInfoReq{
				CommonReq: in.GetCommonReq(),
				GroupId:   in.GroupId,
				MemberId:  in.GetCommonReq().GetUserId(),
			})
		})
		if err != nil {
			l.Errorf("HandleGroupApply GetGroupMemberInfo error: %v", err)
			return &pb.KickGroupMemberResp{CommonResp: pb.NewRetryErrorResp()}, err
		}
		if memberInfo.GetGroupMemberInfo().GetRole() == pb.GroupRole_OWNER && in.MemberId == in.CommonReq.UserId {
			// 解散群
			var commonResp *pb.CommonResp
			var err error
			xtrace.StartFuncSpan(l.ctx, "DismissGroup", func(ctx context.Context) {
				commonResp, err = l.DismissGroup(in)
			})
			if err != nil {
				l.Errorf("HandleGroupApply DismissGroup error: %v", err)
				return &pb.KickGroupMemberResp{CommonResp: pb.NewRetryErrorResp()}, err
			}
			return &pb.KickGroupMemberResp{CommonResp: commonResp}, nil
		}
	}
	// 踢出群成员
	xtrace.StartFuncSpan(l.ctx, "KickGroupMember.Transaction", func(ctx context.Context) {
		tip := ""
		if in.CommonReq.UserId != in.MemberId {
			tip = in.MemberId + l.svcCtx.T(in.CommonReq.Language, "被移出群聊")
		}
		err = xorm.Transaction(l.svcCtx.Mysql(), func(tx *gorm.DB) error {
			// groupmember表
			member := &groupmodel.GroupMember{}
			err := tx.Model(member).Where("groupId = ? and userId = ?", in.GroupId, in.MemberId).Delete(member).Error
			if err != nil {
				l.Errorf("KickGroupMember groupmember delete error: %v", err)
				return err
			}
			return nil
		}, func(tx *gorm.DB) error {
			notice := &noticemodel.Notice{
				ConvId: pb.HiddenConvIdGroup(in.GroupId),
				Options: noticemodel.NoticeOption{
					StorageForClient: false,
					UpdateConvMsg:    false,
				},
				ContentType: pb.NoticeContentType_GroupMemberLeave,
				Content: utils.AnyToBytes(pb.NoticeContent_GroupMemberLeave{
					GroupId: in.GroupId,
					Tip:     tip,
				}),
				Title: "",
				Ext:   nil,
			}
			err = notice.Insert(l.ctx, tx)
			if err != nil {
				l.Errorf("insert notice failed, err: %v", err)
				return err
			}
			return nil
		})
	})
	if err != nil {
		l.Errorf("KickGroupMember error: %v", err)
		return &pb.KickGroupMemberResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	{
		// 预热缓存
		// 刷新订阅
		utils.RetryProxy(context.Background(), 12, 1*time.Second, func() error {
			_, err = NewSyncGroupMemberCountLogic(l.ctx, l.svcCtx).SyncGroupMemberCount(&pb.SyncGroupMemberCountReq{
				CommonReq: in.GetCommonReq(),
				GroupId:   in.GroupId,
			})
			if err != nil {
				l.Errorf("SyncGroupMemberCount failed, err: %v", err)
				return err
			}
			// 删除缓存
			{
				err = groupmodel.FlushGroupsByUserIdCache(l.ctx, l.svcCtx.Redis(), in.MemberId)
				if err != nil {
					l.Errorf("InviteFriendToGroup FlushGroupsByUserIdCache error: %v", err)
					return err
				}
			}
			// 预热缓存
			go xtrace.RunWithTrace(xtrace.TraceIdFromContext(l.ctx), "CacheWarmUp", func(ctx context.Context) {
				_, err = groupmodel.ListGroupsByUserIdFromMysql(ctx, l.svcCtx.Mysql(), l.svcCtx.Redis(), in.MemberId)
				if err != nil {
					l.Errorf("CreateGroup ListGroupsByUserIdFromMysql error: %v", err)
				}
			}, propagation.MapCarrier{
				"group_id": in.GroupId,
			})
			_, err := l.svcCtx.MsgService().FlushUsersSubConv(l.ctx, &pb.FlushUsersSubConvReq{UserIds: []string{in.MemberId}})
			if err != nil {
				l.Errorf("FlushUsersSubConv failed, err: %v", err)
				return err
			}
			_, err = l.svcCtx.NoticeService().GetUserNoticeData(l.ctx, &pb.GetUserNoticeDataReq{
				CommonReq: in.CommonReq,
				ConvId:    pb.HiddenConvIdGroup(in.GroupId),
			})
			if err != nil {
				l.Errorf("SendNoticeData failed, err: %v", err)
			}
			return err
		})
	}
	return &pb.KickGroupMemberResp{}, nil
}

func (l *KickGroupMemberLogic) DismissGroup(in *pb.KickGroupMemberReq) (*pb.CommonResp, error) {
	// 查询group
	mapGroupByIds, err := NewMapGroupByIdsLogic(l.ctx, l.svcCtx).MapGroupByIds(&pb.MapGroupByIdsReq{
		Ids: []string{in.GroupId},
	})
	if err != nil {
		l.Errorf("MapGroupByIds error: %v", err)
		return pb.NewRetryErrorResp(), err
	}
	if len(mapGroupByIds.GroupMap) == 0 {
		l.Errorf("MapGroupByIds failed, groupMap lenght = 0")
		return pb.NewAlertErrorResp(
			l.svcCtx.T(in.CommonReq.Language, "操作失败"),
			l.svcCtx.T(in.CommonReq.Language, "群组不存在"),
		), nil
	}
	groupBytes, ok := mapGroupByIds.GroupMap[in.GroupId]
	if !ok {
		l.Errorf("MapGroupByIds failed, groupMap: %v", mapGroupByIds.GroupMap)
		return pb.NewAlertErrorResp(
			l.svcCtx.T(in.CommonReq.Language, "操作失败"),
			l.svcCtx.T(in.CommonReq.Language, "群组不存在"),
		), nil
	}
	group := groupmodel.GroupFromBytes(groupBytes)
	err = xorm.Transaction(l.svcCtx.Mysql(), func(tx *gorm.DB) error {
		// update dismissTime
		return xorm.Update(tx, group, map[string]interface{}{
			"dismissTime": time.Now().UnixMilli(),
		})
	}, func(tx *gorm.DB) error {
		notice := &noticemodel.Notice{
			ConvId: pb.HiddenConvIdGroup(group.Id),
			Options: noticemodel.NoticeOption{
				StorageForClient: false,
				UpdateConvMsg:    false,
			},
			ContentType: pb.NoticeContentType_DismissGroup,
			Content: utils.AnyToBytes(pb.NoticeContent_DismissGroup{
				GroupId: group.Id,
			}),
			Title: "",
			Ext:   nil,
		}
		err = notice.Insert(l.ctx, tx)
		if err != nil {
			l.Errorf("insert notice failed, err: %v", err)
			return err
		}
		return nil
	}, func(tx *gorm.DB) error {
		err = groupmodel.CleanGroupCache(l.ctx, l.svcCtx.Redis(), group.Id)
		if err != nil {
			l.Errorf("CreateGroup CleanGroupCache error: %v", err)
			return err
		}
		return nil
	})
	if err != nil {
		l.Errorf("Transaction error: %v", err)
		return pb.NewRetryErrorResp(), err
	}
	{
		// 刷新订阅
		utils.RetryProxy(context.Background(), 12, 1*time.Second, func() error {
			_, err := l.svcCtx.MsgService().FlushUsersSubConv(l.ctx, &pb.FlushUsersSubConvReq{UserIds: []string{in.MemberId}})
			if err != nil {
				l.Errorf("FlushUsersSubConv failed, err: %v", err)
				return err
			}
			_, err = l.svcCtx.NoticeService().GetUserNoticeData(l.ctx, &pb.GetUserNoticeDataReq{
				CommonReq: in.CommonReq,
				ConvId:    pb.HiddenConvIdGroup(in.GroupId),
			})
			if err != nil {
				l.Errorf("SendNoticeData failed, err: %v", err)
			}
			return err
		})
	}
	return pb.NewSuccessResp(), nil
}
