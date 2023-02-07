package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"exam/models"
	"exam/pkg/helper"
)

type CategoriesRepo struct {
	db *pgxpool.Pool
}

func NewCategoriesRepo(db *pgxpool.Pool) *CategoriesRepo {
	return &CategoriesRepo{
		db: db,
	}
}

func (f *CategoriesRepo) Create(ctx context.Context, category *models.CreateCategories) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO categorys (
			id,
			name,
			parent_uuid,
			updated_at
		) VALUES ( $1, $2, $3, now() )
	`

	_, err := f.db.Exec(ctx, query,
		id,
		category.Name,
		helper.NewNullString(category.ParentID),
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *CategoriesRepo) GetByPKey(ctx context.Context, pkey *models.CategoriesPrimarKey) (*models.Categories, error) {

	var (
		id          sql.NullString
		name        sql.NullString
		parent_uuid sql.NullString
		created_at  sql.NullString
		updated_at  sql.NullString
	)

	query := `
		SELECT
			id,
			name,
			parent_uuid,
			created_at,
			updated_at
		FROM
			categorys 
		WHERE parent_uuid IS NULL AND id = $1

	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&name,
			&parent_uuid,
			&created_at,
			&updated_at,
		)

	if err != nil {
		return nil, err
	}

	resp := &models.Categories{
		Id:        id.String,
		Name:      name.String,
		ParentID:  parent_uuid.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}

	childQuery := `
		SELECT
			id,
			name,
			parent_uuid,
			created_at,
			updated_at
		FROM
			categorys 
		WHERE parent_uuid IS NULL AND id = $1
	`

	rows, err := f.db.Query(ctx, childQuery, resp.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return resp, nil
		}
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&id,
			&name,
			&parent_uuid,
			&created_at,
			&updated_at,
		)

		if err != nil {
			return nil, err
		}

		resp.ChildCategory = append(resp.ChildCategory, &models.Categories{
			Id:        id.String,
			Name:      name.String,
			ParentID:  parent_uuid.String,
			CreatedAt: created_at.String,
			UpdatedAt: updated_at.String,
		})
	}

	return resp, nil
}

// func (f *CategoriesRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {

// 	var (
// 		resp   = models.GetListCategoryResponse{}
// 		offset = " OFFSET 0"
// 		limit  = " LIMIT 5"
// 	)

// 	if req.Limit > 0 {
// 		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
// 	}

// 	if req.Offset > 0 {
// 		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
// 	}

// 	query := `
// 		SELECT
// 			COUNT(*) OVER(),
// 			id,
// 			name,
// 			parent_uuid,
// 			created_at,
// 			updated_at
// 		FROM
// 			categorys
// 		WHERE parent_uuid IS NULL AND deleted_at IS NULL
// 	`

// 	query += offset + limit

// 	rows, err := f.db.Query(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for rows.Next() {

// 		res := &models.CategoryByParent{}

// 		var (
// 			id          sql.NullString
// 			name        sql.NullString
// 			parent_uuid sql.NullString
// 			created_at  sql.NullString
// 			updated_at  sql.NullString
// 		)

// 		err := rows.Scan(
// 			&resp.Count,
// 			&id,
// 			&name,
// 			&parent_uuid,
// 			&created_at,
// 			&updated_at,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		res.Id = id.String
// 		res.Name = name.String
// 		res.ParentUUID = parent_uuid.String

// 		resp.Categorys = append(resp.Categorys, res)
// 	}

// 	for ind, category := range resp.Categorys {

// 		childQuery := `
// 			SELECT
// 				id,
// 				name,
// 				parent_uuid,
// 				created_at,
// 				updated_at
// 			FROM
// 				categorys
// 			WHERE parent_uuid =  $1 AND deleted_at IS NULL
// 		`

// 		childRows, err := f.db.Query(ctx, childQuery, category.Id)
// 		if err != nil {
// 			return nil, err
// 		}

// 		for childRows.Next() {

// 			var (
// 				id          sql.NullString
// 				name        sql.NullString
// 				parent_uuid sql.NullString
// 				created_at  sql.NullString
// 				updated_at  sql.NullString
// 			)

// 			err := childRows.Scan(
// 				&id,
// 				&name,
// 				&parent_uuid,
// 				&created_at,
// 				&updated_at,
// 			)

// 			if err != nil {
// 				return nil, err
// 			}

