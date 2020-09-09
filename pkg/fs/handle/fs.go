package handle

import (
	"backend/pkg/fs/model"
	"backend/pkg/fs/service"
	"backend/util/common"
	"backend/util/request"
	"encoding/json"
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

// 移除公司收藏
func RemoveCollectionCompany(ctx *gin.Context) {
	code := ctx.DefaultPostForm("code", "")
	if len(code) != 0 {
		service.RemoveCollectionCompany(code)
	}
	request.Success(ctx, nil)
}

/*
	-------- 财报 --------

	Rev  Revenue 收入
	COGS  Cost of Goods Sales 销售成本
	GP  Gross Profit 毛利
	SM  Selling and Marketing 市场推广费用
	GA  General Administration 管理费用
	RD  Research & Development 研发费用
*/

// 获取公司利润表
func GetCompanyProfit(ctx *gin.Context) {
	sc := ctx.DefaultQuery("sc", "")
	if len(sc) == 0 {
		request.Success(ctx, nil)
	}

	var cp common.Params
	if request.ParseParamFail(ctx, &cp) {
		return
	}

	profit := service.GetCompanyProfit(sc, cp)

	response := model.GetProfitTableData(profit)

	request.Success(ctx, response)
}

// 获取公司资产负债表
func GetCompanyBalanceSheet(ctx *gin.Context) {
	sc := ctx.DefaultQuery("sc", "")
	if len(sc) == 0 {
		request.Success(ctx, nil)
	}
	var cp common.Params
	if request.ParseParamFail(ctx, &cp) {
		return
	}

	bs := service.GetCompanyBalanceSheet(sc, cp)

	request.Success(ctx, bs)
}

func SaveEdit(ctx *gin.Context) {
	code := ctx.DefaultPostForm("code", "")
	if len(code) == 0 {
		request.Success(ctx, nil)
	}
	data := ctx.DefaultPostForm("data", "")
	var dataMap map[string]string
	err := json.Unmarshal([]byte(data), &dataMap)
	if request.Fail(ctx, err) {
		return
	}

	service.SaveEdit(dataMap, code)

	request.Success(ctx, nil)
}