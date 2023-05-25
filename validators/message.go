package validators

type MessageRequest struct {
    Message string `json:"message"`
}

type MessageValidator struct {}

func (messageValidator *MessageValidator) ValidateMessage(request MessageRequest) (bool, []string) {
    messages := []string{}

    if request.Message == "" || len(request.Message) < 1 {
        messages = append(messages, "Content is required and must be at least 1 characters")
    }

    return len(messages) == 0, messages
}