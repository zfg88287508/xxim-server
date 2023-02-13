// Code generated by goctl. DO NOT EDIT!
// Source: mgmt.proto

package mgmtservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddMSApiPathReq                          = pb.AddMSApiPathReq
	AddMSApiPathResp                         = pb.AddMSApiPathResp
	AddMSIpWhiteListReq                      = pb.AddMSIpWhiteListReq
	AddMSIpWhiteListResp                     = pb.AddMSIpWhiteListResp
	AddMSMenuReq                             = pb.AddMSMenuReq
	AddMSMenuResp                            = pb.AddMSMenuResp
	AddMSRoleReq                             = pb.AddMSRoleReq
	AddMSRoleResp                            = pb.AddMSRoleResp
	AddMSUserReq                             = pb.AddMSUserReq
	AddMSUserResp                            = pb.AddMSUserResp
	AppLineConfig                            = pb.AppLineConfig
	AppLineConfig_Storage                    = pb.AppLineConfig_Storage
	AppLineConfig_Storage_Cos                = pb.AppLineConfig_Storage_Cos
	AppLineConfig_Storage_Kodo               = pb.AppLineConfig_Storage_Kodo
	AppLineConfig_Storage_Minio              = pb.AppLineConfig_Storage_Minio
	AppLineConfig_Storage_Oss                = pb.AppLineConfig_Storage_Oss
	ConfigMSResp                             = pb.ConfigMSResp
	DeleteMSApiPathReq                       = pb.DeleteMSApiPathReq
	DeleteMSApiPathResp                      = pb.DeleteMSApiPathResp
	DeleteMSIpWhiteListReq                   = pb.DeleteMSIpWhiteListReq
	DeleteMSIpWhiteListResp                  = pb.DeleteMSIpWhiteListResp
	DeleteMSMenuReq                          = pb.DeleteMSMenuReq
	DeleteMSMenuResp                         = pb.DeleteMSMenuResp
	DeleteMSOperationLogReq                  = pb.DeleteMSOperationLogReq
	DeleteMSOperationLogResp                 = pb.DeleteMSOperationLogResp
	DeleteMSRoleReq                          = pb.DeleteMSRoleReq
	DeleteMSRoleResp                         = pb.DeleteMSRoleResp
	DeleteMSUserReq                          = pb.DeleteMSUserReq
	DeleteMSUserResp                         = pb.DeleteMSUserResp
	GetAllMSApiPathListReq                   = pb.GetAllMSApiPathListReq
	GetAllMSApiPathListResp                  = pb.GetAllMSApiPathListResp
	GetAllMSIpWhiteListReq                   = pb.GetAllMSIpWhiteListReq
	GetAllMSIpWhiteListResp                  = pb.GetAllMSIpWhiteListResp
	GetAllMSLoginRecordReq                   = pb.GetAllMSLoginRecordReq
	GetAllMSLoginRecordResp                  = pb.GetAllMSLoginRecordResp
	GetAllMSMenuListReq                      = pb.GetAllMSMenuListReq
	GetAllMSMenuListResp                     = pb.GetAllMSMenuListResp
	GetAllMSOperationLogReq                  = pb.GetAllMSOperationLogReq
	GetAllMSOperationLogResp                 = pb.GetAllMSOperationLogResp
	GetAllMSRoleListReq                      = pb.GetAllMSRoleListReq
	GetAllMSRoleListResp                     = pb.GetAllMSRoleListResp
	GetAllMSUserListReq                      = pb.GetAllMSUserListReq
	GetAllMSUserListResp                     = pb.GetAllMSUserListResp
	GetAppLineConfigReq                      = pb.GetAppLineConfigReq
	GetAppLineConfigResp                     = pb.GetAppLineConfigResp
	GetMSApiPathDetailReq                    = pb.GetMSApiPathDetailReq
	GetMSApiPathDetailResp                   = pb.GetMSApiPathDetailResp
	GetMSIpWhiteListDetailReq                = pb.GetMSIpWhiteListDetailReq
	GetMSIpWhiteListDetailResp               = pb.GetMSIpWhiteListDetailResp
	GetMSMenuDetailReq                       = pb.GetMSMenuDetailReq
	GetMSMenuDetailResp                      = pb.GetMSMenuDetailResp
	GetMSOperationLogDetailReq               = pb.GetMSOperationLogDetailReq
	GetMSOperationLogDetailResp              = pb.GetMSOperationLogDetailResp
	GetMSRoleDetailReq                       = pb.GetMSRoleDetailReq
	GetMSRoleDetailResp                      = pb.GetMSRoleDetailResp
	GetMSUserDetailReq                       = pb.GetMSUserDetailReq
	GetMSUserDetailResp                      = pb.GetMSUserDetailResp
	GetMyMSApiPathListReq                    = pb.GetMyMSApiPathListReq
	GetMyMSApiPathListResp                   = pb.GetMyMSApiPathListResp
	GetMyMSMenuListReq                       = pb.GetMyMSMenuListReq
	GetMyMSMenuListResp                      = pb.GetMyMSMenuListResp
	GetSelfMSUserDetailReq                   = pb.GetSelfMSUserDetailReq
	GetSelfMSUserDetailResp                  = pb.GetSelfMSUserDetailResp
	GetServerAllConfigReq                    = pb.GetServerAllConfigReq
	GetServerAllConfigResp                   = pb.GetServerAllConfigResp
	GetServerAllConfigResp_AppMgmtRpcConfig  = pb.GetServerAllConfigResp_AppMgmtRpcConfig
	GetServerAllConfigResp_CommonConfig      = pb.GetServerAllConfigResp_CommonConfig
	GetServerAllConfigResp_ConnRpcConfig     = pb.GetServerAllConfigResp_ConnRpcConfig
	GetServerAllConfigResp_GroupRpcConfig    = pb.GetServerAllConfigResp_GroupRpcConfig
	GetServerAllConfigResp_ImRpcConfig       = pb.GetServerAllConfigResp_ImRpcConfig
	GetServerAllConfigResp_MgmtConfig        = pb.GetServerAllConfigResp_MgmtConfig
	GetServerAllConfigResp_MobPushConfig     = pb.GetServerAllConfigResp_MobPushConfig
	GetServerAllConfigResp_MsgPulsarConfig   = pb.GetServerAllConfigResp_MsgPulsarConfig
	GetServerAllConfigResp_MsgRpcConfig      = pb.GetServerAllConfigResp_MsgRpcConfig
	GetServerAllConfigResp_MysqlConfig       = pb.GetServerAllConfigResp_MysqlConfig
	GetServerAllConfigResp_NoticeRpcConfig   = pb.GetServerAllConfigResp_NoticeRpcConfig
	GetServerAllConfigResp_RedisConfig       = pb.GetServerAllConfigResp_RedisConfig
	GetServerAllConfigResp_RelationRpcConfig = pb.GetServerAllConfigResp_RelationRpcConfig
	GetServerAllConfigResp_SmsConfig         = pb.GetServerAllConfigResp_SmsConfig
	GetServerAllConfigResp_TelemetryConfig   = pb.GetServerAllConfigResp_TelemetryConfig
	GetServerAllConfigResp_TencentSmsConfig  = pb.GetServerAllConfigResp_TencentSmsConfig
	GetServerAllConfigResp_UserRpcConfig     = pb.GetServerAllConfigResp_UserRpcConfig
	GetServerConfigReq                       = pb.GetServerConfigReq
	GetServerConfigResp                      = pb.GetServerConfigResp
	HealthMSResp                             = pb.HealthMSResp
	LoginMSReq                               = pb.LoginMSReq
	LoginMSResp                              = pb.LoginMSResp
	MSApiPath                                = pb.MSApiPath
	MSIpWhiteList                            = pb.MSIpWhiteList
	MSLoginRecord                            = pb.MSLoginRecord
	MSMenu                                   = pb.MSMenu
	MSOperationLog                           = pb.MSOperationLog
	MSRole                                   = pb.MSRole
	MSUser                                   = pb.MSUser
	SwitchMSUserStatusReq                    = pb.SwitchMSUserStatusReq
	SwitchMSUserStatusResp                   = pb.SwitchMSUserStatusResp
	UpdateAppLineConfigReq                   = pb.UpdateAppLineConfigReq
	UpdateAppLineConfigResp                  = pb.UpdateAppLineConfigResp
	UpdateMSApiPathReq                       = pb.UpdateMSApiPathReq
	UpdateMSApiPathResp                      = pb.UpdateMSApiPathResp
	UpdateMSIpWhiteListReq                   = pb.UpdateMSIpWhiteListReq
	UpdateMSIpWhiteListResp                  = pb.UpdateMSIpWhiteListResp
	UpdateMSMenuReq                          = pb.UpdateMSMenuReq
	UpdateMSMenuResp                         = pb.UpdateMSMenuResp
	UpdateMSRoleReq                          = pb.UpdateMSRoleReq
	UpdateMSRoleResp                         = pb.UpdateMSRoleResp
	UpdateMSUserReq                          = pb.UpdateMSUserReq
	UpdateMSUserResp                         = pb.UpdateMSUserResp
	UpdateServerConfigReq                    = pb.UpdateServerConfigReq
	UpdateServerConfigResp                   = pb.UpdateServerConfigResp

	MgmtService interface {
		AfterConnect(ctx context.Context, in *AfterConnectReq, opts ...grpc.CallOption) (*CommonResp, error)
		AfterDisconnect(ctx context.Context, in *AfterDisconnectReq, opts ...grpc.CallOption) (*CommonResp, error)
		GetServerConfig(ctx context.Context, in *GetServerConfigReq, opts ...grpc.CallOption) (*GetServerConfigResp, error)
		GetServerAllConfig(ctx context.Context, in *GetServerAllConfigReq, opts ...grpc.CallOption) (*GetServerAllConfigResp, error)
		UpdateServerConfig(ctx context.Context, in *UpdateServerConfigReq, opts ...grpc.CallOption) (*UpdateServerConfigResp, error)
		GetAppLineConfig(ctx context.Context, in *GetAppLineConfigReq, opts ...grpc.CallOption) (*GetAppLineConfigResp, error)
		UpdateAppLineConfig(ctx context.Context, in *UpdateAppLineConfigReq, opts ...grpc.CallOption) (*UpdateAppLineConfigResp, error)
		LoginMS(ctx context.Context, in *LoginMSReq, opts ...grpc.CallOption) (*LoginMSResp, error)
		HealthMS(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*HealthMSResp, error)
		ConfigMS(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*ConfigMSResp, error)
		GetAllMSMenuList(ctx context.Context, in *GetAllMSMenuListReq, opts ...grpc.CallOption) (*GetAllMSMenuListResp, error)
		GetMyMSMenuList(ctx context.Context, in *GetMyMSMenuListReq, opts ...grpc.CallOption) (*GetMyMSMenuListResp, error)
		GetMSMenuDetail(ctx context.Context, in *GetMSMenuDetailReq, opts ...grpc.CallOption) (*GetMSMenuDetailResp, error)
		AddMSMenu(ctx context.Context, in *AddMSMenuReq, opts ...grpc.CallOption) (*AddMSMenuResp, error)
		UpdateMSMenu(ctx context.Context, in *UpdateMSMenuReq, opts ...grpc.CallOption) (*UpdateMSMenuResp, error)
		DeleteMSMenu(ctx context.Context, in *DeleteMSMenuReq, opts ...grpc.CallOption) (*DeleteMSMenuResp, error)
		GetAllMSApiPathList(ctx context.Context, in *GetAllMSApiPathListReq, opts ...grpc.CallOption) (*GetAllMSApiPathListResp, error)
		GetMyMSApiPathList(ctx context.Context, in *GetMyMSApiPathListReq, opts ...grpc.CallOption) (*GetMyMSApiPathListResp, error)
		GetMSApiPathDetail(ctx context.Context, in *GetMSApiPathDetailReq, opts ...grpc.CallOption) (*GetMSApiPathDetailResp, error)
		AddMSApiPath(ctx context.Context, in *AddMSApiPathReq, opts ...grpc.CallOption) (*AddMSApiPathResp, error)
		UpdateMSApiPath(ctx context.Context, in *UpdateMSApiPathReq, opts ...grpc.CallOption) (*UpdateMSApiPathResp, error)
		DeleteMSApiPath(ctx context.Context, in *DeleteMSApiPathReq, opts ...grpc.CallOption) (*DeleteMSApiPathResp, error)
		GetAllMSRoleList(ctx context.Context, in *GetAllMSRoleListReq, opts ...grpc.CallOption) (*GetAllMSRoleListResp, error)
		GetMSRoleDetail(ctx context.Context, in *GetMSRoleDetailReq, opts ...grpc.CallOption) (*GetMSRoleDetailResp, error)
		AddMSRole(ctx context.Context, in *AddMSRoleReq, opts ...grpc.CallOption) (*AddMSRoleResp, error)
		UpdateMSRole(ctx context.Context, in *UpdateMSRoleReq, opts ...grpc.CallOption) (*UpdateMSRoleResp, error)
		DeleteMSRole(ctx context.Context, in *DeleteMSRoleReq, opts ...grpc.CallOption) (*DeleteMSRoleResp, error)
		GetAllMSUserList(ctx context.Context, in *GetAllMSUserListReq, opts ...grpc.CallOption) (*GetAllMSUserListResp, error)
		GetMSUserDetail(ctx context.Context, in *GetMSUserDetailReq, opts ...grpc.CallOption) (*GetMSUserDetailResp, error)
		GetSelfMSUserDetail(ctx context.Context, in *GetSelfMSUserDetailReq, opts ...grpc.CallOption) (*GetSelfMSUserDetailResp, error)
		AddMSUser(ctx context.Context, in *AddMSUserReq, opts ...grpc.CallOption) (*AddMSUserResp, error)
		UpdateMSUser(ctx context.Context, in *UpdateMSUserReq, opts ...grpc.CallOption) (*UpdateMSUserResp, error)
		DeleteMSUser(ctx context.Context, in *DeleteMSUserReq, opts ...grpc.CallOption) (*DeleteMSUserResp, error)
		SwitchMSUserStatus(ctx context.Context, in *SwitchMSUserStatusReq, opts ...grpc.CallOption) (*SwitchMSUserStatusResp, error)
		GetAllMSIpWhiteList(ctx context.Context, in *GetAllMSIpWhiteListReq, opts ...grpc.CallOption) (*GetAllMSIpWhiteListResp, error)
		GetMSIpWhiteListDetail(ctx context.Context, in *GetMSIpWhiteListDetailReq, opts ...grpc.CallOption) (*GetMSIpWhiteListDetailResp, error)
		AddMSIpWhiteList(ctx context.Context, in *AddMSIpWhiteListReq, opts ...grpc.CallOption) (*AddMSIpWhiteListResp, error)
		UpdateMSIpWhiteList(ctx context.Context, in *UpdateMSIpWhiteListReq, opts ...grpc.CallOption) (*UpdateMSIpWhiteListResp, error)
		DeleteMSIpWhiteList(ctx context.Context, in *DeleteMSIpWhiteListReq, opts ...grpc.CallOption) (*DeleteMSIpWhiteListResp, error)
		GetAllMSOperationLog(ctx context.Context, in *GetAllMSOperationLogReq, opts ...grpc.CallOption) (*GetAllMSOperationLogResp, error)
		GetMSOperationLogDetail(ctx context.Context, in *GetMSOperationLogDetailReq, opts ...grpc.CallOption) (*GetMSOperationLogDetailResp, error)
		DeleteMSOperationLog(ctx context.Context, in *DeleteMSOperationLogReq, opts ...grpc.CallOption) (*DeleteMSOperationLogResp, error)
		GetAllMSLoginRecord(ctx context.Context, in *GetAllMSLoginRecordReq, opts ...grpc.CallOption) (*GetAllMSLoginRecordResp, error)
	}

	defaultMgmtService struct {
		cli zrpc.Client
	}
)

