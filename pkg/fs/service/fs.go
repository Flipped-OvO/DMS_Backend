package service

import (
	"backend/pkg/fs/model"
	"backend/util/common"
	"backend/util/db"
	"fmt"
)

const (
	companyTB    = "company"
	collectionTB = "companyCollection"
	profitTB     = "profit"
)

// 根据公司或股票代码搜索公司
func SearchCompanyByNameOrCode(s string) (companys []model.Company) {
	searchWord := "%" + s + "%" // 拼装like关键词
	action := db.Action{
		TableName: companyTB,
		Query:     "name LIKE ? OR stock_code LIKE ?",
		Value:     []interface{}{searchWord, searchWord},
	}
	action.QueryAndFind(&companys)

	return
}

// 获取未初始化的公司
func GetUnInitCompany(params common.Params) (companys []model.Company, total int64) {
	limit := params.Limit
	offset := params.Limit * (params.Page - 1)

	action := db.Action{
		TableName: companyTB,
		Query:     "initialization = ?",
		Value:     []interface{}{false},
		Limit:     limit,
		Offset:    offset,
		Total:     &total,
	}

	action.QueryAndPagination(&companys)

	return
}

// 获取收藏的公司
func GetCollectionCompany(params common.Params) (companys []model.Company, total int64) {
	var colls []model.CompanyCollection
	action := db.Action{
		TableName: collectionTB,
		Total:     &total,
	}

	action.Find(&colls)
	codes := make([]string, 0)

	if len(colls) <= 0 {
		return
	}

	for _, coll := range colls {
		codes = append(codes, coll.StockCode)
	}
	companys = getCompanyByCodes(codes, params)
	return
}

// 添加到收藏
func NewCollectionCompany(code string) {
	com := model.CompanyCollection{
		StockCode: code,
	}
	action := db.Action{
		TableName: collectionTB,
	}
	action.InsertOne(&com)
}

// 移除收藏
func RemoveCollectionCompany(code string) {
	action := db.Action{
		TableName: collectionTB,
		Query:     "stock_code = ?",
		Value:     []interface{}{code},
	}

	action.DeleteOne(model.CompanyCollection{})
}

// 获取所有公司
func GetAllCompany() (companys []model.Company) {
	action := db.Action{
		TableName: companyTB,
	}
	action.Find(&companys)

	return
}

// 批量查找公司
func getCompanyByCodes(codes []string, params common.Params) (companys []model.Company) {
	sql := "SELECT * FROM company WHERE stock_code="

	for index, code := range codes {
		if index == 0 {
			sql += fmt.Sprintf("'%s'", code)
		} else {
			sql += fmt.Sprintf("OR stock_code='%s'", code)
		}
	}
	limit := params.Limit
	offset := params.Limit * (params.Page - 1)
	sql += fmt.Sprintf(" LIMIT %d OFFSET %d;", limit, offset)

	action := db.Action{
		TableName: companyTB,
		Sql:       sql,
	}

	fmt.Println(sql)

	action.QueryBySQL(&companys)

	return
}

/*
	-------- 财报 --------
*/

// 获取公司利润表
func GetCompanyProfit(sc string, params common.Params) (profit []model.Profit) {
	action := db.Action{
		TableName: profitTB,
		Query:     "stock_code = ?",
		Value:     []interface{}{sc},
		Offset:    (params.Page - 1) * 5,
		Limit:     5,
		Order:     "standard_time desc",
	}
	action.QueryAndOrderPagination(&profit)
	return
}
