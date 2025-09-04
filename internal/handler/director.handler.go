// package handlers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type directorHandler struct {
// 	sr *repositories.directorRepository
// }

// func NewdirectorHandler(sr *repositories.directorRepository) *directorHandler {
// 	return &directorHandler{
// 		sr: sr,
// 	}
// }

// func (s *directorHandler) Getdirector(ctx *gin.Context) {
// 	page, err := strconv.Atoi(ctx.Query("page"))
// 	if err != nil {
// 		page = 1
// 	}
// 	limit := 4
// 	offset := (page - 1) * limit

// 	directors, err := s.sr.GetdirectorData(ctx.Request.Context(), offset, limit)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"success": false,
// 			"data":    directors,
// 		})
// 		return
// 	}

// 	if len(directors) == 0 {
// 		ctx.JSON(http.StatusNotFound, gin.H{
// 			"success": false,
// 			"data":    []any{},
// 			"page":    page,
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"data":    directors,
// 		"page":    page,
// 	})
// }

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/habibmrizki/gin/internal/repositories"
)

type directorHandler struct {
	sr *repositories.DirectorRepository
}

func NewdirectorHandler(sr *repositories.DirectorRepository) *directorHandler {
	return &directorHandler{
		sr: sr,
	}
}

func (s *directorHandler) Getdirector(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit := 4
	offset := (page - 1) * limit

	directors, err := s.sr.GetdirectorData(ctx.Request.Context(), offset, limit)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal server error",
		})
		return
	}

	if len(directors) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    []any{},
			"page":    page,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    directors,
		"page":    page,
	})
}