func NewMgmtService(cli zrpc.Client) MgmtService {
	return &defaultMgmtService{
		cli: cli,
	}
}

func (m *defaultMgmtService) AfterConnect(ctx context.Context, in *AfterConnectReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.AfterConnect(ctx, in, opts...)
}

func (m *defaultMgmtService) AfterDisconnect(ctx context.Context, in *AfterDisconnectReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.AfterDisconnect(ctx, in, opts...)
}

func (m *defaultMgmtService) GetServerConfig(ctx context.Context, in *GetServerConfigReq, opts ...grpc.CallOption) (*GetServerConfigResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetServerConfig(ctx, in, opts...)
}

func (m *defaultMgmtService) GetServerAllConfig(ctx context.Context, in *GetServerAllConfigReq, opts ...grpc.CallOption) (*GetServerAllConfigResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetServerAllConfig(ctx, in, opts...)
}

func (m *defaultMgmtService) UpdateServerConfig(ctx context.Context, in *UpdateServerConfigReq, opts ...grpc.CallOption) (*UpdateServerConfigResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.UpdateServerConfig(ctx, in, opts...)
}

func (m *defaultMgmtService) GetAppLineConfig(ctx context.Context, in *GetAppLineConfigReq, opts ...grpc.CallOption) (*GetAppLineConfigResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetAppLineConfig(ctx, in, opts...)
}

