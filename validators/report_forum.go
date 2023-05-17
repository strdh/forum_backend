package validators

type ReportForumRequest struct {
    Problem string `json:"problem"`
}

type ReportForumValidator struct {}

func (reportForumValidator *ReportForumValidator) Validate(request ReportForumRequest) (bool, []string) {
    messages := []string{}

    if request.Problem == "" || len(request.Problem) < 10 {
        messages = append(messages, "Cannot and must be at least 10 characters")
    }

    return len(messages) == 0, messages
}