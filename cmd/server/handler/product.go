package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Laura-2950/API-Go-Products.git/API-Go-Products/internal/domain"
	"github.com/Laura-2950/API-Go-Products.git/API-Go-Products/internal/product"
	"github.com/Laura-2950/API-Go-Products.git/API-Go-Products/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService product.IService
}

func (h *ProductHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	prodFounded, err := h.ProductService.GetProductBy(id)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &prodFounded)
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	prodFounded, err := h.ProductService.GetAllProducts()
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, prodFounded)
}

func (h *ProductHandler) NewProduct(ctx *gin.Context) {
	var product *domain.Product

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid product"})
		return
	}
	valid, err := validateEmptys(product)
	if !valid {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	valid, err = validateExpiration(product.Expiration)
	if !valid {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	prodFounded, err := h.ProductService.CreateNewProducts(product)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, prodFounded)
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	erro := h.ProductService.DeleteProduct(id)
	if erro != nil {
		if errApi, ok := erro.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, erro)
		return
	}

	ctx.JSON(http.StatusOK, "product removed successfully")
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(product *domain.Product) (bool, error) {
	switch {
	case product.Name == "" || product.CodeValue == "" || product.Expiration == "":
		return false, errors.New("fields can't be empty")
	case product.Quantity <= 0 || product.Price <= 0:
		if product.Quantity <= 0 {
			return false, errors.New("quantity must be greater than 0")
		}
		if product.Price <= 0 {
			return false, errors.New("price must be greater than 0")
		}
	}
	return true, nil
}

// validateExpiration valida que la fecha de expiracion sea valida
func validateExpiration(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid expiration date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid expiration date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) || (list[1] < 1 || list[1] > 12) || (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}