func (m *defaultMgmtService) UpdateAppLineConfig(ctx context.Context, in *UpdateAppLineConfigReq, opts ...grpc.CallOption) (*UpdateAppLineConfigResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.UpdateAppLineConfig(ctx, in, opts...)
}

func (m *defaultMgmtService) LoginMS(ctx context.Context, in *LoginMSReq, opts ...grpc.CallOption) (*LoginMSResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.LoginMS(ctx, in, opts...)
}

func (m *defaultMgmtService) HealthMS(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*HealthMSResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.HealthMS(ctx, in, opts...)
}

func (m *defaultMgmtService) ConfigMS(ctx context.Context, in *CommonReq, opts ...grpc.CallOption) (*ConfigMSResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.ConfigMS(ctx, in, opts...)
}

func (m *defaultMgmtService) GetAllMSMenuList(ctx context.Context, in *GetAllMSMenuListReq, opts ...grpc.CallOption) (*GetAllMSMenuListResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetAllMSMenuList(ctx, in, opts...)
}

func (m *defaultMgmtService) GetMyMSMenuList(ctx context.Context, in *GetMyMSMenuListReq, opts ...grpc.CallOption) (*GetMyMSMenuListResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetMyMSMenuList(ctx, in, opts...)
}

func (m *defaultMgmtService) GetMSMenuDetail(ctx context.Context, in *GetMSMenuDetailReq, opts ...grpc.CallOption) (*GetMSMenuDetailResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetMSMenuDetail(ctx, in, opts...)
}

