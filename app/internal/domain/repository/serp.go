package repository

import (
	"ratri/internal/domain/model"
)

type SerpRepository interface {
	FetchSerpByTaskID(taskId, offset int) ([]model.SearchPage, error)
	FetchSerpWithIconByTaskID(taskId, offset, top int) ([]model.SerpWithIconQueryResult, error)
	FetchSerpWithRatioByTaskID(taskId, offset, top int) ([]model.SerpWithRatioQueryResult, error)
}
