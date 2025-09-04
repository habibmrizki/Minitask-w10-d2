package routers

import (
	"github.com/gin-gonic/gin"

	handlers "github.com/habibmrizki/gin/internal/handler"
	"github.com/habibmrizki/gin/internal/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitProductRouter(router *gin.Engine, db *pgxpool.Pool) {
	productRouter := router.Group("/products")
	productRepository := repositories.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productRepository)

	productRouter.POST("", productHandler.AddNewProduct)
	productRouter.PATCH("/:id", productHandler.UpdateProduct)
}
