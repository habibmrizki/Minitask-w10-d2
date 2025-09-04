// package routers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jackc/pgx/v5/pgxpool"
// 	"github.com/kodacampmain/koda3_gin/internal/middlewares"
// 	"github.com/kodacampmain/koda3_gin/internal/models"
// )

// func InitRouter(db *pgxpool.Pool) *gin.Engine {
// 	router := gin.Default()
// 	router.Use(middlewares.MyLogger)
// 	router.Use(middlewares.CORSMiddleware)

// 	// config := cors.Config{
// 	// 	AllowOrigins: []string{"http://127.0.0.1:5500", "http://127.0.0.1:3001"},
// 	// 	AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
// 	// 	AllowHeaders: []string{"Authorization", "Content-Type"},
// 	// }

// 	// router.Use(cors.New(config))

// 	// setup routing
// 	InitPingRouter(router)
// 	InitDirectorRouter(router, db)
// 	InitProductRouter(router, db)

// 	// catch all route
// 	router.NoRoute(func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusNotFound, models.Response{
// 			Message: "Rute Salah",
// 			Status:  "Rute Tidak Ditemukan",
// 		})
// 	})

//		return router
//	}
package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	middlewares "github.com/habibmrizki/gin/internal/middlewars"
	"github.com/habibmrizki/gin/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.MyLogger)
	router.Use(middlewares.CORSMiddleware)

	// config := cors.Config{
	// 	AllowOrigins: []string{"http://127.0.0.1:5500", "http://127.0.0.1:3001"},
	// 	AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders: []string{"Authorization", "Content-Type"},
	// }

	// router.Use(cors.New(config))

	// setup routing
	initPingRouter(router)
	InitDirectorRouter(router, db)
	InitProductRouter(router, db)

	// catch all route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: "Rute Salah",
			Status:  "Rute Tidak Ditemukan",
		})
	})

	return router
}
