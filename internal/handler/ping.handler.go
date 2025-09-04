// package handlers

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/habibmrizki/gin/internal/models"
// 	"github.com/habibmrizki/gin/internal/utils"
// )

// type PingHandler struct{}

// func NewPingHandler() *PingHandler {
// 	return &PingHandler{}
// }

// func (p *PingHandler) GetPing(ctx *gin.Context) {
// 	requestId := ctx.GetHeader("X-Request-ID")
// 	contentType := ctx.GetHeader("Content-Type")
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message":     "pong",
// 		"requestId":   requestId,
// 		"contentType": contentType,
// 	})
// }

// func (p *PingHandler) GetPingWithParam(ctx *gin.Context) {
// 	pingId := ctx.Param("id")
// 	param2 := ctx.Param("param2")
// 	q := ctx.Query("q")
// 	log.Printf("%s, %s, %s\n", pingId, param2, q)
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"param":  pingId,
// 		"param2": param2,
// 		"q":      q,
// 	})
// }

//	func (p *PingHandler) PostPing(ctx *gin.Context) {
//		body := models.Body{}
//		// data-binding, memasukkan body ke dalam variabel golang
//		if err := ctx.ShouldBind(&body); err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error":   err.Error(),
//				"success": false,
//			})
//			return
//		}
//		if err := utils.ValidateBody(body); err != nil {
//			ctx.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		log.Println(body)
//		ctx.JSON(http.StatusOK, gin.H{
//			"success": true,
//			"body":    body,
//		})
//	}
package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/habibmrizki/gin/internal/models"
	"github.com/habibmrizki/gin/internal/utils"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (p *PingHandler) GetPing(ctx *gin.Context) {
	requestId := ctx.GetHeader("X-Request-ID")
	contentType := ctx.GetHeader("Content-Type")
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "pong",
		"requestId":   requestId,
		"contentType": contentType,
	})
}

func (p *PingHandler) GetPingWithParam(ctx *gin.Context) {
	pingId := ctx.Param("id")
	param2 := ctx.Param("param2")
	q := ctx.Query("q")
	log.Printf("%s, %s, %s\n", pingId, param2, q)
	ctx.JSON(http.StatusOK, gin.H{
		"param":  pingId,
		"param2": param2,
		"q":      q,
	})
}

func (p *PingHandler) PostPing(ctx *gin.Context) {
	body := models.Body{}
	// data-binding, memasukkan body ke dalam variabel golang
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	if err := utils.ValidateBody(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(body)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"body":    body,
	})
}

func (p *PingHandler) PatchPing(ctx *gin.Context) {
	id := ctx.Param("id")
	body := models.Body{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.ValidateBody(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data for ID " + id + " updated successfully",
		"data":    body,
	})
}
