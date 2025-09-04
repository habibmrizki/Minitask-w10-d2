// package handlers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/habibmrizki/gin/internal/models"
// )

// type ProductHandler struct {
// 	pr *repositories.ProductRepository
// }

// func NewProductHandler(pr *repositories.ProductRepository) *ProductHandler {
// 	return &ProductHandler{
// 		pr: pr,
// 	}
// }

// func (p *ProductHandler) AddNewProduct(ctx *gin.Context) {
// 	var body models.Product
// 	if err := ctx.ShouldBind(&body); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error":   err.Error(),
// 			"success": false,
// 		})
// 		return
// 	}

// 	newProduct, err := p.pr.AddNewProduct(ctx.Request.Context(), body)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"success": false,
// 		})
// 		return
// 	}

// 	// ctag, err := p.pr.InsertNewProduct(ctx.Request.Context(), body)
// 	// if err != nil {
// 	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
// 	// 		"success": false,
// 	// 	})
// 	// 	return
// 	// }
// 	// if ctag.RowsAffected() == 0 {
// 	// 	ctx.JSON(http.StatusConflict, gin.H{
// 	// 		"success": false,
// 	// 	})
// 	// 	return
// 	// }

//		ctx.JSON(http.StatusCreated, gin.H{
//			"success": true,
//			"data":    newProduct,
//		})
//	}
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/habibmrizki/gin/internal/models"
	"github.com/habibmrizki/gin/internal/repositories"
)

type ProductHandler struct {
	pr *repositories.ProductRepository
}

func NewProductHandler(pr *repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{
		pr: pr,
	}
}

func (p *ProductHandler) AddNewProduct(ctx *gin.Context) {
	var body models.Product
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	newProduct, err := p.pr.AddNewProduct(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create product",
		})
		return
	}

	// ctag, err := p.pr.InsertNewProduct(ctx.Request.Context(), body)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"success": false,
	// 	})
	// 	return
	// }
	// if ctag.RowsAffected() == 0 {
	// 	ctx.JSON(http.StatusConflict, gin.H{
	// 		"success": false,
	// 	})
	// 	return
	// }

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    newProduct,
	})
}

func (p *ProductHandler) UpdateProduct(ctx *gin.Context) {
	// Ambil ID dari URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid product ID",
		})
		return
	}

	var body models.Product
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Panggil repository untuk memperbarui produk
	updatedProduct, err := p.pr.UpdateProduct(ctx.Request.Context(), id, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update product",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    updatedProduct,
	})
}
