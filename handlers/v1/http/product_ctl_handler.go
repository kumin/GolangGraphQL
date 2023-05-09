package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kumin/GolangGraphQL/entities"
	"github.com/kumin/GolangGraphQL/repos"
)

type ProductCtlHandler struct {
	productRepo repos.ProductRepo
}

func NewProductCtlHandler(
	productRepo repos.ProductRepo,
) *ProductCtlHandler {
	return &ProductCtlHandler{
		productRepo: productRepo,
	}
}

func (p *ProductCtlHandler) CreateProduct(ctx *gin.Context) {
	var product *entities.Product
	if err := ctx.BindJSON(&product); err != nil {
		return
	}
	prod, err := p.productRepo.CreateProduct(ctx, product)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
	}
	ctx.IndentedJSON(http.StatusOK, prod)
}

func (p *ProductCtlHandler) ListProducts(ctx *gin.Context) {
	page, _ := strconv.ParseInt(ctx.Query("page"), 10, 32)
	limit, _ := strconv.ParseInt(ctx.Query("limit"), 10, 32)
	prods, err := p.productRepo.ListProducts(ctx, int(page), int(limit))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
	}

	ctx.IndentedJSON(http.StatusOK, prods)
}
