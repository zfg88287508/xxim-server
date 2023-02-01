package logic

import (
	"context"
	"github.com/cherish-chat/xxim-server/app/mgmt/internal/svc"
	"github.com/cherish-chat/xxim-server/app/user/usermodel"
	"github.com/cherish-chat/xxim-server/common/pb"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/cherish-chat/xxim-server/common/xorm"
	"github.com/cherish-chat/xxim-server/common/xpwd"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"time"
)

type InitLogic struct {
	svcCtx *svc.ServiceContext
	ctx    context.Context
	logx.Logger
}

func NewInitLogic(svcCtx *svc.ServiceContext) *InitLogic {
	ctx := context.Background()
	return &InitLogic{svcCtx: svcCtx, ctx: ctx, Logger: logx.WithContext(ctx)}
}

func (l *InitLogic) Init() {
	// 获取配置文件中的SuperAdmin
	superAdmin := l.svcCtx.Config.SuperAdmin
	if superAdmin.Id == "" {
		// 说明不需要创建超级管理员
		l.Infof("不需要创建超级管理员")
		return
	}
	// 查询超级管理员是否存在
	userByIds, err := l.svcCtx.UserService().MapUserByIds(l.ctx, &pb.MapUserByIdsReq{
		CommonReq: &pb.CommonReq{},
		Ids:       []string{superAdmin.Id},
	})
	if err != nil {
		l.Errorf("MapUserByIds error: %v", err)
		return
	}
	_, ok := userByIds.Users[superAdmin.Id]
	if ok {
		// 说明超级管理员已经存在
		l.Infof("超级管理员已经存在")
		return
	}
	// 说明超级管理员不存在
	// 创建超级管理员
	l.Infof("创建超级管理员: %s, %s", superAdmin.Id, superAdmin.Password)
	salt := utils.GenId()
	password := xpwd.GeneratePwd(superAdmin.Password, salt)
	// 插入用户表
	user := &usermodel.User{
		Id:           superAdmin.Id,
		Password:     password,
		PasswordSalt: salt,
		Nickname:     "超级管理员",
		Avatar:       utils.AnyRandomInSlice(l.svcCtx.SystemConfigMgr.GetSlice("avatars_default"), ""),
		RegInfo: &usermodel.LoginInfo{
			Time: time.Now().UnixMilli(),
		},
		Role:     usermodel.RoleAdmin,
		RoleType: usermodel.RoleTypeAdminSuper,
	}
	err = xorm.InsertOne(l.svcCtx.Mysql(), user)
	if err != nil {
		log.Fatalf("insert super admin error: %v", err)
	}
}
