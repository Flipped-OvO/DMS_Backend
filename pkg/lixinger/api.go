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
		"q.ps.toi.c",
		"q.ps.oi.c",
		"q.ps.ii.c",
		"q.ps.ep.c",
		"q.ps.faci.c",
		"q.ps.ooi.c",
		"q.ps.toc.c",
		"q.ps.oc.c",
		"q.ps.gp_m.c",
		"q.ps.ie.c",
		"q.ps.face.c",
		"q.ps.s.c",
		"q.ps.ce.c",
		"q.ps.iiicr.c",
		"q.ps.phdrfpip.c",
		"q.ps.rie.c",
		"q.ps.tas.c",
		"q.ps.se.c",
		"q.ps.ae.c",
		"q.ps.rade.c",
		"q.ps.fe.c",
		"q.ps.ieife.c",
		"q.ps.iiife.c",
		"q.ps.se_r.c",
		"q.ps.ae_r.c",
		"q.ps.rade_r.c",
		"q.ps.fe_r.c",
		"q.ps.oe_r.c",
		"q.ps.te_r.c",
		"q.ps.oic.c",
		"q.ps.ivi.c",
		"q.ps.iifaajv.c",
		"q.ps.iftdofamaac.c",
		"q.ps.ei.c",
		"q.ps.nehb.c",
		"q.ps.ciofv.c",
		"q.ps.cilor.c",
		"q.ps.ailor.c",
		"q.ps.oail.c",
		"q.ps.adi.c",
		"q.ps.op.c",
		"q.ps.op_s_r.c",
		"q.ps.op_op_r.c",
		"q.ps.noi.c",
		"q.ps.ncadarg.c",
		"q.ps.noe.c",
		"q.ps.ncadarl.c",
		"q.ps.tp.c",
		"q.ps.rade_tp_r.c",
		"q.ps.ite.c",
		"q.ps.ite_tp_r.c",
		"q.ps.np.c",
		"q.ps.np_s_r.c",
		"q.ps.npfco.c",
		"q.ps.npfdco.c",
		"q.ps.wroe.c",
		"q.ps.npatshaoehopc.c",
		"q.ps.npatoshopc.c",
		"q.ps.npadnrpatoshaopc.c",
		"q.ps.npadnrpatoshaopc_npatoshopc_r.c",
		"q.ps.npatoeihopc.c",
		"q.ps.npatmsh.c",
		"q.ps.natooci.c",
		"q.ps.tci.c",
		"q.ps.tciatshaoehopc.c",
		"q.ps.tciatoshopc.c",
		"q.ps.tciatmsh.c",
		"q.ps.beps.c",
		"q.ps.deps.c",
	}
	balanceSheet = []string{
		"q.bs.ta.c",
		"q.bs.nwc.c",
		"q.bs.tca.c",
		"q.bs.tca_ta_r.c",
		"q.bs.cabb.c",
		"q.bs.cabb_ta_r.c",
		"q.bs.sr.c",
		"q.bs.pwbaofi.c",
		"q.bs.tfa.c",
		"q.bs.cdfa.c",
		"q.bs.nraar.c",
		"q.bs.nr.c",
		"q.bs.ar.c",
		"q.bs.rf.c",
		"q.bs.ats.c",
		"q.bs.pr.c",
		"q.bs.rir.c",
		"q.bs.crorir.c",
		"q.bs.or.c",
		"q.bs.ir.c",
		"q.bs.dr.c",
		"q.bs.fahursa.c",
		"q.bs.i.c",
		"q.bs.ca.c",
		"q.bs.ahfs.c",
		"q.bs.pe.c",
		"q.bs.ncadwioy.c",
		"q.bs.oca.c",
		"q.bs.tca_tcl_r.c",
		"q.bs.q_r.c",
		"q.bs.tnca.c",
		"q.bs.tnca_ta_r.c",
		"q.bs.ah_ta_r.c",
		"q.bs.nclaatc.c",
		"q.bs.cri.c",
		"q.bs.ocri.c",
		"q.bs.ncafsfa.c",
		"q.bs.htmi.c",
		"q.bs.ltar.c",
		"q.bs.ltei.c",
		"q.bs.oeii.c",
		"q.bs.oncfa.c",
		"q.bs.rei.c",
		"q.bs.fa.c",
		"q.bs.dofa.c",
		"q.bs.fa_ta_r.c",
		"q.bs.cip.c",
		"q.bs.es.c",
		"q.bs.cip_fa_r.c",
		"q.bs.pba.c",
		"q.bs.oaga.c",
		"q.bs.pwba.c",
		"q.bs.roua.c",
		"q.bs.ia.c",
		"q.bs.rade.c",
		"q.bs.gw.c",
		"q.bs.gw_toe_r.c",
		"q.bs.ltpe.c",
		"q.bs.dita.c",
		"q.bs.onca.c",
		"q.bs.tl.c",
		"q.bs.lwi.c",
		"q.bs.lwi_ta_r.c",
		"q.bs.tl_ta_r.c",
		"q.bs.tcl.c",
		"q.bs.tcl_tl_r.c",
		"q.bs.stl.c",
		"q.bs.bfcb.c",
		"q.bs.pfbaofi.c",
		"q.bs.tfl.c",
		"q.bs.dfl.c",
		"q.bs.npaap.c",
		"q.bs.np.c",
		"q.bs.ap.c",
		"q.bs.afc.c",
		"q.bs.cl.c",
		"q.bs.fasurpa.c",
		"q.bs.dfcab.c",
		"q.bs.stoa.c",
		"q.bs.ssoa.c",
		"q.bs.sawp.c",
		"q.bs.tp.c",
		"q.bs.oap.c",
		"q.bs.intp.c",
		"q.bs.dp.c",
		"q.bs.facp.c",
		"q.bs.rip.c",
		"q.bs.lhfs.c",
		"q.bs.ncldwioy.c",
		"q.bs.didwioy.c",
		"q.bs.cal.c",
		"q.bs.stbp.c",
		"q.bs.ocl.c",
		"q.bs.tncl.c",
		"q.bs.tncl_tl_r.c",
		"q.bs.icr.c",
		"q.bs.ltl.c",
		"q.bs.bp.c",
		"q.bs.psibp.c",
		"q.bs.pcsibp.c",
		"q.bs.ll.c",
		"q.bs.ltap.c",
		"q.bs.sap.c",
		"q.bs.ltpoe.c",
		"q.bs.ncal.c",
		"q.bs.ltdi.c",
		"q.bs.ditl.c",
		"q.bs.oncl.c",
		"q.bs.toe.c",
		"q.bs.toe_ta_r.c",
		"q.bs.sc.c",
		"q.bs.oei.c",
		"q.bs.psioei.c",
		"q.bs.pcsioei.c",
		"q.bs.capr.c",
		"q.bs.is.c",
		"q.bs.oci.c",
		"q.bs.rr.c",
		"q.bs.surr.c",
		"q.bs.pogr.c",
		"q.bs.rtp.c",
		"q.bs.tetshaoehopc.c",
		"q.bs.tetoshopc.c",
		"q.bs.tetoshopc_ps.c",
		"q.bs.etmsh.c",
		"q.bs.mc.c",
		"q.bs.tsc.c",
		"q.bs.csc.c",
		"q.bs.shn.c",
		"q.bs.shbt1sh_tsc_r.c",
		"q.bs.shbt10sh_tsc_r.c",
		"q.bs.shbt10sh_csc_r.c",
		"q.bs.shbpoof_csc_r.c",
		"q.bs.pe_ttm.c",
		"q.bs.d_pe_ttm.c",
		"q.bs.pb.c",
		"q.bs.pb_wo_gw.c",
		"q.bs.ps_ttm.c",
		"q.bs.dyr.c",
	}
	cashFlow = []string{
		"q.cfs.crfscapls.c",
		"q.cfs.niicdadfbaofi.c",
		"q.cfs.niibfcb.c",
		"q.cfs.niipfofi.c",
		"q.cfs.ndilaatc.c",
		"q.cfs.crfp.c",
		"q.cfs.ncrfrib.c",
		"q.cfs.niiphd.c",
		"q.cfs.ndifahftp.c",
		"q.cfs.crfifac.c",
		"q.cfs.niipfbaofi.c",
		"q.cfs.niifasurpaioa.c",
		"q.cfs.ncrfstoa.c",
		"q.cfs.crfwbot.c",
		"q.cfs.crrtooa.c",
		"q.cfs.stciffoa.c",
		"q.cfs.cpfpcarls.c",
		"q.cfs.niilaatc.c",
		"q.cfs.niibwcbbaofi.c",
		"q.cfs.cpfc.c",
		"q.cfs.niipwbaofi.c",
		"q.cfs.cpfifac.c",
		"q.cfs.cpfphd.c",
		"q.cfs.niifahftp.c",
		"q.cfs.cptofe.c",
		"q.cfs.cpft.c",
		"q.cfs.cprtooa.c",
		"q.cfs.stcoffoa.c",
		"q.cfs.ncffoa.c",
		"q.cfs.crfrci.c",
		"q.cfs.crfii.c",
		"q.cfs.crfdofiaolta.c",
		"q.cfs.ncrfdossaou.c",
		"q.cfs.crrtoia.c",
		"q.cfs.stcifia.c",
		"q.cfs.cpfpfiaolta.c",
		"q.cfs.cpfi.c",
		"q.cfs.niipl.c",
		"q.cfs.ncpfbssaou.c",
		"q.cfs.cprtoia.c",
		"q.cfs.stcoffia.c",
		"q.cfs.ncffia.c",
		"q.cfs.crfai.c",
		"q.cfs.crfamshibss.c",
		"q.cfs.crfl.c",
		"q.cfs.crfib.c",
		"q.cfs.crrtofa.c",
		"q.cfs.stcifffa.c",
		"q.cfs.cpfbrp.c",
		"q.cfs.cpfdapdoip.c",
		"q.cfs.cpfdapomshpbss.c",
		"q.cfs.cprtofa.c",
		"q.cfs.stcofffa.c",
		"q.cfs.ncfffa.c",
		"q.cfs.iocacedtfier.c",
		"q.cfs.niicace.c",
		"q.cfs.bocaceatpb.c",
		"q.cfs.bocaceatpe.c",
		"q.cfs.uril.c",
		"q.cfs.pfai.c",
		"q.cfs.dofaaip.c",
		"q.cfs.aoia.c",
		"q.cfs.aoltpe.c",
		"q.cfs.dipe.c",
		"q.cfs.iiae.c",
		"q.cfs.godofaaonca.c",
		"q.cfs.losofa.c",
		"q.cfs.clofv.c",
	}
	financialIndex = []string{
		"q.m.npatoshopc_ps.c",
		"q.m.npadnrpatoshaopc_ps.c",
		"q.m.tetoshopc_ps.c",
		"q.m.cr_ps.c",
		"q.m.rp_ps.c",
		"q.m.stciffoa_ps.c",
		"q.m.ncffoa_ps.c",
		"q.m.roe_atoshaopc.c",
		"q.m.roe_adnrpatoshaopc.c",
		"q.m.wroe.c",
		"q.m.roe.c",
		"q.m.l.c",
		"q.m.roa.c",
		"q.m.ta_to.c",
		"q.m.np_s_r.c",
		"q.m.gp_r.c",
		"q.m.rota.c",
		"q.m.roic.c",
		"q.m.roc.c",
		"q.m.ats_tor.c",
		"q.m.ca_tor.c",
		"q.m.i_tor.c",
		"q.m.nraar_tor.c",
		"q.m.nr_tor.c",
		"q.m.ar_tor.c",
		"q.m.rf_tor.c",
		"q.m.afc_tor.c",
		"q.m.cl_tor.c",
		"q.m.npaap_tor.c",
		"q.m.np_tor.c",
		"q.m.ap_tor.c",
		"q.m.fa_tor.c",
		"q.m.ats_ds.c",
		"q.m.ca_ds.c",
		"q.m.i_ds.c",
		"q.m.nraar_ds.c",
		"q.m.nr_ds.c",
		"q.m.ar_ds.c",
		"q.m.rf_ds.c",
		"q.m.afc_ds.c",
		"q.m.cl_ds.c",
		"q.m.npaap_ds.c",
		"q.m.np_ds.c",
		"q.m.ap_ds.c",
		"q.m.b_ds.c",
		"q.m.m_ds.c",
		"q.m.fa_ds.c",
		"q.m.tca_ds.c",
		"q.m.toe_ds.c",
		"q.m.tl_ta_r.c",
		"q.m.lwi_ta_r.c",
		"q.m.cabb_tcl_r.c",
		"q.m.c_r.c",
		"q.m.q_r.c",
		"q.m.fa_ta_r.c",
		"q.m.lv_r.c",
		"q.m.fcf.c",
		"q.m.crfscapls_oi_r.c",
		"q.m.ncffoa_op_r.c",
		"q.m.ncffoa_np_r.c",
		"q.m.crfscapls_ta_r.c",
		"q.m.ncffoa_fa_r.c",
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
			service.SaveProfit(fsProfit)
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
			"metricsList": balanceSheet[:70],
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
		params["metricsList"] = balanceSheet[70:]
		r, err = req.Post(url+"/a/stock/fs/non_financial", req.BodyJSON(&params))
		if err != nil {
			log.Println(err)
		}
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

/*
	获取财务指标
*/
func FetchFinancialIndex() {
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
		var res model.ResponseForFinancialIndex
		err = r.ToJSON(&res)
		if err != nil {
			log.Println()
		}
		for _, d := range res.Data {
			fsFi := FS.FinancialIndex{}
			fsFi.New(d)
			service.SaveFinancialIndex(fsFi)
		}
	}
	fmt.Println("财务指标储存完成")
}