// 			category.ChildCategory = append(category.ChildCategory, &models.Category{
// 				Id:         id.String,
// 				Name:       name.String,
// 				ParentUUID: parent_uuid.String,
// 				CreatedAt:  created_at.String,
// 				UpdatedAt:  updated_at.String,
// 			})

// 		}

// 		resp.Categorys[ind] = category
// 	}

// 	return &resp, err
// }

func fetchCategories(categories []*models.Categories, parentUid *string) (result []*models.Categories) {
	result = make([]*models.Categories, 0)

	for _, category := range categories {
		if parentUid == nil {
			if category.ParentID == "" {
				category.ChildCategory = fetchCategories(categories, &category.ParentID)
				result = append(result, category)
			}
		} else {
			if category.ParentID != "" && *&category.ParentID == *parentUid {
				category.ChildCategory = fetchCategories(categories, &category.ParentID)
				result = append(result, category)
			}
		}
	}

	return
}

func (r *CategoriesRepo) Product2GetList(ctx context.Context, req *models.GetListCategoriesRequest) (*models.GetListCategoriesResponse, error) {

	resp := &models.GetListCategoriesResponse{}

	var (
		query               string
		categoryMap         = make(map[string]*models.Categories)
		cateogryByParentMap = make(map[string][]*models.Categories)
		parentCategory      = []*models.Categories{}
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			parent_id,
			created_at,
			updated_at
		FROM 
			categories
		WHERE deleted_at IS NULL
	`

	rows, err := r.db.Query(ctx, query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id       sql.NullString
			name     sql.NullString
			parentId sql.NullString
		)

		err = rows.Scan(
			&id,
			&name,
			&parentId,
		)

		if err != nil {
			return nil, err
		}

		category := &models.Categories{
			Id:       id.String,
			Name:     name.String,
			ParentID: parentId.String,
		}

		if category.ParentID == "" {
			parentCategory = append(parentCategory, category)
		}

	}

	for _, category := range parentCategory {

		categoryMap[category.Id] = category

		if category.ParentID == "" {
			category.ParentID = "parent"
		}

		if _, ok := cateogryByParentMap[category.ParentID]; !ok {
			cateogryByParentMap[category.ParentID] = append(cateogryByParentMap[category.ParentID], category)
		} else {
			cateogryByParentMap[category.ParentID] = append(cateogryByParentMap[category.ParentID], category)
		}
	}

	var getChild func([]*models.Categories) []*models.Categories
	getChild = func(categories []*models.Categories) []*models.Categories {
		resp = &models.GetListCategoriesResponse{}

		for _, category := range categories {
			var c = category

			subCategories := cateogryByParentMap[category.Id]
			if len(subCategories) > 0 {
				resp.Categories = append(resp.Categories, c)
				continue
			}

			children := getChild(subCategories)
			c.ChildCategory = children

			resp.Categories = append(resp.Categories, c)
		}

		return resp.Categories
	}

	resp.Categories = getChild(cateogryByParentMap["parent"])

	return resp, nil
}

func (f *CategoriesRepo) GetList(ctx context.Context, req *models.GetListCategoriesRequest) (*models.GetListCategoriesResponse, error) {

	var (
		resp   = models.GetListCategoriesResponse{}
		offset = ""
		limit  = ""
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			parent_uuid,
			created_at,
			updated_at
		FROM
			categorys 
		WHERE deleted_at IS NULL
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			categoryByParent models.Categories
			category         models.Categories

			id        sql.NullString
			parent_id sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&parent_id,
		)
		if err != nil {
			return nil, err
		}

		if parent_id.Valid {
			// bu child

		} else {
			// bu root
		}

		categoryByParent.Id = id.String
		category.ParentID = parent_id.String

		resp.Categories = append(resp.Categories, &models.Categories{
			Id:       id.String,
			ParentID: parent_id.String,
		})

	}

	return &resp, err
}

func (f *CategoriesRepo) Update(ctx context.Context, id string, req *models.UpdateCategories) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			categorys
		SET
			name = :name,
			parent_uuid = :parent_uuid,
			updated_at = now()
		WHERE id = :id AND deleted_at IS NULL
	`

	params = map[string]interface{}{
		"name":        req.Name,
		"parent_uuid": req.ParentID,
		"id":          id,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *CategoriesRepo) Delete(ctx context.Context, req *models.CategoriesPrimarKey) error {

	_, err := f.db.Exec(ctx, "UPDATE categorys SET deleted_at = now() WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