func (m *defaultMgmtService) AddMSMenu(ctx context.Context, in *AddMSMenuReq, opts ...grpc.CallOption) (*AddMSMenuResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.AddMSMenu(ctx, in, opts...)
}

func (m *defaultMgmtService) UpdateMSMenu(ctx context.Context, in *UpdateMSMenuReq, opts ...grpc.CallOption) (*UpdateMSMenuResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.UpdateMSMenu(ctx, in, opts...)
}

func (m *defaultMgmtService) DeleteMSMenu(ctx context.Context, in *DeleteMSMenuReq, opts ...grpc.CallOption) (*DeleteMSMenuResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.DeleteMSMenu(ctx, in, opts...)
}

func (m *defaultMgmtService) GetAllMSApiPathList(ctx context.Context, in *GetAllMSApiPathListReq, opts ...grpc.CallOption) (*GetAllMSApiPathListResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetAllMSApiPathList(ctx, in, opts...)
}

func (m *defaultMgmtService) GetMyMSApiPathList(ctx context.Context, in *GetMyMSApiPathListReq, opts ...grpc.CallOption) (*GetMyMSApiPathListResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetMyMSApiPathList(ctx, in, opts...)
}

func (m *defaultMgmtService) GetMSApiPathDetail(ctx context.Context, in *GetMSApiPathDetailReq, opts ...grpc.CallOption) (*GetMSApiPathDetailResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetMSApiPathDetail(ctx, in, opts...)
}

