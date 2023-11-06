package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/yanguiyuan/lxy/biz/model/lxyhttp"
	"gorm.io/gen"
)

func AddOk(resp *lxyhttp.CommonResponse) {
	resp.Code = 200
	resp.Message = "添加成功"
}
func SqlAddFailed(c *app.RequestContext, err error) {
	hlog.Error(err)
	resp := new(lxyhttp.CommonResponse)
	resp.Code = consts.StatusConflict
	resp.Message = "添加值已存在"
	c.JSON(consts.StatusConflict, resp)
}
func DeleteOk(resp *lxyhttp.CommonResponse) {
	resp.Code = consts.StatusOK
	resp.Message = "删除成功"
}
func SqlDeleteFailed(c *app.RequestContext, err error, info gen.ResultInfo) {
	resp := new(lxyhttp.CommonResponse)
	resp.Code = consts.StatusInternalServerError
	resp.Message = "删除失败"
	c.JSON(consts.StatusInternalServerError, resp)
}
func SqlUpdateFailed(c *app.RequestContext, err error, info gen.ResultInfo) {
	resp := new(lxyhttp.CommonResponse)
	resp.Code = consts.StatusInternalServerError
	resp.Message = "更新失败"
	c.JSON(consts.StatusInternalServerError, resp)
}
func UpdateOk(resp *lxyhttp.CommonResponse) {
	resp.Code = consts.StatusOK
	resp.Message = "更新成功"
}
func SqlQueryFailed(c *app.RequestContext, err error) {
	resp := new(lxyhttp.CommonResponse)
	resp.Code = consts.StatusInternalServerError
	resp.Message = "查询sql失败"
	c.JSON(consts.StatusInternalServerError, resp)
}
