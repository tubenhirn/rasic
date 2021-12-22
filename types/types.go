package types

import "time"

type CVEReport struct {
	SchemaVersion int       `json:"SchemaVersion"`
	ArtifactName  string    `json:"ArtifactName"`
	ArtifactType  string    `json:"ArtifactType"`
	Metadata      Metadata  `json:"Metadata"`
	Results       []Results `json:"Results"`
}

type Rootfs struct {
	Type    string      `json:"type"`
	DiffIds interface{} `json:"diff_ids"`
}

type Config struct {
}

type ImageConfig struct {
	Architecture string    `json:"architecture"`
	Created      time.Time `json:"created"`
	Os           string    `json:"os"`
	Rootfs       Rootfs    `json:"rootfs"`
	Config       Config    `json:"config"`
}

type Metadata struct {
	ImageConfig ImageConfig `json:"ImageConfig"`
}
type Layer struct {
}

type Nvd struct {
	V2Vector string  `json:"V2Vector"`
	V3Vector string  `json:"V3Vector"`
	V2Score  float64 `json:"V2Score"`
	V3Score  float64 `json:"V3Score"`
}

type Redhat struct {
	V3Vector string  `json:"V3Vector"`
	V3Score  float64 `json:"V3Score"`
}

type Cvss struct {
	Nvd    Nvd    `json:"nvd"`
	Redhat Redhat `json:"redhat"`
}

type Vulnerabilities struct {
	VulnerabilityID  string    `json:"VulnerabilityID"`
	PkgName          string    `json:"PkgName"`
	InstalledVersion string    `json:"InstalledVersion"`
	FixedVersion     string    `json:"FixedVersion"`
	Layer            Layer     `json:"Layer"`
	SeveritySource   string    `json:"SeveritySource"`
	PrimaryURL       string    `json:"PrimaryURL"`
	Title            string    `json:"Title"`
	Description      string    `json:"Description"`
	Severity         string    `json:"Severity"`
	Cvss             Cvss      `json:"CVSS"`
	References       []string  `json:"References"`
	PublishedDate    time.Time `json:"PublishedDate"`
	LastModifiedDate time.Time `json:"LastModifiedDate"`
}

type Results struct {
	Target          string            `json:"Target"`
	Class           string            `json:"Class"`
	Type            string            `json:"Type"`
	Vulnerabilities []Vulnerabilities `json:"Vulnerabilities"`
}

