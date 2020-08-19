package model

import (
	"github.com/jinzhu/gorm/dialects/postgres"
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
		StockCode                   string  `gorm:"column:stock_code"`                    // 股票代码
		Currency                    string  `gorm:"column:currency"`                      // 货币
		StandardDate                string  `gorm:"column:standar_date"`                  // 标准日期
		StandarTime                 int64   `gorm:"column:standar_time"`                  // 标准日期时间戳
		ReportDate                  string  `gorm:"column:report_date"`                   // 报告时间
		ReportType                  string  `gorm:"column:report_type"`                   // 报告类型
		Toi                         int64   `gorm:"column:toi"`                           // 营业总收入
		Oi                          int64   `gorm:"column:oi"`                            // 营业收入
		Ii                          int64   `gorm:"column:ii"`                            // 利息收入
		Ep                          int64   `gorm:"column:ep"`                            // 已赚保费
		Faci                        int64   `gorm:"column:faci"`                          // 手续费及佣金收入
		Ooi                         int64   `gorm:"column:ooi"`                           // 其他业务收入
		Toc                         int64   `gorm:"column:toc"`                           // 营业总成本
		Oc                          int64   `gorm:"column:oc"`                            // 营业成本
		GpM                         int64   `gorm:"column:gp_m"`                          // 毛利率
		Ie                          int64   `gorm:"column:ie"`                            // 利息支出
		Face                        int64   `gorm:"column:face"`                          // 手续费及佣金支出
		S                           int64   `gorm:"column:s"`                             // 退保金
		Ce                          int64   `gorm:"column:ce"`                            // 保险合同赔付支出
		Iiicr                       int64   `gorm:"column:iiicr"`                         // 提取保险责任准备金净额
		Phdrfpip                    int64   `gorm:"column:phdrfpip"`                      // 保单红利支出
		Rie                         int64   `gorm:"column:rie"`                           // 分保费用
		Tas                         int64   `gorm:"column:tas"`                           // 税金及附加
		Se                          int64   `gorm:"column:se"`                            // 销售费用
		Ae                          int64   `gorm:"column:ae"`                            // 管理费用
		Rade                        int64   `gorm:"column:rade"`                          // 研发费用
		Fe                          int64   `gorm:"column:fe"`                            // 财务费用
		Ieife                       int64   `gorm:"column:ieife"`                         // 利息费用
		Iiife                       int64   `gorm:"column:iiife"`                         // 利息收入
		SeR                         float64 `gorm:"column:se_r"`                          // 销售费用率
		AeR                         float64 `gorm:"column:ae_r"`                          // 管理费用率
		RadeR                       float64 `gorm:"column:rade_r"`                        // 研发费用率
		FeR                         float64 `gorm:"column:fe_r"`                          // 财务费用率
		OeR                         float64 `gorm:"column:oe_r"`                          // 总营业费用率
		TeR                         float64 `gorm:"column:te_r"`                          // 三项费用率
		Oic                         int64   `gorm:"column:oic"`                           // 其他收益
		Ivi                         int64   `gorm:"column:ivi"`                           // 投资收益
		Iifaajv                     int64   `gorm:"column:iifaajv"`                       // 对联营企业及合营企业的投资收益
		Iftdofamaac                 int64   `gorm:"column:iftdofamaac"`                   // 以摊余成本计量的金融资产终止确认产生的投资收益
		Ei                          int64   `gorm:"column:ei"`                            // 汇兑收益
		Nehb                        int64   `gorm:"column:nehb"`                          // 净敞口套期收益
		Ciofv                       int64   `gorm:"column:ciofv"`                         // 公允价值变动收益
		Cilor                       int64   `gorm:"column:cilor"`                         // 信用减值损失
		Ailor                       int64   `gorm:"column:ailor"`                         // 资产减值损失
		Oail                        int64   `gorm:"column:oail"`                          // 其他资产减值损失
		Adi                         int64   `gorm:"column:adi"`                           // 资产处置收益
		Op                          int64   `gorm:"column:op"`                            // 营业利润
		OpSR                        float64 `gorm:"column:op_s_r"`                        // 营业利润率
		OpOpR                       float64 `gorm:"column:op_op_r"`                       // 其他营业利润占比
		Noi                         int64   `gorm:"column:noi"`                           // 营业外收入
		Ncadarg                     int64   `gorm:"column:ncadarg"`                       // 非流动资产毁损报废利得
		Noe                         int64   `gorm:"column:noe"`                           // 营业外支出
		Ncadarl                     int64   `gorm:"column:ncadarl"`                       // 非流动资产毁损报废损失
		Tp                          int64   `gorm:"column:tp"`                            // 利润总额
		RadeTpR                     float64 `gorm:"column:rade_tp_r"`                     // 研发费占利润总额比值
		Ite                         int64   `gorm:"column:ite"`                           // 所得税费用
		IteTpR                      float64 `gorm:"column:ite_tp_r"`                      // 有效税率
		Np                          int64   `gorm:"column:np"`                            // 净利润
		NpSR                        float64 `gorm:"column:np_s_r"`                        // 净利润率
		Npfco                       int64   `gorm:"column:npfco"`                         // 持续经营净利润
		Npfdco                      int64   `gorm:"column:npfdco"`                        // 终止经营净利润
		Wroe                        int64   `gorm:"column:wroe"`                          // 归属于母公司普通股股东的加权ROE
		Npatshaoehopc               int64   `gorm:"column:npatshaoehopc"`                 // 归属于母公司股东及其他权益持有者的净利润
		Npatoshopc                  int64   `gorm:"column:npatoshopc"`                    // 归属于母公司普通股股东的净利润
		Npadnrpatoshaopc            int64   `gorm:"column:npadnrpatoshaopc"`              // 归属于母公司普通股股东的扣除非经常性损益的净利润
		NpadnrpatoshaopcNpatoshopcR float64 `gorm:"column:npadnrpatoshaopc_npatoshopc_r"` // 扣非净利润占比
		Npatoeihopc                 int64   `gorm:"column:npatoeihopc"`                   // 归属于母公司其他权益工具持有者的净利润
		Npatmsh                     int64   `gorm:"column:npatmsh"`                       // 少数股东损益
		Natooci                     int64   `gorm:"column:natooci"`                       // 其他综合收益的税后净额
		Tci                         int64   `gorm:"column:tci"`                           // 综合收益总额
		Tciatshaoehopc              int64   `gorm:"column:tciatshaoehopc"`                // 归属于母公司股东及其他权益持有者的综合收益总额
		Tciatoshopc                 int64   `gorm:"column:tciatoshopc"`                   // 归属于母公司普通股股东的综合收益总额
		Tciatmsh                    int64   `gorm:"column:tciatmsh"`                      // 归属于少数股东的综合收益总额
		Beps                        int64   `gorm:"column:beps"`                          // 基本每股收益
		Deps                        int64   `gorm:"column:deps"`                          // 稀释每股收益
	}

	// 资产负债表
	BalanceSheet struct {
		Index        int64   `gorm:"column:index;primary_key"`
		StockCode    string  `json:"stockCode" gorm:"column:stock_code;"` // 股票代码
		Currency     string  `gorm:"column:currency"`                     // 货币
		Ta           int64   `gorm:"column:ta"`                           // 资产合计
		Nwc          int64   `gorm:"column:nwc"`                          // 净营运资本
		Tca          int64   `gorm:"column:tca"`                          // 流动资产合计
		TcaTaR       float64 `gorm:"column:tca_ta_r"`                     // 流动资产占比
		Cabb         int64   `gorm:"column:cabb"`                         // 货币资金
		CabbTaR      float64 `gorm:"column:cabb_ta_r"`                    // 货币资金占比
		Sr           int64   `gorm:"column:sr"`                           // 结算备付金
		Pwbaofi      int64   `gorm:"column:pwbaofi"`                      // 拆出资金
		Tfa          int64   `gorm:"column:tfa"`                          // 交易性金融资产
		Cdfa         int64   `gorm:"column:cdfa"`                         // 衍生金融资产(流动)
		Nraar        int64   `gorm:"column:nraar"`                        // 应收票据及应收账款
		Nr           int64   `gorm:"column:nr"`                           // 应收票据
		Ar           int64   `gorm:"column:ar"`                           // 应收账款
		Rf           int64   `gorm:"column:rf"`                           // 应收款项融资
		Ats          int64   `gorm:"column:ats"`                          // 预付款项
		Pr           int64   `gorm:"column:pr"`                           // 应收保费
		Rir          int64   `gorm:"column:rir"`                          // 应收分保账款
		Crorir       int64   `gorm:"column:crorir"`                       // 应收分保合同准备金
		Or           int64   `gorm:"column:or"`                           // 其他应收款
		Ir           int64   `gorm:"column:ir"`                           // 应收利息
		Dr           int64   `gorm:"column:dr"`                           // 应收股利
		Fahursa      int64   `gorm:"column:fahursa"`                      // 买入返售金融资产
		I            int64   `gorm:"column:i"`                            // 存货
		Ca           int64   `gorm:"column:ca"`                           // 合同资产
		Ahfs         int64   `gorm:"column:ahfs"`                         // 持有待售资产
		Pe           int64   `gorm:"column:pe"`                           // 待摊费用
		Ncadwioy     int64   `gorm:"column:ncadwioy"`                     // 一年内到期的非流动资产
		Oca          int64   `gorm:"column:oca"`                          // 其他流动资产
		TcaTclR      float64 `gorm:"column:tca_tcl_r"`                    // 流动比率
		QR           float64 `gorm:"column:q_r"`                          // 速动比率
		Tnca         int64   `gorm:"column:tnca"`                         // 非流动资产合计
		TncaTaR      float64 `gorm:"column:tnca_ta_r"`                    // 非流动资产占比
		AhTaR        float64 `gorm:"column:ah_ta_r"`                      // 重资产占比
		Nclaatc      int64   `gorm:"column:nclaatc"`                      // 发放贷款及垫款(非流动)
		Cri          int64   `gorm:"column:cri"`                          // 债权投资
		Ocri         int64   `gorm:"column:ocri"`                         // 其他债权投资
		Ncafsfa      int64   `gorm:"column:ncafsfa"`                      // 可供出售金融资产(非流动)
		Htmi         int64   `gorm:"column:htmi"`                         // 持有至到期投资
		Ltar         int64   `gorm:"column:ltar"`                         // 长期应收款
		Ltei         int64   `gorm:"column:ltei"`                         // 长期股权投资
		Oeii         int64   `gorm:"column:oeii"`                         // 其他权益工具投资
		Oncfa        int64   `gorm:"column:oncfa"`                        // 其他非流动金融资产
		Rei          int64   `gorm:"column:rei"`                          // 投资性房地产
		Fa           int64   `gorm:"column:fa"`                           // 固定资产
		Dofa         int64   `gorm:"column:dofa"`                         // 固定资产清理
		FaTaR        float64 `gorm:"column:fa_ta_r"`                      // 固定资产占总资产比率
		Cip          int64   `gorm:"column:cip"`                          // 在建工程
		Es           int64   `gorm:"column:es"`                           // 工程物资
		CipFaR       float64 `gorm:"column:cip_fa_r"`                     // 在建工程占固定资产比率
		Pba          int64   `gorm:"column:pba"`                          // 生产性生物资产
		Oaga         int64   `gorm:"column:oaga"`                         // 油气资产
		Pwba         int64   `gorm:"column:pwba"`                         // 公益性生物资产
		Roua         int64   `gorm:"column:roua"`                         // 使用权资产
		Ia           int64   `gorm:"column:ia"`                           // 无形资产
		Rade         int64   `gorm:"column:rade"`                         // 开发支出
		Gw           int64   `gorm:"column:gw"`                           // 商誉
		GwToeR       float64 `gorm:"column:gw_toe_r"`                     // 商誉占净资产比率
		Ltpe         int64   `gorm:"column:ltpe"`                         // 长期待摊费用
		Dita         int64   `gorm:"column:dita"`                         // 递延所得税资产
		Onca         int64   `gorm:"column:onca"`                         // 其他非流动资产
		Tl           int64   `gorm:"column:tl"`                           // 负债合计
		Lwi          int64   `gorm:"column:lwi"`                          // 有息负债
		LwiTaR       float64 `gorm:"column:lwi_ta_r"`                     // 有息负债率
		TlTaR        float64 `gorm:"column:tl_ta_r"`                      // 资产负债率
		Tcl          int64   `gorm:"column:tcl"`                          // 流动负债合计
		TclTlR       float64 `gorm:"column:tcl_tl_r"`                     // 流动负债占比
		Stl          int64   `gorm:"column:stl"`                          // 短期借款
		Bfcb         int64   `gorm:"column:bfcb"`                         // 向中央银行借款
		Pfbaofi      int64   `gorm:"column:pfbaofi"`                      // 拆入资金
		Tfl          int64   `gorm:"column:tfl"`                          // 交易性金融负债
		Dfl          int64   `gorm:"column:dfl"`                          // 衍生金融负债
		Npaap        int64   `gorm:"column:npaap"`                        // 应付票据及应付账款
		Np           int64   `gorm:"column:np"`                           // 应付票据
		Ap           int64   `gorm:"column:ap"`                           // 应付账款
		Afc          int64   `gorm:"column:afc"`                          // 预收账款
		Cl           int64   `gorm:"column:cl"`                           // 合同负债
		Fasurpa      int64   `gorm:"column:fasurpa"`                      // 卖出回购金融资产
		Dfcab        int64   `gorm:"column:dfcab"`                        // 吸收存款及同业存放
		Stoa         int64   `gorm:"column:stoa"`                         // 代理买卖证券款
		Ssoa         int64   `gorm:"column:ssoa"`                         // 代理承销证券款
		Sawp         int64   `gorm:"column:sawp"`                         // 应付职工薪酬
		Tp           int64   `gorm:"column:tp"`                           // 应交税费
		Oap          int64   `gorm:"column:oap"`                          // 其他应付款
		Intp         int64   `gorm:"column:intp"`                         // 应付利息
		Dp           int64   `gorm:"column:dp"`                           // 应付股利
		Facp         int64   `gorm:"column:facp"`                         // 应付手续费及佣金
		Rip          int64   `gorm:"column:rip"`                          // 应付分保账款
		Lhfs         int64   `gorm:"column:lhfs"`                         // 持有待售负债
		Ncldwioy     int64   `gorm:"column:ncldwioy"`                     // 一年内到期的非流动负债
		Didwioy      int64   `gorm:"column:didwioy"`                      // 一年内到期的递延收益
		Cal          int64   `gorm:"column:cal"`                          // 预计负债(流动)
		Stbp         int64   `gorm:"column:stbp"`                         // 短期应付债券
		Ocl          int64   `gorm:"column:ocl"`                          // 其他流动负债
		Tncl         int64   `gorm:"column:tncl"`                         // 非流动负债合计
		TnclTlR      float64 `gorm:"column:tncl_tl_r"`                    // 非流动负债占比
		Icr          int64   `gorm:"column:icr"`                          // 保险合同准备金
		Ltl          int64   `gorm:"column:ltl"`                          // 长期借款
		Bp           int64   `gorm:"column:bp"`                           // 应付债券
		Psibp        int64   `gorm:"column:psibp"`                        // 优先股
		Pcsibp       int64   `gorm:"column:pcsibp"`                       // 永续债
		Ll           int64   `gorm:"column:ll"`                           // 租赁负债
		Ltap         int64   `gorm:"column:ltap"`                         // 长期应付款
		Sap          int64   `gorm:"column:sap"`                          // 专项应付款
		Ltpoe        int64   `gorm:"column:ltpoe"`                        // 长期应付职工薪酬
		Ncal         int64   `gorm:"column:ncal"`                         // 预计负债(非流动)
		Ltdi         int64   `gorm:"column:ltdi"`                         // 长期递延收益
		Ditl         int64   `gorm:"column:ditl"`                         // 递延所得税负债
		Oncl         int64   `gorm:"column:oncl"`                         // 其他非流动负债
		Toe          int64   `gorm:"column:toe"`                          // 股东权益合计
		ToeTaR       float64 `gorm:"column:toe_ta_r"`                     // 股东权益占比
		Sc           int64   `gorm:"column:sc"`                           // 股本
		Oei          int64   `gorm:"column:oei"`                          // 其他权益工具
		Psioei       int64   `gorm:"column:psioei"`                       // 优先股
		Pcsioei      int64   `gorm:"column:pcsioei"`                      // 永续债
		Capr         int64   `gorm:"column:capr"`                         // 资本公积
		Is           int64   `gorm:"column:is"`                           // 库存股
		Oci          int64   `gorm:"column:oci"`                          // 其他综合收益
		Rr           int64   `gorm:"column:rr"`                           // 专项储备
		Surr         int64   `gorm:"column:surr"`                         // 盈余公积
		Pogr         int64   `gorm:"column:pogr"`                         // 一般风险准备金
		Rtp          int64   `gorm:"column:rtp"`                          // 未分配利润
		Tetshaoehopc int64   `gorm:"column:tetshaoehopc"`                 // 归属于母公司股东及其他权益持有者的权益合计
		Tetoshopc    int64   `gorm:"column:tetoshopc"`                    // 归属于母公司普通股股东权益合计
		TetoshopcPs  int64   `gorm:"column:tetoshopc_ps"`                 // 归属于母公司普通股股东的每股股东权益
		Etmsh        int64   `gorm:"column:etmsh"`                        // 少数股东权益
		Mc           int64   `gorm:"column:mc"`                           // 市值
		Tsc          int64   `gorm:"column:tsc"`                          // 总股本
		Csc          int64   `gorm:"column:csc"`                          // 流通股本
		Shn          int64   `gorm:"column:shn"`                          // 股东人数(季度)
		Shbt1shTscR  float64 `gorm:"column:shbt1sh_tsc_r"`                // 第一大股东持仓占总股本比例
		Shbt10shTscR float64 `gorm:"column:shbt10sh_tsc_r"`               // 前十大股东持仓占总股本比例
		Shbt10shCscR float64 `gorm:"column:shbt10sh_csc_r"`               // 前十大流通股东持仓占流通股本比例
		ShbpoofCscR  float64 `gorm:"column:shbpoof_csc_r"`                // 公募基金持仓占流通股本比例
		PeTtm        int64   `gorm:"column:pe_ttm"`                       // PE-TTM
		DPeTtm       int64   `gorm:"column:d_pe_ttm"`                     // PE-TTM(扣非)
		Pb           int64   `gorm:"column:pb"`                           // PB
		PbWoGw       int64   `gorm:"column:pb_wo_gw"`                     // PB(不含商誉)
		PsTtm        int64   `gorm:"column:ps_ttm"`                       // PS-TTM
		Dyr          int64   `gorm:"column:dyr"`                          // 股息率
	}

	// 现金流量表
	CashFlow struct {
		Index          int64  `gorm:"column:index;primary_key"`
		StockCode      string `json:"stockCode" gorm:"column:stock_code;"` // 股票代码
		Currency       string `gorm:"column:currency"`                     // 货币
		Crfscapls      int64  `gorm:"column:crfscapls"`                    // 销售商品、提供劳务收到的现金
		Niicdadfbaofi  int64  `gorm:"column:niicdadfbaofi"`                // 客户存款和同业及其他金融机构存放款项净增加额
		Niibfcb        int64  `gorm:"column:niibfcb"`                      // 向中央银行借款净增加额
		Niipfofi       int64  `gorm:"column:niipfofi"`                     // 向其他金融机构拆入资金净增加额
		Ndilaatc       int64  `gorm:"column:ndilaatc"`                     // 发放贷款及垫款的净减少额
		Crfp           int64  `gorm:"column:crfp"`                         // 收到原保险合同保费取得的现金
		Ncrfrib        int64  `gorm:"column:ncrfrib"`                      // 收到再保险业务现金净额
		Niiphd         int64  `gorm:"column:niiphd"`                       // 保户储金及投资款净增加额
		Ndifahftp      int64  `gorm:"column:ndifahftp"`                    // 为交易目的而持有的金融资产净减少额
		Crfifac        int64  `gorm:"column:crfifac"`                      // 收取利息、手续费及佣金的现金
		Niipfbaofi     int64  `gorm:"column:niipfbaofi"`                   // 拆入资金净增加额
		Niifasurpaioa  int64  `gorm:"column:niifasurpaioa"`                // 回购业务资金净增加额
		Ncrfstoa       int64  `gorm:"column:ncrfstoa"`                     // 代理买卖证券收到的现金净额
		Crfwbot        int64  `gorm:"column:crfwbot"`                      // 收到的税费返还
		Crrtooa        int64  `gorm:"column:crrtooa"`                      // 收到的其他与经营活动有关现金
		Stciffoa       int64  `gorm:"column:stciffoa"`                     // 经营活动现金流入小计
		Cpfpcarls      int64  `gorm:"column:cpfpcarls"`                    // 购买商品、接收劳务支付的现金
		Niilaatc       int64  `gorm:"column:niilaatc"`                     // 发放贷款和垫款的净增加额
		Niibwcbbaofi   int64  `gorm:"column:niibwcbbaofi"`                 // 存放中央银行和同业及其他金融机构款项净增加额
		Cpfc           int64  `gorm:"column:cpfc"`                         // 支付原保险合同赔付等款项的现金
		Niipwbaofi     int64  `gorm:"column:niipwbaofi"`                   // 拆出资金增加额
		Cpfifac        int64  `gorm:"column:cpfifac"`                      // 支付利息、手续费及佣金的现金
		Npfphd         int64  `gorm:"column:cpfphd"`                       // 支付保单红利的现金
		Niifahftp      int64  `gorm:"column:niifahftp"`                    // 为交易目的而持有的金融资产净增加额
		Cptofe         int64  `gorm:"column:cptofe"`                       // 支付给职工及为职工支付的现金
		Cpft           int64  `gorm:"column:cpft"`                         // 支付的各种税费
		Cprtooa        int64  `gorm:"column:cprtooa"`                      // 支付的其它与经营活动有关现金
		Stcoffoa       int64  `gorm:"column:stcoffoa"`                     // 经营活动现金流出小计
		Ncffoa         int64  `gorm:"column:ncffoa"`                       // 经营活动产生的现金流量净额
		Crfrci         int64  `gorm:"column:crfrci"`                       // 收回投资收到的现金
		Crfii          int64  `gorm:"column:crfii"`                        // 取得投资收益所收到的现金
		Crfdofiaolta   int64  `gorm:"column:crfdofiaolta"`                 // 处置固定资产、无形资产和其他长期资产收到的现金
		Ncrfdossaou    int64  `gorm:"column:ncrfdossaou"`                  // 处置子公司或其他营业单位收到的现金净额
		Crrtoia        int64  `gorm:"column:crrtoia"`                      // 收到的其他与投资活动相关的现金
		Stcifia        int64  `gorm:"column:stcifia"`                      // 投资活动现金流入小计
		Cpfpfiaolta    int64  `gorm:"column:cpfpfiaolta"`                  // 购建固定资产、无形资产及其他长期资产所支付的现金
		Cpfi           int64  `gorm:"column:cpfi"`                         // 投资所支付的现金
		Niipl          int64  `gorm:"column:niipl"`                        // 质押贷款净增加额
		Ncpfbssaou     int64  `gorm:"column:ncpfbssaou"`                   // 取得子公司及其营业单位支付的现金净额
		Cprtoia        int64  `gorm:"column:cprtoia"`                      // 支付的其他与投资活动有关的现金
		Stcoffia       int64  `gorm:"column:stcoffia"`                     // 投资活动现金流出小计
		Ncffia         int64  `gorm:"column:ncffia"`                       // 投资活动产生的现金流量净额
		Crfai          int64  `gorm:"column:crfai"`                        // 吸收投资收到的现金
		Crfamshibss    int64  `gorm:"column:crfamshibss"`                  // 子公司吸收少数股东投资收到的现金
		Crfl           int64  `gorm:"column:crfl"`                         // 取得借款收到的现金
		Crfib          int64  `gorm:"column:crfib"`                        // 发行债券收到的现金
		Crrtofa        int64  `gorm:"column:crrtofa"`                      // 收到的其他与筹资活动有关的现金
		Stcifffa       int64  `gorm:"column:stcifffa"`                     // 筹资活动产生的现金流入小计
		Cpfbrp         int64  `gorm:"column:cpfbrp"`                       // 偿付债务支付的现金
		Cpfdapdoip     int64  `gorm:"column:cpfdapdoip"`                   // 分配股利、利润或偿付利息所支付的现金
		Cpfdapomshpbss int64  `gorm:"column:cpfdapomshpbss"`               // 子公司支付少数股东股利和利润
		Cprtofa        int64  `gorm:"column:cprtofa"`                      // 支付的其他与筹资活动有关的现金
		Stcofffa       int64  `gorm:"column:stcofffa"`                     // 筹资活动产生的现金流出小计
		Ncfffa         int64  `gorm:"column:ncfffa"`                       // 筹资活动产生的现金流量净额
		Iocacedtfier   int64  `gorm:"column:iocacedtfier"`                 // 汇率变动对现金及现金等价物的影响
		Niicace        int64  `gorm:"column:niicace"`                      // 现金及现金等价物的净增加额
		Bocaceatpb     int64  `gorm:"column:bocaceatpb"`                   // 期初现金及现金等价物的余额
		Bocaceatpe     int64  `gorm:"column:bocaceatpe"`                   // 期末现金及现金等价物净余额
		Uril           int64  `gorm:"column:uril"`                         // 未确认的投资损失
		Pfai           int64  `gorm:"column:pfai"`                         // 资产减值准备
		Dofaaip        int64  `gorm:"column:dofaaip"`                      // 固定资产折旧、油气资产折耗、生产性物资折旧
		Aoia           int64  `gorm:"column:aoia"`                         // 无形资产摊销
		Aoltpe         int64  `gorm:"column:aoltpe"`                       // 长期待摊费用摊销
		Dipe           int64  `gorm:"column:dipe"`                         // 待摊费用的减少
		Iiae           int64  `gorm:"column:iiae"`                         // 预提费用的增加
		Godofaaonca    int64  `gorm:"column:godofaaonca"`                  // 处置固定资产、无形资产和其他长期资产的损失
		Losofa         int64  `gorm:"column:losofa"`                       // 固定资产报废损失
		Clofv          int64  `gorm:"column:clofv"`                        // 公允价值变动损失
	}

	// 财务指标
	FinancialIndex struct {
		Index              int64   `gorm:"column:index;primary_key"`
		StockCode          string  `json:"stockCode" gorm:"column:stock_code;"` // 股票代码
		Currency           string  `gorm:"column:currency"`                     // 货币
		NpatoshopcPs       int64   `gorm:"column:npatoshopc_ps"`                // 归属于母公司普通股股东的每股收益
		NpadnrpatoshaopcPs int64   `gorm:"column:npadnrpatoshaopc_ps"`          // 归属于母公司普通股股东的每股扣非收益
		TetoshopcPs        int64   `gorm:"column:tetoshopc_ps"`                 // 归属于母公司普通股股东的每股股东权益
		CrPs               int64   `gorm:"column:cr_ps"`                        // 每股资本公积
		RpPs               int64   `gorm:"column:rp_ps"`                        // 每股未分配利润
		StciffoaPs         int64   `gorm:"column:stciffoa_ps"`                  // 每股经营活动产生的现金流量
		NcffoaPs           int64   `gorm:"column:ncffoa_ps"`                    // 每股经营活动产生的现金流量净额
		RoeAtoshaopc       int64   `gorm:"column:roe_atoshaopc"`                // 归属于母公司普通股股东的ROE
		RoeAdnrpatoshaopc  int64   `gorm:"column:roe_adnrpatoshaopc"`           // 归属于母公司普通股股东的扣非ROE
		Wroe               int64   `gorm:"column:wroe"`                         // 归属于母公司普通股股东的加权ROE
		Roe                int64   `gorm:"column:roe"`                          // 净资产收益率(ROE)
		L                  int64   `gorm:"column:l"`                            // 杠杆倍数
		Roa                int64   `gorm:"column:roa"`                          // 总资产收益率(ROA)
		TaTo               int64   `gorm:"column:ta_to"`                        // 资产周转率
		NpSR               float64 `gorm:"column:np_s_r"`                       // 净利润率
		GpR                float64 `gorm:"column:gp_r"`                         // 毛利率(GM)
		Rota               int64   `gorm:"column:rota"`                         // 有形资产回报率(ROTA)
		Roic               int64   `gorm:"column:roic"`                         // ROIC
		Roc                int64   `gorm:"column:roc"`                          // 资本回报率(ROC)
		AtsTor             int64   `gorm:"column:ats_tor"`                      // 预付账款周转率
		CaTor              int64   `gorm:"column:ca_tor"`                       // 合同资产周转率
		ITor               int64   `gorm:"column:i_tor"`                        // 存货周转率
		NraarTor           int64   `gorm:"column:nraar_tor"`                    // 应收票据和应收账款周转率
		NrTor              int64   `gorm:"column:nr_tor"`                       // 应收票据周转率
		ArTor              int64   `gorm:"column:ar_tor"`                       // 应收账款周转率
		RfTor              int64   `gorm:"column:rf_tor"`                       // 应收款项融资周转率
		AfcTor             int64   `gorm:"column:afc_tor"`                      // 预收账款周转率
		ClTor              int64   `gorm:"column:cl_tor"`                       // 合同负债周转率
		NpaapTor           int64   `gorm:"column:npaap_tor"`                    // 应付票据和应付账款周转率
		NpTor              int64   `gorm:"column:np_tor"`                       // 应付票据周转率
		ApTor              int64   `gorm:"column:ap_tor"`                       // 应付账款周转率
		FaTor              int64   `gorm:"column:fa_tor"`                       // 固定资产周转率
		AtsDs              int64   `gorm:"column:ats_ds"`                       // 预付账款周转天数
		CaDs               int64   `gorm:"column:ca_ds"`                        // 合同资产周转天数
		IDs                int64   `gorm:"column:i_ds"`                         // 存货周转天数
		NraarDs            int64   `gorm:"column:nraar_ds"`                     // 应收票据和应收账款周转天数
		NrDs               int64   `gorm:"column:nr_ds"`                        // 应收票据周转天数
		ArDs               int64   `gorm:"column:ar_ds"`                        // 应收账款周转天数
		RfDs               int64   `gorm:"column:rf_ds"`                        // 应收款项融资周转天数
		AfcDs              int64   `gorm:"column:afc_ds"`                       // 预收账款周转天数
		ClDs               int64   `gorm:"column:cl_ds"`                        // 合同负债周转天数
		NpaapDs            int64   `gorm:"column:npaap_ds"`                     // 应付票据和应付账款周转天数
		NpDs               int64   `gorm:"column:np_ds"`                        // 应付票据周转天数
		ApDs               int64   `gorm:"column:ap_ds"`                        // 应付账款周转天数
		BDs                int64   `gorm:"column:b_ds"`                         // 营业周转天数
		MDs                int64   `gorm:"column:m_ds"`                         // 净现金周转天数(CCC)
		FaDs               int64   `gorm:"column:fa_ds"`                        // 固定资产周转天数
		TcaDs              int64   `gorm:"column:tca_ds"`                       // 流动资产周转天数
		ToeDs              int64   `gorm:"column:toe_ds"`                       // 股东权益周转天数
		TlTaR              float64 `gorm:"column:tl_ta_r"`                      // 资产负债率
		LwiTaR             float64 `gorm:"column:lwi_ta_r"`                     // 有息负债率
		CabbTclR           float64 `gorm:"column:cabb_tcl_r"`                   // 货币资金占流动负债比率
		CR                 float64 `gorm:"column:c_r"`                          // 流动比率
		GR                 float64 `gorm:"column:q_r"`                          // 速动比率
		FaTaR              float64 `gorm:"column:fa_ta_r"`                      // 固定资产占总资产比率
		LvR                float64 `gorm:"column:lv_r"`                         // 清算价值比率
		Fcf                int64   `gorm:"column:fcf"`                          // 自由现金流量
		CrfscaplsOiR       float64 `gorm:"column:crfscapls_oi_r"`               // 销售商品提供劳务收到的现金对营业收入的比率
		NcffoaOpR          float64 `gorm:"column:ncffoa_op_r"`                  // 经营活动产生的现金流量净额对营业利润的比率
		NcffoaNpR          float64 `gorm:"column:ncffoa_np_r"`                  // 经营活动产生的现金流量净额对净利润的比率
		CrfscaplsTaR       float64 `gorm:"column:crfscapls_ta_r"`               // 销售商品提供劳务收到的现金对总资产的比率
		NcffoaFaR          float64 `gorm:"column:ncffoa_fa_r"`                  // 经营活动产生的现金流量净额对固定资产的比率
	}
)
