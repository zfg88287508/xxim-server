// Code generated by goctl. DO NOT EDIT.
// Source: conversation.proto

package subscriptionservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ConversationSettingUpdateReq                            = pb.ConversationSettingUpdateReq
	ConversationSettingUpdateResp                           = pb.ConversationSettingUpdateResp
	CountCreateGroupReq                                     = pb.CountCreateGroupReq
	CountCreateGroupResp                                    = pb.CountCreateGroupResp
	CountFriendReq                                          = pb.CountFriendReq
	CountFriendResp                                         = pb.CountFriendResp
	CountJoinGroupReq                                       = pb.CountJoinGroupReq
	CountJoinGroupResp                                      = pb.CountJoinGroupResp
	DeleteUserSubscriptionReq                               = pb.DeleteUserSubscriptionReq
	DeleteUserSubscriptionResp                              = pb.DeleteUserSubscriptionResp
	FriendApplyHandleReq                                    = pb.FriendApplyHandleReq
	FriendApplyHandleResp                                   = pb.FriendApplyHandleResp
	FriendApplyReq                                          = pb.FriendApplyReq
	FriendApplyResp                                         = pb.FriendApplyResp
	GroupCreateReq                                          = pb.GroupCreateReq
	GroupCreateResp                                         = pb.GroupCreateResp
	GroupSubscribeReq                                       = pb.GroupSubscribeReq
	GroupSubscribeResp                                      = pb.GroupSubscribeResp
	ListFriendApplyReq                                      = pb.ListFriendApplyReq
	ListFriendApplyReq_Filter                               = pb.ListFriendApplyReq_Filter
	ListFriendApplyReq_Option                               = pb.ListFriendApplyReq_Option
	ListFriendApplyResp                                     = pb.ListFriendApplyResp
	ListFriendApplyResp_FriendApply                         = pb.ListFriendApplyResp_FriendApply
	ListGroupSubscribersReq                                 = pb.ListGroupSubscribersReq
	ListGroupSubscribersReq_Filter                          = pb.ListGroupSubscribersReq_Filter
	ListGroupSubscribersReq_Option                          = pb.ListGroupSubscribersReq_Option
	ListGroupSubscribersResp                                = pb.ListGroupSubscribersResp
	ListGroupSubscribersResp_Subscriber                     = pb.ListGroupSubscribersResp_Subscriber
	ListJoinedConversationsReq                              = pb.ListJoinedConversationsReq
	ListJoinedConversationsReq_Filter                       = pb.ListJoinedConversationsReq_Filter
	ListJoinedConversationsReq_Filter_SettingKV             = pb.ListJoinedConversationsReq_Filter_SettingKV
	ListJoinedConversationsReq_Option                       = pb.ListJoinedConversationsReq_Option
	ListJoinedConversationsResp                             = pb.ListJoinedConversationsResp
	ListJoinedConversationsResp_Conversation                = pb.ListJoinedConversationsResp_Conversation
	ListJoinedConversationsResp_Conversation_SelfMemberInfo = pb.ListJoinedConversationsResp_Conversation_SelfMemberInfo
	ListSubscriptionSubscribersReq                          = pb.ListSubscriptionSubscribersReq
	ListSubscriptionSubscribersReq_Filter                   = pb.ListSubscriptionSubscribersReq_Filter
	ListSubscriptionSubscribersReq_Option                   = pb.ListSubscriptionSubscribersReq_Option
	ListSubscriptionSubscribersResp                         = pb.ListSubscriptionSubscribersResp
	ListSubscriptionSubscribersResp_Subscriber              = pb.ListSubscriptionSubscribersResp_Subscriber
	SubscriptionAfterOfflineReq                             = pb.SubscriptionAfterOfflineReq
	SubscriptionAfterOfflineResp                            = pb.SubscriptionAfterOfflineResp
	SubscriptionAfterOnlineReq                              = pb.SubscriptionAfterOnlineReq
	SubscriptionAfterOnlineResp                             = pb.SubscriptionAfterOnlineResp
	SubscriptionSubscribeReq                                = pb.SubscriptionSubscribeReq
	SubscriptionSubscribeResp                               = pb.SubscriptionSubscribeResp
	UpsertUserSubscriptionReq                               = pb.UpsertUserSubscriptionReq
	UpsertUserSubscriptionResp                              = pb.UpsertUserSubscriptionResp
	UserSubscription                                        = pb.UserSubscription

	SubscriptionService interface {
		// SubscriptionSubscribe 订阅号订阅
		SubscriptionSubscribe(ctx context.Context, in *SubscriptionSubscribeReq, opts ...grpc.CallOption) (*SubscriptionSubscribeResp, error)
		// SubscriptionAfterOnline 订阅号在做用户上线后的操作
		SubscriptionAfterOnline(ctx context.Context, in *SubscriptionAfterOnlineReq, opts ...grpc.CallOption) (*SubscriptionAfterOnlineResp, error)
		// SubscriptionAfterOffline 订阅号在做用户下线后的操作
		SubscriptionAfterOffline(ctx context.Context, in *SubscriptionAfterOfflineReq, opts ...grpc.CallOption) (*SubscriptionAfterOfflineResp, error)
		// UpsertUserSubscription 更新用户订阅的订阅号
		UpsertUserSubscription(ctx context.Context, in *UpsertUserSubscriptionReq, opts ...grpc.CallOption) (*UpsertUserSubscriptionResp, error)
		// DeleteUserSubscription 删除用户订阅的订阅号
		DeleteUserSubscription(ctx context.Context, in *DeleteUserSubscriptionReq, opts ...grpc.CallOption) (*DeleteUserSubscriptionResp, error)
		// ListSubscriptionSubscribers 列出订阅号订阅者
		ListSubscriptionSubscribers(ctx context.Context, in *ListSubscriptionSubscribersReq, opts ...grpc.CallOption) (*ListSubscriptionSubscribersResp, error)
	}

	defaultSubscriptionService struct {
		cli zrpc.Client
	}
)

