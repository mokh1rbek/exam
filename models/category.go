package models

type CategoriesPrimarKey struct {
	Id       string `json:"category_id"`
	ParentID string `json:"parent_id"`
}

type CreateCategories struct {
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}
type Categories struct {
	Id            string        `json:"category_id"`
	Name          string        `json:"name"`
	ParentID      string        `json:"parent_id"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
	ChildCategory []*Categories `json:"childs"`
}

type UpdateCategories struct {
	Id       string `json:"category_id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

type GetListCategoriesRequest struct {
	Limit  int32
	Offset int32
}

type GetListCategoriesResponse struct {
	Count      int32               `json:"count"`
	Categories []*ChildsByCategory `json:"categorys"`
}

type ChildsByCategory struct {
	Id            string              `json:"category_id"`
	Name          string              `json:"name"`
	ParentID      string              `json:"parent_id"`
	ChildCategory []*ChildsByCategory `json:"childs"`
}
