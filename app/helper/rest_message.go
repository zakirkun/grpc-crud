package helper

import "github.com/zakirkun/grpc-crud/app/domain/types"

func RestMessage(code int, message string, data interface{}) types.GeneralResponse {
	rest := types.GeneralResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return rest
}
