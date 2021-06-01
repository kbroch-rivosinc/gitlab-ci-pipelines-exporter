package schemas

import (
	"hash/crc32"
	"strconv"
)

var (
	defaultProjectOutputSparseStatusMetrics                               = true
	defaultProjectPullEnvironmentsEnabled                                 = false
	defaultProjectPullEnvironmentsRegexp                                  = `.*`
	defaultProjectPullRefsRegexp                                          = `^(main|master)$`
	defaultProjectPullRefsMaxAgeSeconds                              uint = 0
	defaultProjectPullRefsFromPipelinesEnabled                            = false
	defaultProjectPullRefsFromPipelinesDepth                              = 100
	defaultProjectPullRefsFromMergeRequestsEnabled                        = false
	defaultProjectPullRefsFromMergeRequestsDepth                          = 1
	defaultProjectPullPipelineJobsEnabled                                 = false
	defaultProjectPullPipelineJobsFromChildPipelinesEnabled               = true
	defaultProjectPullPipelineJobsRunnerDescriptionEnabled                = true
	defaultProjectPullPipelineJobsRunnerDescriptionAggregationRegexp      = `shared-runners-manager-(\d*)\.gitlab\.com`
	defaultProjectPullPipelineVariablesEnabled                            = false
	defaultProjectPullPipelineVariablesRegexp                             = `.*`
)

// ProjectParameters for the fetching configuration of Projects and Wildcards
type ProjectParameters struct {
	// From handles ProjectPullParameters configuration
	Pull ProjectPull `yaml:"pull"`

	// Whether or not to export all pipeline/job statuses (being 0) or solely the one of the last job (being 1).
	OutputSparseStatusMetricsValue *bool `yaml:"output_sparse_status_metrics"`
}

// ProjectPull ..
type ProjectPull struct {
	Environments ProjectPullEnvironments `yaml:"environments"`
	Refs         ProjectPullRefs         `yaml:"refs"`
	Pipeline     ProjectPullPipeline     `yaml:"pipeline"`
}

// ProjectPullEnvironments ..
type ProjectPullEnvironments struct {
	// Whether to pull environments/deployments or not for this project
	EnabledValue *bool `yaml:"enabled"`

	// Regular expression to filter environments to fetch by their names
	RegexpValue *string `yaml:"regexp"`
}

// ProjectPullRefs ..
type ProjectPullRefs struct {
	// Regular expression to filter refs to fetch (defaults to '.*')
	RegexpValue *string `yaml:"regexp"`

	// If the age of the most recent pipeline for the ref is greater than this value, the ref won't get exported
	MaxAgeSecondsValue *uint `yaml:"max_age_seconds"`

	// From handles ProjectPullRefsFromParameters configuration
	From ProjectPullRefsFrom `yaml:"from"`
}

// ProjectPullRefsFrom ..
type ProjectPullRefsFrom struct {
	// Pipelines defines whether or not to fetch refs from historical pipelines
	Pipelines ProjectPullRefsFromPipelines `yaml:"pipelines"`

	// MergeRequests defines whether or not to fetch refs from merge requests
	MergeRequests ProjectPullRefsFromMergeRequests `yaml:"merge_requests"`
}

// ProjectPullRefsFromParameters ..
type ProjectPullRefsFromParameters struct {
	EnabledValue *bool `yaml:"enabled"`
	DepthValue   *int  `yaml:"depth"`
}

// ProjectPullRefsFromPipelines ..
type ProjectPullRefsFromPipelines ProjectPullRefsFromParameters

// ProjectPullRefsFromMergeRequests ..
type ProjectPullRefsFromMergeRequests ProjectPullRefsFromParameters

// ProjectPullPipeline ..
type ProjectPullPipeline struct {
	Jobs      ProjectPullPipelineJobs      `yaml:"jobs"`
	Variables ProjectPullPipelineVariables `yaml:"variables"`
}

// ProjectPullPipelineJobs ..
type ProjectPullPipelineJobs struct {
	// Enabled set to true will pull pipeline jobs related metrics
	EnabledValue *bool `yaml:"enabled"`

	// Pull pipeline jobs from child/downstream pipelines
	FromChildPipelines ProjectPullPipelineJobsFromChildPipelines `yaml:"from_child_pipelines"`

	// Configure the export of the runner description which ran the job
	RunnerDescription ProjectPullPipelineJobsRunnerDescription `yaml:"runner_description"`
}

