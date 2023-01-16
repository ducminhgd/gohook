package schema

type Event string

const (
	HTTPEventRequestHeader string = "X-Gitlab-Event"

	// X-Gitlab-Event
	PushEvent               Event = "Push Hook"
	TagEvent                Event = "Tag Push Hook"
	IssuesEvent             Event = "Issue Hook"
	ConfidentialIssuesEvent Event = "Confidential Issue Hook"
	CommentEvent            Event = "Note Hook"
	MergeRequestEvent       Event = "Merge Request Hook"
	WikiPageEvent           Event = "Wiki Page Hook"
	PipelineEvent           Event = "Pipeline Hook"
	BuildEvent              Event = "Build Hook"
	JobEvent                Event = "Job Hook"
	SystemHookEvent         Event = "System Hook"

	// Object Kind
	ObjectKindPush         string = "push"
	ObjectKindTag          string = "tag_push"
	ObjectKindMergeRequest string = "merge_request"
	ObjectKindBuild        string = "build"
	ObjectKindComment      string = "note"
	ObjectKindWiki         string = "wiki_page"
	ObjectKindPipeline     string = "pipeline"
	ObjectKindJob          string = ObjectKindBuild
	ObjectKindDeployment   string = "deployment"
	ObjectKindFeatureFlag  string = "feature_flag"

	IssueStateOpened string = "opened"
	IssueStateClosed string = "closed"
	IssueStateLocked string = "locked"
	IssueStateMerged string = "merged"
	IssueStateAll    string = "all"

	MRStateOpened string = "opened"
	MRStateClosed string = "closed"
	MRStateLocked string = "locked"
	MRStateMerged string = "merged"
	MRStateAll    string = "all"

	MRStatusUnchecked             string = "unchecked"
	MRStatusChecking              string = "checking"
	MRStatusCanBeMerged           string = "can_be_merged"
	MRStatusCannotBeMerged        string = "cannot_be_merged"
	MRStatusCannotBeMergedRecheck string = "cannot_be_merged_recheck"

	MemberEventAdd    string = "user_add_to_group"
	MemberEventUpdate string = "user_update_for_group"
	MemberEventDelete string = "user_remove_from_group"
)
