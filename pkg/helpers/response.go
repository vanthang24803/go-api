package helpers

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func Send(status int, data interface{}) Response {
	resp := Response{
		Status: status,
		Result: data,
	}
	return resp
}
