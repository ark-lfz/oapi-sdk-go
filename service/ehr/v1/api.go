// Package ehr code generated by oapi sdk gen
package ehr

import (
	"bytes"
	"context"
	"net/http"

	"github.com/feishu/oapi-sdk-go/core"
)

/**
构建业务域服务实例
**/
func NewService(httpClient *http.Client, config *core.Config) *EhrService {
	e := &EhrService{httpClient: httpClient, config: config}
	e.Attachment = &attachment{service: e}
	e.Employee = &employee{service: e}
	return e
}

/**
业务域服务定义
**/
type EhrService struct {
	httpClient *http.Client
	config     *core.Config
	Attachment *attachment
	Employee   *employee
}

/**
资源服务定义
**/
type attachment struct {
	service *EhrService
}
type employee struct {
	service *EhrService
}

/**
资源服务方法定义
**/
func (a *attachment) Get(ctx context.Context, req *GetAttachmentReq, options ...core.RequestOptionFunc) (*GetAttachmentResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, a.service.config, http.MethodGet,
		"/open-apis/ehr/v1/attachments/:token", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetAttachmentResp{RawResponse: rawResp}
	// 如果是下载，则设置响应结果
	if rawResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(rawResp.RawBody)
		resp.FileName = core.FileNameByHeader(rawResp.Header)
		return resp, err
	}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *employee) List(ctx context.Context, req *ListEmployeeReq, options ...core.RequestOptionFunc) (*ListEmployeeResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, e.service.config, http.MethodGet,
		"/open-apis/ehr/v1/employees", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListEmployeeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (e *employee) ListEmployee(ctx context.Context, req *ListEmployeeReq, options ...core.RequestOptionFunc) (*ListEmployeeIterator, error) {
	return &ListEmployeeIterator{
		ctx:      ctx,
		req:      req,
		listFunc: e.List,
		options:  options,
		limit:    req.Limit}, nil
}
