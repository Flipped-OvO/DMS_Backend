package service

import (
	FS "backend/pkg/fs/model"
	"backend/util/db"
	"fmt"
	"time"
)

const (
	companyTB        = "company"
	profitTB         = "profit"
	balanceTB        = "balanceSheet"
	cashFlowTB       = "cashFlow"
	financialIndexTB = "financialIndex"
)

// 储存公司信息
func SaveCompanyDatas(companys []FS.Company) {
	sql := "INSERT INTO company (stock_code, market, ipo_date, area_code, name, fs_type, initialzation) VALUES "
	for index, c := range companys {
		t, _ := time.Parse(time.RFC3339, c.IpoDate)
		if index == len(companys)-1 {
			sql += fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', %v);", c.StockCode, c.Market, t.Format("2006-01-02"), c.AreaCode, c.Name, c.FsType, false)
		} else {
			sql += fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', %v), ", c.StockCode, c.Market, t.Format("2006-01-02"), c.AreaCode, c.Name, c.FsType, false)
		}
	}

	action := db.Action{
		TableName: companyTB,
		Sql:       sql,
	}

	action.Exce()
}

/*
	储存利润表
*/
func SaveProfit(profit FS.Profit) {
	action := db.Action{
		TableName: profitTB,
	}
	action.InsertOne(&profit)
}

/*
	储存资产负债表
*/
func SaveBalanceSheet(bs FS.BalanceSheet) {
	action := db.Action{
		TableName: balanceTB,
	}
	action.InsertOne(&bs)
}

/*
	储存现金流量表
*/
func SaveCashFlow(cf FS.CashFlow) {
	action := db.Action{
		TableName: cashFlowTB,
	}
	action.InsertOne(&cf)
}

/*
	储存财务指标
*/func SaveFinancialIndex(fi FS.FinancialIndex) {
	action := db.Action{
		TableName: financialIndexTB,
	}
	action.InsertOne(&fi)
}
