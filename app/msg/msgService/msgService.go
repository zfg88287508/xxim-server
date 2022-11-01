// Code generated by goctl. DO NOT EDIT!
// Source: msg.proto

package msgservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BatchSendMsgReq           = pb.BatchSendMsgReq
	GetGroupMsgListBySeqReq   = pb.GetGroupMsgListBySeqReq
	GetGroupMsgListBySeqResp  = pb.GetGroupMsgListBySeqResp
	GetSingleMsgListBySeqReq  = pb.GetSingleMsgListBySeqReq
	GetSingleMsgListBySeqResp = pb.GetSingleMsgListBySeqResp
	MsgData                   = pb.MsgData
	MsgDataList               = pb.MsgDataList
	MsgData_OfflinePush       = pb.MsgData_OfflinePush
	MsgData_Options           = pb.MsgData_Options
	MsgData_Receiver          = pb.MsgData_Receiver
	MsgMQBody                 = pb.MsgMQBody
	SendMsgListReq            = pb.SendMsgListReq

	MsgService interface {
		InsertMsgDataList(ctx context.Context, in *MsgDataList, opts ...grpc.CallOption) (*CommonResp, error)
		SendMsgListSync(ctx context.Context, in *SendMsgListReq, opts ...grpc.CallOption) (*CommonResp, error)
		SendMsgListAsync(ctx context.Context, in *SendMsgListReq, opts ...grpc.CallOption) (*CommonResp, error)
		BatchSendMsgSync(ctx context.Context, in *BatchSendMsgReq, opts ...grpc.CallOption) (*CommonResp, error)
		BatchSendMsgAsync(ctx context.Context, in *BatchSendMsgReq, opts ...grpc.CallOption) (*CommonResp, error)
		// GetSingleMsgListBySeq 通过seq拉取一个单聊会话的消息
		GetSingleMsgListBySeq(ctx context.Context, in *GetSingleMsgListBySeqReq, opts ...grpc.CallOption) (*GetSingleMsgListBySeqResp, error)
		// GetGroupMsgListBySeq 通过seq拉取一个群聊会话的消息
		GetGroupMsgListBySeq(ctx context.Context, in *GetGroupMsgListBySeqReq, opts ...grpc.CallOption) (*GetGroupMsgListBySeqResp, error)
	}

	defaultMsgService struct {
		cli zrpc.Client
	}
)

func NewMsgService(cli zrpc.Client) MsgService {
	return &defaultMsgService{
		cli: cli,
	}
}

func (m *defaultMsgService) InsertMsgDataList(ctx context.Context, in *MsgDataList, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewMsgServiceClient(m.cli.Conn())
	return client.InsertMsgDataList(ctx, in, opts...)
}

func (m *defaultMsgService) SendMsgListSync(ctx context.Context, in *SendMsgListReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewMsgServiceClient(m.cli.Conn())
	return client.SendMsgListSync(ctx, in, opts...)
}

func (m *defaultMsgService) SendMsgListAsync(ctx context.Context, in *SendMsgListReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewMsgServiceClient(m.cli.Conn())
	return client.SendMsgListAsync(ctx, in, opts...)
}

func (m *defaultMsgService) BatchSendMsgSync(ctx context.Context, in *BatchSendMsgReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewMsgServiceClient(m.cli.Conn())
	return client.BatchSendMsgSync(ctx, in, opts...)
}

func (m *defaultMsgService) BatchSendMsgAsync(ctx context.Context, in *BatchSendMsgReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewMsgServiceClient(m.cli.Conn())
	return client.BatchSendMsgAsync(ctx, in, opts...)
}

// GetSingleMsgListBySeq 通过seq拉取一个单聊会话的消息
func (m *defaultMsgService) GetSingleMsgListBySeq(ctx context.Context, in *GetSingleMsgListBySeqReq, opts ...grpc.CallOption) (*GetSingleMsgListBySeqResp, error) {
	client := pb.NewMsgServiceClient(m.cli.Conn())
	return client.GetSingleMsgListBySeq(ctx, in, opts...)
}

// GetGroupMsgListBySeq 通过seq拉取一个群聊会话的消息
func (m *defaultMsgService) GetGroupMsgListBySeq(ctx context.Context, in *GetGroupMsgListBySeqReq, opts ...grpc.CallOption) (*GetGroupMsgListBySeqResp, error) {
	client := pb.NewMsgServiceClient(m.cli.Conn())
	return client.GetGroupMsgListBySeq(ctx, in, opts...)
}