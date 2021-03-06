package mysql_test

import (
	"ratri/domain/model"
	"testing"
)

func TestCumulateSerpViewingTime(t *testing.T) {
	p := model.SerpDwellTimeLogParam{
		UserID:      42,
		TaskID:      5,
		ConditionID: 5,
	}

	if err := logDao.CumulateSerpDwellTime(p); err != nil {
		t.Fatal(err)
	}
}

func TestCumulatePageViewingTime(t *testing.T) {
	p := model.PageDwellTimeLogParam{
		UserID:      42,
		TaskID:      5,
		ConditionID: 5,
		PageID:      432,
	}

	if err := logDao.CumulatePageDwellTime(p); err != nil {
		t.Fatal(err)
	}
}

func TestStoreSerpEventLog(t *testing.T) {
	p := model.SearchPageEventLogParam{
		User:        42,
		TaskID:      5,
		ConditionID: 5,
		Time:        142,
		Page:        2,
		Rank:        5,
		IsVisible:   true,
		Event:       "click",
	}
	if err := logDao.StoreSerpEventLog(p); err != nil {
		t.Fatal(err)
	}
}

func TestStoreSearchSession(t *testing.T) {
	p := model.SearchSessionParam{
		UserID:      42,
		TaskID:      5,
		ConditionID: 5,
	}

	if err := logDao.StoreSearchSeeion(p); err != nil {
		t.Fatal(err)
	}
}
