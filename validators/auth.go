package validators

import (
    "strings"
    "unicode"
    "regexp"
    "xyzforum/config"
)

type RegisterRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Email string `json:"email"`
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type AuthValidator struct {}

func (authValidator *AuthValidator) ValidateRegister(request RegisterRequest) (bool, []string) {
    messages := []string{}

    if request.Username == "" || len(request.Username) < 3 {
        messages = append(messages, "Username is required and must be at least 3 characters")
    } else {
        // Validate username format using regular expression
        usernameRegex := regexp.MustCompile("^[a-z0-9_]+$")
        if !usernameRegex.MatchString(request.Username) {
            messages = append(messages, "Username must contain only lowercase letters, numbers[0-9], and underscores")
        } else {
            result, _ := config.DB.Query("SELECT username FROM users WHERE username = ?", request.Username)
            if result.Next() {
                messages = append(messages, "Username is already taken")
            }
        }
    }

    if request.Email != "" {
        pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	    regex := regexp.MustCompile(pattern)
	    isValid := regex.MatchString(request.Email)

        if !isValid {
            messages = append(messages, "Email is not valid")
        } else {
            result, _ := config.DB.Query("SELECT email FROM users WHERE email = ?", request.Email)
            if result.Next() {
                messages = append(messages, "Email is already taken")
            }
        }
    } else {
        messages = append(messages, "Email is required")
    }

    if request.Password == "" || len(request.Password) < 6 {
        messages = append(messages, "Password is required and must be at least 6 characters")
    } else {
        hasLower := false
        hasUpper := false
        hasSpecial := false
        hasNumber := false

        for _, char := range request.Password {
            if unicode.IsLower(char) {
                hasLower = true
            } else if unicode.IsUpper(char) {
                hasUpper = true
            } else if strings.ContainsAny(string(char), "!@#$%^&*()") {
                hasSpecial = true
            } else if unicode.IsNumber(char) {
                hasNumber = true
            }
        }

        if !hasLower || !hasUpper || !hasSpecial || !hasNumber {
            messages = append(messages, "Password must contain at least one lowercase letter, uppercase letter, number, and special character")
        }
    }

    if len(messages) > 0 {
        return false, messages
    }

    return true, messages
}

func (authValidator *AuthValidator) ValidateLogin(request LoginRequest) (bool, []string) {
    messages := []string{}

    if request.Username == "" {
        messages = append(messages, "Username is required")
    }

    if request.Password == "" {
        messages = append(messages, "Password is required")
    }

    if len(messages) > 0 {
        return false, messages
    }

    return true, messages
}

