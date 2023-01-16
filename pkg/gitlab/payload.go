package gitlab

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
	Repository       Repository       `json:"resposity"`
	Assignees        []User           `json:"assignees"`
	Assignee         User             `json:"assignee"`
	Labels           []Label          `json:"labels"`
	Changes          Change           `json:"changes"`
}

type CommentEventPayload struct {
	ObjectKind       string           `json:"object_kind"`
	EventType        string           `json:"event_type"`
	User             User             `json:"user"`
	ProjectID        int32            `json:"project_id"`
	Project          Project          `json:"project"`
	Repository       Repository       `json:"repository"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`

	// Comment on a commit
	Commit *Commit `json:"commit"`

	// Comment on a merge request
	MergeRequest *MergeRequest `json:"merge_request"`

	// Comment on an issue
	Issue *Issue `json:"issue"`

	// Comment on a code snippet
	Snippet *Snippet `json:"snippet"`
}

type MergeRequestEventPayload struct {
	ObjectKind       string           `json:"object_kind"`
	EventType        string           `json:"event_type"`
	User             User             `json:"user"`
	Project          Project          `json:"project"`
	Repository       Repository       `json:"repository"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Labels           []Label          `json:"labels"`
	Changes          Change           `json:"changes"`
	Assignees        []User           `json:"assignees"`
	Reviewers        []User           `json:"reviewers"`
}

type WikiPageEventPlayload struct {
	ObjectKind string  `json:"object_kind"`
	User       User    `json:"user"`
	Project    Project `json:"project"`
	Wiki       struct {
		WebURL            string `json:"web_url"`
		GitSSHURL         string `json:"git_ssh_url"`
		GitHTTPURL        string `json:"git_http_url"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch"`
	} `json:"wiki"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}

type PipelineEventPayload struct {
	ObjectKind       string           `json:"object_kind"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	MergeRequest     MergeRequest     `json:"merge_request"`
	User             User             `json:"user"`
	Project          Project          `json:"project"`
	Commit           Commit           `json:"commit"`
	SourcePipeline   struct {
		Project struct {
			ID                int32  `json:"id"`
			WebURL            string `json:"web_url"`
			PathWithNamespace string `json:"path_with_namespace"`
		}
		PipelineID int32 `json:"pipeline_id"`
		JobID      int32 `json:"job_id"`
	} `json:"source_pipeline"`
	Builds []Build `json:"builds"`
}

type JobEventPayload struct {
	ObjectKind         string     `json:"object_kind"`
	Ref                string     `json:"ref"`
	Tag                bool       `json:"tag"`
	BeforeSHA          string     `json:"before_sha"`
	SHA                string     `json:"sha"`
	BuildID            int32      `json:"build_id"`
	BuildName          string     `json:"build_name"`
	BuildState         string     `json:"build_state"`
	BuildStatus        string     `json:"build_status"`
	BuildCreatedAt     customTime `json:"build_created_at"`
	BuildStartedAt     customTime `json:"build_started_at"`
	BuildFinishedAt    customTime `json:"build_finished_at"`
	BuildDuration      int32      `json:"build_duration"`
	BuildAllowFailure  bool       `json:"build_allow_failure"`
	BuildFailureReason string     `json:"build_failure_reason"`
	RetriesCount       int32      `json:"retries_count"`
	PipelineId         int32      `json:"pipeline_id"`
	ProjectId          int32      `json:"project_id"`
	ProjectName        string     `json:"project_name"`
	User               User       `json:"user"`
	Commit             Commit     `json:"commit"`
	Repository         Repository `json:"repository"`
	Runner             Runner     `json:"runner"`
	Environment        string     `json:"environment"`
}

type DeploymentPayload struct {
	ObjectKind             string     `json:"object_kind"`
	Status                 string     `json:"status"`
	StatusChangedAt        customTime `json:"status_changed_at"`
	DeploymentID           int32      `json:"deployment_id"`
	DeployableID           int32      `json:"deployable_id"`
	DeployableURL          string     `json:"deployable_url"`
	Environment            string     `json:"environment"`
	EnvironmentSlug        string     `json:"environment_slug"`
	EnvironmentExternalURL string     `json:"environment_external_url"`
	Project                Project    `json:"project"`
	ShortSHA               string     `json:"short_sha"`
	User                   User       `json:"user"`
	UserURL                string     `json:"user_url"`
	CommitURL              string     `json:"commit_url"`
	CommitTitle            string     `json:"commit_title"`
}

type MemberEventPayload struct {
	CreatedAt    customTime `json:"created_at"`
	UpdatedAt    customTime `json:"updated_at"`
	GroupName    string     `json:"group_name"`
	GroupPath    string     `json:"group_path"`
	GroupID      int32      `json:"group_id"`
	UserUsername string     `json:"user_username"`
	UserName     string     `json:"user_name"`
	UserEmail    string     `json:"user_email"`
	UserID       int32      `json:"user_id"`
	GroupAccess  string     `json:"group_access"`
	GroupPlan    string     `json:"group_plan"`
	ExpiresAt    customTime `json:"expires_at"`
	EventName    string     `json:"event_name"`
}

type SubGroupEventPayload struct {
	CreatedAt      customTime `json:"created_at"`
	UpdatedAt      customTime `json:"updated_at"`
	EventName      string     `json:"event_name"`
	Name           string     `json:"name"`
	Path           string     `json:"path"`
	FullPath       string     `json:"full_path"`
	GroupID        int32      `json:"group_id"`
	ParentGroupID  int32      `json:"parent_group_id"`
	ParentName     string     `json:"parent_name"`
	ParentPath     string     `json:"parent_path"`
	ParentFullPath string     `json:"parent_full_path"`
}

type FeatureFlagEventPayload struct {
	ObjectKind       string  `json:"object_kind"`
	Project          Project `json:"project"`
	User             User    `json:"user"`
	UserURL          string  `json:"user_url"`
	ObjectAttributes struct {
		ID          int32  `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Active      bool   `json:"active"`
	} `json:"object_attributes"`
}

type ReleaseEventPayload struct {
	ID          int32      `json:"id"`
	CreatedAt   customTime `json:"created_at"`
	Description string     `json:"description"`
	Name        string     `json:"name"`
	ReleasedAt  customTime `json:"released_at"`
	Tag         string     `json:"tag"`
	ObjectKind  string     `json:"object_kind"`
	Project     Project    `json:"project"`
	URL         string     `json:"url"`
	Action      string     `json:"action"`
	Assets      struct {
		Count int32 `json:"count"`
		Links []struct {
			ID       int32  `json:"id"`
			External bool   `json:"external"`
			LinkType string `json:"link_type"`
			Name     string `json:"name"`
			URL      string `json:"url"`
		} `json:"links"`
		Sources []struct {
			Format string `json:"format"`
			URL    string `json:"url"`
		} `json:"sources"`
	} `json:"assets"`
	Commit Commit `json:"commit"`
}
