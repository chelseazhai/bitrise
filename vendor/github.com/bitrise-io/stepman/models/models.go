package models

import (
	"time"

	envmanModels "github.com/bitrise-io/envman/models"
)

// GlobalStepInfoModel ...
type GlobalStepInfoModel struct {
	RemovalDate    string `json:"removal_date,omitempty" yaml:"removal_date,omitempty"`
	DeprecateNotes string `json:"deprecate_notes,omitempty" yaml:"deprecate_notes,omitempty"`
}

// StepSourceModel ...
type StepSourceModel struct {
	Git    string `json:"git,omitempty" yaml:"git,omitempty"`
	Commit string `json:"commit,omitempty" yaml:"commit,omitempty"`
}

// DependencyModel ...
type DependencyModel struct {
	Manager string `json:"manager,omitempty" yaml:"manager,omitempty"`
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
}

// BrewDepModel ...
type BrewDepModel struct {
	// Name is the package name for Brew
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// BinName is the binary's name, if it doesn't match the package's name.
	// Can be used for e.g. calling `which`.
	// E.g. in case of "AWS CLI" the package is `awscli` and the binary is `aws`.
	// If BinName is empty Name will be used as BinName too.
	BinName string `json:"bin_name,omitempty" yaml:"bin_name,omitempty"`
}

// AptGetDepModel ...
type AptGetDepModel struct {
	// Name is the package name for Apt-get
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// BinName is the binary's name, if it doesn't match the package's name.
	// Can be used for e.g. calling `which`.
	// E.g. in case of "AWS CLI" the package is `awscli` and the binary is `aws`.
	// If BinName is empty Name will be used as BinName too.
	BinName string `json:"bin_name,omitempty" yaml:"bin_name,omitempty"`
}

// CheckOnlyDepModel ...
type CheckOnlyDepModel struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// DepsModel ...
type DepsModel struct {
	Brew      []BrewDepModel      `json:"brew,omitempty" yaml:"brew,omitempty"`
	AptGet    []AptGetDepModel    `json:"apt_get,omitempty" yaml:"apt_get,omitempty"`
	CheckOnly []CheckOnlyDepModel `json:"check_only,omitempty" yaml:"check_only,omitempty"`
}

// BashStepToolkitModel ...
type BashStepToolkitModel struct {
	EntryFile string `json:"entry_file,omitempty" yaml:"entry_file,omitempty"`
}

// GoStepToolkitModel ...
type GoStepToolkitModel struct {
	// PackageName - required
	PackageName string `json:"package_name" yaml:"package_name"`
}

// StepToolkitModel ...
type StepToolkitModel struct {
	Bash *BashStepToolkitModel `json:"bash,omitempty" yaml:"bash,omitempty"`
	Go   *GoStepToolkitModel   `json:"go,omitempty" yaml:"go,omitempty"`
}

