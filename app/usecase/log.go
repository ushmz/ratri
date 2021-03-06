//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"bytes"
	"fmt"
	"ratri/domain/model"
	repo "ratri/domain/repository"
	"ratri/domain/store"
)

// LogUsecase : Abstract operations that log usecase should have
type LogUsecase interface {
	CumulateSerpDwellTime(*model.SerpDwellTimeLogParam) error
	CumulatePageDwellTime(*model.PageDwellTimeLogParam) error
	StoreSerpEventLog(*model.SearchPageEventLogParam) error
	StoreSearchSeeion(*model.SearchSessionParam) error

	ExportSerpDwellTimeLog(bool, store.FileType) (*bytes.Buffer, error)
	ExportPageDwellTimeLog(bool, store.FileType) (*bytes.Buffer, error)
	ExportSerpEventLog(bool, store.FileType) (*bytes.Buffer, error)
	ExportSearchSeeion(bool, store.FileType) (*bytes.Buffer, error)
}

// LogImpl : Implemention of log usecase
type LogImpl struct {
	repository repo.LogRepository
	store      store.LogStore
}

// NewLogUsecase : Return new log usecase
func NewLogUsecase(logRepository repo.LogRepository, store store.LogStore) LogUsecase {
	return &LogImpl{repository: logRepository, store: store}
}

// CumulateSerpDwellTime : Count up dwell time in SERP
func (l *LogImpl) CumulateSerpDwellTime(p *model.SerpDwellTimeLogParam) error {
	if l == nil {
		return model.ErrNilReceiver
	}

	if err := l.repository.CumulateSerpDwellTime(p); err != nil {
		return fmt.Errorf("Try to store SERP dwell time log: %w", err)
	}
	return nil
}

// CumulatePageDwellTime : Count up dwell time in result page
func (l *LogImpl) CumulatePageDwellTime(p *model.PageDwellTimeLogParam) error {
	if l == nil {
		return model.ErrNilReceiver
	}
	if err := l.repository.CumulatePageDwellTime(p); err != nil {
		return fmt.Errorf("Try to store search result page dwell time log: %w", err)
	}
	return nil
}

// StoreSerpEventLog : Store events in SERP
func (l *LogImpl) StoreSerpEventLog(p *model.SearchPageEventLogParam) error {
	if l == nil {
		return model.ErrNilReceiver
	}
	if err := l.repository.StoreSerpEventLog(p); err != nil {
		return fmt.Errorf("Try to storee SERP event log: %w", err)
	}
	return nil
}

// StoreSearchSeeion : Store search task session
func (l *LogImpl) StoreSearchSeeion(p *model.SearchSessionParam) error {
	if l == nil {
		return model.ErrNilReceiver
	}
	if err := l.repository.StoreSearchSeeion(p); err != nil {
		return fmt.Errorf("Try to store search session log: %w", err)
	}
	return nil
}

// ExportSerpDwellTimeLog : Export SERP dwell time logs
func (l *LogImpl) ExportSerpDwellTimeLog(header bool, filetype store.FileType) (*bytes.Buffer, error) {
	if l == nil {
		return nil, model.ErrNilReceiver
	}

	data, err := l.repository.FetchAllSerpDwellTimeLogs()
	if err != nil {
		return nil, fmt.Errorf("Try te get SERP dwell time logs: %w", err)
	}

	b, err := l.store.ExportSerpDwellTimeLog(data, header, filetype)
	if err != nil {
		return nil, fmt.Errorf("Try to write log data to buffer: %w", err)
	}

	return b, nil

}

// ExportPageDwellTimeLog : Export result page dwell time logs
func (l *LogImpl) ExportPageDwellTimeLog(header bool, filetype store.FileType) (*bytes.Buffer, error) {
	if l == nil {
		return nil, model.ErrNilReceiver
	}

	data, err := l.repository.FetchAllPageDwellTimeLogs()
	if err != nil {
		return nil, fmt.Errorf("Try to get search result page dwell time logs: %w", err)
	}

	b, err := l.store.ExportPageDwellTimeLog(data, header, filetype)
	if err != nil {
		return nil, fmt.Errorf("Try to write log data to buffer: %w", err)
	}

	return b, nil
}

// ExportSerpEventLog : Export SERP event logs
func (l *LogImpl) ExportSerpEventLog(header bool, filetype store.FileType) (*bytes.Buffer, error) {
	if l == nil {
		return nil, model.ErrNilReceiver
	}

	data, err := l.repository.FetchAllSerpEventLogs()
	if err != nil {
		return nil, fmt.Errorf("Try to get SERP event logs: %w", err)
	}

	b, err := l.store.ExportSerpEventLog(data, header, filetype)
	if err != nil {
		return nil, fmt.Errorf("Try to write log data to buffer: %w", err)
	}

	return b, nil
}

// ExportSearchSeeion : Export search session logs
func (l *LogImpl) ExportSearchSeeion(header bool, filetype store.FileType) (*bytes.Buffer, error) {
	if l == nil {
		return nil, model.ErrNilReceiver
	}

	data, err := l.repository.FetchAllSearchSessions()
	if err != nil {
		return nil, fmt.Errorf("Try to get search session logs: %w", err)
	}

	b, err := l.store.ExportSearchSessionLog(data, header, filetype)
	if err != nil {
		return nil, fmt.Errorf("Try to write log data to buffer: %w", err)
	}

	return b, nil
}
