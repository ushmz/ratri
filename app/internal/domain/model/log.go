package model

// SerpViewingLogParamWithTime : Struct for task viewing time log request body
// This struct is deprecated on the same reason with `repository.StoreTaskTimeLog()`
type SerpViewingLogParamWithTime struct {
	// UserId : The ID of user (worker)
	UserId int `db:"user_id" json:"user"`

	// TaskId : The Id of task that user working.
	TaskId int `db:"task_id" json:"task"`

	// ConditionId : User's condition Id that means group and task category.
	ConditionId int `db:"condition_id" json:"condition"`

	// TimeOnPage : User's page viewing time.
	TimeOnPage int `db:"time_on_page" json:"time"`
}

// SerpViewingLogParam : Struct for task viewing time log request body
type SerpViewingLogParam struct {
	// UserId : The ID of user (worker)
	UserId int `db:"user_id" json:"user"`

	// TaskId : The Id of task that user working.
	TaskId int `db:"task_id" json:"task"`

	// ConditionId : User's condition Id that means group and task category.
	ConditionId int `db:"condition_id" json:"condition"`
}

// PageViewingLogParam : Struct for each search result page viewing time log request body
type PageViewingLogParam struct {
	// UserId : The ID of user (worker)
	UserId int `db:"user_id" json:"user"`

	// TaskId : The Id of task that user working.
	TaskId int `db:"task_id" json:"task"`

	// ConditionId : User's condition Id that means group and task category.
	ConditionId int `db:"condition_id" json:"condition"`

	// PageId : Page ID that user view.
	PageId int `db:"page_id" json:"page"`
}

// SearchPageEventLogParam : Struct for page click log request body.
type SearchPageEventLogParam struct {
	// Id : The ID of each log record.
	Id string `db:"id" json:"id"`

	// Uid : The ID of user (worker)
	User int `db:"user_id" json:"user"`

	// TaskId : The Id of task that user working.
	TaskId int `db:"task_id" json:"task"`

	// ConditionId : User's condition Id that means group and task category.
	ConditionId int `db:"condition_id" json:"condition"`

	// Time : User's page viewing time.
	Time int `db:"time_on_page" json:"time"`

	// Page : The Id of page that user clicked.
	Page int `db:"serp_page" json:"page"`

	// Rank : Search result rank that user clicked.
	Rank int `db:"serp_rank" json:"rank"`

	// IsVisible : Risk is visible or not.
	IsVisible bool `db:"is_visible" json:"visible"`

	// Event : It is expected to be "click", "hover" or "paginate"
	Event string `db:"event" json:"event"`
}

// SearchSession : Struct fot search session request body.
type SearchSession struct {
	// UserId : Allocated ID of user (worker)
	UserId int `db:"user_id" json:"user"`

	// TaskId : The Id of task that user working.
	TaskId int `db:"task_id" json:"task"`

	// ConditionId : User's condition Id that means group and task category.
	ConditionId int `db:"condition_id" json:"condition"`
}