// StepModel ...
type StepModel struct {
	Title       *string `json:"title,omitempty" yaml:"title,omitempty"`
	Summary     *string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	//
	Website       *string `json:"website,omitempty" yaml:"website,omitempty"`
	SourceCodeURL *string `json:"source_code_url,omitempty" yaml:"source_code_url,omitempty"`
	SupportURL    *string `json:"support_url,omitempty" yaml:"support_url,omitempty"`
	// auto-generated at share
	PublishedAt *time.Time        `json:"published_at,omitempty" yaml:"published_at,omitempty"`
	Source      *StepSourceModel  `json:"source,omitempty" yaml:"source,omitempty"`
	AssetURLs   map[string]string `json:"asset_urls,omitempty" yaml:"asset_urls,omitempty"`
	//
	HostOsTags          []string          `json:"host_os_tags,omitempty" yaml:"host_os_tags,omitempty"`
	ProjectTypeTags     []string          `json:"project_type_tags,omitempty" yaml:"project_type_tags,omitempty"`
	TypeTags            []string          `json:"type_tags,omitempty" yaml:"type_tags,omitempty"`
	Dependencies        []DependencyModel `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Toolkit             *StepToolkitModel `json:"toolkit,omitempty" yaml:"toolkit,omitempty"`
	Deps                *DepsModel        `json:"deps,omitempty" yaml:"deps,omitempty"`
	IsRequiresAdminUser *bool             `json:"is_requires_admin_user,omitempty" yaml:"is_requires_admin_user,omitempty"`
	// IsAlwaysRun : if true then this step will always run,
	//  even if a previous step fails.
	IsAlwaysRun *bool `json:"is_always_run,omitempty" yaml:"is_always_run,omitempty"`
	// IsSkippable : if true and this step fails the build will still continue.
	//  If false then the build will be marked as failed and only those
	//  steps will run which are marked with IsAlwaysRun.
	IsSkippable *bool `json:"is_skippable,omitempty" yaml:"is_skippable,omitempty"`
	// RunIf : only run the step if the template example evaluates to true
	RunIf   *string `json:"run_if,omitempty" yaml:"run_if,omitempty"`
	Timeout *int    `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	//
	Inputs  []envmanModels.EnvironmentItemModel `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Outputs []envmanModels.EnvironmentItemModel `json:"outputs,omitempty" yaml:"outputs,omitempty"`
}

// StepGroupInfoModel ...
type StepGroupInfoModel struct {
	RemovalDate    string            `json:"removal_date,omitempty" yaml:"removal_date,omitempty"`
	DeprecateNotes string            `json:"deprecate_notes,omitempty" yaml:"deprecate_notes,omitempty"`
	AssetURLs      map[string]string `json:"asset_urls,omitempty" yaml:"asset_urls,omitempty"`
}

// StepGroupModel ...
type StepGroupModel struct {
	Info                StepGroupInfoModel   `json:"info,omitempty" yaml:"info,omitempty"`
	LatestVersionNumber string               `json:"latest_version_number,omitempty" yaml:"latest_version_number,omitempty"`
	Versions            map[string]StepModel `json:"versions,omitempty" yaml:"versions,omitempty"`
}

// StepHash ...
type StepHash map[string]StepGroupModel

// DownloadLocationModel ...
type DownloadLocationModel struct {
	Type string `json:"type"`
	Src  string `json:"src"`
}

// StepCollectionModel ...
type StepCollectionModel struct {
	FormatVersion         string                  `json:"format_version" yaml:"format_version"`
	GeneratedAtTimeStamp  int64                   `json:"generated_at_timestamp" yaml:"generated_at_timestamp"`
	SteplibSource         string                  `json:"steplib_source" yaml:"steplib_source"`
	DownloadLocations     []DownloadLocationModel `json:"download_locations" yaml:"download_locations"`
	AssetsDownloadBaseURI string                  `json:"assets_download_base_uri" yaml:"assets_download_base_uri"`
	Steps                 StepHash                `json:"steps" yaml:"steps"`
}

// EnvInfoModel ...
type EnvInfoModel struct {
	Key          string   `json:"key,omitempty" yaml:"key,omitempty"`
	Title        string   `json:"title,omitempty" yaml:"title,omitempty"`
	Description  string   `json:"description,omitempty" yaml:"description,omitempty"`
	ValueOptions []string `json:"value_options,omitempty" yaml:"value_options,omitempty"`
	DefaultValue string   `json:"default_value,omitempty" yaml:"default_value,omitempty"`
	IsExpand     bool     `json:"is_expand" yaml:"is_expand"`
}

// StepInfoModel ...
type StepInfoModel struct {
	ID            string              `json:"step_id,omitempty" yaml:"step_id,omitempty"`
	Title         string              `json:"step_title,omitempty" yaml:"step_title,omitempty"`
	Version       string              `json:"step_version,omitempty" yaml:"step_version,omitempty"`
	Latest        string              `json:"latest_version,omitempty" yaml:"latest_version,omitempty"`
	Description   string              `json:"description,omitempty" yaml:"description,omitempty"`
	Source        string              `json:"source,omitempty" yaml:"source,omitempty"`
	StepLib       string              `json:"steplib,omitempty" yaml:"steplib,omitempty"`
	SupportURL    string              `json:"support_url,omitempty" yaml:"support_url,omitempty"`
	SourceCodeURL string              `json:"source_code_url,omitempty" yaml:"source_code_url,omitempty"`
	Inputs        []EnvInfoModel      `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Outputs       []EnvInfoModel      `json:"outputs,omitempty" yaml:"outputs,omitempty"`
	GlobalInfo    GlobalStepInfoModel `json:"global_info,omitempty" yaml:"global_info,omitempty"`
}

// StepListModel ...
type StepListModel struct {
	StepLib string   `json:"steplib,omitempty" yaml:"steplib,omitempty"`
	Steps   []string `json:"steps,omitempty" yaml:"steps,omitempty"`
}

// StepLibURIModel ...
type StepLibURIModel struct {
	URI string `json:"uri,omitempty" yaml:"uri,omitempty"`
}

// StepLibURIsModel ...
type StepLibURIsModel struct {
	StepLibURIs []StepLibURIModel `json:"steplibs,omitempty" yaml:"steplibs,omitempty"`
}
