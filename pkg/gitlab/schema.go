package gitlab

import (
	"strings"
	"time"
)

// Custom time struct for Gitlab only
type customTime struct {
	time.Time
}

func (t *customTime) UnmarshalJSON(b []byte) (err error) {
	layout := []string{
		"2006-01-02 15:04:05 MST",
		"2006-01-02 15:04:05 Z07:00",
		"2006-01-02 15:04:05 Z0700",
		time.RFC3339,
	}
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	for _, l := range layout {
		t.Time, err = time.Parse(l, s)
		if err == nil {
			break
		}
	}
	return
}

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
	ID        string     `json:"id"`
	Message   string     `json:"message"`
	Title     string     `json:"title"`
	Timestamp customTime `json:"timestamp"`
	URL       string     `json:"url"`
	Author    NameEmail  `json:"author"`
	Added     []string   `json:"added"`
	Modified  []string   `json:"modified"`
	Removed   []string   `json:"removed"`
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

type ProjectWithoutID struct {
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

type Project struct {
	ID int32 `json:"id"`
	ProjectWithoutID
}

type Label struct {
	ID          int32      `json:"id"`
	Title       string     `json:"title"`
	Color       string     `json:"color"`
	ProjectID   int32      `json:"project_id"`
	CreatedAt   customTime `json:"created_at"`
	UpdatedAt   customTime `json:"updated_at"`
	Template    bool       `json:"template"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	GroupID     int32      `json:"group_id"`
}

type ObjectAttributes struct {
	ID                  int32      `json:"name"`
	Title               string     `json:"title"`
	AssigneeIDs         []int32    `json:"assignee_ids"`
	AssigneeID          int32      `json:"assignee_id"`
	AuthorID            int32      `json:"author_id"`
	ProjectID           int32      `json:"project_id"`
	CreatedAt           customTime `json:"created_at"`
	UpdatedAt           customTime `json:"updated_at"`
	UpdatedByID         int32      `json:"updated_by_id"`
	LastEditedAt        customTime `json:"last_edited_at"`
	LastEditedByID      int32      `json:"last_edited_by_id"`
	RelativePostion     int32      `json:"relative_position"`
	Description         string     `json:"description"`
	MilestoneID         int32      `json:"milestone_id"`
	StateID             int32      `json:"state_id"`
	Confidential        bool       `json:"confidential"`
	DiscussionLocked    bool       `json:"discussion_locked"`
	DueDate             customTime `json:"due_date"`
	MovedToID           int32      `json:"moved_to_id"`
	DuplicatedToID      int32      `json:"duplicated_to_id"`
	TimeEstimate        int32      `json:"time_estimate"`
	TotalTimeSpent      int32      `json:"total_time_spent"`
	TimeChange          int32      `json:"time_change"`
	HumanTotalTimeSpent *string    `json:"human_total_time_spent"`
	HumanTimeEstimate   *string    `json:"human_time_estimate"`
	HumanTimeChange     *string    `json:"human_time_change"`
	Weight              *int32     `json:"weight"`
	IID                 int32      `json:"iid"`
	URL                 string     `json:"url"`
	State               string     `json:"state"`
	Action              string     `json:"action"`
	Serverity           string     `json:"serverity"`
	EscalationStatus    string     `json:"escalation_status"`
	EscalationPolicy    struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	} `json:"escalation_policy"`
	Labels []Label `json:"labels"`
}

type MergeRequest struct {
	ID              int32            `json:"id"`
	TargetBranch    string           `json:"target_branch"`
	SourceBranch    string           `json:"source_bracnh"`
	SourceProjectID int32            `json:"source_project_id"`
	AuthorID        int32            `json:"author_id"`
	AssigneeID      int32            `json:"assignee_id"`
	Title           string           `json:"title"`
	CreatedAt       customTime       `json:"created_at"`
	UpdatedAt       customTime       `json:"updated_at"`
	MilestoneID     int32            `json:"milestone_id"`
	State           string           `json:"state"`
	MergeStatus     string           `json:"merge_status"`
	TargetProjectID int32            `json:"target_project_id"`
	IID             int32            `json:"iid"`
	Description     string           `json:"description"`
	Position        int32            `json:"postition"`
	Labels          []Label          `json:"labels"`
	Source          ProjectWithoutID `json:"source"`
	Target          ProjectWithoutID `json:"target"`
	LastCommit      Commit           `json:"last_commit"`
	WorkInProgress  bool             `json:"work_in_progress"`
	Assignee        struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"assignee"`
	DetailedMergeStatus string `json:"detailed_merge_status"`
}

type Issue struct {
	ID          int32      `json:"id"`
	Title       string     `json:"title"`
	AssigneeIDs []int32    `json:"assignee_ids"`
	AuthorID    int32      `json:"author_id"`
	ProjectID   int32      `json:"project_id"`
	CreatedAt   customTime `json:"created_at"`
	UpdatedAt   customTime `json:"updated_at"`
	Position    int32      `json:"postition"`
	BranchName  string     `json:"branch_name"`
	Description string     `json:"description"`
	MilestoneID int32      `json:"milestone_id"`
	State       string     `json:"state"`
	IID         int32      `json:"iid"`
	Labels      []Label    `json:"labels"`
}

type Snippet struct {
	ID        int32      `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	AuthorID  int32      `json:"author_id"`
	ProjectID int32      `json:"project_id"`
	CreatedAt customTime `json:"created_at"`
	UpdatedAt customTime `json:"updated_at"`
	FileName  string     `json:"file_name"`
	//ExpiresAt string `json:"expires_at"` // This field is described is `null`, may be it is deprecated so this library currently does not support this field
	Type            string `json:"type"`
	VisibilityLevel int32  `json:"visibility_level"`
}

type Change struct {
	UpdatedByID struct {
		Previous int32 `json:"previous"`
		Current  int32 `json:"current"`
	} `json:"updated_by_id"`
	UpdatedAt struct {
		Previous customTime `json:"previous"`
		Current  customTime `json:"current"`
	} `json:"updated_at"`
	Labels struct {
		Previous []Label `json:"previous"`
		Current  []Label `json:"current"`
	} `json:"labels"`
}

type Build struct {
	ID             int32      `json:"id"`
	Stage          string     `json:"stage"`
	Name           string     `json:"name"`
	Status         string     `json:"status"`
	CreatedAt      customTime `json:"created_at"`
	StartedAt      customTime `json:"started_at"`
	FinishedAt     customTime `json:"finished_at"`
	Duration       int32      `json:"duration"`
	QueuedDuration int32      `json:"queued_duration"`
	FailureReason  string     `json:"failure_reason"`
	When           string     `json:"when"`
	Manual         bool       `json:"manual"`
	AllowFailure   bool       `json:"allow_failure"`
	User           User       `json:"user"`
	Runner         string     `json:"runner"`
	ArtifactsFile  struct {
		Filename string `json:"filename"`
		Size     int32  `json:"size"`
	} `json:"artifacts_file"`
	Environment struct {
		Name           string `json:"name"`
		Action         string `json:"action"`
		DeploymentTier string `json:"deployment_tier"`
	} `json:"environment"`
}

type Runner struct {
	Active      bool     `json:"active"`
	RunnerType  string   `json:"runner_type"`
	IsShared    bool     `json:"is_shared"`
	Id          int32    `json:"id"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
