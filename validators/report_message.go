package validators

type ReportMessageRequest struct {
    Problem string `json:"problem"`
}

type ReportMessageValidator struct {}

func (reportMessageValidator *ReportMessageValidator) Validate(request ReportMessageRequest) (bool, []string) {
    messages := []string{}

    if request.Problem == "" || len(request.Problem) < 10 {
        messages = append(messages, "Cannot be null and must be at least 10 characters")
    }

    return len(messages) == 0, messages
}