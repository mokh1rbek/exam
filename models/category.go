package models

type CategoryPrimarKey struct {
	Id         string `json:"category_id"`
	ParentUUID string `json:"parent_uuid"`
}

type CreateCategory struct {
	Name       string `json:"name"`
	ParentUUID string `json:"parent_uuid"`
}
type Category struct {
	Id         string `json:"category_id"`
	Name       string `json:"name"`
	ParentUUID string `json:"parent_uuid"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type UpdateCategory struct {
	Id         string `json:"category_id"`
	Name       string `json:"name"`
	ParentUUID string `json:"parent_uuid"`
}

type GetListCategoryRequest struct {
	Limit  int32
	Offset int32
}

type GetListCategoryResponse struct {
	Count     int32               `json:"count"`
	Categorys []*CategoryByParent `json:"categorys"`
}

type CategoryByParent struct {
	Id            string      `json:"id"`
	Name          string      `json:"name"`
	ParentUUID    string      `json:"parent_uuid"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	ChildCategory []*Category `json:"childs"`
}
