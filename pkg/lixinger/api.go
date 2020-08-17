package lixinger

import (
	FS "backend/pkg/fs/model"
	"backend/util/db"
	"fmt"
	"github.com/imroc/req"
	"log"
	"time"
)

const (
	token = "af34c8c6-ea4b-41d0-8e13-59c0cd5091e3"
	url   = "https://open.lixinger.com/api"
)
const (
	nonFinancial int = iota // 非金融
	bank                    // 银行
	insurance               // 股票
	security                // 证券
)

type (
	response struct {
		Code    int
		Message string
		Data    []FS.Company
	}
)

var fsType = [4]string{"non_financial", "bank", "insurance", "security"}

// 每日更新公司数据并储存

// 获取A股公司数据
func FetchACompanyData() {
	param := req.Param{
		"token":  token,
		"fsType": fsType[nonFinancial],
	}

	r, err := req.Post(url+"/a/stock", param)
	if err != nil {
		log.Println(err)
	}

	var res response
	err = r.ToJSON(&res)
	if err != nil {
		log.Println(err)
	}
	if res.Code == 0 {
		saveCompanyDatas(res.Data)
	}
}

func saveCompanyDatas(companys []FS.Company) {
	sql := "INSERT INTO company (stock_code, market, ipo_date, area_code, name, fs_type, initialzation) VALUES "
	for index, c := range companys {
		t, _ := time.Parse(time.RFC3339, c.IpoDate)
		if index == len(companys)-1 {
			sql += fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', %v);", c.StockCode, c.Market, t.Format("2006-01-02"), c.AreaCode, c.Name, c.FsType, false)
		} else {
			sql += fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', %v), ", c.StockCode, c.Market, t.Format("2006-01-02"), c.AreaCode, c.Name, c.FsType, false)
		}
	}

	coll := db.NewConnection("company")
	defer coll.Close()
	fmt.Println(sql)
	coll.Exec(sql)
}
