package schema

import "time"

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

// Tag Event has the same payload with Push Event
type TagEventPayload PushEventPayload

type IssueEventPayload struct {
	ObjectKind       string           `json:"object_kind"`
	EventType        string           `json:"event_type"`
	User             User             `json:"user"`
	ProjectID        int32            `json:"project_id"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Repository       Repository
	Assignees        []User
	Assignee         User
	Labels           []Label
	Changes          struct {
		UpdatedByID struct {
			Previous int32 `json:"previous"`
			Current  int32 `json:"current"`
		} `json:"updated_by_id"`
		UpdatedAt struct {
			Previous time.Time `json:"previous"`
			Current  time.Time `json:"current"`
		} `json:"updated_at"`
		Labels struct {
			Previous []Label `json:"previous"`
			Current  []Label `json:"current"`
		} `json:"labels"`
	} `json:"changes"`
}
