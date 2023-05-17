package validators

type ReportForumRequest struct {
    Problem string `json:"problem"`
}

type ReportForumValidator struct {}

func (reportForumValidator *ReportForumValidator) Validate(request ReportForumRequest) (bool, []string) {
    messages := []string{}

    if request.Problem == "" || len(request.Problem) < 1 {
        messages = append(messages, "Content is required and must be at least 1 characters")
    }

    return len(messages) == 0, messages
}