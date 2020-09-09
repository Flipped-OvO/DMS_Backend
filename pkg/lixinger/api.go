package lixinger

import (
	FS "backend/pkg/fs/model"
	FSService "backend/pkg/fs/service"
	"backend/pkg/lixinger/model"
	"backend/pkg/lixinger/service"
	"fmt"
	"github.com/imroc/req"
	"log"
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
	companyResponse struct {
		Code    int
		Message string
		Data    []FS.Company
	}
)

var (
	fsType          = [4]string{"non_financial", "bank", "insurance", "security"}
	standardDate    = []string{"03-31", "2017-06-30", "09-30", "12-31"}
	incomeStatement = []string{
		"q.ps.toi.t",
		"q.ps.ii.t",
		"q.ps.toc.t",
		"q.ps.ie.t",
		"q.ps.tas.t",
		"q.ps.se.t",
		"q.ps.ae.t",
		"q.ps.rade.t",
		"q.ps.ieife.t",
		"q.ps.iiife.t",
		"q.ps.ivi.t",
		"q.ps.iifaajv.t",
		"q.ps.iftdofamaac.t",
		"q.ps.ei.t",
		"q.ps.ciofv.t",
		"q.ps.ailor.t",
		"q.ps.oail.t",
		"q.ps.ite.t",
		"q.ps.np.t",
	}
	balanceSheet = []string{
		"q.bs.ta.t",
		"q.bs.cabb.t",
		"q.bs.tfa.t",
		"q.bs.nr.t",
		"q.bs.ar.t",
		"q.bs.ats.t",
		"q.bs.or.t",
		"q.bs.fahursa.t",
		"q.bs.i.t",
		"q.bs.oca.t",
		"q.bs.ncafsfa.t",
		"q.bs.htmi.t",
		"q.bs.ltar.t",
		"q.bs.ltei.t",
		"q.bs.rei.t",
		"q.bs.fa.t",
		"q.bs.cip.t",
		"q.bs.es.t",
		"q.bs.pba.t",
		"q.bs.ia.t",
		"q.bs.gw.t",
		"q.bs.ltpe.t",
		"q.bs.onca.t",
		"q.bs.tl.t",
		"q.bs.lwi.t",
		"q.bs.stl.t",
		"q.bs.bfcb.t",
		"q.bs.np.t",
		"q.bs.ap.t",
		"q.bs.afc.t",
		"q.bs.oap.t",
		"q.bs.ltl.t",
		"q.bs.sap.t",
		"q.bs.ltap.t",
		"q.bs.mc.t",
	}
	cashFlow = []string{
		"q.cfs.crfscapls.t",
		"q.cfs.cpfpcarls.t",
		"q.cfs.ncffoa.t",
		"q.cfs.cpfpfiaolta.t",
		"q.cfs.cpfi.t",
		"q.cfs.ncffia.t",
		"q.cfs.cpfbrp.t",
		"q.cfs.cpfdapdoip.t",
		"q.cfs.ncfffa.t",
		"q.cfs.niicace.t",
		"q.cfs.bocaceatpe.t",
	}
)

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

	var res companyResponse
	err = r.ToJSON(&res)
	if err != nil {
		log.Println(err)
	}
	if res.Code == 0 {
		service.SaveCompanyDatas(res.Data)
	}
}

/*
	获取利润表
*/
func FetchIncomStatement() {
	cps := FSService.GetAllCompany()
	for _, cp := range cps {
		if len(cp.IpoDate) < 10 {
			return
		}
		startDate := cp.IpoDate[:5] + standardDate[0]
		params := req.Param{
			"token":       token,
			"startDate":   startDate,
			"stockCodes":  []string{cp.StockCode},
			"metricsList": incomeStatement,
		}
		r, err := req.Post(url+"/a/stock/fs/non_financial", req.BodyJSON(&params))
		if err != nil {
			log.Println(err)
		}
		var res model.ResponseForProfit
		err = r.ToJSON(&res)
		if err != nil {
			log.Println()
		}
		for _, d := range res.Data {
			fsProfit := FS.Profit{}
			fsProfit.New(d)
			//service.SaveProfit(fsProfit)
		}
	}
	fmt.Println("利润表储存完成")
}

/*
	获取资产负债表
*/
func FetchBalanceSheet() {
	cps := FSService.GetAllCompany()
	for _, cp := range cps {
		if len(cp.IpoDate) < 10 {
			return
		}
		startDate := cp.IpoDate[:5] + standardDate[0]
		params := req.Param{
			"token":       token,
			"startDate":   startDate,
			"stockCodes":  []string{cp.StockCode},
			"metricsList": balanceSheet,
		}
		r, err := req.Post(url+"/a/stock/fs/non_financial", req.BodyJSON(&params))
		if err != nil {
			log.Println(err)
		}
		var res model.ResponseForBalanceSheet
		err = r.ToJSON(&res)
		if err != nil {
			log.Println()
		}
		for _, d := range res.Data {
			fsBs := FS.BalanceSheet{}
			fsBs.New(d)
			service.SaveBalanceSheet(fsBs)
		}
	}
	fmt.Println("资产负债表储存完成")
}

/*
	获取现金流量表
*/
func FetchCashFlow() {
	cps := FSService.GetAllCompany()
	for _, cp := range cps {
		if len(cp.IpoDate) < 10 {
			return
		}
		startDate := cp.IpoDate[:5] + standardDate[0]
		params := req.Param{
			"token":       token,
			"startDate":   startDate,
			"stockCodes":  []string{cp.StockCode},
			"metricsList": cashFlow,
		}
		r, err := req.Post(url+"/a/stock/fs/non_financial", req.BodyJSON(&params))
		if err != nil {
			log.Println(err)
		}
		var res model.ResponseForCashFlow
		err = r.ToJSON(&res)
		if err != nil {
			log.Println()
		}
		for _, d := range res.Data {
			fsCf := FS.CashFlow{}
			fsCf.New(d)
			service.SaveCashFlow(fsCf)
		}
	}
	fmt.Println("现金流量表储存完成")
}
