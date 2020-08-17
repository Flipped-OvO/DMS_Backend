package handle

import (
	"backend/pkg/fs/model"
	"backend/pkg/fs/service"
	"backend/util/common"
	"backend/util/request"
	"github.com/gin-gonic/gin"
)

// 根据输入 搜索公司
func SearchCompanyByNameOrCode(ctx *gin.Context) {
	s := ctx.DefaultQuery("s", "")
	if len(s) == 0 {
		request.Success(ctx, []model.Company{})
		return
	}
	cs := service.SearchCompanyByNameOrCode(s)
	request.Success(ctx, cs)
}

// 获取未初始化的公司
func GetUnInitCompany(ctx *gin.Context) {
	var cp common.Params
	if request.ParseParamFail(ctx, &cp) {
		return
	}
	cs, total := service.GetUnInitCompany(cp)
	request.Success(ctx, cs, total)
}

// 获取收藏的公司
func GetCollectionCompany(ctx *gin.Context) {
	var cp common.Params
	if request.ParseParamFail(ctx, &cp) {
		return
	}
	cs, total := service.GetCollectionCompany(cp)
	request.Success(ctx, cs, total)
}

// 添加收藏
func NewCollectionCompany(ctx *gin.Context) {
	code := ctx.DefaultPostForm("code", "")
	if len(code) != 0 {
		service.NewCollectionCompany(code)
	}
	request.Success(ctx, nil)
}

func RemoveCollectionCompany(ctx *gin.Context) {
	code := ctx.DefaultPostForm("code", "")
	if len(code) != 0 {
		service.RemoveCollectionCompany(code)
	}
	request.Success(ctx, nil)
}
