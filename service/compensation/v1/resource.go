// Code generated by Lark OpenAPI.

package larkcompensation

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"net/http"
)

type V1 struct {
	Archive      *archive      // archive
	ChangeReason *changeReason // change_reason
	Indicator    *indicator    // indicator
	Item         *item         // item
	ItemCategory *itemCategory // item_category
	Plan         *plan         // plan
}

func New(config *larkcore.Config) *V1 {
	return &V1{
		Archive:      &archive{config: config},
		ChangeReason: &changeReason{config: config},
		Indicator:    &indicator{config: config},
		Item:         &item{config: config},
		ItemCategory: &itemCategory{config: config},
		Plan:         &plan{config: config},
	}
}

type archive struct {
	config *larkcore.Config
}
type changeReason struct {
	config *larkcore.Config
}
type indicator struct {
	config *larkcore.Config
}
type item struct {
	config *larkcore.Config
}
type itemCategory struct {
	config *larkcore.Config
}
type plan struct {
	config *larkcore.Config
}

// Query
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=query&project=compensation&resource=archive&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/compensationv1/query_archive.go
func (a *archive) Query(ctx context.Context, req *QueryArchiveReq, options ...larkcore.RequestOptionFunc) (*QueryArchiveResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/compensation/v1/archives/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryArchiveResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// List
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=list&project=compensation&resource=change_reason&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/compensationv1/list_changeReason.go
func (c *changeReason) List(ctx context.Context, req *ListChangeReasonReq, options ...larkcore.RequestOptionFunc) (*ListChangeReasonResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/compensation/v1/change_reasons"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, c.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListChangeReasonResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, c.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *changeReason) ListByIterator(ctx context.Context, req *ListChangeReasonReq, options ...larkcore.RequestOptionFunc) (*ListChangeReasonIterator, error) {
	return &ListChangeReasonIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.List,
		options:  options,
		limit:    req.Limit}, nil
}

// List
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=list&project=compensation&resource=indicator&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/compensationv1/list_indicator.go
func (i *indicator) List(ctx context.Context, req *ListIndicatorReq, options ...larkcore.RequestOptionFunc) (*ListIndicatorResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/compensation/v1/indicators"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListIndicatorResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *indicator) ListByIterator(ctx context.Context, req *ListIndicatorReq, options ...larkcore.RequestOptionFunc) (*ListIndicatorIterator, error) {
	return &ListIndicatorIterator{
		ctx:      ctx,
		req:      req,
		listFunc: i.List,
		options:  options,
		limit:    req.Limit}, nil
}

// List
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=list&project=compensation&resource=item&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/compensationv1/list_item.go
func (i *item) List(ctx context.Context, req *ListItemReq, options ...larkcore.RequestOptionFunc) (*ListItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/compensation/v1/items"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListItemResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *item) ListByIterator(ctx context.Context, req *ListItemReq, options ...larkcore.RequestOptionFunc) (*ListItemIterator, error) {
	return &ListItemIterator{
		ctx:      ctx,
		req:      req,
		listFunc: i.List,
		options:  options,
		limit:    req.Limit}, nil
}

// List
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=list&project=compensation&resource=item_category&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/compensationv1/list_itemCategory.go
func (i *itemCategory) List(ctx context.Context, req *ListItemCategoryReq, options ...larkcore.RequestOptionFunc) (*ListItemCategoryResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/compensation/v1/item_categories"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListItemCategoryResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *itemCategory) ListByIterator(ctx context.Context, req *ListItemCategoryReq, options ...larkcore.RequestOptionFunc) (*ListItemCategoryIterator, error) {
	return &ListItemCategoryIterator{
		ctx:      ctx,
		req:      req,
		listFunc: i.List,
		options:  options,
		limit:    req.Limit}, nil
}

// List
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=list&project=compensation&resource=plan&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/compensationv1/list_plan.go
func (p *plan) List(ctx context.Context, req *ListPlanReq, options ...larkcore.RequestOptionFunc) (*ListPlanResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/compensation/v1/plans"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListPlanResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, p.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *plan) ListByIterator(ctx context.Context, req *ListPlanReq, options ...larkcore.RequestOptionFunc) (*ListPlanIterator, error) {
	return &ListPlanIterator{
		ctx:      ctx,
		req:      req,
		listFunc: p.List,
		options:  options,
		limit:    req.Limit}, nil
}