//  code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package lark

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/acs"
	"github.com/larksuite/oapi-sdk-go/v3/service/admin"
	"github.com/larksuite/oapi-sdk-go/v3/service/application"
	"github.com/larksuite/oapi-sdk-go/v3/service/approval"
	"github.com/larksuite/oapi-sdk-go/v3/service/attendance"
	"github.com/larksuite/oapi-sdk-go/v3/service/auth"
	"github.com/larksuite/oapi-sdk-go/v3/service/authen"
	"github.com/larksuite/oapi-sdk-go/v3/service/baike"
	"github.com/larksuite/oapi-sdk-go/v3/service/bitable"
	"github.com/larksuite/oapi-sdk-go/v3/service/block"
	"github.com/larksuite/oapi-sdk-go/v3/service/calendar"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact"
	"github.com/larksuite/oapi-sdk-go/v3/service/corehr"
	"github.com/larksuite/oapi-sdk-go/v3/service/document_ai"
	"github.com/larksuite/oapi-sdk-go/v3/service/docx"
	"github.com/larksuite/oapi-sdk-go/v3/service/drive"
	"github.com/larksuite/oapi-sdk-go/v3/service/ehr"
	"github.com/larksuite/oapi-sdk-go/v3/service/event"
	"github.com/larksuite/oapi-sdk-go/v3/service/ext"
	"github.com/larksuite/oapi-sdk-go/v3/service/gray_test_open_sg"
	"github.com/larksuite/oapi-sdk-go/v3/service/helpdesk"
	"github.com/larksuite/oapi-sdk-go/v3/service/hire"
	"github.com/larksuite/oapi-sdk-go/v3/service/human_authentication"
	"github.com/larksuite/oapi-sdk-go/v3/service/im"
	"github.com/larksuite/oapi-sdk-go/v3/service/lingo"
	"github.com/larksuite/oapi-sdk-go/v3/service/mail"
	"github.com/larksuite/oapi-sdk-go/v3/service/mdm"
	"github.com/larksuite/oapi-sdk-go/v3/service/meeting_room"
	"github.com/larksuite/oapi-sdk-go/v3/service/okr"
	"github.com/larksuite/oapi-sdk-go/v3/service/optical_char_recognition"
	"github.com/larksuite/oapi-sdk-go/v3/service/passport"
	"github.com/larksuite/oapi-sdk-go/v3/service/personal_settings"
	"github.com/larksuite/oapi-sdk-go/v3/service/report"
	"github.com/larksuite/oapi-sdk-go/v3/service/search"
	"github.com/larksuite/oapi-sdk-go/v3/service/security_and_compliance"
	"github.com/larksuite/oapi-sdk-go/v3/service/sheets"
	"github.com/larksuite/oapi-sdk-go/v3/service/speech_to_text"
	"github.com/larksuite/oapi-sdk-go/v3/service/task"
	"github.com/larksuite/oapi-sdk-go/v3/service/tenant"
	"github.com/larksuite/oapi-sdk-go/v3/service/translation"
	"github.com/larksuite/oapi-sdk-go/v3/service/vc"
	"github.com/larksuite/oapi-sdk-go/v3/service/verification"
	"github.com/larksuite/oapi-sdk-go/v3/service/wiki"
	"github.com/larksuite/oapi-sdk-go/v3/service/workplace"
)

type Client struct {
	config                 *larkcore.Config
	Attendance             *attendance.Service
	Auth                   *auth.Service
	Contact                *contact.Service
	GrayTestOpenSg         *gray_test_open_sg.Service
	Helpdesk               *helpdesk.Service
	HumanAuthentication    *human_authentication.Service
	Report                 *report.Service
	SpeechToText           *speech_to_text.Service
	Bitable                *bitable.Service
	Calendar               *calendar.Service
	DocumentAi             *document_ai.Service
	Mdm                    *mdm.Service
	SecurityAndCompliance  *security_and_compliance.Service
	Translation            *translation.Service
	Verification           *verification.Service
	Corehr                 *corehr.Service
	Hire                   *hire.Service
	Im                     *im.Service
	Tenant                 *tenant.Service
	Admin                  *admin.Service
	Block                  *block.Service
	Event                  *event.Service
	Lingo                  *lingo.Service
	Okr                    *okr.Service
	Search                 *search.Service
	Approval               *approval.Service
	Docx                   *docx.Service
	Ehr                    *ehr.Service
	MeetingRoom            *meeting_room.Service
	OpticalCharRecognition *optical_char_recognition.Service
	Sheets                 *sheets.Service
	Task                   *task.Service
	Workplace              *workplace.Service
	Passport               *passport.Service
	Vc                     *vc.Service
	Acs                    *acs.Service
	Application            *application.Service
	Authen                 *authen.Service
	Baike                  *baike.Service
	Drive                  *drive.Service
	Mail                   *mail.Service
	PersonalSettings       *personal_settings.Service
	Wiki                   *wiki.Service
	Ext                    *larkext.ExtService
}

