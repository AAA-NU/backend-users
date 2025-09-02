package requests

type SaveUserRequest struct {
	TelegramID string `json:"tgID"`
	Language   string `json:"language"`
}
