//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"ratri/internal/domain/model"
	repo "ratri/internal/domain/repository"
)

type Log interface {
	CumulateSerpViewingTime(*model.SerpViewingLogParam) error
	CumulatePageViewingTime(*model.PageViewingLogParam) error
	StoreSerpEventLog(*model.SearchPageEventLogParam) error
	StoreSearchSeeion(*model.SearchSession) error
}

type LogImpl struct {
	repository repo.LogRepository
}

func NewLogUsecase(logRepository repo.LogRepository) Log {
	return &LogImpl{repository: logRepository}
}

func (l *LogImpl) CumulateSerpViewingTime(p *model.SerpViewingLogParam) error {
	return l.repository.CumulateSerpViewingTime(p)
}

func (l *LogImpl) CumulatePageViewingTime(p *model.PageViewingLogParam) error {
	return l.repository.CumulatePageViewingTime(p)
}

func (l *LogImpl) StoreSerpEventLog(p *model.SearchPageEventLogParam) error {
	return l.repository.StoreSerpEventLog(p)
}

func (l *LogImpl) StoreSearchSeeion(p *model.SearchSession) error {
	return l.repository.StoreSearchSeeion(p)
}
