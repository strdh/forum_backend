package validators

type ForumRequest struct {
    Title string `json:"Title"`
    Description string `json:"description"`
}

type SearchRequest struct {
    Search string `json:"search"`
}

type ForumValidator struct {}

func (forumValidator *ForumValidator) ValidateForum(request ForumRequest) (bool, []string) {
    messages := []string{}

    if request.Title == "" || len(request.Title) < 5 {
        messages = append(messages, "Title is required and must be at least 5 characters")
    }

    if request.Description == "" || len(request.Description) < 15 {
        messages = append(messages, "Description is required and must be at least 15 characters")
    }

    return len(messages) == 0, messages
}

func (forumValidator *ForumValidator) ValidateSearch(request SearchRequest) (bool, []string) {
    messages := []string{}

    if request.Search == "" || len(request.Search) < 3 {
        messages = append(messages, "Search is required and must be at least 3 characters")
    }

    return len(messages) == 0, messages
}