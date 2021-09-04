package utils

import (
	"blog-api-golang/types"
)

func GetErrorMessage(message string) types.ErrorRespone {
	return types.ErrorRespone{Status: "fail", Message: message}
}

func GetSuccessMessage(data interface{}) types.SuccessRespone {
	return types.SuccessRespone{Status: "success", Data: data}
}
