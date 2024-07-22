package repository

import "github.com/aleroxac/b2c/internal/dto"

type ReportRespository interface {
	Create(report *dto.Report) error
	Read() (*dto.Report, error)
	Update(report *dto.Report) error
}
