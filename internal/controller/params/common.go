package params

type Common400resp struct {
	Code  string `json:"code"`
	Msg   string `json:"msg"`
	Error string `json:"error"`
}

type Common500resp struct {
	Code  string `json:"code"`
	Msg   string `json:"msg"`
	Error string `json:"error"`
}