func (m *defaultMgmtService) AddMSApiPath(ctx context.Context, in *AddMSApiPathReq, opts ...grpc.CallOption) (*AddMSApiPathResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.AddMSApiPath(ctx, in, opts...)
}

func (m *defaultMgmtService) UpdateMSApiPath(ctx context.Context, in *UpdateMSApiPathReq, opts ...grpc.CallOption) (*UpdateMSApiPathResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.UpdateMSApiPath(ctx, in, opts...)
}

func (m *defaultMgmtService) DeleteMSApiPath(ctx context.Context, in *DeleteMSApiPathReq, opts ...grpc.CallOption) (*DeleteMSApiPathResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.DeleteMSApiPath(ctx, in, opts...)
}

func (m *defaultMgmtService) GetAllMSRoleList(ctx context.Context, in *GetAllMSRoleListReq, opts ...grpc.CallOption) (*GetAllMSRoleListResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetAllMSRoleList(ctx, in, opts...)
}

func (m *defaultMgmtService) GetMSRoleDetail(ctx context.Context, in *GetMSRoleDetailReq, opts ...grpc.CallOption) (*GetMSRoleDetailResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetMSRoleDetail(ctx, in, opts...)
}

func (m *defaultMgmtService) AddMSRole(ctx context.Context, in *AddMSRoleReq, opts ...grpc.CallOption) (*AddMSRoleResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.AddMSRole(ctx, in, opts...)
}

func (m *defaultMgmtService) UpdateMSRole(ctx context.Context, in *UpdateMSRoleReq, opts ...grpc.CallOption) (*UpdateMSRoleResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.UpdateMSRole(ctx, in, opts...)
}

