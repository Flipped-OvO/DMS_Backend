package request

type HTTPError struct {
	Status int    `json:"status" example:"http status code"`
	Msg    string `json:"msg" example:"status bad request"`
}

func (e HTTPError) Error() string {
	return e.Msg
}
