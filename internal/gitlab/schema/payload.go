package schema

type PushEventPayload struct {
	ObjectKind        string     `json:"object_kind"`
	EventName         string     `json:"event_name"`
	BeforeSHA         string     `json:"before"`
	AfterSHA          string     `json:"after"`
	Ref               string     `json:"ref"`
	CheckoutSHA       string     `json:"checkout_sha"`
	UserID            int32      `json:"user_id"`
	UserName          string     `json:"user_name"`
	UserUsername      string     `json:"user_username"`
	UserEmail         string     `json:"user_email"`
	UserAvatar        string     `json:"user_avatar"`
	ProjectID         int32      `json:"project_id"`
	Project           Project    `json:"project"`
	Repository        Repository `json:"repository"`
	Commits           []Commit   `json:"commits"`
	TotalCommitsCount int32      `json:"total_commits_count"`
}
