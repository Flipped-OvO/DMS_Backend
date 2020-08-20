package model

import (
	"backend/pkg/lixinger/model"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type (
	// 公司
	Company struct {
		StockCode string `json:"stockCode" gorm:"column:stock_code;primary_key"` // 股票代码
		Market    string `json:"market" gorm:"column:market"`                    // 市场类型 a
		IpoDate   string `json:"ipoDate" gorm:"column:ipo_date"`                 // IPO时间
		AreaCode  string `json:"areaCode" gorm:"column:area_code"`               // 区域代码 cn
		Name      string `json:"name" gorm:"column:name"`                        // 公司名称
		FsType    string `json:"fsType" gorm:"column:fs_type"`                   // 财报类型

		// TODO 公司个性化配置项
		Initialization bool           `json:"initialization" gorm:"column:initialization"` //是否初始化
		Structure      postgres.Jsonb `json:"structure" gorm:"column:structure"`
	}

	// 公司收藏
	CompanyCollection struct {
		StockCode string `json:"stockCode" gorm:"column:stock_code;primary_key"` // 股票代码
	}

	// 利润表
	Profit struct {
		Index                       int64   `gorm:"column:index;primary_key"`
		StockCode                   string  `json:"stockCode" gorm:"column:stock_code"`                                        // 股票代码
		Currency                    string  `json:"currency" gorm:"column:currency"`                                           // 货币
		StandardDate                string  `json:"standardDate" gorm:"column:standard_date"`                                  // 标准日期
		StandardTime                int64   `json:"standardTime" gorm:"column:standard_time"`                                  // 标准日期时间戳
		ReportDate                  string  `json:"reportDate" gorm:"column:report_date"`                                      // 报告时间
		ReportType                  string  `json:"reportType" gorm:"column:report_type"`                                      // 报告类型
		Toi                         float64 `json:"toi" gorm:"column:toi"`                                                     // 营业总收入
		Oi                          float64 `json:"oi" gorm:"column:oi"`                                                       // 营业收入
		Ii                          float64 `json:"ii" gorm:"column:ii"`                                                       // 利息收入
		Ep                          float64 `json:"ep" gorm:"column:ep"`                                                       // 已赚保费
		Faci                        float64 `json:"faci" gorm:"column:faci"`                                                   // 手续费及佣金收入
		Ooi                         float64 `json:"ooi" gorm:"column:ooi"`                                                     // 其他业务收入
		Toc                         float64 `json:"toc" gorm:"column:toc"`                                                     // 营业总成本
		Oc                          float64 `json:"oc" gorm:"column:oc"`                                                       // 营业成本
		GpM                         float64 `json:"gp_m" gorm:"column:gp_m"`                                                   // 毛利率
		Ie                          float64 `json:"ie" gorm:"column:ie"`                                                       // 利息支出
		Face                        float64 `json:"face" gorm:"column:face"`                                                   // 手续费及佣金支出
		S                           float64 `json:"s" gorm:"column:s"`                                                         // 退保金
		Ce                          float64 `json:"ce" gorm:"column:ce"`                                                       // 保险合同赔付支出
		Iiicr                       float64 `json:"iiicr" gorm:"column:iiicr"`                                                 // 提取保险责任准备金净额
		Phdrfpip                    float64 `json:"phdrfpip" gorm:"column:phdrfpip"`                                           // 保单红利支出
		Rie                         float64 `json:"rie" gorm:"column:rie"`                                                     // 分保费用
		Tas                         float64 `json:"tas" gorm:"column:tas"`                                                     // 税金及附加
		Se                          float64 `json:"se" gorm:"column:se"`                                                       // 销售费用
		Ae                          float64 `json:"ae" gorm:"column:ae"`                                                       // 管理费用
		Rade                        float64 `json:"rade" gorm:"column:rade"`                                                   // 研发费用
		Fe                          float64 `json:"fe" gorm:"column:fe"`                                                       // 财务费用
		Ieife                       float64 `json:"ieife" gorm:"column:ieife"`                                                 // 利息费用
		Iiife                       float64 `json:"iiife" gorm:"column:iiife"`                                                 // 利息收入
		SeR                         float64 `json:"se_r" gorm:"column:se_r"`                                                   // 销售费用率
		AeR                         float64 `json:"ae_r" gorm:"column:ae_r"`                                                   // 管理费用率
		RadeR                       float64 `json:"rade_r" gorm:"column:rade_r"`                                               // 研发费用率
		FeR                         float64 `json:"fe_r" gorm:"column:fe_r"`                                                   // 财务费用率
		OeR                         float64 `json:"oe_r" gorm:"column:oe_r"`                                                   // 总营业费用率
		TeR                         float64 `json:"te_r" gorm:"column:te_r"`                                                   // 三项费用率
		Oic                         float64 `json:"oic" gorm:"column:oic"`                                                     // 其他收益
		Ivi                         float64 `json:"ivi" gorm:"column:ivi"`                                                     // 投资收益
		Iifaajv                     float64 `json:"iifaajv" gorm:"column:iifaajv"`                                             // 对联营企业及合营企业的投资收益
		Iftdofamaac                 float64 `json:"iftdofamaac" gorm:"column:iftdofamaac"`                                     // 以摊余成本计量的金融资产终止确认产生的投资收益
		Ei                          float64 `json:"ei" gorm:"column:ei"`                                                       // 汇兑收益
		Nehb                        float64 `json:"nehb" gorm:"column:nehb"`                                                   // 净敞口套期收益
		Ciofv                       float64 `json:"ciofv" gorm:"column:ciofv"`                                                 // 公允价值变动收益
		Cilor                       float64 `json:"cilor" gorm:"column:cilor"`                                                 // 信用减值损失
		Ailor                       float64 `json:"ailor" gorm:"column:ailor"`                                                 // 资产减值损失
		Oail                        float64 `json:"oail" gorm:"column:oail"`                                                   // 其他资产减值损失
		Adi                         float64 `json:"adi" gorm:"column:adi"`                                                     // 资产处置收益
		Op                          float64 `json:"op" gorm:"column:op"`                                                       // 营业利润
		OpSR                        float64 `json:"op_s_r" gorm:"column:op_s_r"`                                               // 营业利润率
		OpOpR                       float64 `json:"op_op_r" gorm:"column:op_op_r"`                                             // 其他营业利润占比
		Noi                         float64 `json:"noi" gorm:"column:noi"`                                                     // 营业外收入
		Ncadarg                     float64 `json:"ncadarg" gorm:"column:ncadarg"`                                             // 非流动资产毁损报废利得
		Noe                         float64 `json:"noe" gorm:"column:noe"`                                                     // 营业外支出
		Ncadarl                     float64 `json:"ncadarl" gorm:"column:ncadarl"`                                             // 非流动资产毁损报废损失
		Tp                          float64 `json:"tp" gorm:"column:tp"`                                                       // 利润总额
		RadeTpR                     float64 `json:"rade_tp_r" gorm:"column:rade_tp_r"`                                         // 研发费占利润总额比值
		Ite                         float64 `json:"ite" gorm:"column:ite"`                                                     // 所得税费用
		IteTpR                      float64 `json:"ite_tp_r" gorm:"column:ite_tp_r"`                                           // 有效税率
		Np                          float64 `json:"np" gorm:"column:np"`                                                       // 净利润
		NpSR                        float64 `json:"np_s_r" gorm:"column:np_s_r"`                                               // 净利润率
		Npfco                       float64 `json:"npfco" gorm:"column:npfco"`                                                 // 持续经营净利润
		Npfdco                      float64 `json:"npfdco" gorm:"column:npfdco"`                                               // 终止经营净利润
		Wroe                        float64 `json:"wroe" gorm:"column:wroe"`                                                   // 归属于母公司普通股股东的加权ROE
		Npatshaoehopc               float64 `json:"npatshaoehopc" gorm:"column:npatshaoehopc"`                                 // 归属于母公司股东及其他权益持有者的净利润
		Npatoshopc                  float64 `json:"npatoshopc" gorm:"column:npatoshopc"`                                       // 归属于母公司普通股股东的净利润
		Npadnrpatoshaopc            float64 `json:"npadnrpatoshaopc" gorm:"column:npadnrpatoshaopc"`                           // 归属于母公司普通股股东的扣除非经常性损益的净利润
		NpadnrpatoshaopcNpatoshopcR float64 `json:"npadnrpatoshaopc_npatoshopc_r" gorm:"column:npadnrpatoshaopc_npatoshopc_r"` // 扣非净利润占比
		Npatoeihopc                 float64 `json:"npatoeihopc" gorm:"column:npatoeihopc"`                                     // 归属于母公司其他权益工具持有者的净利润
		Npatmsh                     float64 `json:"npatmsh" gorm:"column:npatmsh"`                                             // 少数股东损益
		Natooci                     float64 `json:"natooci" gorm:"column:natooci"`                                             // 其他综合收益的税后净额
		Tci                         float64 `json:"tci" gorm:"column:tci"`                                                     // 综合收益总额
		Tciatshaoehopc              float64 `json:"tciatshaoehopc" gorm:"column:tciatshaoehopc"`                               // 归属于母公司股东及其他权益持有者的综合收益总额
		Tciatoshopc                 float64 `json:"tciatoshopc" gorm:"column:tciatoshopc"`                                     // 归属于母公司普通股股东的综合收益总额
		Tciatmsh                    float64 `json:"tciatmsh" gorm:"column:tciatmsh"`                                           // 归属于少数股东的综合收益总额
		Beps                        float64 `json:"beps" gorm:"column:beps"`                                                   // 基本每股收益
		Deps                        float64 `json:"deps" gorm:"column:deps"`                                                   // 稀释每股收益
	}

	// 资产负债表
	BalanceSheet struct {
		Index        int64   `gorm:"column:index;primary_key"`
		StockCode    string  `json:"stockCode" gorm:"column:stock_code"`          // 股票代码
		Currency     string  `json:"currency" gorm:"column:currency"`             // 货币
		StandardDate string  `json:"standardDate" gorm:"column:standard_date"`    // 标准日期
		StandardTime int64   `json:"standardTime" gorm:"column:standard_time"`    // 标准日期时间戳
		ReportDate   string  `json:"reportDate" gorm:"column:report_date"`        // 报告时间
		ReportType   string  `json:"reportType" gorm:"column:report_type"`        // 报告类型
		Ta           float64 `json:"ta" gorm:"column:ta"`                         // 资产合计
		Nwc          float64 `json:"nwc" gorm:"column:nwc"`                       // 净营运资本
		Tca          float64 `json:"tca" gorm:"column:tca"`                       // 流动资产合计
		TcaTaR       float64 `json:"tca_ta_r" gorm:"column:tca_ta_r"`             // 流动资产占比
		Cabb         float64 `json:"cabb" gorm:"column:cabb"`                     // 货币资金
		CabbTaR      float64 `json:"cabb_ta_r" gorm:"column:cabb_ta_r"`           // 货币资金占比
		Sr           float64 `json:"sr" gorm:"column:sr"`                         // 结算备付金
		Pwbaofi      float64 `json:"pwbaofi" gorm:"column:pwbaofi"`               // 拆出资金
		Tfa          float64 `json:"tfa" gorm:"column:tfa"`                       // 交易性金融资产
		Cdfa         float64 `json:"cdfa" gorm:"column:cdfa"`                     // 衍生金融资产(流动)
		Nraar        float64 `json:"nraar" gorm:"column:nraar"`                   // 应收票据及应收账款
		Nr           float64 `json:"nr" gorm:"column:nr"`                         // 应收票据
		Ar           float64 `json:"ar" gorm:"column:ar"`                         // 应收账款
		Rf           float64 `json:"rf" gorm:"column:rf"`                         // 应收款项融资
		Ats          float64 `json:"ats" gorm:"column:ats"`                       // 预付款项
		Pr           float64 `json:"pr" gorm:"column:pr"`                         // 应收保费
		Rir          float64 `json:"rir" gorm:"column:rir"`                       // 应收分保账款
		Crorir       float64 `json:"crorir" gorm:"column:crorir"`                 // 应收分保合同准备金
		Or           float64 `json:"or" gorm:"column:or"`                         // 其他应收款
		Ir           float64 `json:"ir" gorm:"column:ir"`                         // 应收利息
		Dr           float64 `json:"dr" gorm:"column:dr"`                         // 应收股利
		Fahursa      float64 `json:"fahursa" gorm:"column:fahursa"`               // 买入返售金融资产
		I            float64 `json:"i" gorm:"column:i"`                           // 存货
		Ca           float64 `json:"ca" gorm:"column:ca"`                         // 合同资产
		Ahfs         float64 `json:"ahfs" gorm:"column:ahfs"`                     // 持有待售资产
		Pe           float64 `json:"pe" gorm:"column:pe"`                         // 待摊费用
		Ncadwioy     float64 `json:"ncadwioy" gorm:"column:ncadwioy"`             // 一年内到期的非流动资产
		Oca          float64 `json:"oca" gorm:"column:oca"`                       // 其他流动资产
		TcaTclR      float64 `json:"tca_tcl_r" gorm:"column:tca_tcl_r"`           // 流动比率
		QR           float64 `json:"q_r" gorm:"column:q_r"`                       // 速动比率
		Tnca         float64 `json:"tnca" gorm:"column:tnca"`                     // 非流动资产合计
		TncaTaR      float64 `json:"tnca_ta_r" gorm:"column:tnca_ta_r"`           // 非流动资产占比
		AhTaR        float64 `json:"ah_ta_r" gorm:"column:ah_ta_r"`               // 重资产占比
		Nclaatc      float64 `json:"nclaatc" gorm:"column:nclaatc"`               // 发放贷款及垫款(非流动)
		Cri          float64 `json:"cri" gorm:"column:cri"`                       // 债权投资
		Ocri         float64 `json:"ocri" gorm:"column:ocri"`                     // 其他债权投资
		Ncafsfa      float64 `json:"ncafsfa" gorm:"column:ncafsfa"`               // 可供出售金融资产(非流动)
		Htmi         float64 `json:"htmi" gorm:"column:htmi"`                     // 持有至到期投资
		Ltar         float64 `json:"ltar" gorm:"column:ltar"`                     // 长期应收款
		Ltei         float64 `json:"ltei" gorm:"column:ltei"`                     // 长期股权投资
		Oeii         float64 `json:"oeii" gorm:"column:oeii"`                     // 其他权益工具投资
		Oncfa        float64 `json:"oncfa" gorm:"column:oncfa"`                   // 其他非流动金融资产
		Rei          float64 `json:"rei" gorm:"column:rei"`                       // 投资性房地产
		Fa           float64 `json:"fa" gorm:"column:fa"`                         // 固定资产
		Dofa         float64 `json:"dofa" gorm:"column:dofa"`                     // 固定资产清理
		FaTaR        float64 `json:"fa_ta_r" gorm:"column:fa_ta_r"`               // 固定资产占总资产比率
		Cip          float64 `json:"cip" gorm:"column:cip"`                       // 在建工程
		Es           float64 `json:"es" gorm:"column:es"`                         // 工程物资
		CipFaR       float64 `json:"cip_fa_r" gorm:"column:cip_fa_r"`             // 在建工程占固定资产比率
		Pba          float64 `json:"pba" gorm:"column:pba"`                       // 生产性生物资产
		Oaga         float64 `json:"oaga" gorm:"column:oaga"`                     // 油气资产
		Pwba         float64 `json:"pwba" gorm:"column:pwba"`                     // 公益性生物资产
		Roua         float64 `json:"roua" gorm:"column:roua"`                     // 使用权资产
		Ia           float64 `json:"ia" gorm:"column:ia"`                         // 无形资产
		Rade         float64 `json:"rade" gorm:"column:rade"`                     // 开发支出
		Gw           float64 `json:"gw" gorm:"column:gw"`                         // 商誉
		GwToeR       float64 `json:"gw_toe_r" gorm:"column:gw_toe_r"`             // 商誉占净资产比率
		Ltpe         float64 `json:"ltpe" gorm:"column:ltpe"`                     // 长期待摊费用
		Dita         float64 `json:"dita" gorm:"column:dita"`                     // 递延所得税资产
		Onca         float64 `json:"onca" gorm:"column:onca"`                     // 其他非流动资产
		Tl           float64 `json:"tl" gorm:"column:tl"`                         // 负债合计
		Lwi          float64 `json:"lwi" gorm:"column:lwi"`                       // 有息负债
		LwiTaR       float64 `json:"lwi_ta_r" gorm:"column:lwi_ta_r"`             // 有息负债率
		TlTaR        float64 `json:"tl_ta_r" gorm:"column:tl_ta_r"`               // 资产负债率
		Tcl          float64 `json:"tcl" gorm:"column:tcl"`                       // 流动负债合计
		TclTlR       float64 `json:"tcl_tl_r" gorm:"column:tcl_tl_r"`             // 流动负债占比
		Stl          float64 `json:"stl" gorm:"column:stl"`                       // 短期借款
		Bfcb         float64 `json:"bfcb" gorm:"column:bfcb"`                     // 向中央银行借款
		Pfbaofi      float64 `json:"pfbaofi" gorm:"column:pfbaofi"`               // 拆入资金
		Tfl          float64 `json:"tfl" gorm:"column:tfl"`                       // 交易性金融负债
		Dfl          float64 `json:"dfl" gorm:"column:dfl"`                       // 衍生金融负债
		Npaap        float64 `json:"npaap" gorm:"column:npaap"`                   // 应付票据及应付账款
		Np           float64 `json:"np" gorm:"column:np"`                         // 应付票据
		Ap           float64 `json:"ap" gorm:"column:ap"`                         // 应付账款
		Afc          float64 `json:"afc" gorm:"column:afc"`                       // 预收账款
		Cl           float64 `json:"cl" gorm:"column:cl"`                         // 合同负债
		Fasurpa      float64 `json:"fasurpa" gorm:"column:fasurpa"`               // 卖出回购金融资产
		Dfcab        float64 `json:"dfcab" gorm:"column:dfcab"`                   // 吸收存款及同业存放
		Stoa         float64 `json:"stoa" gorm:"column:stoa"`                     // 代理买卖证券款
		Ssoa         float64 `json:"ssoa" gorm:"column:ssoa"`                     // 代理承销证券款
		Sawp         float64 `json:"sawp" gorm:"column:sawp"`                     // 应付职工薪酬
		Tp           float64 `json:"tp" gorm:"column:tp"`                         // 应交税费
		Oap          float64 `json:"oap" gorm:"column:oap"`                       // 其他应付款
		Intp         float64 `json:"intp" gorm:"column:intp"`                     // 应付利息
		Dp           float64 `json:"dp" gorm:"column:dp"`                         // 应付股利
		Facp         float64 `json:"facp" gorm:"column:facp"`                     // 应付手续费及佣金
		Rip          float64 `json:"rip" gorm:"column:rip"`                       // 应付分保账款
		Lhfs         float64 `json:"lhfs" gorm:"column:lhfs"`                     // 持有待售负债
		Ncldwioy     float64 `json:"ncldwioy" gorm:"column:ncldwioy"`             // 一年内到期的非流动负债
		Didwioy      float64 `json:"didwioy" gorm:"column:didwioy"`               // 一年内到期的递延收益
		Cal          float64 `json:"cal" gorm:"column:cal"`                       // 预计负债(流动)
		Stbp         float64 `json:"stbp" gorm:"column:stbp"`                     // 短期应付债券
		Ocl          float64 `json:"ocl" gorm:"column:ocl"`                       // 其他流动负债
		Tncl         float64 `json:"tncl" gorm:"column:tncl"`                     // 非流动负债合计
		TnclTlR      float64 `json:"tncl_tl_r" gorm:"column:tncl_tl_r"`           // 非流动负债占比
		Icr          float64 `json:"icr" gorm:"column:icr"`                       // 保险合同准备金
		Ltl          float64 `json:"ltl" gorm:"column:ltl"`                       // 长期借款
		Bp           float64 `json:"bp" gorm:"column:bp"`                         // 应付债券
		Psibp        float64 `json:"psibp" gorm:"column:psibp"`                   // 优先股
		Pcsibp       float64 `json:"pcsibp" gorm:"column:pcsibp"`                 // 永续债
		Ll           float64 `json:"ll" gorm:"column:ll"`                         // 租赁负债
		Ltap         float64 `json:"ltap" gorm:"column:ltap"`                     // 长期应付款
		Sap          float64 `json:"sap" gorm:"column:sap"`                       // 专项应付款
		Ltpoe        float64 `json:"ltpoe" gorm:"column:ltpoe"`                   // 长期应付职工薪酬
		Ncal         float64 `json:"ncal" gorm:"column:ncal"`                     // 预计负债(非流动)
		Ltdi         float64 `json:"ltdi" gorm:"column:ltdi"`                     // 长期递延收益
		Ditl         float64 `json:"ditl" gorm:"column:ditl"`                     // 递延所得税负债
		Oncl         float64 `json:"oncl" gorm:"column:oncl"`                     // 其他非流动负债
		Toe          float64 `json:"toe" gorm:"column:toe"`                       // 股东权益合计
		ToeTaR       float64 `json:"toe_ta_r" gorm:"column:toe_ta_r"`             // 股东权益占比
		Sc           float64 `json:"sc" gorm:"column:sc"`                         // 股本
		Oei          float64 `json:"oei" gorm:"column:oei"`                       // 其他权益工具
		Psioei       float64 `json:"psioei" gorm:"column:psioei"`                 // 优先股
		Pcsioei      float64 `json:"pcsioei" gorm:"column:pcsioei"`               // 永续债
		Capr         float64 `json:"capr" gorm:"column:capr"`                     // 资本公积
		Is           float64 `json:"is" gorm:"column:is"`                         // 库存股
		Oci          float64 `json:"oci" gorm:"column:oci"`                       // 其他综合收益
		Rr           float64 `json:"rr" gorm:"column:rr"`                         // 专项储备
		Surr         float64 `json:"surr" gorm:"column:surr"`                     // 盈余公积
		Pogr         float64 `json:"pogr" gorm:"column:pogr"`                     // 一般风险准备金
		Rtp          float64 `json:"rtp" gorm:"column:rtp"`                       // 未分配利润
		Tetshaoehopc float64 `json:"tetshaoehopc" gorm:"column:tetshaoehopc"`     // 归属于母公司股东及其他权益持有者的权益合计
		Tetoshopc    float64 `json:"tetoshopc" gorm:"column:tetoshopc"`           // 归属于母公司普通股股东权益合计
		TetoshopcPs  float64 `json:"tetoshopc_ps" gorm:"column:tetoshopc_ps"`     // 归属于母公司普通股股东的每股股东权益
		Etmsh        float64 `json:"etmsh" gorm:"column:etmsh"`                   // 少数股东权益
		Mc           float64 `json:"mc" gorm:"column:mc"`                         // 市值
		Tsc          float64 `json:"tsc" gorm:"column:tsc"`                       // 总股本
		Csc          float64 `json:"csc" gorm:"column:csc"`                       // 流通股本
		Shn          float64 `json:"shn" gorm:"column:shn"`                       // 股东人数(季度)
		Shbt1shTscR  float64 `json:"shbt1sh_tsc_r" gorm:"column:shbt1sh_tsc_r"`   // 第一大股东持仓占总股本比例
		Shbt10shTscR float64 `json:"shbt10sh_tsc_r" gorm:"column:shbt10sh_tsc_r"` // 前十大股东持仓占总股本比例
		Shbt10shCscR float64 `json:"shbt10sh_csc_r" gorm:"column:shbt10sh_csc_r"` // 前十大流通股东持仓占流通股本比例
		ShbpoofCscR  float64 `json:"shbpoof_csc_r" gorm:"column:shbpoof_csc_r"`   // 公募基金持仓占流通股本比例
		PeTtm        float64 `json:"pe_ttm" gorm:"column:pe_ttm"`                 // PE-TTM
		DPeTtm       float64 `json:"d_pe_ttm" gorm:"column:d_pe_ttm"`             // PE-TTM(扣非)
		Pb           float64 `json:"pb" gorm:"column:pb"`                         // PB
		PbWoGw       float64 `json:"pb_wo_gw" gorm:"column:pb_wo_gw"`             // PB(不含商誉)
		PsTtm        float64 `json:"ps_ttm" gorm:"column:ps_ttm"`                 // PS-TTM
		Dyr          float64 `json:"dyr" gorm:"column:dyr"`                       // 股息率
	}

	// 现金流量表
	CashFlow struct {
		Index          int64   `gorm:"column:index;primary_key"`
		StockCode      string  `json:"stockCode" gorm:"column:stock_code;"`         // 股票代码
		Currency       string  `gorm:"column:currency"`                             // 货币
		StandardDate   string  `json:"standardDate" gorm:"column:standard_date"`    // 标准日期
		StandardTime   int64   `json:"standardTime" gorm:"column:standard_time"`    // 标准日期时间戳
		ReportDate     string  `json:"reportDate" gorm:"column:report_date"`        // 报告时间
		ReportType     string  `json:"reportType" gorm:"column:report_type"`        // 报告类型
		Crfscapls      float64 `json:"crfscapls" gorm:"column:crfscapls"`           // 销售商品、提供劳务收到的现金
		Niicdadfbaofi  float64 `json:"niicdadfbaofi" gorm:"column:niicdadfbaofi"`   // 客户存款和同业及其他金融机构存放款项净增加额
		Niibfcb        float64 `json:"niibfcb" gorm:"column:niibfcb"`               // 向中央银行借款净增加额
		Niipfofi       float64 `json:"niipfofi" gorm:"column:niipfofi"`             // 向其他金融机构拆入资金净增加额
		Ndilaatc       float64 `json:"ndilaatc" gorm:"column:ndilaatc"`             // 发放贷款及垫款的净减少额
		Crfp           float64 `json:"crfp" gorm:"column:crfp"`                     // 收到原保险合同保费取得的现金
		Ncrfrib        float64 `json:"ncrfrib" gorm:"column:ncrfrib"`               // 收到再保险业务现金净额
		Niiphd         float64 `json:"niiphd" gorm:"column:niiphd"`                 // 保户储金及投资款净增加额
		Ndifahftp      float64 `json:"ndifahftp" gorm:"column:ndifahftp"`           // 为交易目的而持有的金融资产净减少额
		Crfifac        float64 `json:"crfifac" gorm:"column:crfifac"`               // 收取利息、手续费及佣金的现金
		Niipfbaofi     float64 `json:"niipfbaofi" gorm:"column:niipfbaofi"`         // 拆入资金净增加额
		Niifasurpaioa  float64 `json:"niifasurpaioa" gorm:"column:niifasurpaioa"`   // 回购业务资金净增加额
		Ncrfstoa       float64 `json:"ncrfstoa" gorm:"column:ncrfstoa"`             // 代理买卖证券收到的现金净额
		Crfwbot        float64 `json:"crfwbot" gorm:"column:crfwbot"`               // 收到的税费返还
		Crrtooa        float64 `json:"crrtooa" gorm:"column:crrtooa"`               // 收到的其他与经营活动有关现金
		Stciffoa       float64 `json:"stciffoa" gorm:"column:stciffoa"`             // 经营活动现金流入小计
		Cpfpcarls      float64 `json:"cpfpcarls" gorm:"column:cpfpcarls"`           // 购买商品、接收劳务支付的现金
		Niilaatc       float64 `json:"niilaatc" gorm:"column:niilaatc"`             // 发放贷款和垫款的净增加额
		Niibwcbbaofi   float64 `json:"niibwcbbaofi" gorm:"column:niibwcbbaofi"`     // 存放中央银行和同业及其他金融机构款项净增加额
		Cpfc           float64 `json:"cpfc" gorm:"column:cpfc"`                     // 支付原保险合同赔付等款项的现金
		Niipwbaofi     float64 `json:"niipwbaofi" gorm:"column:niipwbaofi"`         // 拆出资金增加额
		Cpfifac        float64 `json:"cpfifac" gorm:"column:cpfifac"`               // 支付利息、手续费及佣金的现金
		Npfphd         float64 `json:"cpfphd" gorm:"column:cpfphd"`                 // 支付保单红利的现金
		Niifahftp      float64 `json:"niifahftp" gorm:"column:niifahftp"`           // 为交易目的而持有的金融资产净增加额
		Cptofe         float64 `json:"cptofe" gorm:"column:cptofe"`                 // 支付给职工及为职工支付的现金
		Cpft           float64 `json:"cpft" gorm:"column:cpft"`                     // 支付的各种税费
		Cprtooa        float64 `json:"cprtooa" gorm:"column:cprtooa"`               // 支付的其它与经营活动有关现金
		Stcoffoa       float64 `json:"stcoffoa" gorm:"column:stcoffoa"`             // 经营活动现金流出小计
		Ncffoa         float64 `json:"ncffoa" gorm:"column:ncffoa"`                 // 经营活动产生的现金流量净额
		Crfrci         float64 `json:"crfrci" gorm:"column:crfrci"`                 // 收回投资收到的现金
		Crfii          float64 `json:"crfii" gorm:"column:crfii"`                   // 取得投资收益所收到的现金
		Crfdofiaolta   float64 `json:"crfdofiaolta" gorm:"column:crfdofiaolta"`     // 处置固定资产、无形资产和其他长期资产收到的现金
		Ncrfdossaou    float64 `json:"ncrfdossaou" gorm:"column:ncrfdossaou"`       // 处置子公司或其他营业单位收到的现金净额
		Crrtoia        float64 `json:"crrtoia" gorm:"column:crrtoia"`               // 收到的其他与投资活动相关的现金
		Stcifia        float64 `json:"stcifia" gorm:"column:stcifia"`               // 投资活动现金流入小计
		Cpfpfiaolta    float64 `json:"cpfpfiaolta" gorm:"column:cpfpfiaolta"`       // 购建固定资产、无形资产及其他长期资产所支付的现金
		Cpfi           float64 `json:"cpfi" gorm:"column:cpfi"`                     // 投资所支付的现金
		Niipl          float64 `json:"niipl" gorm:"column:niipl"`                   // 质押贷款净增加额
		Ncpfbssaou     float64 `json:"ncpfbssaou" gorm:"column:ncpfbssaou"`         // 取得子公司及其营业单位支付的现金净额
		Cprtoia        float64 `json:"cprtoia" gorm:"column:cprtoia"`               // 支付的其他与投资活动有关的现金
		Stcoffia       float64 `json:"stcoffia" gorm:"column:stcoffia"`             // 投资活动现金流出小计
		Ncffia         float64 `json:"ncffia" gorm:"column:ncffia"`                 // 投资活动产生的现金流量净额
		Crfai          float64 `json:"crfai" gorm:"column:crfai"`                   // 吸收投资收到的现金
		Crfamshibss    float64 `json:"crfamshibss" gorm:"column:crfamshibss"`       // 子公司吸收少数股东投资收到的现金
		Crfl           float64 `json:"crfl" gorm:"column:crfl"`                     // 取得借款收到的现金
		Crfib          float64 `json:"crfib" gorm:"column:crfib"`                   // 发行债券收到的现金
		Crrtofa        float64 `json:"crrtofa" gorm:"column:crrtofa"`               // 收到的其他与筹资活动有关的现金
		Stcifffa       float64 `json:"stcifffa" gorm:"column:stcifffa"`             // 筹资活动产生的现金流入小计
		Cpfbrp         float64 `json:"cpfbrp" gorm:"column:cpfbrp"`                 // 偿付债务支付的现金
		Cpfdapdoip     float64 `json:"cpfdapdoip" gorm:"column:cpfdapdoip"`         // 分配股利、利润或偿付利息所支付的现金
		Cpfdapomshpbss float64 `json:"cpfdapomshpbss" gorm:"column:cpfdapomshpbss"` // 子公司支付少数股东股利和利润
		Cprtofa        float64 `json:"cprtofa" gorm:"column:cprtofa"`               // 支付的其他与筹资活动有关的现金
		Stcofffa       float64 `json:"stcofffa" gorm:"column:stcofffa"`             // 筹资活动产生的现金流出小计
		Ncfffa         float64 `json:"ncfffa" gorm:"column:ncfffa"`                 // 筹资活动产生的现金流量净额
		Iocacedtfier   float64 `json:"iocacedtfier" gorm:"column:iocacedtfier"`     // 汇率变动对现金及现金等价物的影响
		Niicace        float64 `json:"niicace" gorm:"column:niicace"`               // 现金及现金等价物的净增加额
		Bocaceatpb     float64 `json:"bocaceatpb" gorm:"column:bocaceatpb"`         // 期初现金及现金等价物的余额
		Bocaceatpe     float64 `json:"bocaceatpe" gorm:"column:bocaceatpe"`         // 期末现金及现金等价物净余额
		Uril           float64 `json:"uril" gorm:"column:uril"`                     // 未确认的投资损失
		Pfai           float64 `json:"pfai" gorm:"column:pfai"`                     // 资产减值准备
		Dofaaip        float64 `json:"dofaaip" gorm:"column:dofaaip"`               // 固定资产折旧、油气资产折耗、生产性物资折旧
		Aoia           float64 `json:"aoia" gorm:"column:aoia"`                     // 无形资产摊销
		Aoltpe         float64 `json:"aoltpe" gorm:"column:aoltpe"`                 // 长期待摊费用摊销
		Dipe           float64 `json:"dipe" gorm:"column:dipe"`                     // 待摊费用的减少
		Iiae           float64 `json:"iiae" gorm:"column:iiae"`                     // 预提费用的增加
		Godofaaonca    float64 `json:"godofaaonca" gorm:"column:godofaaonca"`       // 处置固定资产、无形资产和其他长期资产的损失
		Losofa         float64 `json:"losofa" gorm:"column:losofa"`                 // 固定资产报废损失
		Clofv          float64 `json:"clofv" gorm:"column:clofv"`                   // 公允价值变动损失
	}

	// 财务指标
	FinancialIndex struct {
		Index              int64   `gorm:"column:index;primary_key"`
		StockCode          string  `json:"stockCode" gorm:"column:stock_code;"`                   // 股票代码
		Currency           string  `gorm:"column:currency"`                                       // 货币
		StandardDate       string  `json:"standardDate" gorm:"column:standard_date"`              // 标准日期
		StandardTime       int64   `json:"standardTime" gorm:"column:standard_time"`              // 标准日期时间戳
		ReportDate         string  `json:"reportDate" gorm:"column:report_date"`                  // 报告时间
		ReportType         string  `json:"reportType" gorm:"column:report_type"`                  // 报告类型
		NpatoshopcPs       float64 `json:"npatoshopc_ps" gorm:"column:npatoshopc_ps"`             // 归属于母公司普通股股东的每股收益
		NpadnrpatoshaopcPs float64 `json:"npadnrpatoshaopc_ps" gorm:"column:npadnrpatoshaopc_ps"` // 归属于母公司普通股股东的每股扣非收益
		TetoshopcPs        float64 `json:"tetoshopc_ps" gorm:"column:tetoshopc_ps"`               // 归属于母公司普通股股东的每股股东权益
		CrPs               float64 `json:"cr_ps" gorm:"column:cr_ps"`                             // 每股资本公积
		RpPs               float64 `json:"rp_ps" gorm:"column:rp_ps"`                             // 每股未分配利润
		StciffoaPs         float64 `json:"stciffoa_ps" gorm:"column:stciffoa_ps"`                 // 每股经营活动产生的现金流量
		NcffoaPs           float64 `json:"ncffoa_ps" gorm:"column:ncffoa_ps"`                     // 每股经营活动产生的现金流量净额
		RoeAtoshaopc       float64 `json:"roe_atoshaopc" gorm:"column:roe_atoshaopc"`             // 归属于母公司普通股股东的ROE
		RoeAdnrpatoshaopc  float64 `json:"roe_adnrpatoshaopc" gorm:"column:roe_adnrpatoshaopc"`   // 归属于母公司普通股股东的扣非ROE
		Wroe               float64 `json:"wroe" gorm:"column:wroe"`                               // 归属于母公司普通股股东的加权ROE
		Roe                float64 `json:"roe" gorm:"column:roe"`                                 // 净资产收益率(ROE)
		L                  float64 `json:"l" gorm:"column:l"`                                     // 杠杆倍数
		Roa                float64 `json:"roa" gorm:"column:roa"`                                 // 总资产收益率(ROA)
		TaTo               float64 `json:"ta_to" gorm:"column:ta_to"`                             // 资产周转率
		NpSR               float64 `json:"np_s_r" gorm:"column:np_s_r"`                           // 净利润率
		GpR                float64 `json:"gp_r" gorm:"column:gp_r"`                               // 毛利率(GM)
		Rota               float64 `json:"rota" gorm:"column:rota"`                               // 有形资产回报率(ROTA)
		Roic               float64 `json:"roic" gorm:"column:roic"`                               // ROIC
		Roc                float64 `json:"roc" gorm:"column:roc"`                                 // 资本回报率(ROC)
		AtsTor             float64 `json:"ats_tor" gorm:"column:ats_tor"`                         // 预付账款周转率
		CaTor              float64 `json:"ca_tor" gorm:"column:ca_tor"`                           // 合同资产周转率
		ITor               float64 `json:"i_tor" gorm:"column:i_tor"`                             // 存货周转率
		NraarTor           float64 `json:"nraar_tor" gorm:"column:nraar_tor"`                     // 应收票据和应收账款周转率
		NrTor              float64 `json:"nr_tor" gorm:"column:nr_tor"`                           // 应收票据周转率
		ArTor              float64 `json:"ar_tor" gorm:"column:ar_tor"`                           // 应收账款周转率
		RfTor              float64 `json:"rf_tor" gorm:"column:rf_tor"`                           // 应收款项融资周转率
		AfcTor             float64 `json:"afc_tor" gorm:"column:afc_tor"`                         // 预收账款周转率
		ClTor              float64 `json:"cl_tor" gorm:"column:cl_tor"`                           // 合同负债周转率
		NpaapTor           float64 `json:"npaap_tor" gorm:"column:npaap_tor"`                     // 应付票据和应付账款周转率
		NpTor              float64 `json:"np_tor" gorm:"column:np_tor"`                           // 应付票据周转率
		ApTor              float64 `json:"ap_tor" gorm:"column:ap_tor"`                           // 应付账款周转率
		FaTor              float64 `json:"fa_tor" gorm:"column:fa_tor"`                           // 固定资产周转率
		AtsDs              float64 `json:"ats_ds" gorm:"column:ats_ds"`                           // 预付账款周转天数
		CaDs               float64 `json:"ca_ds" gorm:"column:ca_ds"`                             // 合同资产周转天数
		IDs                float64 `json:"i_ds" gorm:"column:i_ds"`                               // 存货周转天数
		NraarDs            float64 `json:"nraar_ds" gorm:"column:nraar_ds"`                       // 应收票据和应收账款周转天数
		NrDs               float64 `json:"nr_ds" gorm:"column:nr_ds"`                             // 应收票据周转天数
		ArDs               float64 `json:"ar_ds" gorm:"column:ar_ds"`                             // 应收账款周转天数
		RfDs               float64 `json:"rf_ds" gorm:"column:rf_ds"`                             // 应收款项融资周转天数
		AfcDs              float64 `json:"afc_ds" gorm:"column:afc_ds"`                           // 预收账款周转天数
		ClDs               float64 `json:"cl_ds" gorm:"column:cl_ds"`                             // 合同负债周转天数
		NpaapDs            float64 `json:"npaap_ds" gorm:"column:npaap_ds"`                       // 应付票据和应付账款周转天数
		NpDs               float64 `json:"np_ds" gorm:"column:np_ds"`                             // 应付票据周转天数
		ApDs               float64 `json:"ap_ds" gorm:"column:ap_ds"`                             // 应付账款周转天数
		BDs                float64 `json:"b_ds" gorm:"column:b_ds"`                               // 营业周转天数
		MDs                float64 `json:"m_ds" gorm:"column:m_ds"`                               // 净现金周转天数(CCC)
		FaDs               float64 `json:"fa_ds" gorm:"column:fa_ds"`                             // 固定资产周转天数
		TcaDs              float64 `json:"tca_ds" gorm:"column:tca_ds"`                           // 流动资产周转天数
		ToeDs              float64 `json:"toe_ds" gorm:"column:toe_ds"`                           // 股东权益周转天数
		TlTaR              float64 `json:"tl_ta_r" gorm:"column:tl_ta_r"`                         // 资产负债率
		LwiTaR             float64 `json:"lwi_ta_r" gorm:"column:lwi_ta_r"`                       // 有息负债率
		CabbTclR           float64 `json:"cabb_tcl_r" gorm:"column:cabb_tcl_r"`                   // 货币资金占流动负债比率
		CR                 float64 `json:"c_r" gorm:"column:c_r"`                                 // 流动比率
		GR                 float64 `json:"q_r" gorm:"column:q_r"`                                 // 速动比率
		FaTaR              float64 `json:"fa_ta_r" gorm:"column:fa_ta_r"`                         // 固定资产占总资产比率
		LvR                float64 `json:"lv_r" gorm:"column:lv_r"`                               // 清算价值比率
		Fcf                float64 `json:"fcf" gorm:"column:fcf"`                                 // 自由现金流量
		CrfscaplsOiR       float64 `json:"crfscapls_oi_r" gorm:"column:crfscapls_oi_r"`           // 销售商品提供劳务收到的现金对营业收入的比率
		NcffoaOpR          float64 `json:"ncffoa_op_r" gorm:"column:ncffoa_op_r"`                 // 经营活动产生的现金流量净额对营业利润的比率
		NcffoaNpR          float64 `json:"ncffoa_np_r" gorm:"column:ncffoa_np_r"`                 // 经营活动产生的现金流量净额对净利润的比率
		CrfscaplsTaR       float64 `json:"crfscapls_ta_r" gorm:"column:crfscapls_ta_r"`           // 销售商品提供劳务收到的现金对总资产的比率
		NcffoaFaR          float64 `json:"ncffoa_fa_r" gorm:"column:ncffoa_fa_r"`                 // 经营活动产生的现金流量净额对固定资产的比率
	}
)

func (p *Profit) New(resData model.ProfitResponse) {
	data := resData.GetMapData()
	byteData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(byteData, &p)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *BalanceSheet) New(resData model.BalanceSheetResponse) {
	data := resData.GetMapData()
	byteData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(byteData, &s)
	if err != nil {
		log.Println(err)
		return
	}
}

func (f *CashFlow) New(resData model.CashFlowResponse) {
	data := resData.GetMapData()
	byteData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(byteData, &f)
	if err != nil {
		log.Println(err)
		return
	}
}

func (i *FinancialIndex) New(resData model.FinancialIndexResponse) {
	data := resData.GetMapData()
	byteData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(byteData, &i)
	if err != nil {
		log.Println(err)
		return
	}
}
