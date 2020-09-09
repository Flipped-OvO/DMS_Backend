package model

import (
	"backend/pkg/lixinger/model"
	"backend/util"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"reflect"
	"strings"
)

var (
	profitTags = map[string]string{
		"Toi":         "default_Rev",
		"Toc":         "default_COGS",
		"GP":          "default_GP",
		"Se":          "default_SM",
		"Ae":          "default_GA",
		"Rade":        "default_RD",
		"IsOthers1":   "default_isOthers1",
		"OP":          "default_OP",
		"Np":          "default_NI",
		"IsFI":        "default_isFI",
		"IsSI":        "default_isSI",
		"IsFC":        "default_isFC",
		"IsTax":       "default_isTax",
		"IsEI":        "default_isEI",
		"IsOthers2":   "default_isOthers2",
		"Ii":          "hidden_isFI_1",
		"Iiife":       "hidden_isFI_2",
		"Ie":          "hidden_isFC_1",
		"Ieife":       "hidden_isFC_2",
		"Ivi":         "hidden_isEI_1",
		"Iifaajv":     "hidden_isEI_2",
		"Ailor":       "hidden_isEI_3",
		"Oail":        "hidden_isEI_4",
		"Iftdofamaac": "hidden_isSI_1",
		"Ei":          "hidden_isSI_2",
		"Ciofv":       "hidden_isSI_3",
		"Tas":         "hidden_isTax_1",
		"Ite":         "hidden_isTax_2",
	}

	profitFormula = map[string]string{
		"GP":        "=Rev-COGS",
		"OP":        "=GP-SM-GA-RD-isOthers1",
		"isFI":      "=0",
		"isSI":      "=0",
		"isFC":      "=0",
		"isTax":     "=0",
		"isEI":      "=0",
		"isOthers1": "=0",
		"isOthers2": "=OP-NI-(isFI+isSI+isFC+isTax+isEI)",
	}

	profitFormulaTag = map[string]string{
		"Rev":       "Toi",
		"COGS":      "Toc",
		"GP":        "GP",
		"SM":        "Se",
		"GA":        "Ae",
		"RD":        "Rade",
		"isOthers1": "IsOthers1",
		"OP":        "OP",
		"NI":        "Np",
		"isFI":      "IsFI",
		"isSI":      "IsSI",
		"isFC":      "IsFC",
		"isTax":     "IsTax",
		"isEI":      "IsEI",
		"isOthers2": "IsOthers2",
		"isFI_1":    "Ii",
		"isFI_2":    "Iiife",
		"isFC_1":    "Ie",
		"isFC_2":    "Ieife",
		"isEI_1":    "Ivi",
		"isEI_2":    "Iifaajv",
		"isEI_3":    "Ailor",
		"isEI_4":    "Oail",
		"isSI_1":    "Iftdofamaac",
		"isSI_2":    "Ei",
		"isSI_3":    "Ciofv",
		"isTax_1":   "Tas",
		"isTax_2":   "Ite",
	}
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
		Index        int64   `json:"index" gorm:"column:index;primary_key"`
		StockCode    string  `json:"stockCode" gorm:"column:stock_code"`       // 股票代码
		Currency     string  `json:"currency" gorm:"column:currency"`          // 货币
		StandardDate string  `json:"standardDate" gorm:"column:standard_date"` // 标准日期
		StandardTime int64   `json:"standardTime" gorm:"column:standard_time"` // 标准日期时间戳
		ReportDate   string  `json:"reportDate" gorm:"column:report_date"`     // 报告时间
		ReportType   string  `json:"reportType" gorm:"column:report_type"`     // 报告类型
		Toi          float64 `json:"toi" gorm:"column:toi"`                    // 营业总收入
		Ii           float64 `json:"ii" gorm:"column:ii"`                      // 利息收入
		Toc          float64 `json:"toc" gorm:"column:toc"`                    // 营业总成本
		Ie           float64 `json:"ie" gorm:"column:ie"`                      // 利息支出
		Tas          float64 `json:"tas" gorm:"column:tas"`                    // 税金及附加
		Se           float64 `json:"se" gorm:"column:se"`                      // 销售费用
		Ae           float64 `json:"ae" gorm:"column:ae"`                      // 管理费用
		Rade         float64 `json:"rade" gorm:"column:rade"`                  // 研发费用
		Ieife        float64 `json:"ieife" gorm:"column:ieife"`                // 利息费用
		Iiife        float64 `json:"iiife" gorm:"column:iiife"`                // 利息收入
		Ivi          float64 `json:"ivi" gorm:"column:ivi"`                    // 投资收益
		Iifaajv      float64 `json:"iifaajv" gorm:"column:iifaajv"`            // 对联营企业及合营企业的投资收益
		Iftdofamaac  float64 `json:"iftdofamaac" gorm:"column:iftdofamaac"`    // 以摊余成本计量的金融资产终止确认产生的投资收益
		Ei           float64 `json:"ei" gorm:"column:ei"`                      // 汇兑收益
		Ciofv        float64 `json:"ciofv" gorm:"column:ciofv"`                // 公允价值变动收益
		Ailor        float64 `json:"ailor" gorm:"column:ailor"`                // 资产减值损失
		Oail         float64 `json:"oail" gorm:"column:oail"`                  // 其他资产减值损失
		Ite          float64 `json:"ite" gorm:"column:ite"`                    // 所得税费用
		Np           float64 `json:"np" gorm:"column:np"`                      // 净利润
		IsFI         float64 `json:"is_fi" gorm:"column:is_fi"`
		FIF          string  `json:"fif" gorm:"column:fif"`
		IsSI         float64 `json:"is_si" gorm:"column:is_si"`
		SIF          string  `json:"sif" gorm:"column:sif"`
		IsFC         float64 `json:"is_fc" gorm:"column:is_fc"`
		FCF          string  `json:"fcf" gorm:"column:fcf"`
		IsTax        float64 `json:"is_tax" gorm:"column:is_tax"`
		TaxF         string  `json:"taxf" gorm:"column:taxf"`
		IsEI         float64 `json:"is_ei" gorm:"column:is_ei"`
		EIF          string  `json:"eif" gorm:"column:eif"`
		OP           float64 `json:"op" gorm:"column:op"`
		OPF          string  `json:"opf" gorm:"column:opf"`
		GP           float64 `json:"gp" gorm:"column:gp"`
		GPF          string  `json:"gpf" gorm:"column:gpf"`
		IsOthers1    float64 `json:"is_others_1" gorm:"column:is_others1"`
		O1F          string  `json:"o1f" gorm:"column:o1f"`
		IsOthers2    float64 `json:"is_others_2" gorm:"column:is_others2"`
		O2F          string  `json:"o2f" gorm:"column:o2f"`
	}

	// 资产负债表
	BalanceSheet struct {
		Index        int64   `json:"index" gorm:"column:index;primary_key"`
		StockCode    string  `json:"stockCode" gorm:"column:stock_code"`       // 股票代码
		Currency     string  `json:"currency" gorm:"column:currency"`          // 货币
		StandardDate string  `json:"standardDate" gorm:"column:standard_date"` // 标准日期
		StandardTime int64   `json:"standardTime" gorm:"column:standard_time"` // 标准日期时间戳
		ReportDate   string  `json:"reportDate" gorm:"column:report_date"`     // 报告时间
		ReportType   string  `json:"reportType" gorm:"column:report_type"`     // 报告类型
		Ta           float64 `json:"ta" gorm:"column:ta"`                      // 资产合计
		Cabb         float64 `json:"cabb" gorm:"column:cabb"`                  // 货币资金
		Tfa          float64 `json:"tfa" gorm:"column:tfa"`                    // 交易性金融资产
		Nr           float64 `json:"nr" gorm:"column:nr"`                      // 应收票据
		Ar           float64 `json:"ar" gorm:"column:ar"`                      // 应收账款
		Ats          float64 `json:"ats" gorm:"column:ats"`                    // 预付款项
		Or           float64 `json:"or" gorm:"column:or"`                      // 其他应收款
		Fahursa      float64 `json:"fahursa" gorm:"column:fahursa"`            // 买入返售金融资产
		I            float64 `json:"i" gorm:"column:i"`                        // 存货
		Oca          float64 `json:"oca" gorm:"column:oca"`                    // 其他流动资产
		Ncafsfa      float64 `json:"ncafsfa" gorm:"column:ncafsfa"`            // 可供出售金融资产(非流动)
		Htmi         float64 `json:"htmi" gorm:"column:htmi"`                  // 持有至到期投资
		Ltar         float64 `json:"ltar" gorm:"column:ltar"`                  // 长期应收款
		Ltei         float64 `json:"ltei" gorm:"column:ltei"`                  // 长期股权投资
		Rei          float64 `json:"rei" gorm:"column:rei"`                    // 投资性房地产
		Fa           float64 `json:"fa" gorm:"column:fa"`                      // 固定资产
		Cip          float64 `json:"cip" gorm:"column:cip"`                    // 在建工程
		Es           float64 `json:"es" gorm:"column:es"`                      // 工程物资
		Pba          float64 `json:"pba" gorm:"column:pba"`                    // 生产性生物资产
		Ia           float64 `json:"ia" gorm:"column:ia"`                      // 无形资产
		Gw           float64 `json:"gw" gorm:"column:gw"`                      // 商誉
		Ltpe         float64 `json:"ltpe" gorm:"column:ltpe"`                  // 长期待摊费用
		Onca         float64 `json:"onca" gorm:"column:onca"`                  // 其他非流动资产
		Tl           float64 `json:"tl" gorm:"column:tl"`                      // 负债合计
		Lwi          float64 `json:"lwi" gorm:"column:lwi"`                    // 有息负债
		Stl          float64 `json:"stl" gorm:"column:stl"`                    // 短期借款
		Bfcb         float64 `json:"bfcb" gorm:"column:bfcb"`                  // 向中央银行借款
		Np           float64 `json:"np" gorm:"column:np"`                      // 应付票据
		Ap           float64 `json:"ap" gorm:"column:ap"`                      // 应付账款
		Afc          float64 `json:"afc" gorm:"column:afc"`                    // 预收账款
		Oap          float64 `json:"oap" gorm:"column:oap"`                    // 其他应付款
		Ltl          float64 `json:"ltl" gorm:"column:ltl"`                    // 长期借款
		Sap          float64 `json:"sap" gorm:"column:sap"`                    // 专项应付款
		Ltap         float64 `json:"ltap" gorm:"column:ltap"`                  // 长期应付款
		Mc           float64 `json:"mc" gorm:"column:mc"`                      // 市值
	}

	// 现金流量表
	CashFlow struct {
		Index        int64   `json:"index" gorm:"column:index;primary_key"`
		StockCode    string  `json:"stockCode" gorm:"column:stock_code;"`      // 股票代码
		Currency     string  `gorm:"column:currency"`                          // 货币
		StandardDate string  `json:"standardDate" gorm:"column:standard_date"` // 标准日期
		StandardTime int64   `json:"standardTime" gorm:"column:standard_time"` // 标准日期时间戳
		ReportDate   string  `json:"reportDate" gorm:"column:report_date"`     // 报告时间
		ReportType   string  `json:"reportType" gorm:"column:report_type"`     // 报告类型
		Crfscapls    float64 `json:"crfscapls" gorm:"column:crfscapls"`        // 销售商品、提供劳务收到的现金
		Cpfpcarls    float64 `json:"cpfpcarls" gorm:"column:cpfpcarls"`        // 购买商品、接收劳务支付的现金
		Ncffoa       float64 `json:"ncffoa" gorm:"column:ncffoa"`              // 经营活动产生的现金流量净额
		Cpfpfiaolta  float64 `json:"cpfpfiaolta" gorm:"column:cpfpfiaolta"`    // 购建固定资产、无形资产及其他长期资产所支付的现金
		Cpfi         float64 `json:"cpfi" gorm:"column:cpfi"`                  // 投资所支付的现金
		Ncffia       float64 `json:"ncffia" gorm:"column:ncffia"`              // 投资活动产生的现金流量净额
		Cpfbrp       float64 `json:"cpfbrp" gorm:"column:cpfbrp"`              // 偿付债务支付的现金
		Cpfdapdoip   float64 `json:"cpfdapdoip" gorm:"column:cpfdapdoip"`      // 分配股利、利润或偿付利息所支付的现金
		Ncfffa       float64 `json:"ncfffa" gorm:"column:ncfffa"`              // 筹资活动产生的现金流量净额
		Niicace      float64 `json:"niicace" gorm:"column:niicace"`            // 现金及现金等价物的净增加额
		Bocaceatpe   float64 `json:"bocaceatpe" gorm:"column:bocaceatpe"`      // 期末现金及现金等价物净余额
	}

	ResponseData struct {
		Tag     string        `json:"tag"`
		Field   string        `json:"field"`
		Formula string        `json:"Z"`
		Data    []interface{} `json:"data"`
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
	p.GPF = profitFormula["GP"]
	p.GP = eval(p.GPF, *p)
	fmt.Println(p.GP, p.GPF)
	//TODO 获取时候根据默认公式计算出值， 保存到数据库， web可更改公式，取值时默认取已有值
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

func GetProfitTableData(slice []Profit) []ResponseData {

	p := reflect.ValueOf(Profit{})

	response := make([]ResponseData, p.NumField())

	for _, s := range slice {
		t := reflect.TypeOf(s)
		v := reflect.ValueOf(s)
		for i := 0; i < v.NumField(); i++ {
			name := t.Field(i).Name

			response[i].Field = name
			response[i].Tag = profitTags[name]
			formula := profitFormula[name]
			// ToDO 如果有公式 就计算， 没有公式就取值
			if formula != "" {
				response[i].Formula = formula
				result := eval(formula, s)
				if response[i].Data == nil {
					response[i].Data = []interface{}{result}
				} else {
					response[i].Data = append(response[i].Data, result)
				}
			} else {
				if response[i].Data == nil {
					response[i].Data = []interface{}{v.Field(i).Interface()}
				} else {
					response[i].Data = append(response[i].Data, v.Field(i).Interface())
				}
			}
		}
	}
	return response
}

func eval(formula string, data interface{}) float64 {

	formula = strings.Replace(formula, "=", "", 1)
	temp := strings.ReplaceAll(formula, "+", ",+,")
	temp = strings.ReplaceAll(temp, "-", ",-,")
	temp = strings.ReplaceAll(temp, "*", ",*,")
	temp = strings.ReplaceAll(temp, "/", ",/,")
	temp = strings.ReplaceAll(temp, "(", ",(,")
	temp = strings.ReplaceAll(temp, ")", ",),")
	tempArr := strings.Split(temp, ",")

	v := reflect.ValueOf(data)
	for i, d := range tempArr {
		switch d {
		case "+", "-", "*", "/", "(", ")":
			continue
		default:
			field := v.FieldByName(profitFormulaTag[d])
			value := util.FloatToString(field.Interface().(float64))
			tempArr[i] = value
		}
	}

	tempArr = calculateBrackets(tempArr)
	tempArr = calulateOrder(tempArr)
	result := calculate(tempArr)
	fmt.Println(result)
	return util.ParseFloat(result[0])
}

func handleTag(tag string) string {
	t := strings.ReplaceAll(tag, "default_", "")
	t = strings.ReplaceAll(t, "hidden_", "")
	return t
}

// 正式计算
func calculate(exp []string) []string {
	temp := exp
	if len(exp) == 1 {
		return exp
	}
c:
	for i, s := range temp {
		switch s {
		case "+":
			result := util.ParseFloat(temp[i-1]) + util.ParseFloat(temp[i+1])
			temp = append([]string{util.FloatToString(result)}, temp[i+2:]...)
			break c
		case "-":
			result := util.ParseFloat(temp[i-1]) - util.ParseFloat(temp[i+1])
			temp = append([]string{util.FloatToString(result)}, temp[i+2:]...)
			break c
		case "*":
			result := util.ParseFloat(temp[i-1]) * util.ParseFloat(temp[i+1])
			temp = append([]string{util.FloatToString(result)}, temp[i+2:]...)
		case "/":
			result := util.ParseFloat(temp[i-1]) / util.ParseFloat(temp[i+1])
			temp = append([]string{util.FloatToString(result)}, temp[i+2:]...)
		}
	}

	if len(temp) != 1 {
		temp = calculate(temp)
	}

	return temp
}

// 判断括号优先级
func calculateBrackets(exp []string) []string {
	startIndex := 0
	endIndex := 0
	moreBrackets := false
	for i, s := range exp {
		switch s {
		case "(":
			if startIndex != 0 {
				moreBrackets = true
			} else {
				startIndex = i
			}
		case ")":
			if endIndex != 0 {
				moreBrackets = true
			} else {
				endIndex = i
			}
		}
	}
	if startIndex == 0 {
		return exp
	}

	subExp := exp[startIndex+1 : endIndex]
	result := calculate(subExp)
	temp := append(exp[:startIndex], result...)
	temp = append(temp, exp[endIndex+1:]...)

	if moreBrackets {
		temp = calculateBrackets(temp)
	}

	return temp
}

// 判断乘除优先级
func calulateOrder(exp []string) (result []string) {
	needC := false
n:
	for _, s := range exp {
		switch s {
		case "*", "/":
			needC = true
			break n
		}
	}
	if needC {
	c:
		for i, s := range exp {
			switch s {
			case "*", "/":
				result = calulate2(exp, i)
				break c
			}
		}
		result = calulateOrder(result)
	} else {
		result = exp
	}

	return
}

// 乘除替换数字
func calulate2(exp []string, index int) (result []string) {
	subExp := exp[index-1 : index+2]
	r := calculate(subExp)
	result = append(exp[:index-1], r...)
	result = append(result, exp[index+2:]...)
	return
}