// ProjectPullPipelineJobsFromChildPipelines ..
type ProjectPullPipelineJobsFromChildPipelines struct {
	// Enabled set to true will pull pipeline jobs from child/downstream pipelines related metrics
	EnabledValue *bool `yaml:"enabled"`
}

// ProjectPullPipelineJobsRunnerDescription ..
type ProjectPullPipelineJobsRunnerDescription struct {
	// Enabled set to true will export the description of the runner which ran the job
	EnabledValue *bool `yaml:"enabled"`

	// Regular expression to be able to reduce the cardinality of the exported value when necessary
	AggregationRegexpValue *string `yaml:"aggregation_regexp"`
}

// ProjectPullPipelineVariables ..
type ProjectPullPipelineVariables struct {
	// Enabled set to true will attempt to retrieve variables included in the pipeline
	EnabledValue *bool `yaml:"enabled"`

	// Regexp to filter pipeline variables values to fetch (defaults to '.*')
	RegexpValue *string `yaml:"regexp"`
}

// UpdateProjectDefaults ..
func UpdateProjectDefaults(d ProjectParameters) {
	if d.OutputSparseStatusMetricsValue != nil {
		defaultProjectOutputSparseStatusMetrics = *d.OutputSparseStatusMetricsValue
	}

	if d.Pull.Environments.EnabledValue != nil {
		defaultProjectPullEnvironmentsEnabled = *d.Pull.Environments.EnabledValue
	}

	if d.Pull.Environments.RegexpValue != nil {
		defaultProjectPullEnvironmentsRegexp = *d.Pull.Environments.RegexpValue
	}

	if d.Pull.Refs.RegexpValue != nil {
		defaultProjectPullRefsRegexp = *d.Pull.Refs.RegexpValue
	}

	if d.Pull.Refs.MaxAgeSecondsValue != nil {
		defaultProjectPullRefsMaxAgeSeconds = *d.Pull.Refs.MaxAgeSecondsValue
	}

	if d.Pull.Refs.From.Pipelines.EnabledValue != nil {
		defaultProjectPullRefsFromPipelinesEnabled = *d.Pull.Refs.From.Pipelines.EnabledValue
	}

	if d.Pull.Refs.From.Pipelines.DepthValue != nil {
		defaultProjectPullRefsFromPipelinesDepth = *d.Pull.Refs.From.Pipelines.DepthValue
	}

	if d.Pull.Refs.From.MergeRequests.EnabledValue != nil {
		defaultProjectPullRefsFromMergeRequestsEnabled = *d.Pull.Refs.From.MergeRequests.EnabledValue
	}

	if d.Pull.Refs.From.MergeRequests.DepthValue != nil {
		defaultProjectPullRefsFromMergeRequestsDepth = *d.Pull.Refs.From.MergeRequests.DepthValue
	}

	if d.Pull.Pipeline.Jobs.EnabledValue != nil {
		defaultProjectPullPipelineJobsEnabled = *d.Pull.Pipeline.Jobs.EnabledValue
	}

	if d.Pull.Pipeline.Jobs.FromChildPipelines.EnabledValue != nil {
		defaultProjectPullPipelineJobsFromChildPipelinesEnabled = *d.Pull.Pipeline.Jobs.FromChildPipelines.EnabledValue
	}

	if d.Pull.Pipeline.Jobs.RunnerDescription.EnabledValue != nil {
		defaultProjectPullPipelineJobsRunnerDescriptionEnabled = *d.Pull.Pipeline.Jobs.RunnerDescription.EnabledValue
	}

	if d.Pull.Pipeline.Jobs.RunnerDescription.AggregationRegexpValue != nil {
		defaultProjectPullPipelineJobsRunnerDescriptionAggregationRegexp = *d.Pull.Pipeline.Jobs.RunnerDescription.AggregationRegexpValue
	}

	if d.Pull.Pipeline.Variables.EnabledValue != nil {
		defaultProjectPullPipelineVariablesEnabled = *d.Pull.Pipeline.Variables.EnabledValue
	}

	if d.Pull.Pipeline.Variables.RegexpValue != nil {
		defaultProjectPullPipelineVariablesRegexp = *d.Pull.Pipeline.Variables.RegexpValue
	}
}

