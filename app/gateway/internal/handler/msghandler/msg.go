package msghandler

import (
	"context"
	"github.com/cherish-chat/xxim-server/app/gateway/internal/logic"
	"github.com/cherish-chat/xxim-server/app/gateway/internal/svc"
	"github.com/cherish-chat/xxim-server/app/gateway/internal/wrapper"
	"github.com/cherish-chat/xxim-server/common/pb"
	"github.com/cherish-chat/xxim-server/common/utils"
	"google.golang.org/grpc"
	"time"
)

// SendMsgConfig ...
func SendMsgConfig[REQ *pb.SendMsgListReq, RESP *pb.SendMsgListResp](svcCtx *svc.ServiceContext) wrapper.Config[*pb.SendMsgListReq, *pb.SendMsgListResp] {
	if svcCtx.Config.EnablePulsar {
		return wrapper.Config[*pb.SendMsgListReq, *pb.SendMsgListResp]{
			Do: func(ctx context.Context, in *pb.SendMsgListReq, opts ...grpc.CallOption) (*pb.SendMsgListResp, error) {
				requestTime := time.Now()
				resp, err := svcCtx.MsgService().SendMsgListAsync(ctx, in, opts...)
				go logic.NewApiLogLogic(ctx, svcCtx).ApiLog(in.GetCommonReq(), "SendMsg", resp.GetCommonResp(), utils.AnyToString(in), utils.AnyToString(resp), requestTime, time.Now(), err)
				return resp, err
			},
			NewRequest: func() *pb.SendMsgListReq {
				return &pb.SendMsgListReq{}
			},
		}
	} else {
		return wrapper.Config[*pb.SendMsgListReq, *pb.SendMsgListResp]{
			Do: func(ctx context.Context, in *pb.SendMsgListReq, opts ...grpc.CallOption) (*pb.SendMsgListResp, error) {
				requestTime := time.Now()
				resp, err := svcCtx.MsgService().SendMsgListSync(ctx, in, opts...)
				go logic.NewApiLogLogic(ctx, svcCtx).ApiLog(in.GetCommonReq(), "SendMsg", resp.GetCommonResp(), utils.AnyToString(in), utils.AnyToString(resp), requestTime, time.Now(), err)
				return resp, err
			},
			NewRequest: func() *pb.SendMsgListReq {
				return &pb.SendMsgListReq{}
			},
		}
	}
}

// GetMsgListByConvIdConfig ...
func GetMsgListByConvIdConfig[REQ *pb.GetMsgListByConvIdReq, RESP *pb.GetMsgListResp](svcCtx *svc.ServiceContext) wrapper.Config[*pb.GetMsgListByConvIdReq, *pb.GetMsgListResp] {
	return wrapper.Config[*pb.GetMsgListByConvIdReq, *pb.GetMsgListResp]{
		Do: func(ctx context.Context, in *pb.GetMsgListByConvIdReq, opts ...grpc.CallOption) (*pb.GetMsgListResp, error) {
			requestTime := time.Now()
			resp, err := svcCtx.MsgService().GetMsgListByConvId(ctx, in, opts...)
			go logic.NewApiLogLogic(ctx, svcCtx).ApiLog(in.GetCommonReq(), "GetMsgListByConvId", resp.GetCommonResp(), utils.AnyToString(in), utils.AnyToString(resp), requestTime, time.Now(), err)
			return resp, err
		},
		NewRequest: func() *pb.GetMsgListByConvIdReq {
			return &pb.GetMsgListByConvIdReq{}
		},
	}
}

// GetMsgByIdConfig ...
func GetMsgByIdConfig[REQ *pb.GetMsgByIdReq, RESP *pb.GetMsgByIdResp](svcCtx *svc.ServiceContext) wrapper.Config[*pb.GetMsgByIdReq, *pb.GetMsgByIdResp] {
	return wrapper.Config[*pb.GetMsgByIdReq, *pb.GetMsgByIdResp]{
		Do: func(ctx context.Context, in *pb.GetMsgByIdReq, opts ...grpc.CallOption) (*pb.GetMsgByIdResp, error) {
			requestTime := time.Now()
			resp, err := svcCtx.MsgService().GetMsgById(ctx, in, opts...)
			go logic.NewApiLogLogic(ctx, svcCtx).ApiLog(in.GetCommonReq(), "GetMsgById", resp.GetCommonResp(), utils.AnyToString(in), utils.AnyToString(resp), requestTime, time.Now(), err)
			return resp, err
		},
		NewRequest: func() *pb.GetMsgByIdReq {
			return &pb.GetMsgByIdReq{}
		},
	}
}
