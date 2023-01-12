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
)
