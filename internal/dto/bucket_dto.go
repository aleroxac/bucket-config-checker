package dto

type Action struct {
	Type string
}

type Condition struct {
	IsLive                  int
	NumNewerVersions        int
	DaysSinceNonCurrentTime int
}

type Rule struct {
	Action    Action
	Condition Condition
}

type LifeCycleConfig struct {
	Rules []Rule
}

type Location struct {
	Type string
	Name string
}

type LoggingConfig struct {
	LogBucket       string
	LogObjectPrefix string
}

type IAMPolicy struct {
	Role    string
	Members []string
}

type PublicAccessConfig struct {
	IsPublic       bool
	PreventionType string
}

type Bucket struct {
	Path       string
	Name       string
	Project    string
	Class      string
	LifeCycle  LifeCycleConfig
	Location   Location
	Logging    LoggingConfig
	Public     PublicAccessConfig
	Versioning bool
}
