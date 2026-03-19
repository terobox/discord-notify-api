package dto

// SendMessageRequest represents the incoming request payload
type SendMessageRequest struct {
	ChannelID string `json:"channel_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Source    string `json:"source,omitempty"` // e.g. "monitoring", "ci/cd", "cron-job"
}

// SendMessageResponse represents the API response
type SendMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
