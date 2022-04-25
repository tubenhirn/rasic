package types

import "time"

// types used by gitlab.com api

type GitlabProjects []struct {
	ID                                        int                       `json:"id,omitempty"`
	Description                               string                    `json:"description,omitempty"`
	Name                                      string                    `json:"name,omitempty"`
	NameWithNamespace                         string                    `json:"name_with_namespace,omitempty"`
	Path                                      string                    `json:"path,omitempty"`
	PathWithNamespace                         string                    `json:"path_with_namespace,omitempty"`
	CreatedAt                                 time.Time                 `json:"created_at,omitempty"`
	DefaultBranch                             string                    `json:"default_branch,omitempty"`
	TagList                                   []interface{}             `json:"tag_list,omitempty"`
	Topics                                    []interface{}             `json:"topics,omitempty"`
	SSHURLToRepo                              string                    `json:"ssh_url_to_repo,omitempty"`
	HTTPURLToRepo                             string                    `json:"http_url_to_repo,omitempty"`
	WebURL                                    string                    `json:"web_url,omitempty"`
	ReadmeURL                                 string                    `json:"readme_url,omitempty"`
	AvatarURL                                 string                    `json:"avatar_url,omitempty"`
	ForksCount                                int                       `json:"forks_count,omitempty"`
	StarCount                                 int                       `json:"star_count,omitempty"`
	LastActivityAt                            time.Time                 `json:"last_activity_at,omitempty"`
	Namespace                                 Namespace                 `json:"namespace,omitempty"`
	ContainerRegistryImagePrefix              string                    `json:"container_registry_image_prefix,omitempty"`
	Links                                     LinksMergeRequest         `json:"_links,omitempty"`
	PackagesEnabled                           bool                      `json:"packages_enabled,omitempty"`
	EmptyRepo                                 bool                      `json:"empty_repo,omitempty"`
	Archived                                  bool                      `json:"archived,omitempty"`
	Visibility                                string                    `json:"visibility,omitempty"`
	ResolveOutdatedDiffDiscussions            bool                      `json:"resolve_outdated_diff_discussions,omitempty"`
	ContainerExpirationPolicy                 ContainerExpirationPolicy `json:"container_expiration_policy,omitempty"`
	IssuesEnabled                             bool                      `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled                      bool                      `json:"merge_requests_enabled,omitempty"`
	WikiEnabled                               bool                      `json:"wiki_enabled,omitempty"`
	JobsEnabled                               bool                      `json:"jobs_enabled,omitempty"`
	SnippetsEnabled                           bool                      `json:"snippets_enabled,omitempty"`
	ContainerRegistryEnabled                  bool                      `json:"container_registry_enabled,omitempty"`
	ServiceDeskEnabled                        bool                      `json:"service_desk_enabled,omitempty"`
	ServiceDeskAddress                        string                    `json:"service_desk_address,omitempty"`
	CanCreateMergeRequestIn                   bool                      `json:"can_create_merge_request_in,omitempty"`
	IssuesAccessLevel                         string                    `json:"issues_access_level,omitempty"`
	RepositoryAccessLevel                     string                    `json:"repository_access_level,omitempty"`
	MergeRequestsAccessLevel                  string                    `json:"merge_requests_access_level,omitempty"`
	ForkingAccessLevel                        string                    `json:"forking_access_level,omitempty"`
	WikiAccessLevel                           string                    `json:"wiki_access_level,omitempty"`
	BuildsAccessLevel                         string                    `json:"builds_access_level,omitempty"`
	SnippetsAccessLevel                       string                    `json:"snippets_access_level,omitempty"`
	PagesAccessLevel                          string                    `json:"pages_access_level,omitempty"`
	OperationsAccessLevel                     string                    `json:"operations_access_level,omitempty"`
	AnalyticsAccessLevel                      string                    `json:"analytics_access_level,omitempty"`
	ContainerRegistryAccessLevel              string                    `json:"container_registry_access_level,omitempty"`
	EmailsDisabled                            interface{}               `json:"emails_disabled,omitempty"`
	SharedRunnersEnabled                      bool                      `json:"shared_runners_enabled,omitempty"`
	LfsEnabled                                bool                      `json:"lfs_enabled,omitempty"`
	CreatorID                                 int                       `json:"creator_id,omitempty"`
	ImportStatus                              string                    `json:"import_status,omitempty"`
	OpenIssuesCount                           int                       `json:"open_issues_count,omitempty"`
	CiDefaultGitDepth                         int                       `json:"ci_default_git_depth,omitempty"`
	CiForwardDeploymentEnabled                bool                      `json:"ci_forward_deployment_enabled,omitempty"`
	CiJobTokenScopeEnabled                    bool                      `json:"ci_job_token_scope_enabled,omitempty"`
	PublicJobs                                bool                      `json:"public_jobs,omitempty"`
	BuildTimeout                              int                       `json:"build_timeout,omitempty"`
	AutoCancelPendingPipelines                string                    `json:"auto_cancel_pending_pipelines,omitempty"`
	BuildCoverageRegex                        interface{}               `json:"build_coverage_regex,omitempty"`
	CiConfigPath                              string                    `json:"ci_config_path,omitempty"`
	SharedWithGroups                          []interface{}             `json:"shared_with_groups,omitempty"`
	OnlyAllowMergeIfPipelineSucceeds          bool                      `json:"only_allow_merge_if_pipeline_succeeds,omitempty"`
	AllowMergeOnSkippedPipeline               interface{}               `json:"allow_merge_on_skipped_pipeline,omitempty"`
	RestrictUserDefinedVariables              bool                      `json:"restrict_user_defined_variables,omitempty"`
	RequestAccessEnabled                      bool                      `json:"request_access_enabled,omitempty"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool                      `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	RemoveSourceBranchAfterMerge              bool                      `json:"remove_source_branch_after_merge,omitempty"`
	PrintingMergeRequestLinkEnabled           bool                      `json:"printing_merge_request_link_enabled,omitempty"`
	MergeMethod                               string                    `json:"merge_method,omitempty"`
	SquashOption                              string                    `json:"squash_option,omitempty"`
	SuggestionCommitMessage                   interface{}               `json:"suggestion_commit_message,omitempty"`
	MergeCommitTemplate                       interface{}               `json:"merge_commit_template,omitempty"`
	SquashCommitTemplate                      interface{}               `json:"squash_commit_template,omitempty"`
	AutoDevopsEnabled                         bool                      `json:"auto_devops_enabled,omitempty"`
	AutoDevopsDeployStrategy                  string                    `json:"auto_devops_deploy_strategy,omitempty"`
	AutocloseReferencedIssues                 bool                      `json:"autoclose_referenced_issues,omitempty"`
	KeepLatestArtifact                        bool                      `json:"keep_latest_artifact,omitempty"`
	ApprovalsBeforeMerge                      int                       `json:"approvals_before_merge,omitempty"`
	Mirror                                    bool                      `json:"mirror,omitempty"`
	ExternalAuthorizationClassificationLabel  string                    `json:"external_authorization_classification_label,omitempty"`
	MarkedForDeletionAt                       interface{}               `json:"marked_for_deletion_at,omitempty"`
	MarkedForDeletionOn                       interface{}               `json:"marked_for_deletion_on,omitempty"`
	RequirementsEnabled                       bool                      `json:"requirements_enabled,omitempty"`
	SecurityAndComplianceEnabled              bool                      `json:"security_and_compliance_enabled,omitempty"`
	ComplianceFrameworks                      []interface{}             `json:"compliance_frameworks,omitempty"`
	IssuesTemplate                            interface{}               `json:"issues_template,omitempty"`
	MergeRequestsTemplate                     interface{}               `json:"merge_requests_template,omitempty"`
	MergePipelinesEnabled                     bool                      `json:"merge_pipelines_enabled,omitempty"`
	MergeTrainsEnabled                        bool                      `json:"merge_trains_enabled,omitempty"`
}

type Namespace struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Kind      string `json:"kind,omitempty"`
	FullPath  string `json:"full_path,omitempty"`
	ParentID  int    `json:"parent_id,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	WebURL    string `json:"web_url,omitempty"`
}

