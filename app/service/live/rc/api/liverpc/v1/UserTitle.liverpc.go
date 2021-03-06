// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/UserTitle.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// ===================
// UserTitle Interface
// ===================

type UserTitleRPCClient interface {
	// * 根据头衔id获取头衔名称，头衔不存在不返回
	//
	GetTitleByIds(ctx context.Context, req *UserTitleGetTitleByIdsReq, opts ...liverpc.CallOption) (resp *UserTitleGetTitleByIdsResp, err error)

	// * 移动端获取佩戴的头衔
	//
	GetMobileTitle(ctx context.Context, req *UserTitleGetMobileTitleReq, opts ...liverpc.CallOption) (resp *UserTitleGetMobileTitleResp, err error)

	// * 获取弹幕头衔(需要登录态)
	//
	GetCommentTitle(ctx context.Context, req *UserTitleGetCommentTitleReq, opts ...liverpc.CallOption) (resp *UserTitleGetCommentTitleResp, err error)

	// * 添加头衔
	//
	Add(ctx context.Context, req *UserTitleAddReq, opts ...liverpc.CallOption) (resp *UserTitleAddResp, err error)

	// * 个人中心我的头衔列表
	//
	GetAll(ctx context.Context, req *UserTitleGetAllReq, opts ...liverpc.CallOption) (resp *UserTitleGetAllResp, err error)
}

// =========================
// UserTitle Live Rpc Client
// =========================

type userTitleRPCClient struct {
	client *liverpc.Client
}

// NewUserTitleRPCClient creates a client that implements the UserTitleRPCClient interface.
func NewUserTitleRPCClient(client *liverpc.Client) UserTitleRPCClient {
	return &userTitleRPCClient{
		client: client,
	}
}

func (c *userTitleRPCClient) GetTitleByIds(ctx context.Context, in *UserTitleGetTitleByIdsReq, opts ...liverpc.CallOption) (*UserTitleGetTitleByIdsResp, error) {
	out := new(UserTitleGetTitleByIdsResp)
	err := doRPCRequest(ctx, c.client, 1, "UserTitle.getTitleByIds", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTitleRPCClient) GetMobileTitle(ctx context.Context, in *UserTitleGetMobileTitleReq, opts ...liverpc.CallOption) (*UserTitleGetMobileTitleResp, error) {
	out := new(UserTitleGetMobileTitleResp)
	err := doRPCRequest(ctx, c.client, 1, "UserTitle.getMobileTitle", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTitleRPCClient) GetCommentTitle(ctx context.Context, in *UserTitleGetCommentTitleReq, opts ...liverpc.CallOption) (*UserTitleGetCommentTitleResp, error) {
	out := new(UserTitleGetCommentTitleResp)
	err := doRPCRequest(ctx, c.client, 1, "UserTitle.getCommentTitle", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTitleRPCClient) Add(ctx context.Context, in *UserTitleAddReq, opts ...liverpc.CallOption) (*UserTitleAddResp, error) {
	out := new(UserTitleAddResp)
	err := doRPCRequest(ctx, c.client, 1, "UserTitle.add", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTitleRPCClient) GetAll(ctx context.Context, in *UserTitleGetAllReq, opts ...liverpc.CallOption) (*UserTitleGetAllResp, error) {
	out := new(UserTitleGetAllResp)
	err := doRPCRequest(ctx, c.client, 1, "UserTitle.getAll", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}
