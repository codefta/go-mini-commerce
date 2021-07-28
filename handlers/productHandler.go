package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fathisiddiqi/go-mini-commerce/models"
	"github.com/fathisiddiqi/go-mini-commerce/utils"
	"github.com/gorilla/mux"
)

func (a *API) GetProducts(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if limit < 1 {
		limit = 5
	}

	products, err := a.tempStorage.GetAllTempProducts("productsAll")
	if err != nil || len(products) > 0 {
		products, err = a.storage.GetProducts(limit)
		if err != nil {
			utils.WriteResp(w, utils.NewInternalServerErrorResp(err))
			return 
		}

		err = a.tempStorage.SetAllTempProducts("productsAll", products)
		if err != nil {
			utils.WriteResp(w, utils.NewInternalServerErrorResp(err))
			return 
		}
	}


	utils.WriteResp(w, utils.NewSuccessResp(http.StatusOK, map[string]interface{}{
		"product": products,
	}))
}

func (a *API) GetProductById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, err := a.storage.CheckProductIfExist(id)
	if err != nil {
		utils.WriteResp(w, utils.NewNotFoundErrorResp(err.Error()))
		return
	}

	idString := strconv.Itoa(id)

	product, err := a.tempStorage.GetTempProduct("product"+idString)
	if err != nil || product == nil {
		product, err = a.storage.GetProductById(id)
		if err != nil {
			utils.WriteResp(w, utils.NewInternalServerErrorResp(err))
			return
		}
		
		err = a.tempStorage.SetTempProduct("product"+idString, *product)
		if err != nil {
			utils.WriteResp(w, utils.NewInternalServerErrorResp(err))
			return 
		}
	}

	utils.WriteResp(w, utils.NewSuccessResp(http.StatusOK, map[string]interface{}{
		"product": product,
	}))
}

func (a *API) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var body models.ProductForm

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.WriteResp(w, utils.NewBadRequestErrorResp(err.Error()))
		return
	}

	err = a.validate.ProductValidator(body)
	if err != nil {
		var data interface{} = utils.ValidatorError(err)
		fmt.Println(data)
		utils.WriteResp(w, utils.NewValidationErrorResp(data))
		return
	}

	product := models.Product{
		ProductName: body.ProductName,
		Description: body.Description,
		Price: body.Price,
		Categories: body.Categories,
	}
	
	productCreated, err := a.storage.PostProduct(product)
	if err != nil {
		utils.WriteResp(w, utils.NewBadRequestErrorResp(err.Error()))
		return
	}

	a.tempStorage.DeleteAllTempProductsData()

	utils.WriteResp(w, utils.NewSuccessResp(http.StatusCreated, map[string]interface{}{
		"new_product": productCreated,
	}))
}

func (a *API) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var body models.ProductForm
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.WriteResp(w, utils.NewBadRequestErrorResp(err.Error()))
		return
	}

	err = a.validate.ProductValidator(body)
	if err != nil {
		var data interface{} = utils.ValidatorError(err)
		utils.WriteResp(w, utils.NewValidationErrorResp(data))
		return
	}

	if err != nil {
		utils.WriteResp(w, utils.NewBadRequestErrorResp(err.Error()))
		return
	}

	_, err = a.storage.CheckProductIfExist(id)
	if err != nil {
		utils.WriteResp(w, utils.NewNotFoundErrorResp(err.Error()))
		return
	}

	product := models.Product{
		ProductName: body.ProductName,
		Description: body.Description,
		Price: body.Price,
		Categories: body.Categories,
	}
	
	productUpdated, err := a.storage.UpdateProduct(id, product)
	if err != nil {
		utils.WriteResp(w, utils.NewInternalServerErrorResp(err))
		return
	}

	a.tempStorage.DeleteAllTempProductsData()

	utils.WriteResp(w, utils.NewSuccessResp(http.StatusOK, map[string]interface{}{
		"updated_product": productUpdated, 
	}))
}

func (a *API) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, err := a.storage.CheckProductIfExist(id)
	if err != nil {
		utils.WriteResp(w, utils.NewNotFoundErrorResp(err.Error()))
		return
	}

	err = a.storage.DeleteProduct(id)
	if err != nil {
		utils.WriteResp(w, utils.NewInternalServerErrorResp(err))
		return
	}

	a.tempStorage.DeleteAllTempProductsData()

	utils.WriteResp(w, utils.NewSuccessResp(http.StatusNoContent, map[string]interface{}{}))
}