type Projects []struct {
	ID                                        int                       `json:"id"`
	Description                               string                    `json:"description"`
	Name                                      string                    `json:"name"`
	NameWithNamespace                         string                    `json:"name_with_namespace"`
	Path                                      string                    `json:"path"`
	PathWithNamespace                         string                    `json:"path_with_namespace"`
	CreatedAt                                 time.Time                 `json:"created_at"`
	DefaultBranch                             string                    `json:"default_branch"`
	TagList                                   []interface{}             `json:"tag_list"`
	Topics                                    []interface{}             `json:"topics"`
	SSHURLToRepo                              string                    `json:"ssh_url_to_repo"`
	HTTPURLToRepo                             string                    `json:"http_url_to_repo"`
	WebURL                                    string                    `json:"web_url"`
	ReadmeURL                                 string                    `json:"readme_url"`
	AvatarURL                                 string                    `json:"avatar_url"`
	ForksCount                                int                       `json:"forks_count"`
	StarCount                                 int                       `json:"star_count"`
	LastActivityAt                            time.Time                 `json:"last_activity_at"`
	Namespace                                 Namespace                 `json:"namespace"`
	ContainerRegistryImagePrefix              string                    `json:"container_registry_image_prefix"`
	Links                                     LinksMergeRequest         `json:"_links"`
	PackagesEnabled                           bool                      `json:"packages_enabled"`
	EmptyRepo                                 bool                      `json:"empty_repo"`
	Archived                                  bool                      `json:"archived"`
	Visibility                                string                    `json:"visibility"`
	ResolveOutdatedDiffDiscussions            bool                      `json:"resolve_outdated_diff_discussions"`
	ContainerExpirationPolicy                 ContainerExpirationPolicy `json:"container_expiration_policy"`
	IssuesEnabled                             bool                      `json:"issues_enabled"`
	MergeRequestsEnabled                      bool                      `json:"merge_requests_enabled"`
	WikiEnabled                               bool                      `json:"wiki_enabled"`
	JobsEnabled                               bool                      `json:"jobs_enabled"`
	SnippetsEnabled                           bool                      `json:"snippets_enabled"`
	ContainerRegistryEnabled                  bool                      `json:"container_registry_enabled"`
	ServiceDeskEnabled                        bool                      `json:"service_desk_enabled"`
	ServiceDeskAddress                        string                    `json:"service_desk_address"`
	CanCreateMergeRequestIn                   bool                      `json:"can_create_merge_request_in"`
	IssuesAccessLevel                         string                    `json:"issues_access_level"`
	RepositoryAccessLevel                     string                    `json:"repository_access_level"`
	MergeRequestsAccessLevel                  string                    `json:"merge_requests_access_level"`
	ForkingAccessLevel                        string                    `json:"forking_access_level"`
	WikiAccessLevel                           string                    `json:"wiki_access_level"`
	BuildsAccessLevel                         string                    `json:"builds_access_level"`
	SnippetsAccessLevel                       string                    `json:"snippets_access_level"`
	PagesAccessLevel                          string                    `json:"pages_access_level"`
	OperationsAccessLevel                     string                    `json:"operations_access_level"`
	AnalyticsAccessLevel                      string                    `json:"analytics_access_level"`
	ContainerRegistryAccessLevel              string                    `json:"container_registry_access_level"`
	EmailsDisabled                            interface{}               `json:"emails_disabled"`
	SharedRunnersEnabled                      bool                      `json:"shared_runners_enabled"`
	LfsEnabled                                bool                      `json:"lfs_enabled"`
	CreatorID                                 int                       `json:"creator_id"`
	ImportStatus                              string                    `json:"import_status"`
	OpenIssuesCount                           int                       `json:"open_issues_count"`
	CiDefaultGitDepth                         int                       `json:"ci_default_git_depth"`
	CiForwardDeploymentEnabled                bool                      `json:"ci_forward_deployment_enabled"`
	CiJobTokenScopeEnabled                    bool                      `json:"ci_job_token_scope_enabled"`
	PublicJobs                                bool                      `json:"public_jobs"`
	BuildTimeout                              int                       `json:"build_timeout"`
	AutoCancelPendingPipelines                string                    `json:"auto_cancel_pending_pipelines"`
	BuildCoverageRegex                        interface{}               `json:"build_coverage_regex"`
	CiConfigPath                              string                    `json:"ci_config_path"`
	SharedWithGroups                          []interface{}             `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds          bool                      `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline               interface{}               `json:"allow_merge_on_skipped_pipeline"`
	RestrictUserDefinedVariables              bool                      `json:"restrict_user_defined_variables"`
	RequestAccessEnabled                      bool                      `json:"request_access_enabled"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool                      `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool                      `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestLinkEnabled           bool                      `json:"printing_merge_request_link_enabled"`
	MergeMethod                               string                    `json:"merge_method"`
	SquashOption                              string                    `json:"squash_option"`
	SuggestionCommitMessage                   interface{}               `json:"suggestion_commit_message"`
	MergeCommitTemplate                       interface{}               `json:"merge_commit_template"`
	SquashCommitTemplate                      interface{}               `json:"squash_commit_template"`
	AutoDevopsEnabled                         bool                      `json:"auto_devops_enabled"`
	AutoDevopsDeployStrategy                  string                    `json:"auto_devops_deploy_strategy"`
	AutocloseReferencedIssues                 bool                      `json:"autoclose_referenced_issues"`
	KeepLatestArtifact                        bool                      `json:"keep_latest_artifact"`
	ApprovalsBeforeMerge                      int                       `json:"approvals_before_merge"`
	Mirror                                    bool                      `json:"mirror"`
	ExternalAuthorizationClassificationLabel  string                    `json:"external_authorization_classification_label"`
	MarkedForDeletionAt                       interface{}               `json:"marked_for_deletion_at"`
	MarkedForDeletionOn                       interface{}               `json:"marked_for_deletion_on"`
	RequirementsEnabled                       bool                      `json:"requirements_enabled"`
	SecurityAndComplianceEnabled              bool                      `json:"security_and_compliance_enabled"`
	ComplianceFrameworks                      []interface{}             `json:"compliance_frameworks"`
	IssuesTemplate                            interface{}               `json:"issues_template"`
	MergeRequestsTemplate                     interface{}               `json:"merge_requests_template"`
	MergePipelinesEnabled                     bool                      `json:"merge_pipelines_enabled"`
	MergeTrainsEnabled                        bool                      `json:"merge_trains_enabled"`
}

type Namespace struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Kind      string `json:"kind"`
	FullPath  string `json:"full_path"`
	ParentID  int    `json:"parent_id"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

type LinksMergeRequest struct {
	Self          string `json:"self"`
	Issues        string `json:"issues"`
	MergeRequests string `json:"merge_requests"`
	RepoBranches  string `json:"repo_branches"`
	Labels        string `json:"labels"`
	Events        string `json:"events"`
	Members       string `json:"members"`
}

type ContainerExpirationPolicy struct {
	Cadence       string      `json:"cadence"`
	Enabled       bool        `json:"enabled"`
	KeepN         int         `json:"keep_n"`
	OlderThan     string      `json:"older_than"`
	NameRegex     string      `json:"name_regex"`
	NameRegexKeep interface{} `json:"name_regex_keep"`
	NextRunAt     time.Time   `json:"next_run_at"`
}

type Issues []struct {
	ID                   int                  `json:"id"`
	Iid                  int                  `json:"iid"`
	ProjectID            int                  `json:"project_id"`
	Title                string               `json:"title"`
	Description          string               `json:"description"`
	State                string               `json:"state"`
	CreatedAt            time.Time            `json:"created_at"`
	UpdatedAt            time.Time            `json:"updated_at"`
	ClosedAt             interface{}          `json:"closed_at"`
	ClosedBy             interface{}          `json:"closed_by"`
	Labels               []string             `json:"labels"`
	Milestone            interface{}          `json:"milestone"`
	Assignees            []interface{}        `json:"assignees"`
	Author               Author               `json:"author"`
	Type                 string               `json:"type"`
	Assignee             interface{}          `json:"assignee"`
	UserNotesCount       int                  `json:"user_notes_count"`
	MergeRequestsCount   int                  `json:"merge_requests_count"`
	Upvotes              int                  `json:"upvotes"`
	Downvotes            int                  `json:"downvotes"`
	DueDate              interface{}          `json:"due_date"`
	Confidential         bool                 `json:"confidential"`
	DiscussionLocked     interface{}          `json:"discussion_locked"`
	IssueType            string               `json:"issue_type"`
	WebURL               string               `json:"web_url"`
	TimeStats            TimeStats            `json:"time_stats"`
	TaskCompletionStatus TaskCompletionStatus `json:"task_completion_status"`
	Weight               interface{}          `json:"weight"`
	BlockingIssuesCount  int                  `json:"blocking_issues_count"`
	HasTasks             bool                 `json:"has_tasks"`
	Links                Links                `json:"_links"`
	References           References           `json:"references"`
	MovedToID            interface{}          `json:"moved_to_id"`
	ServiceDeskReplyTo   interface{}          `json:"service_desk_reply_to"`
	EpicIid              interface{}          `json:"epic_iid"`
	Epic                 interface{}          `json:"epic"`
}

type Author struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

type TimeStats struct {
	TimeEstimate        int         `json:"time_estimate"`
	TotalTimeSpent      int         `json:"total_time_spent"`
	HumanTimeEstimate   interface{} `json:"human_time_estimate"`
	HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
}

type TaskCompletionStatus struct {
	Count          int `json:"count"`
	CompletedCount int `json:"completed_count"`
}

type Links struct {
	Self       string `json:"self"`
	Notes      string `json:"notes"`
	AwardEmoji string `json:"award_emoji"`
	Project    string `json:"project"`
}

type References struct {
	Short    string `json:"short"`
	Relative string `json:"relative"`
	Full     string `json:"full"`
}

type Project struct {
	ID                int           `json:"id"`
	Description       string        `json:"description"`
	Name              string        `json:"name"`
	NameWithNamespace string        `json:"name_with_namespace"`
	Path              string        `json:"path"`
	PathWithNamespace string        `json:"path_with_namespace"`
	CreatedAt         time.Time     `json:"created_at"`
	DefaultBranch     string        `json:"default_branch"`
	TagList           []interface{} `json:"tag_list"`
	Topics            []interface{} `json:"topics"`
	SSHURLToRepo      string        `json:"ssh_url_to_repo"`
	HTTPURLToRepo     string        `json:"http_url_to_repo"`
	WebURL            string        `json:"web_url"`
	ReadmeURL         string        `json:"readme_url"`
	AvatarURL         string        `json:"avatar_url"`
	ForksCount        int           `json:"forks_count"`
	StarCount         int           `json:"star_count"`
	LastActivityAt    time.Time     `json:"last_activity_at"`
	Namespace         struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Path      string `json:"path"`
		Kind      string `json:"kind"`
		FullPath  string `json:"full_path"`
		ParentID  int    `json:"parent_id"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"namespace"`
	ContainerRegistryImagePrefix string `json:"container_registry_image_prefix"`
	Links                        struct {
		Self          string `json:"self"`
		Issues        string `json:"issues"`
		MergeRequests string `json:"merge_requests"`
		RepoBranches  string `json:"repo_branches"`
		Labels        string `json:"labels"`
		Events        string `json:"events"`
		Members       string `json:"members"`
	} `json:"_links"`
	PackagesEnabled                bool   `json:"packages_enabled"`
	EmptyRepo                      bool   `json:"empty_repo"`
	Archived                       bool   `json:"archived"`
	Visibility                     string `json:"visibility"`
	ResolveOutdatedDiffDiscussions bool   `json:"resolve_outdated_diff_discussions"`
	ContainerExpirationPolicy      struct {
		Cadence       string    `json:"cadence"`
		Enabled       bool      `json:"enabled"`
		KeepN         int       `json:"keep_n"`
		OlderThan     string    `json:"older_than"`
		NameRegex     string    `json:"name_regex"`
		NameRegexKeep string    `json:"name_regex_keep"`
		NextRunAt     time.Time `json:"next_run_at"`
	} `json:"container_expiration_policy"`
	IssuesEnabled                             bool          `json:"issues_enabled"`
	MergeRequestsEnabled                      bool          `json:"merge_requests_enabled"`
	WikiEnabled                               bool          `json:"wiki_enabled"`
	JobsEnabled                               bool          `json:"jobs_enabled"`
	SnippetsEnabled                           bool          `json:"snippets_enabled"`
	ContainerRegistryEnabled                  bool          `json:"container_registry_enabled"`
	ServiceDeskEnabled                        bool          `json:"service_desk_enabled"`
	ServiceDeskAddress                        string        `json:"service_desk_address"`
	CanCreateMergeRequestIn                   bool          `json:"can_create_merge_request_in"`
	IssuesAccessLevel                         string        `json:"issues_access_level"`
	RepositoryAccessLevel                     string        `json:"repository_access_level"`
	MergeRequestsAccessLevel                  string        `json:"merge_requests_access_level"`
	ForkingAccessLevel                        string        `json:"forking_access_level"`
	WikiAccessLevel                           string        `json:"wiki_access_level"`
	BuildsAccessLevel                         string        `json:"builds_access_level"`
	SnippetsAccessLevel                       string        `json:"snippets_access_level"`
	PagesAccessLevel                          string        `json:"pages_access_level"`
	OperationsAccessLevel                     string        `json:"operations_access_level"`
	AnalyticsAccessLevel                      string        `json:"analytics_access_level"`
	ContainerRegistryAccessLevel              string        `json:"container_registry_access_level"`
	EmailsDisabled                            interface{}   `json:"emails_disabled"`
	SharedRunnersEnabled                      bool          `json:"shared_runners_enabled"`
	LfsEnabled                                bool          `json:"lfs_enabled"`
	CreatorID                                 int           `json:"creator_id"`
	ImportStatus                              string        `json:"import_status"`
	ImportError                               interface{}   `json:"import_error"`
	OpenIssuesCount                           int           `json:"open_issues_count"`
	RunnersToken                              string        `json:"runners_token"`
	CiDefaultGitDepth                         int           `json:"ci_default_git_depth"`
	CiForwardDeploymentEnabled                bool          `json:"ci_forward_deployment_enabled"`
	CiJobTokenScopeEnabled                    bool          `json:"ci_job_token_scope_enabled"`
	PublicJobs                                bool          `json:"public_jobs"`
	BuildGitStrategy                          string        `json:"build_git_strategy"`
	BuildTimeout                              int           `json:"build_timeout"`
	AutoCancelPendingPipelines                string        `json:"auto_cancel_pending_pipelines"`
	BuildCoverageRegex                        interface{}   `json:"build_coverage_regex"`
	CiConfigPath                              string        `json:"ci_config_path"`
	SharedWithGroups                          []interface{} `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds          bool          `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline               bool          `json:"allow_merge_on_skipped_pipeline"`
	RestrictUserDefinedVariables              bool          `json:"restrict_user_defined_variables"`
	RequestAccessEnabled                      bool          `json:"request_access_enabled"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool          `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool          `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestLinkEnabled           bool          `json:"printing_merge_request_link_enabled"`
	MergeMethod                               string        `json:"merge_method"`
	SquashOption                              string        `json:"squash_option"`
	SuggestionCommitMessage                   string        `json:"suggestion_commit_message"`
	MergeCommitTemplate                       interface{}   `json:"merge_commit_template"`
	SquashCommitTemplate                      interface{}   `json:"squash_commit_template"`
	AutoDevopsEnabled                         bool          `json:"auto_devops_enabled"`
	AutoDevopsDeployStrategy                  string        `json:"auto_devops_deploy_strategy"`
	AutocloseReferencedIssues                 bool          `json:"autoclose_referenced_issues"`
	KeepLatestArtifact                        bool          `json:"keep_latest_artifact"`
	ApprovalsBeforeMerge                      int           `json:"approvals_before_merge"`
	Mirror                                    bool          `json:"mirror"`
	ExternalAuthorizationClassificationLabel  string        `json:"external_authorization_classification_label"`
	MarkedForDeletionAt                       interface{}   `json:"marked_for_deletion_at"`
	MarkedForDeletionOn                       interface{}   `json:"marked_for_deletion_on"`
	RequirementsEnabled                       bool          `json:"requirements_enabled"`
	SecurityAndComplianceEnabled              bool          `json:"security_and_compliance_enabled"`
	ComplianceFrameworks                      []interface{} `json:"compliance_frameworks"`
	IssuesTemplate                            interface{}   `json:"issues_template"`
	MergeRequestsTemplate                     string        `json:"merge_requests_template"`
	MergePipelinesEnabled                     bool          `json:"merge_pipelines_enabled"`
	MergeTrainsEnabled                        bool          `json:"merge_trains_enabled"`
	Permissions                               struct {
		ProjectAccess interface{} `json:"project_access"`
		GroupAccess   struct {
			AccessLevel       int `json:"access_level"`
			NotificationLevel int `json:"notification_level"`
		} `json:"group_access"`
	} `json:"permissions"`
}
