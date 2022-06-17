// Package optical_char_recognition code generated by oapi sdk gen
package optical_char_recognition

import (
	"github.com/feishu/oapi-sdk-go/core"
)

/**生成枚举值 **/

/**生成数据类型 **/

type Image struct {
}

/**生成请求和响应结果类型，以及请求对象的Builder构造器 **/

type BasicRecognizeImageReqBodyBuilder struct {
	image     string
	imageFlag bool
}

// 生成body的New构造器
func NewBasicRecognizeImageReqBodyBuilder() *BasicRecognizeImageReqBodyBuilder {
	builder := &BasicRecognizeImageReqBodyBuilder{}
	return builder
}

/*1.2 生成body的builder属性方法*/
func (builder *BasicRecognizeImageReqBodyBuilder) Image(image string) *BasicRecognizeImageReqBodyBuilder {
	builder.image = image
	builder.imageFlag = true
	return builder
}

/*1.3 生成body的build方法*/
func (builder *BasicRecognizeImageReqBodyBuilder) Build() *BasicRecognizeImageReqBody {
	req := &BasicRecognizeImageReqBody{}
	if builder.imageFlag {
		req.Image = &builder.image

	}
	return req
}

/**上传文件path开始**/
type BasicRecognizeImagePathReqBodyBuilder struct {
	image     string
	imageFlag bool
}

// 生成body的New构造器
func NewBasicRecognizeImagePathReqBodyBuilder() *BasicRecognizeImagePathReqBodyBuilder {
	builder := &BasicRecognizeImagePathReqBodyBuilder{}
	return builder
}

/*1.2 生成body的builder属性方法*/
func (builder *BasicRecognizeImagePathReqBodyBuilder) Image(image string) *BasicRecognizeImagePathReqBodyBuilder {
	builder.image = image
	builder.imageFlag = true
	return builder
}

/*1.3 生成body的build方法*/
func (builder *BasicRecognizeImagePathReqBodyBuilder) Build() (*BasicRecognizeImageReqBody, error) {
	req := &BasicRecognizeImageReqBody{}
	if builder.imageFlag {
		req.Image = &builder.image

	}
	return req, nil
}

/**上传文件path结束**/

/*1.4 生成请求的builder结构体*/
type BasicRecognizeImageReqBuilder struct {
	body     *BasicRecognizeImageReqBody
	bodyFlag bool
}

// 生成请求的New构造器
func NewBasicRecognizeImageReqBuilder() *BasicRecognizeImageReqBuilder {
	builder := &BasicRecognizeImageReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *BasicRecognizeImageReqBuilder) Body(body *BasicRecognizeImageReqBody) *BasicRecognizeImageReqBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *BasicRecognizeImageReqBuilder) Build() *BasicRecognizeImageReq {
	req := &BasicRecognizeImageReq{}
	if builder.bodyFlag {
		req.Body = builder.body

	}
	return req
}

type BasicRecognizeImageReqBody struct {
	Image *string `json:"image,omitempty"`
}

type BasicRecognizeImageReq struct {
	Body *BasicRecognizeImageReqBody `body:""`
}

type BasicRecognizeImageRespData struct {
	TextList []string `json:"text_list,omitempty"`
}

type BasicRecognizeImageResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *BasicRecognizeImageRespData `json:"data"`
}

func (resp *BasicRecognizeImageResp) Success() bool {
	return resp.Code == 0
}

/**生成消息事件结构体 **/

/* 生成请求的builder构造器*/
/*1.1 生成body的builder结构体*/
