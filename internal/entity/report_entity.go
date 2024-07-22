package entity

import "github.com/aleroxac/b2c/internal/dto"

func NewReport(checks []dto.Check) *dto.Report {
	return &dto.Report{
		Checks: checks,
	}
}
