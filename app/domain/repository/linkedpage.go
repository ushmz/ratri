//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import "ratri/domain/model"

type LinkedPageRepository interface {
	Get(linkedPageId int) (model.LinkedPage, error)
	GetBySearchPageIds(pageId []int, taskId, top int) (*[]model.SearchPageWithLinkedPage, error)
	// [TODO] Better name
	GetRatioBySearchPageIds(pageId []int, taskId int) (*[]model.SearchPageWithLinkedPageRatio, error)
	Select(linkedPageIds []int) (*[]model.LinkedPage, error)
	List(offset, limit int) (*[]model.LinkedPage, error)
	// Following methods are not implemented
	// because this `LinkedPage` resource should not be created by API.
	// Create(model.LinkedPage) (model.LinkedPage, error)
	// Update(model.LinkedPage) (model.LinkedPage, error)
	// Delete(linkedPageId int) error
}