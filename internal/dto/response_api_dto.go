package dto

type ResponseApiDto struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}