// Project holds information about a GitLab project
type Project struct {
	// ProjectParameters holds parameters specific to this project
	ProjectParameters `yaml:",inline"`

	// Name is actually what is commonly referred as path_with_namespace on GitLab
	Name string `yaml:"name"`
}

// ProjectKey ..
type ProjectKey string

// Key ..
func (p Project) Key() ProjectKey {
	return ProjectKey(strconv.Itoa(int(crc32.ChecksumIEEE([]byte(p.Name)))))
}

// Projects ..
type Projects map[ProjectKey]Project

// OutputSparseStatusMetrics ...
func (p *ProjectParameters) OutputSparseStatusMetrics() bool {
	if p.OutputSparseStatusMetricsValue != nil {
		return *p.OutputSparseStatusMetricsValue
	}

	return defaultProjectOutputSparseStatusMetrics
}

// Enabled ...
func (p *ProjectPullEnvironments) Enabled() bool {
	if p.EnabledValue != nil {
		return *p.EnabledValue
	}

	return defaultProjectPullEnvironmentsEnabled
}

// Regexp ...
func (p *ProjectPullEnvironments) Regexp() string {
	if p.RegexpValue != nil {
		return *p.RegexpValue
	}

	return defaultProjectPullEnvironmentsRegexp
}

// Regexp ...
func (p *ProjectPullRefs) Regexp() string {
	if p.RegexpValue != nil {
		return *p.RegexpValue
	}

	return defaultProjectPullRefsRegexp
}

// MaxAgeSeconds ...
func (p *ProjectPullRefs) MaxAgeSeconds() uint {
	if p.MaxAgeSecondsValue != nil {
		return *p.MaxAgeSecondsValue
	}

	return defaultProjectPullRefsMaxAgeSeconds
}

// Enabled ...
func (p *ProjectPullRefsFromPipelines) Enabled() bool {
	if p.EnabledValue != nil {
		return *p.EnabledValue
	}

	return defaultProjectPullRefsFromPipelinesEnabled
}

// Depth ...
func (p *ProjectPullRefsFromPipelines) Depth() int {
	if p.DepthValue != nil {
		return *p.DepthValue
	}

	return defaultProjectPullRefsFromPipelinesDepth
}

// Enabled ...
func (p *ProjectPullRefsFromMergeRequests) Enabled() bool {
	if p.EnabledValue != nil {
		return *p.EnabledValue
	}

	return defaultProjectPullRefsFromMergeRequestsEnabled
}

// Depth ...
func (p *ProjectPullRefsFromMergeRequests) Depth() int {
	if p.DepthValue != nil {
		return *p.DepthValue
	}

	return defaultProjectPullRefsFromMergeRequestsDepth
}

// Enabled ...
func (p *ProjectPullPipelineJobs) Enabled() bool {
	if p.EnabledValue != nil {
		return *p.EnabledValue
	}

	return defaultProjectPullPipelineJobsEnabled
}

// Enabled ...
func (p *ProjectPullPipelineJobsFromChildPipelines) Enabled() bool {
	if p.EnabledValue != nil {
		return *p.EnabledValue
	}

	return defaultProjectPullPipelineJobsFromChildPipelinesEnabled
}

// Enabled ...
func (p *ProjectPullPipelineJobsRunnerDescription) Enabled() bool {
	if p.EnabledValue != nil {
		return *p.EnabledValue
	}

	return defaultProjectPullPipelineJobsRunnerDescriptionEnabled
}

// AggregationRegexp ...
func (p *ProjectPullPipelineJobsRunnerDescription) AggregationRegexp() string {
	if p.AggregationRegexpValue != nil {
		return *p.AggregationRegexpValue
	}

	return defaultProjectPullPipelineJobsRunnerDescriptionAggregationRegexp
}

// Enabled ...
func (p *ProjectPullPipelineVariables) Enabled() bool {
	if p.EnabledValue != nil {
		return *p.EnabledValue
	}

	return defaultProjectPullPipelineVariablesEnabled
}

// Regexp ...
func (p *ProjectPullPipelineVariables) Regexp() string {
	if p.RegexpValue != nil {
		return *p.RegexpValue
	}

	return defaultProjectPullPipelineVariablesRegexp
}