func (m *defaultMgmtService) DeleteMSRole(ctx context.Context, in *DeleteMSRoleReq, opts ...grpc.CallOption) (*DeleteMSRoleResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.DeleteMSRole(ctx, in, opts...)
}

func (m *defaultMgmtService) GetAllMSUserList(ctx context.Context, in *GetAllMSUserListReq, opts ...grpc.CallOption) (*GetAllMSUserListResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetAllMSUserList(ctx, in, opts...)
}

func (m *defaultMgmtService) GetMSUserDetail(ctx context.Context, in *GetMSUserDetailReq, opts ...grpc.CallOption) (*GetMSUserDetailResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetMSUserDetail(ctx, in, opts...)
}

func (m *defaultMgmtService) GetSelfMSUserDetail(ctx context.Context, in *GetSelfMSUserDetailReq, opts ...grpc.CallOption) (*GetSelfMSUserDetailResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetSelfMSUserDetail(ctx, in, opts...)
}

func (m *defaultMgmtService) AddMSUser(ctx context.Context, in *AddMSUserReq, opts ...grpc.CallOption) (*AddMSUserResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.AddMSUser(ctx, in, opts...)
}

func (m *defaultMgmtService) UpdateMSUser(ctx context.Context, in *UpdateMSUserReq, opts ...grpc.CallOption) (*UpdateMSUserResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.UpdateMSUser(ctx, in, opts...)
}

func (m *defaultMgmtService) DeleteMSUser(ctx context.Context, in *DeleteMSUserReq, opts ...grpc.CallOption) (*DeleteMSUserResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.DeleteMSUser(ctx, in, opts...)
}

func (m *defaultMgmtService) SwitchMSUserStatus(ctx context.Context, in *SwitchMSUserStatusReq, opts ...grpc.CallOption) (*SwitchMSUserStatusResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.SwitchMSUserStatus(ctx, in, opts...)
}

func (m *defaultMgmtService) GetAllMSIpWhiteList(ctx context.Context, in *GetAllMSIpWhiteListReq, opts ...grpc.CallOption) (*GetAllMSIpWhiteListResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetAllMSIpWhiteList(ctx, in, opts...)
}

func (m *defaultMgmtService) GetMSIpWhiteListDetail(ctx context.Context, in *GetMSIpWhiteListDetailReq, opts ...grpc.CallOption) (*GetMSIpWhiteListDetailResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetMSIpWhiteListDetail(ctx, in, opts...)
}

func (m *defaultMgmtService) AddMSIpWhiteList(ctx context.Context, in *AddMSIpWhiteListReq, opts ...grpc.CallOption) (*AddMSIpWhiteListResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.AddMSIpWhiteList(ctx, in, opts...)
}

func (m *defaultMgmtService) UpdateMSIpWhiteList(ctx context.Context, in *UpdateMSIpWhiteListReq, opts ...grpc.CallOption) (*UpdateMSIpWhiteListResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.UpdateMSIpWhiteList(ctx, in, opts...)
}

func (m *defaultMgmtService) DeleteMSIpWhiteList(ctx context.Context, in *DeleteMSIpWhiteListReq, opts ...grpc.CallOption) (*DeleteMSIpWhiteListResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.DeleteMSIpWhiteList(ctx, in, opts...)
}

func (m *defaultMgmtService) GetAllMSOperationLog(ctx context.Context, in *GetAllMSOperationLogReq, opts ...grpc.CallOption) (*GetAllMSOperationLogResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetAllMSOperationLog(ctx, in, opts...)
}

func (m *defaultMgmtService) GetMSOperationLogDetail(ctx context.Context, in *GetMSOperationLogDetailReq, opts ...grpc.CallOption) (*GetMSOperationLogDetailResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetMSOperationLogDetail(ctx, in, opts...)
}

func (m *defaultMgmtService) DeleteMSOperationLog(ctx context.Context, in *DeleteMSOperationLogReq, opts ...grpc.CallOption) (*DeleteMSOperationLogResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.DeleteMSOperationLog(ctx, in, opts...)
}

func (m *defaultMgmtService) GetAllMSLoginRecord(ctx context.Context, in *GetAllMSLoginRecordReq, opts ...grpc.CallOption) (*GetAllMSLoginRecordResp, error) {
	client := pb.NewMgmtServiceClient(m.cli.Conn())
	return client.GetAllMSLoginRecord(ctx, in, opts...)
}
