package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"meaile-server/meaile-user/ai/model"
	model2 "meaile-server/meaile-user/model"
	"net/http"
	"time"
)

type AiServiceImpl struct {
}

// 调用DeepSeek API
func (ai *AiServiceImpl) CallDeepSeekAPI(apiKey string, req model.DeepSeekRequest) *model2.Response {
	requestBody, _ := json.Marshal(req)

	client := &http.Client{Timeout: 30 * time.Second}
	request, _ := http.NewRequest(
		"POST",
		"https://api.deepseek.com/v1/chat/completions",
		bytes.NewBuffer(requestBody),
	)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(request)
	if err != nil {
		return &model2.Response{
			Code: model2.FAILED,
			Msg:  fmt.Sprintf("API请求失败: %v", err),
			Data: err,
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return &model2.Response{
			Code: model2.FAILED,
			Msg:  fmt.Sprintf("API返回错误状态码: %d, 响应: %s", resp.StatusCode, string(body)),
			Data: nil,
		}
	}

	var result model.DeepSeekResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return &model2.Response{
			Code: model2.FAILED,
			Msg:  fmt.Sprintf("响应解析失败: %v", err),
			Data: nil,
		}
	}
	// 处理API错误
	if len(result.Choices) == 0 {
		return &model2.Response{
			Code: model2.FAILED,
			Msg:  "AI服务返回空响应",
			Data: result.Error.Message,
		}
	}
	return &model2.Response{
		Code: model2.SUCCESS,
		Msg:  "请求成功",
		Data: result.Choices[0].Message.Content,
	}
}
