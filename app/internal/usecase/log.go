//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"errors"
	"ratri/internal/domain/model"
	repo "ratri/internal/domain/repository"
	"ratri/internal/domain/store"
)

type LogUsecase interface {
	CumulateSerpViewingTime(*model.SerpViewingLogParam) error
	CumulatePageViewingTime(*model.PageViewingLogParam) error
	StoreSerpEventLog(*model.SearchPageEventLogParam) error
	StoreSearchSeeion(*model.SearchSessionParam) error

	ExportSerpViewingTimeLog(bool, store.FileType) ([]byte, error)
	ExportPageViewingTimeLog(bool, store.FileType) ([]byte, error)
	ExportSerpEventLog(bool, store.FileType) ([]byte, error)
	ExportSearchSeeion(bool, store.FileType) ([]byte, error)
}

type LogImpl struct {
	dbRepo repo.LogRepository
	store  store.LogStore
}

func NewLogUsecase(logRepository repo.LogRepository, store store.LogStore) LogUsecase {
	return &LogImpl{dbRepo: logRepository, store: store}
}

func (l *LogImpl) CumulateSerpViewingTime(p *model.SerpViewingLogParam) error {
	return l.dbRepo.CumulateSerpViewingTime(p)
}

func (l *LogImpl) CumulatePageViewingTime(p *model.PageViewingLogParam) error {
	return l.dbRepo.CumulatePageViewingTime(p)
}

func (l *LogImpl) StoreSerpEventLog(p *model.SearchPageEventLogParam) error {
	return l.dbRepo.StoreSerpEventLog(p)
}

func (l *LogImpl) StoreSearchSeeion(p *model.SearchSessionParam) error {
	return l.dbRepo.StoreSearchSeeion(p)
}

func (l *LogImpl) ExportSerpViewingTimeLog(header bool, filetype store.FileType) ([]byte, error) {
	if l == nil {
		return nil, errors.New("LogImpl is nil")
	}

	data, err := l.dbRepo.FetchAllSerpViewingTimeLogs()
	if err != nil {
		return nil, err
	}

	b, err := l.store.ExportSerpViewingTimeLog(data, header, filetype)
	if err != nil {
		return nil, err
	}

	return b, nil

}

func (l *LogImpl) ExportPageViewingTimeLog(header bool, filetype store.FileType) ([]byte, error) {
	if l == nil {
		return nil, errors.New("LogImpl is nil")
	}

	data, err := l.dbRepo.FetchAllPageViewingTimeLogs()
	if err != nil {
		return nil, err
	}

	b, err := l.store.ExportPageViewingTimeLog(data, header, filetype)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (l *LogImpl) ExportSerpEventLog(header bool, filetype store.FileType) ([]byte, error) {
	if l == nil {
		return nil, errors.New("LogImpl is nil")
	}

	data, err := l.dbRepo.FetchAllSerpEventLogs()
	if err != nil {
		return nil, err
	}

	b, err := l.store.ExportSerpEventLog(data, header, filetype)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (l *LogImpl) ExportSearchSeeion(header bool, filetype store.FileType) ([]byte, error) {
	if l == nil {
		return nil, errors.New("LogImpl is nil")
	}

	data, err := l.dbRepo.FetchAllSearchSessions()
	if err != nil {
		return nil, err
	}

	b, err := l.store.ExportSearchSessionLog(data, header, filetype)
	if err != nil {
		return nil, err
	}

	return b, nil
}
