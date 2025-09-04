// package routers

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// func InitStudentRouter(router *gin.Engine, db *pgxpool.Pool) {
// 	studentRouter := router.Group("/students")
// 	sr := repositories.NewStudentRepository(db)
// 	sh := handlers.NewStudentHandler(sr)

//		studentRouter.GET("", sh.GetStudent)
//	}
package routers

import (
	"github.com/gin-gonic/gin"

	handlers "github.com/habibmrizki/gin/internal/handler"
	"github.com/habibmrizki/gin/internal/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDirectorRouter(router *gin.Engine, db *pgxpool.Pool) {
	directorRouter := router.Group("/directors")
	sr := repositories.NewdirectorRepository(db)
	sh := handlers.NewdirectorHandler(sr)

	directorRouter.GET("", sh.Getdirector)
}
