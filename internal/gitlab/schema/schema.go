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

type Label struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	Color       string    `json:"color"`
	ProjectID   int32     `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Template    bool      `json:"template"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	GroupID     int32     `json:"group_id"`
}

type ObjectAttributes struct {
	ID                  int32     `json:"name"`
	Title               string    `json:"title"`
	AssigneeIDs         []int32   `json:"assignee_ids"`
	AssigneeID          int32     `json:"assignee_id"`
	AuthorID            int32     `json:"author_id"`
	ProjectID           int32     `json:"project_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	UpdatedByID         int32     `json:"updated_by_id"`
	LastEditedAt        time.Time `json:"last_edited_at"`
	LastEditedByID      int32     `json:"last_edited_by_id"`
	RelativePostion     int32     `json:"relative_position"`
	Description         string    `json:"description"`
	MilestoneID         int32     `json:"milestone_id"`
	StateID             int32     `json:"state_id"`
	Confidential        bool      `json:"confidential"`
	DiscussionLocked    bool      `json:"discussion_locked"`
	DueDate             time.Time `json:"due_date"`
	MovedToID           int32     `json:"moved_to_id"`
	DuplicatedToID      int32     `json:"duplicated_to_id"`
	TimeEstimate        int32     `json:"time_estimate"`
	TotalTimeSpent      int32     `json:"total_time_spent"`
	TimeChange          int32     `json:"time_change"`
	HumanTotalTimeSpent *string   `json:"human_total_time_spent"`
	HumanTimeEstimate   *string   `json:"human_time_estimate"`
	HumanTimeChange     *string   `json:"human_time_change"`
	Weight              *int32    `json:"weight"`
	IID                 int32     `json:"iid"`
	URL                 string    `json:"url"`
	State               string    `json:"state"`
	Action              string    `json:"action"`
	Serverity           string    `json:"serverity"`
	EscalationStatus    string    `json:"escalation_status"`
	EscalationPolicy    struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	} `json:"escalation_policy"`
	Labels []Label `json:"labels"`
}
