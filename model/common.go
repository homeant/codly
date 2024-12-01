package model

type StatusMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SuccessMessage(message string) *StatusMessage {
	return &StatusMessage{
		Success: true,
		Message: message,
	}
}

func ErrorMessage(message string) *StatusMessage {
	return &StatusMessage{
		Success: false,
		Message: message,
	}
}
