package controller

import (
	"github.com/gin-gonic/gin"
	"meaile-server/meaile-user/ai/model"
	impl2 "meaile-server/meaile-user/ai/service/impl"
	"meaile-server/meaile-user/global"
	"net/http"
)

// 处理聊天请求
func HandleChat(ctx *gin.Context) {
	// 绑定用户请求
	var userReq model.UserRequest
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求格式"})
		return
	}

	// 构造系统提示词（食谱领域定制）
	systemPrompt := `你是一位专业营养师和厨艺专家，请遵循以下原则：
	1. 根据用户需求提供精准的食谱建议
	2. 考虑食材的季节性和易得性
	3. 标注营养成分和卡路里
	4. 给出明确的制作步骤
	5. 适合家庭制作的方案`

	// 构建DeepSeek请求
	deepSeekReq := model.DeepSeekRequest{
		Model: "deepseek-chat",
		Messages: []model.ChatMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userReq.Message},
		},
	}
	aiService := impl2.AiServiceImpl{}

	// 调用DeepSeek API
	apiKey := global.ServerConfig.AiConfig.DeepSeekKey
	response := aiService.CallDeepSeekAPI(apiKey, deepSeekReq)
	ctx.JSON(http.StatusOK, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": response.Data,
	})
	return
}
