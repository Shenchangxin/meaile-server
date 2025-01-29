package model

type DeepSeekRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type UserRequest struct {
	Message string `json:"message" binding:"required"`
}
