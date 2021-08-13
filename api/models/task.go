package models

type Task struct {
	// Id : The ID of task
	Id int `db:"id" json:"id"`

	// Query : Search query for this task.
	Query string `db:"query" json:"query"`

	// Title : Title of this task.
	Title string `db:"title" json:"title"`

	// Description : Description text of task.
	Description string `db:"description" json:"description"`

	// SearchUrl : Url used in this task.
	SearchUrl string `db:"search_url" json:"searchUrl"`
}

// GroupCounts : Struct for group count
type GroupCounts struct {
	GroupId int `db:"group_id" json:"groupId"`
	Count   int `db:"count" json:"count"`
}
