package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/janainamai/study-api-go/internal/dto"
	"github.com/janainamai/study-api-go/internal/entity"
	"github.com/janainamai/study-api-go/internal/infra/database"
	entityPkg "github.com/janainamai/study-api-go/pkg/entity"
)

type ProductHandler struct {
	productDB database.ProductDatabaseInterface
}

func NewProductHandler(productDB database.ProductDatabaseInterface) *ProductHandler {
	return &ProductHandler{
		productDB: productDB,
	}
}

// Create product
// @Summary 	Create a product
// @Description Create a product
// @Tags 		products
// @Accept 		json
// @Produce		json
// @Param 		request body dto.CreateProductInput true "product request"
// @Success 	201
// @Failure 	500 {object} Error
// @Router 		/products [post]
// @Security	ApiKeyAuth
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	product, err := entity.NewProduct(input.Name, input.Price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = h.productDB.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Get product by ID
// @Summary 	Get a product by ID
// @Description Get a product by ID
// @Tags 		products
// @Accept 		json
// @Produce		json
// @Param 		id path string true "product id" Format(uuid)
// @Success 	200 {object} entity.Product
// @Success 	404 {object} Error
// @Failure 	500 {object} Error
// @Router 		/products/{id} [get]
// @Security	ApiKeyAuth
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: "id required"}
		json.NewEncoder(w).Encode(err)
		return
	}

	product, err := h.productDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: "product not found"}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Get all products
// @Summary 	Get all products
// @Description Get all products
// @Tags 		products
// @Accept 		json
// @Produce		json
// @Param 		page query string false "page number"
// @Param 		limit query string false "limit"
// @Success 	200 {array} entity.Product
// @Success 	404 {object} Error
// @Failure 	500 {object} Error
// @Router 		/products [get]
// @Security	ApiKeyAuth
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limit := r.URL.Query().Get("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.productDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Update a product
// @Summary 	Update product by ID
// @Description Update product by ID
// @Tags 		products
// @Accept 		json
// @Produce		json
// @Param 		id path string true "product id" Format(uuid)
// @Param 		request body dto.CreateProductInput true "product request"
// @Success 	200
// @Success 	400 {object} Error
// @Success 	404 {object} Error
// @Failure 	500 {object} Error
// @Router 		/products/{id} [put]
// @Security	ApiKeyAuth
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: "id required"}
		json.NewEncoder(w).Encode(err)
		return
	}

	product, err := h.productDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: "product not found"}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = h.productDB.Update(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Delete product by ID
// @Summary 	Delete a product by ID
// @Description Delete a product by ID
// @Tags 		products
// @Accept 		json
// @Produce		json
// @Param 		id path string true "product id" Format(uuid)
// @Success 	200
// @Success 	404 {object} Error
// @Failure 	500 {object} Error
// @Router 		/products/{id} [delete]
// @Security	ApiKeyAuth
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: "id required"}
		json.NewEncoder(w).Encode(err)
		return
	}

	product, err := h.productDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: "product not found"}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = h.productDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