type ClientOptionFunc func(config *larkcore.Config)

func WithAppType(appType larkcore.AppType) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.AppType = appType
	}
}

func WithMarketplaceApp() ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.AppType = larkcore.AppTypeMarketplace
	}
}

func WithEnableTokenCache(enableTokenCache bool) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.EnableTokenCache = enableTokenCache
	}
}

func WithLogger(logger larkcore.Logger) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.Logger = logger
	}
}

func WithOpenBaseUrl(baseUrl string) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.BaseUrl = baseUrl
	}
}

func WithLogLevel(logLevel larkcore.LogLevel) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.LogLevel = logLevel
	}
}

func WithTokenCache(cache larkcore.Cache) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.TokenCache = cache
	}
}

func WithLogReqAtDebug(printReqRespLog bool) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.LogReqAtDebug = printReqRespLog
	}
}

func WithHttpClient(httpClient larkcore.HttpClient) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.HttpClient = httpClient
	}
}

func WithHelpdeskCredential(helpdeskID, helpdeskToken string) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.HelpDeskId = helpdeskID
		config.HelpDeskToken = helpdeskToken
		if helpdeskID != "" && helpdeskToken != "" {
			config.HelpdeskAuthToken = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", helpdeskID, helpdeskToken)))
		}
	}
}

func WithSerialization(serializable larkcore.Serializable) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.Serializable = serializable
	}
}

func WithReqTimeout(reqTimeout time.Duration) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.ReqTimeout = reqTimeout
	}
}

// 设置每次请求都会携带的 header
func WithHeaders(header http.Header) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.Header = header
	}
}

func NewClient(appId, appSecret string, options ...ClientOptionFunc) *Client {
	// 构建配置
	config := &larkcore.Config{
		BaseUrl:          FeishuBaseUrl,
		AppId:            appId,
		AppSecret:        appSecret,
		EnableTokenCache: true,
		AppType:          larkcore.AppTypeSelfBuilt,
	}
	for _, option := range options {
		option(config)
	}

	// 构建日志器
	larkcore.NewLogger(config)

	// 构建缓存
	larkcore.NewCache(config)

	// 创建序列化器
	larkcore.NewSerialization(config)

	// 创建httpclient
	larkcore.NewHttpClient(config)

	// 创建sdk-client，并初始化服务
	client := &Client{config: config}
	initService(client, config)

	// 触发重推 appTicket,如果是 ISV 的话
	resendAppTicketIfNeed(client)
	return client
}

func initService(client *Client, config *larkcore.Config) {
	client.Attendance = attendance.NewService(config)
	client.Auth = auth.NewService(config)
	client.Contact = contact.NewService(config)
	client.GrayTestOpenSg = gray_test_open_sg.NewService(config)
	client.Helpdesk = helpdesk.NewService(config)
	client.HumanAuthentication = human_authentication.NewService(config)
	client.Report = report.NewService(config)
	client.SpeechToText = speech_to_text.NewService(config)
	client.Bitable = bitable.NewService(config)
	client.Calendar = calendar.NewService(config)
	client.DocumentAi = document_ai.NewService(config)
	client.Mdm = mdm.NewService(config)
	client.SecurityAndCompliance = security_and_compliance.NewService(config)
	client.Translation = translation.NewService(config)
	client.Verification = verification.NewService(config)
	client.Corehr = corehr.NewService(config)
	client.Hire = hire.NewService(config)
	client.Im = im.NewService(config)
	client.Tenant = tenant.NewService(config)
	client.Admin = admin.NewService(config)
	client.Block = block.NewService(config)
	client.Event = event.NewService(config)
	client.Lingo = lingo.NewService(config)
	client.Okr = okr.NewService(config)
	client.Search = search.NewService(config)
	client.Approval = approval.NewService(config)
	client.Docx = docx.NewService(config)
	client.Ehr = ehr.NewService(config)
	client.MeetingRoom = meeting_room.NewService(config)
	client.OpticalCharRecognition = optical_char_recognition.NewService(config)
	client.Sheets = sheets.NewService(config)
	client.Task = task.NewService(config)
	client.Workplace = workplace.NewService(config)
	client.Passport = passport.NewService(config)
	client.Vc = vc.NewService(config)
	client.Acs = acs.NewService(config)
	client.Application = application.NewService(config)
	client.Authen = authen.NewService(config)
	client.Baike = baike.NewService(config)
	client.Drive = drive.NewService(config)
	client.Mail = mail.NewService(config)
	client.PersonalSettings = personal_settings.NewService(config)
	client.Wiki = wiki.NewService(config)
	client.Ext = larkext.NewService(config)
}

