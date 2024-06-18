package schemas

type ErrorResponse struct {
	Detail string `json:"detail"`
	Code   int    `json:"code"`
}
