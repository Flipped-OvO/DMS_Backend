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
)

// 根据公司或股票代码搜索公司
func SearchCompanyByNameOrCode(s string) (companys []model.Company) {
	conn := db.NewConnection(companyTB)
	defer conn.Close()
	searchWord := "%" + s + "%" // 拼装like关键词
	conn.Where("name LIKE ? OR stock_code LIKE ?", searchWord, searchWord).Find(&companys)
	return
}

// 获取未初始化的公司
func GetUnInitCompany(params common.Params) (companys []model.Company, total int64) {
	conn := db.NewConnection(companyTB)
	defer conn.Close()
	limit := params.Limit
	offset := params.Limit * (params.Page - 1)
	conn.Where("initialization = ?", false).Limit(limit).Offset(offset).Find(&companys)
	conn.Where("initialization = ?", false).Count(&total)
	return
}

// 获取收藏的公司
func GetCollectionCompany(params common.Params) (companys []model.Company, total int64) {
	conn := db.NewConnection(collectionTB)
	defer conn.Close()
	var colls []model.CompanyCollection
	conn.Find(&colls)
	codes := make([]string, 0)
	for _, coll := range colls {
		codes = append(codes, coll.StockCode)
	}
	companys = getCompanyByCodes(codes, params)
	total = int64(len(colls))
	return
}

// 添加到收藏
func NewCollectionCompany(code string) {
	conn := db.NewConnection(collectionTB)
	defer conn.Close()
	com := model.CompanyCollection{
		StockCode: code,
	}
	conn.Create(&com)
}

// 移除收藏
func RemoveCollectionCompany(code string) {
	conn := db.NewConnection(collectionTB)
	defer conn.Close()
	conn.Where("stock_code = ?", code).Delete(model.CompanyCollection{})
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
	conn := db.NewConnection(companyTB)
	defer conn.Close()
	limit := params.Limit
	offset := params.Limit * (params.Page - 1)
	sql += fmt.Sprintf(" LIMIT %d OFFSET %d;", limit, offset)
	conn.Raw(sql).Scan(&companys)
	return
}
