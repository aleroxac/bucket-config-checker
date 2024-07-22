package dto

import "time"

type Failure struct {
	What  string
	When  time.Time
	Why   string
	Where string
}

type Scan struct {
	Path     string
	Name     string
	Project  string
	Failures []Failure
}

type Check struct {
	Pass []string
	Fail []Scan
}

type Report struct {
	Checks []Check
}
