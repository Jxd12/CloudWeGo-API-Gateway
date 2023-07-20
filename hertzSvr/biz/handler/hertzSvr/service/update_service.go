// Code generated by hertz generator.

package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	service "hertzSvr/biz/model/hertzSvr/service"
)

// Update .
// @router /update [GET]
func Update(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.UpdateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(service.UpdateResp)
	err = provider.UpdateIDL(req.GetIdl(), map[string]string{})
	if err != nil {
		panic("Error: fail to update idl " + err.Error())
	}
	resp.Success = true
	resp.Message = "success"

	c.JSON(consts.StatusOK, resp)
}