func NewSubscriptionService(cli zrpc.Client) SubscriptionService {
	return &defaultSubscriptionService{
		cli: cli,
	}
}

// SubscriptionSubscribe 订阅号订阅
func (m *defaultSubscriptionService) SubscriptionSubscribe(ctx context.Context, in *SubscriptionSubscribeReq, opts ...grpc.CallOption) (*SubscriptionSubscribeResp, error) {
	client := pb.NewSubscriptionServiceClient(m.cli.Conn())
	return client.SubscriptionSubscribe(ctx, in, opts...)
}

// SubscriptionAfterOnline 订阅号在做用户上线后的操作
func (m *defaultSubscriptionService) SubscriptionAfterOnline(ctx context.Context, in *SubscriptionAfterOnlineReq, opts ...grpc.CallOption) (*SubscriptionAfterOnlineResp, error) {
	client := pb.NewSubscriptionServiceClient(m.cli.Conn())
	return client.SubscriptionAfterOnline(ctx, in, opts...)
}

// SubscriptionAfterOffline 订阅号在做用户下线后的操作
func (m *defaultSubscriptionService) SubscriptionAfterOffline(ctx context.Context, in *SubscriptionAfterOfflineReq, opts ...grpc.CallOption) (*SubscriptionAfterOfflineResp, error) {
	client := pb.NewSubscriptionServiceClient(m.cli.Conn())
	return client.SubscriptionAfterOffline(ctx, in, opts...)
}

// UpsertUserSubscription 更新用户订阅的订阅号
func (m *defaultSubscriptionService) UpsertUserSubscription(ctx context.Context, in *UpsertUserSubscriptionReq, opts ...grpc.CallOption) (*UpsertUserSubscriptionResp, error) {
	client := pb.NewSubscriptionServiceClient(m.cli.Conn())
	return client.UpsertUserSubscription(ctx, in, opts...)
}

// DeleteUserSubscription 删除用户订阅的订阅号
func (m *defaultSubscriptionService) DeleteUserSubscription(ctx context.Context, in *DeleteUserSubscriptionReq, opts ...grpc.CallOption) (*DeleteUserSubscriptionResp, error) {
	client := pb.NewSubscriptionServiceClient(m.cli.Conn())
	return client.DeleteUserSubscription(ctx, in, opts...)
}

// ListSubscriptionSubscribers 列出订阅号订阅者
func (m *defaultSubscriptionService) ListSubscriptionSubscribers(ctx context.Context, in *ListSubscriptionSubscribersReq, opts ...grpc.CallOption) (*ListSubscriptionSubscribersResp, error) {
	client := pb.NewSubscriptionServiceClient(m.cli.Conn())
	return client.ListSubscriptionSubscribers(ctx, in, opts...)
}
