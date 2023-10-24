package common

type successResponse struct {
	Data interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *successResponse {
	return &successResponse{
		Data: data,
	}
}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return NewSuccessResponse(data)
}
