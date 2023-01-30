package status

var (
	Success      = &Status{Code: 0, Msg: "success"}
	Error        = &Status{Code: 500, Msg: "server error"}
	InvalidParam = &Status{Code: 400, Msg: "invalid params"}

	// 100xx 基本问题
	AccessErr = &Status{Code: 10000, Msg: "no authority"}
)

type StandardResponse struct {
	Sta  *Status     `json:"status"`
	Data interface{} `json:"data"`
}

type Status struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
