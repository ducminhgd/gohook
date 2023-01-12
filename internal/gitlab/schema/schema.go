package schema

import "time"

// NameEmail contains only name and email of an user
type NameEmail struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	NameEmail
	ID        int32  `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

type Commit struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Title     string    `json:"title"`
	Timestamp time.Time `json:"timestamp"`
	URL       string    `json:"url"`
	Author    NameEmail `json:"author"`
	Added     []string  `json:"added"`
	Modified  []string  `json:"modified"`
	Removed   []string  `json:"removed"`
}

type Repository struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitHTTPURL      string `json:"git_http_url"`
	GitSSHURL       string `json:"git_ssh_url"`
	VisibilityLevel int32  `json:"visibility_level"`
}

type Project struct {
	ID                int32  `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitHTTPURL        string `json:"git_http_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int32  `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}
