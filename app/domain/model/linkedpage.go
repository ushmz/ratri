package model

// LinkedPage : Linked page information with icon URL.
type LinkedPage struct {
	// ID : ID of linked page.
	ID int `db:"id" json:"id"`

	// Title : The title of linked page.
	Title string `db:"title" json:"title"`

	// URL : URL of the linked page.
	URL string `db:"url" json:"url"`

	// Icon : Favicon url of the page.
	Icon string `db:"icon_path" json:"icon"`

	// Category : Category name of linked page.
	Category string `db:"category" json:"category"`
}

// CategoryCount : Distribution information for each categories.
type CategoryCount struct {
	// Category : Category name.
	Category string `json:"category"`
	// Count : Total number of pages.
	Count int `json:"count"`
}