type LinksMergeRequest struct {
	Self          string `json:"self,omitempty"`
	Issues        string `json:"issues,omitempty"`
	MergeRequests string `json:"merge_requests,omitempty"`
	RepoBranches  string `json:"repo_branches,omitempty"`
	Labels        string `json:"labels,omitempty"`
	Events        string `json:"events,omitempty"`
	Members       string `json:"members,omitempty"`
}

type ContainerExpirationPolicy struct {
	Cadence       string      `json:"cadence,omitempty"`
	Enabled       bool        `json:"enabled,omitempty"`
	KeepN         int         `json:"keep_n,omitempty"`
	OlderThan     string      `json:"older_than,omitempty"`
	NameRegex     string      `json:"name_regex,omitempty"`
	NameRegexKeep interface{} `json:"name_regex_keep,omitempty"`
	NextRunAt     time.Time   `json:"next_run_at,omitempty"`
}

type GitlabIssues []struct {
	ID                   int                  `json:"id,omitempty"`
	Iid                  int                  `json:"iid,omitempty"`
	ProjectID            int                  `json:"project_id,omitempty"`
	Title                string               `json:"title,omitempty"`
	Description          string               `json:"description,omitempty"`
	State                string               `json:"state,omitempty"`
	CreatedAt            time.Time            `json:"created_at,omitempty"`
	UpdatedAt            time.Time            `json:"updated_at,omitempty"`
	ClosedAt             interface{}          `json:"closed_at,omitempty"`
	ClosedBy             interface{}          `json:"closed_by,omitempty"`
	Labels               []string             `json:"labels,omitempty"`
	Milestone            interface{}          `json:"milestone,omitempty"`
	Assignees            []interface{}        `json:"assignees,omitempty"`
	Author               Author               `json:"author,omitempty"`
	Type                 string               `json:"type,omitempty"`
	Assignee             interface{}          `json:"assignee,omitempty"`
	UserNotesCount       int                  `json:"user_notes_count,omitempty"`
	MergeRequestsCount   int                  `json:"merge_requests_count,omitempty"`
	Upvotes              int                  `json:"upvotes,omitempty"`
	Downvotes            int                  `json:"downvotes,omitempty"`
	DueDate              interface{}          `json:"due_date,omitempty"`
	Confidential         bool                 `json:"confidential,omitempty"`
	DiscussionLocked     interface{}          `json:"discussion_locked,omitempty"`
	IssueType            string               `json:"issue_type,omitempty"`
	WebURL               string               `json:"web_url,omitempty"`
	TimeStats            TimeStats            `json:"time_stats,omitempty"`
	TaskCompletionStatus TaskCompletionStatus `json:"task_completion_status,omitempty"`
	Weight               interface{}          `json:"weight,omitempty"`
	BlockingIssuesCount  int                  `json:"blocking_issues_count,omitempty"`
	HasTasks             bool                 `json:"has_tasks,omitempty"`
	Links                Links                `json:"_links,omitempty"`
	References           References           `json:"references,omitempty"`
	MovedToID            interface{}          `json:"moved_to_id,omitempty"`
	ServiceDeskReplyTo   interface{}          `json:"service_desk_reply_to,omitempty"`
	EpicIid              interface{}          `json:"epic_iid,omitempty"`
	Epic                 interface{}          `json:"epic,omitempty"`
}

