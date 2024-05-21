package controller

import (
	"crud-golang/api/dto"
	"crud-golang/internal/infra/database"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type CategoryController struct {
	cr database.CategoryRepo
}

func NewCategoryController(categoryRepository database.CategoryRepo) *CategoryController {
	return &CategoryController{
		cr: categoryRepository,
	}
}

func (c *CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryDTO dto.CategoryDTO

	if err := json.NewDecoder(r.Body).Decode(&categoryDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	category := categoryDTO.ParseToEntity()

	category, err := c.cr.Save(category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.FromEntity(*category))
}

func (c *CategoryController) GetCategory(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	category, err := c.cr.FindByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.FromEntity(*category))
}

func (c *CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	err = c.cr.Delete(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	var categoryDTO dto.CategoryDTO
	if err := json.NewDecoder(r.Body).Decode(&categoryDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	category := categoryDTO.ParseToEntity()
	category.ID = uint(id)
	_, err = c.cr.Update(category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *CategoryController) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	var page int
	var size = 10
	var err error
	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				map[string]string{"error": err.Error()},
			)
			return
		}
	}
	if sizeStr := r.URL.Query().Get("size"); sizeStr != "" {
		size, err = strconv.Atoi(sizeStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				map[string]string{"error": err.Error()},
			)
			return
		}
	}

	categories, err := c.cr.FindAll(page, size)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}