func resendAppTicketIfNeed(client *Client) {
	defer func() {
		if err := recover(); err != nil {
			client.config.Logger.Error(context.Background(), fmt.Sprintf("resendAppTicketIfNeed, error:%v", err))
		}
	}()

	if client.config.AppType == larkcore.AppTypeMarketplace {
		ctx := context.Background()
		resp, err := client.ResendAppTicket(ctx, &larkcore.ResendAppTicketReq{
			AppID:     client.config.AppId,
			AppSecret: client.config.AppSecret,
		})
		if err != nil {
			client.config.Logger.Info(ctx, fmt.Sprintf("resend appTicket error:%v", err))
			return
		}
		client.config.Logger.Info(ctx, fmt.Sprintf("resend appTicket resp:%v", resp))

	}
}

func (cli *Client) Post(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return cli.Do(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   httpPath,
		Body:                      body,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{accessTokeType},
	}, options...)
}

func (cli *Client) Do(ctx context.Context, apiReq *larkcore.ApiReq, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return larkcore.Request(ctx, apiReq, cli.config, options...)
}
func (cli *Client) Get(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return cli.Do(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodGet,
		ApiPath:                   httpPath,
		Body:                      body,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{accessTokeType},
	}, options...)
}

func (cli *Client) Delete(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return cli.Do(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodDelete,
		ApiPath:                   httpPath,
		Body:                      body,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{accessTokeType},
	}, options...)
}

func (cli *Client) Put(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return cli.Do(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPut,
		ApiPath:                   httpPath,
		Body:                      body,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{accessTokeType},
	}, options...)
}

func (cli *Client) Patch(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return cli.Do(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPatch,
		ApiPath:                   httpPath,
		Body:                      body,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{accessTokeType},
	}, options...)
}

func (cli *Client) GetAppAccessTokenBySelfBuiltApp(ctx context.Context, req *larkcore.SelfBuiltAppAccessTokenReq) (*larkcore.AppAccessTokenResp, error) {
	rawResp, err := larkcore.Request(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   larkcore.AppAccessTokenInternalUrlPath,
		Body:                      req,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeNone},
	}, cli.config)

	if err != nil {
		return nil, err
	}
	resp := &larkcore.AppAccessTokenResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return nil, err
	}
	resp.ApiResp = rawResp

	return resp, nil
}

func (cli *Client) GetAppAccessTokenByMarketplaceApp(ctx context.Context, req *larkcore.MarketplaceAppAccessTokenReq) (*larkcore.AppAccessTokenResp, error) {
	rawResp, err := larkcore.Request(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   larkcore.AppAccessTokenUrlPath,
		Body:                      req,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeNone},
	}, cli.config)

	if err != nil {
		return nil, err
	}
	resp := &larkcore.AppAccessTokenResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return nil, err
	}
	resp.ApiResp = rawResp

	return resp, nil
}

func (cli *Client) GetTenantAccessTokenBySelfBuiltApp(ctx context.Context, req *larkcore.SelfBuiltTenantAccessTokenReq) (*larkcore.TenantAccessTokenResp, error) {
	rawResp, err := larkcore.Request(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   larkcore.TenantAccessTokenInternalUrlPath,
		Body:                      req,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeNone},
	}, cli.config)

	if err != nil {
		return nil, err
	}
	resp := &larkcore.TenantAccessTokenResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return nil, err
	}
	resp.ApiResp = rawResp

	return resp, nil
}

func (cli *Client) GetTenantAccessTokenByMarketplaceApp(ctx context.Context, req *larkcore.MarketplaceTenantAccessTokenReq) (*larkcore.TenantAccessTokenResp, error) {
	rawResp, err := larkcore.Request(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   larkcore.TenantAccessTokenUrlPath,
		Body:                      req,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeNone},
	}, cli.config)

	if err != nil {
		return nil, err
	}
	resp := &larkcore.TenantAccessTokenResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return nil, err
	}
	resp.ApiResp = rawResp

	return resp, nil
}

func (cli *Client) ResendAppTicket(ctx context.Context, req *larkcore.ResendAppTicketReq) (*larkcore.ResendAppTicketResp, error) {
	rawResp, err := larkcore.Request(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   larkcore.ApplyAppTicketPath,
		Body:                      req,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeNone},
	}, cli.config)

	if err != nil {
		return nil, err
	}
	resp := &larkcore.ResendAppTicketResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return nil, err
	}
	resp.ApiResp = rawResp

	return resp, nil
}

var FeishuBaseUrl = "https://open.feishu.cn"
var LarkBaseUrl = "https://open.larksuite.com"