type Author struct {
	ID        int    `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Name      string `json:"name,omitempty"`
	State     string `json:"state,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	WebURL    string `json:"web_url,omitempty"`
}

type TimeStats struct {
	TimeEstimate        int         `json:"time_estimate,omitempty"`
	TotalTimeSpent      int         `json:"total_time_spent,omitempty"`
	HumanTimeEstimate   interface{} `json:"human_time_estimate,omitempty"`
	HumanTotalTimeSpent interface{} `json:"human_total_time_spent,omitempty"`
}

type TaskCompletionStatus struct {
	Count          int `json:"count,omitempty"`
	CompletedCount int `json:"completed_count,omitempty"`
}

type Links struct {
	Self       string `json:"self,omitempty"`
	Notes      string `json:"notes,omitempty"`
	AwardEmoji string `json:"award_emoji,omitempty"`
	Project    string `json:"project,omitempty"`
}

type References struct {
	Short    string `json:"short,omitempty"`
	Relative string `json:"relative,omitempty"`
	Full     string `json:"full,omitempty"`
}

type GitlabProject struct {
	ID                int           `json:"id,omitempty"`
	Description       string        `json:"description,omitempty"`
	Name              string        `json:"name,omitempty"`
	NameWithNamespace string        `json:"name_with_namespace,omitempty"`
	Path              string        `json:"path,omitempty"`
	PathWithNamespace string        `json:"path_with_namespace,omitempty"`
	CreatedAt         time.Time     `json:"created_at,omitempty"`
	DefaultBranch     string        `json:"default_branch,omitempty"`
	TagList           []interface{} `json:"tag_list,omitempty"`
	Topics            []interface{} `json:"topics,omitempty"`
	SSHURLToRepo      string        `json:"ssh_url_to_repo,omitempty"`
	HTTPURLToRepo     string        `json:"http_url_to_repo,omitempty"`
	WebURL            string        `json:"web_url,omitempty"`
	ReadmeURL         string        `json:"readme_url,omitempty"`
	AvatarURL         string        `json:"avatar_url,omitempty"`
	ForksCount        int           `json:"forks_count,omitempty"`
	StarCount         int           `json:"star_count,omitempty"`
	LastActivityAt    time.Time     `json:"last_activity_at,omitempty"`
	Namespace         struct {
		ID        int    `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Kind      string `json:"kind,omitempty"`
		FullPath  string `json:"full_path,omitempty"`
		ParentID  int    `json:"parent_id,omitempty"`
		AvatarURL string `json:"avatar_url,omitempty"`
		WebURL    string `json:"web_url,omitempty"`
	} `json:"namespace,omitempty"`
	ContainerRegistryImagePrefix string `json:"container_registry_image_prefix,omitempty"`
	Links                        struct {
		Self          string `json:"self,omitempty"`
		Issues        string `json:"issues,omitempty"`
		MergeRequests string `json:"merge_requests,omitempty"`
		RepoBranches  string `json:"repo_branches,omitempty"`
		Labels        string `json:"labels,omitempty"`
		Events        string `json:"events,omitempty"`
		Members       string `json:"members,omitempty"`
	} `json:"_links,omitempty"`
	PackagesEnabled                bool   `json:"packages_enabled,omitempty"`
	EmptyRepo                      bool   `json:"empty_repo,omitempty"`
	Archived                       bool   `json:"archived,omitempty"`
	Visibility                     string `json:"visibility,omitempty"`
	ResolveOutdatedDiffDiscussions bool   `json:"resolve_outdated_diff_discussions,omitempty"`
	ContainerExpirationPolicy      struct {
		Cadence       string    `json:"cadence,omitempty"`
		Enabled       bool      `json:"enabled,omitempty"`
		KeepN         int       `json:"keep_n,omitempty"`
		OlderThan     string    `json:"older_than,omitempty"`
		NameRegex     string    `json:"name_regex,omitempty"`
		NameRegexKeep string    `json:"name_regex_keep,omitempty"`
		NextRunAt     time.Time `json:"next_run_at,omitempty"`
	} `json:"container_expiration_policy,omitempty"`
	IssuesEnabled                             bool          `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled                      bool          `json:"merge_requests_enabled,omitempty"`
	WikiEnabled                               bool          `json:"wiki_enabled,omitempty"`
	JobsEnabled                               bool          `json:"jobs_enabled,omitempty"`
	SnippetsEnabled                           bool          `json:"snippets_enabled,omitempty"`
	ContainerRegistryEnabled                  bool          `json:"container_registry_enabled,omitempty"`
	ServiceDeskEnabled                        bool          `json:"service_desk_enabled,omitempty"`
	ServiceDeskAddress                        string        `json:"service_desk_address,omitempty"`
	CanCreateMergeRequestIn                   bool          `json:"can_create_merge_request_in,omitempty"`
	IssuesAccessLevel                         string        `json:"issues_access_level,omitempty"`
	RepositoryAccessLevel                     string        `json:"repository_access_level,omitempty"`
	MergeRequestsAccessLevel                  string        `json:"merge_requests_access_level,omitempty"`
	ForkingAccessLevel                        string        `json:"forking_access_level,omitempty"`
	WikiAccessLevel                           string        `json:"wiki_access_level,omitempty"`
	BuildsAccessLevel                         string        `json:"builds_access_level,omitempty"`
	SnippetsAccessLevel                       string        `json:"snippets_access_level,omitempty"`
	PagesAccessLevel                          string        `json:"pages_access_level,omitempty"`
	OperationsAccessLevel                     string        `json:"operations_access_level,omitempty"`
	AnalyticsAccessLevel                      string        `json:"analytics_access_level,omitempty"`
	ContainerRegistryAccessLevel              string        `json:"container_registry_access_level,omitempty"`
	EmailsDisabled                            interface{}   `json:"emails_disabled,omitempty"`
	SharedRunnersEnabled                      bool          `json:"shared_runners_enabled,omitempty"`
	LfsEnabled                                bool          `json:"lfs_enabled,omitempty"`
	CreatorID                                 int           `json:"creator_id,omitempty"`
	ImportStatus                              string        `json:"import_status,omitempty"`
	ImportError                               interface{}   `json:"import_error,omitempty"`
	OpenIssuesCount                           int           `json:"open_issues_count,omitempty"`
	RunnersToken                              string        `json:"runners_token,omitempty"`
	CiDefaultGitDepth                         int           `json:"ci_default_git_depth,omitempty"`
	CiForwardDeploymentEnabled                bool          `json:"ci_forward_deployment_enabled,omitempty"`
	CiJobTokenScopeEnabled                    bool          `json:"ci_job_token_scope_enabled,omitempty"`
	PublicJobs                                bool          `json:"public_jobs,omitempty"`
	BuildGitStrategy                          string        `json:"build_git_strategy,omitempty"`
	BuildTimeout                              int           `json:"build_timeout,omitempty"`
	AutoCancelPendingPipelines                string        `json:"auto_cancel_pending_pipelines,omitempty"`
	BuildCoverageRegex                        interface{}   `json:"build_coverage_regex,omitempty"`
	CiConfigPath                              string        `json:"ci_config_path,omitempty"`
	SharedWithGroups                          []interface{} `json:"shared_with_groups,omitempty"`
	OnlyAllowMergeIfPipelineSucceeds          bool          `json:"only_allow_merge_if_pipeline_succeeds,omitempty"`
	AllowMergeOnSkippedPipeline               bool          `json:"allow_merge_on_skipped_pipeline,omitempty"`
	RestrictUserDefinedVariables              bool          `json:"restrict_user_defined_variables,omitempty"`
	RequestAccessEnabled                      bool          `json:"request_access_enabled,omitempty"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool          `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	RemoveSourceBranchAfterMerge              bool          `json:"remove_source_branch_after_merge,omitempty"`
	PrintingMergeRequestLinkEnabled           bool          `json:"printing_merge_request_link_enabled,omitempty"`
	MergeMethod                               string        `json:"merge_method,omitempty"`
	SquashOption                              string        `json:"squash_option,omitempty"`
	SuggestionCommitMessage                   string        `json:"suggestion_commit_message,omitempty"`
	MergeCommitTemplate                       interface{}   `json:"merge_commit_template,omitempty"`
	SquashCommitTemplate                      interface{}   `json:"squash_commit_template,omitempty"`
	AutoDevopsEnabled                         bool          `json:"auto_devops_enabled,omitempty"`
	AutoDevopsDeployStrategy                  string        `json:"auto_devops_deploy_strategy,omitempty"`
	AutocloseReferencedIssues                 bool          `json:"autoclose_referenced_issues,omitempty"`
	KeepLatestArtifact                        bool          `json:"keep_latest_artifact,omitempty"`
	ApprovalsBeforeMerge                      int           `json:"approvals_before_merge,omitempty"`
	Mirror                                    bool          `json:"mirror,omitempty"`
	ExternalAuthorizationClassificationLabel  string        `json:"external_authorization_classification_label,omitempty"`
	MarkedForDeletionAt                       interface{}   `json:"marked_for_deletion_at,omitempty"`
	MarkedForDeletionOn                       interface{}   `json:"marked_for_deletion_on,omitempty"`
	RequirementsEnabled                       bool          `json:"requirements_enabled,omitempty"`
	SecurityAndComplianceEnabled              bool          `json:"security_and_compliance_enabled,omitempty"`
	ComplianceFrameworks                      []interface{} `json:"compliance_frameworks,omitempty"`
	IssuesTemplate                            interface{}   `json:"issues_template,omitempty"`
	MergeRequestsTemplate                     string        `json:"merge_requests_template,omitempty"`
	MergePipelinesEnabled                     bool          `json:"merge_pipelines_enabled,omitempty"`
	MergeTrainsEnabled                        bool          `json:"merge_trains_enabled,omitempty"`
	Permissions                               struct {
		ProjectAccess interface{} `json:"project_access,omitempty"`
		GroupAccess   struct {
			AccessLevel       int `json:"access_level,omitempty"`
			NotificationLevel int `json:"notification_level,omitempty"`
		} `json:"group_access,omitempty"`
	} `json:"permissions,omitempty"`
}

type GitlabIssue struct {
	ProjectID          int           `json:"project_id,omitempty"`
	ID                 int           `json:"id,omitempty"`
	CreatedAt          time.Time     `json:"created_at,omitempty"`
	Iid                int           `json:"iid,omitempty"`
	Title              string        `json:"title,omitempty"`
	State              string        `json:"state,omitempty"`
	Assignees          []interface{} `json:"assignees,omitempty"`
	Assignee           interface{}   `json:"assignee,omitempty"`
	Type               string        `json:"type,omitempty"`
	Labels             []string      `json:"labels,omitempty"`
	Upvotes            int           `json:"upvotes,omitempty"`
	Downvotes          int           `json:"downvotes,omitempty"`
	MergeRequestsCount int           `json:"merge_requests_count,omitempty"`
	Author             struct {
		Name      string      `json:"name,omitempty"`
		AvatarURL interface{} `json:"avatar_url,omitempty"`
		State     string      `json:"state,omitempty"`
		WebURL    string      `json:"web_url,omitempty"`
		ID        int         `json:"id,omitempty"`
		Username  string      `json:"username,omitempty"`
	} `json:"author,omitempty"`
	Description    string      `json:"description,omitempty"`
	UpdatedAt      time.Time   `json:"updated_at,omitempty"`
	ClosedAt       interface{} `json:"closed_at,omitempty"`
	ClosedBy       interface{} `json:"closed_by,omitempty"`
	Milestone      interface{} `json:"milestone,omitempty"`
	Subscribed     bool        `json:"subscribed,omitempty"`
	UserNotesCount int         `json:"user_notes_count,omitempty"`
	DueDate        interface{} `json:"due_date,omitempty"`
	WebURL         string      `json:"web_url,omitempty"`
	References     struct {
		Short    string `json:"short,omitempty"`
		Relative string `json:"relative,omitempty"`
		Full     string `json:"full,omitempty"`
	} `json:"references,omitempty"`
	TimeStats struct {
		TimeEstimate        int         `json:"time_estimate,omitempty"`
		TotalTimeSpent      int         `json:"total_time_spent,omitempty"`
		HumanTimeEstimate   interface{} `json:"human_time_estimate,omitempty"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent,omitempty"`
	} `json:"time_stats,omitempty"`
	Confidential     bool   `json:"confidential,omitempty"`
	DiscussionLocked bool   `json:"discussion_locked,omitempty"`
	IssueType        string `json:"issue_type,omitempty"`
	Severity         string `json:"severity,omitempty"`
	Links            struct {
		Self       string `json:"self,omitempty"`
		Notes      string `json:"notes,omitempty"`
		AwardEmoji string `json:"award_emoji,omitempty"`
		Project    string `json:"project,omitempty"`
	} `json:"_links,omitempty"`
	TaskCompletionStatus struct {
		Count          int `json:"count,omitempty"`
		CompletedCount int `json:"completed_count,omitempty"`
	} `json:"task_completion_status,omitempty"`
}

type GitlabRepositories []struct {
	CleanupPolicyStartedAt string `json:"cleanup_policy_started_at,omitempty"`
	CreatedAt              string `json:"created_at,omitempty"`
	ID                     int    `json:"id,omitempty"`
	Location               string `json:"location,omitempty"`
	Name                   string `json:"name,omitempty"`
	Path                   string `json:"path,omitempty"`
	ProjectID              int    `json:"project_id,omitempty"`
}

type GitlabRepository struct {
	CleanupPolicyStartedAt string `json:"cleanup_policy_started_at,omitempty"`
	CreatedAt              string `json:"created_at,omitempty"`
	ID                     int    `json:"id,omitempty"`
	Location               string `json:"location,omitempty"`
	Name                   string `json:"name,omitempty"`
	Path                   string `json:"path,omitempty"`
	ProjectID              int    `json:"project_id,omitempty"`
	Size                   int    `json:"size,omitempty"`
	Tags                   []struct {
		Location string `json:"location,omitempty"`
		Name     string `json:"name,omitempty"`
		Path     string `json:"path,omitempty"`
	} `json:"tags,omitempty"`
	TagsCount int `json:"tags_count,omitempty"`
}

type GitlabLabel struct {
	ClosedIssuesCount      int64  `json:"closed_issues_count,omitempty"`
	Color                  string `json:"color,omitempty"`
	Description            string `json:"description,omitempty"`
	DescriptionHTML        string `json:"description_html,omitempty"`
	ID                     int64  `json:"id,omitempty"`
	IsProjectLabel         bool   `json:"is_project_label,omitempty"`
	Name                   string `json:"name,omitempty"`
	OpenIssuesCount        int64  `json:"open_issues_count,omitempty"`
	OpenMergeRequestsCount int64  `json:"open_merge_requests_count,omitempty"`
	Priority               int64  `json:"priority,omitempty"`
	Subscribed             bool   `json:"subscribed,omitempty"`
	TextColor              string `json:"text_color,omitempty"`
}

type GitlabLabel struct {
	ClosedIssuesCount      int64  `json:"closed_issues_count"`
	Color                  string `json:"color"`
	Description            string `json:"description"`
	DescriptionHTML        string `json:"description_html"`
	ID                     int64  `json:"id"`
	IsProjectLabel         bool   `json:"is_project_label"`
	Name                   string `json:"name"`
	OpenIssuesCount        int64  `json:"open_issues_count"`
	OpenMergeRequestsCount int64  `json:"open_merge_requests_count"`
	Priority               int64  `json:"priority"`
	Subscribed             bool   `json:"subscribed"`
	TextColor              string `json:"text_color"`
}
