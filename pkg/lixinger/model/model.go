package model

import "time"

type (
	ResponseForProfit struct {
		Code    int
		Message string
		Data    []ProfitResponse
	}

	ResponseForBalanceSheet struct {
		Code    int
		Message string
		Data    []BalanceSheetResponse
	}

	ResponseForCashFlow struct {
		Code    int
		Message string
		Data    []CashFlowResponse
	}

	ResponseForFinancialIndex struct {
		Code    int
		Message string
		Data    []FinancialIndexResponse
	}

	ResponseData interface {
		GetMapData() map[string]interface{}
	}

	ProfitResponse struct {
		StockCode    string
		Currency     string
		StandardDate string
		ReportDate   string
		ReportType   string
		Q            ps
	}

	BalanceSheetResponse struct {
		StockCode    string
		Currency     string
		StandardDate string
		ReportDate   string
		ReportType   string
		Q            bs
	}

	CashFlowResponse struct {
		StockCode    string
		Currency     string
		StandardDate string
		ReportDate   string
		ReportType   string
		Q            cfs
	}

	FinancialIndexResponse struct {
		StockCode    string
		Currency     string
		StandardDate string
		ReportDate   string
		ReportType   string
		Q            m
	}

	ps struct {
		Ps map[string]T
	}

	bs struct {
		Bs map[string]T
	}

	cfs struct {
		Cfs map[string]T
	}

	m struct {
		M map[string]T
	}

	T struct {
		T interface{}
	}
)

func (p ProfitResponse) GetMapData() map[string]interface{} {
	data := make(map[string]interface{})
	data["stockCode"] = p.StockCode
	data["currency"] = p.Currency
	st, _ := time.Parse(time.RFC3339, p.StandardDate)
	data["standardDate"] = st.Format("2006-01-02")
	data["standardTime"] = st.Unix()
	rt, _ := time.Parse(time.RFC3339, p.ReportDate)
	data["reportDate"] = rt.Format("2006-01-02")
	data["reportType"] = p.ReportType
	for k, v := range p.Q.Ps {
		data[k] = v.T
	}

	return data
}

func (b BalanceSheetResponse) GetMapData() map[string]interface{} {
	data := make(map[string]interface{})
	data["stockCode"] = b.StockCode
	data["currency"] = b.Currency
	st, _ := time.Parse(time.RFC3339, b.StandardDate)
	data["standardDate"] = st.Format("2006-01-02")
	data["standardTime"] = st.Unix()
	rt, _ := time.Parse(time.RFC3339, b.ReportDate)
	data["reportDate"] = rt.Format("2006-01-02")
	data["reportType"] = b.ReportType
	for k, v := range b.Q.Bs {
		data[k] = v.T
	}

	return data
}

func (r CashFlowResponse) GetMapData() map[string]interface{} {
	data := make(map[string]interface{})
	data["stockCode"] = r.StockCode
	data["currency"] = r.Currency
	st, _ := time.Parse(time.RFC3339, r.StandardDate)
	data["standardDate"] = st.Format("2006-01-02")
	data["standardTime"] = st.Unix()
	rt, _ := time.Parse(time.RFC3339, r.ReportDate)
	data["reportDate"] = rt.Format("2006-01-02")
	data["reportType"] = r.ReportType
	for k, v := range r.Q.Cfs {
		data[k] = v.T
	}

	return data
}

func (r FinancialIndexResponse) GetMapData() map[string]interface{} {
	data := make(map[string]interface{})
	data["stockCode"] = r.StockCode
	data["currency"] = r.Currency
	st, _ := time.Parse(time.RFC3339, r.StandardDate)
	data["standardDate"] = st.Format("2006-01-02")
	data["standardTime"] = st.Unix()
	rt, _ := time.Parse(time.RFC3339, r.ReportDate)
	data["reportDate"] = rt.Format("2006-01-02")
	data["reportType"] = r.ReportType
	for k, v := range r.Q.M {
		data[k] = v.T
	}

	return data
}
