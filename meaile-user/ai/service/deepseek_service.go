package service

import (
	"meaile-server/meaile-user/ai/model"
	model2 "meaile-server/meaile-user/model"
)

type AiService interface {
	CallDeepSeekAPI(apiKey string, req model.DeepSeekRequest) *model2.Response
}
