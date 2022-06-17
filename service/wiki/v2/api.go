// Package wiki code generated by oapi sdk gen
package wiki

import (
	"context"
	"net/http"

	"github.com/feishu/oapi-sdk-go/core"
)

/**
构建业务域服务实例
**/
func NewService(httpClient *http.Client, config *core.Config) *WikiService {
	w := &WikiService{httpClient: httpClient, config: config}
	w.Space = &space{service: w}
	w.SpaceMember = &spaceMember{service: w}
	w.SpaceNode = &spaceNode{service: w}
	w.SpaceSetting = &spaceSetting{service: w}
	w.Task = &task{service: w}
	return w
}

/**
业务域服务定义
**/
type WikiService struct {
	httpClient   *http.Client
	config       *core.Config
	Space        *space
	SpaceMember  *spaceMember
	SpaceNode    *spaceNode
	SpaceSetting *spaceSetting
	Task         *task
}

/**
资源服务定义
**/
type space struct {
	service *WikiService
}
type spaceMember struct {
	service *WikiService
}
type spaceNode struct {
	service *WikiService
}
type spaceSetting struct {
	service *WikiService
}
type task struct {
	service *WikiService
}

/**
资源服务方法定义
**/
func (s *space) Create(ctx context.Context, req *CreateSpaceReq, options ...core.RequestOptionFunc) (*CreateSpaceResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/wiki/v2/spaces", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpaceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *space) Get(ctx context.Context, req *GetSpaceReq, options ...core.RequestOptionFunc) (*GetSpaceResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/wiki/v2/spaces/:space_id", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetSpaceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *space) GetNode(ctx context.Context, req *GetNodeSpaceReq, options ...core.RequestOptionFunc) (*GetNodeSpaceResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/wiki/v2/spaces/get_node", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetNodeSpaceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *space) List(ctx context.Context, req *ListSpaceReq, options ...core.RequestOptionFunc) (*ListSpaceResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/wiki/v2/spaces", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListSpaceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (s *space) ListSpace(ctx context.Context, req *ListSpaceReq, options ...core.RequestOptionFunc) (*ListSpaceIterator, error) {
	return &ListSpaceIterator{
		ctx:      ctx,
		req:      req,
		listFunc: s.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (s *spaceMember) Create(ctx context.Context, req *CreateSpaceMemberReq, options ...core.RequestOptionFunc) (*CreateSpaceMemberResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/wiki/v2/spaces/:space_id/members", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpaceMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spaceMember) Delete(ctx context.Context, req *DeleteSpaceMemberReq, options ...core.RequestOptionFunc) (*DeleteSpaceMemberResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodDelete,
		"/open-apis/wiki/v2/spaces/:space_id/members/:member_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteSpaceMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spaceNode) Copy(ctx context.Context, req *CopySpaceNodeReq, options ...core.RequestOptionFunc) (*CopySpaceNodeResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/wiki/v2/spaces/:space_id/nodes/:node_token/copy", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CopySpaceNodeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spaceNode) Create(ctx context.Context, req *CreateSpaceNodeReq, options ...core.RequestOptionFunc) (*CreateSpaceNodeResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/wiki/v2/spaces/:space_id/nodes", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpaceNodeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spaceNode) List(ctx context.Context, req *ListSpaceNodeReq, options ...core.RequestOptionFunc) (*ListSpaceNodeResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodGet,
		"/open-apis/wiki/v2/spaces/:space_id/nodes", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListSpaceNodeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (s *spaceNode) ListSpaceNode(ctx context.Context, req *ListSpaceNodeReq, options ...core.RequestOptionFunc) (*ListSpaceNodeIterator, error) {
	return &ListSpaceNodeIterator{
		ctx:      ctx,
		req:      req,
		listFunc: s.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (s *spaceNode) Move(ctx context.Context, req *MoveSpaceNodeReq, options ...core.RequestOptionFunc) (*MoveSpaceNodeResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/wiki/v2/spaces/:space_id/nodes/:node_token/move", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &MoveSpaceNodeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spaceNode) MoveDocsToWiki(ctx context.Context, req *MoveDocsToWikiSpaceNodeReq, options ...core.RequestOptionFunc) (*MoveDocsToWikiSpaceNodeResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/wiki/v2/spaces/:space_id/nodes/move_docs_to_wiki", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &MoveDocsToWikiSpaceNodeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spaceNode) UpdateTitle(ctx context.Context, req *UpdateTitleSpaceNodeReq, options ...core.RequestOptionFunc) (*UpdateTitleSpaceNodeResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/wiki/v2/spaces/:space_id/nodes/:node_token/update_title", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateTitleSpaceNodeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spaceSetting) Update(ctx context.Context, req *UpdateSpaceSettingReq, options ...core.RequestOptionFunc) (*UpdateSpaceSettingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodPut,
		"/open-apis/wiki/v2/spaces/:space_id/setting", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateSpaceSettingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Get(ctx context.Context, req *GetTaskReq, options ...core.RequestOptionFunc) (*GetTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodGet,
		"/open-apis/wiki/v2/tasks/:task_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
