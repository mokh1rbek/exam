package storage_test

import (
	"context"
	"testing"

	"exam/models"

	faker "github.com/bxcodec/faker/v3"
	"github.com/google/go-cmp/cmp"
)

func TestCategoryCreate(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateCategories
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CreateCategories{
				Name: faker.Name(),
			},
			WantErr: false,
		},
		{
			Name: "case 2",
			Input: &models.CreateCategories{
				Name: faker.Name(),
			},
			WantErr: true,
		},
		{
			Name: "case 3",
			Input: &models.CreateCategories{
				Name: faker.Name(),
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			got, err := categoryRepo.Create(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			if got == "" {
				t.Errorf("%s: got: %v", tc.Name, err)
				return
			}
		})
	}

}

func TestCategoryGetById(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CategoriesPrimarKey
		Output  *models.Categories
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CategoriesPrimarKey{
				Id: "",
			},
			Output: &models.Categories{
				Id:        "",
				Name:      "",
				CreatedAt: "",
				UpdatedAt: "",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			got, err := categoryRepo.GetByPKey(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			comparer := cmp.Comparer(func(x, y models.Categories) bool {
				return x.Name == y.Name
			})

			if diff := cmp.Diff(tc.Output, got, comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}

func TestCategoryUpdate(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.UpdateCategories
		Output  *models.Categories
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.UpdateCategories{
				Name: "",
			},
			Output: &models.Categories{
				Name: "",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			_, err := categoryRepo.Update(
				context.Background(),
				"",
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			res, err := categoryRepo.GetByPKey(
				context.Background(),
				&models.CategoriesPrimarKey{
					Id: tc.Input.Id,
				},
			)

			comparer := cmp.Comparer(func(x, y models.Categories) bool {
				return x.Name == y.Name
			})

			if diff := cmp.Diff(tc.Output, res, comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}

func TestCategoryDelete(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CategoriesPrimarKey
		Want    string
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CategoriesPrimarKey{
				Id: "",
			},
			Want:    "no rows in result set",
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			err := categoryRepo.Delete(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			_, err = categoryRepo.GetByPKey(
				context.Background(),
				&models.CategoriesPrimarKey{
					Id: tc.Input.Id,
				},
			)

			comparer := cmp.Comparer(func(x, y string) bool {
				return x == y
			})

			if diff := cmp.Diff(tc.Want, err.Error(), comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}
