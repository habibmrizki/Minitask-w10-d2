package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	// _ "github.com/joho/godotenv/autoload"
)

type Director struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"nama_director"`
}

type Product struct {
	Name    string `json:"name" binding:"required"`
	PromoId *int   `json:"promo_id"`
	Price   int    `json:"price,omitempty" binding:"required"`
	Id      int    `json:"id,omitempty"`
}

func main() {
	// milih load salah satu antara autoload atua manual
	// manual load env
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to loadn env\ncaiuse", err.Error())
		return
	}

	// inisialisai db
	// psql string
	// postgress://username:password@host:port/namadb
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")
	connstring := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(connstring)

	db, err := pgxpool.New(context.Background(), connstring)

	if err != nil {
		log.Println("Failed to connet to database\nCause: ")
		return
	}
	defer db.Close()

	// testing koneksi db
	if err := db.Ping(context.Background()); err != nil {
		log.Println("Ping to DB failed\nCause: ", err.Error())
		return
	}
	log.Println("DB connected")

	// inisialsisai engine gin dengan opsi default
	// bisa menggunakan gin.New bednaya nggak ada logger
	router := gin.Default()
	// setup roating

	router.Use(MyLogger)
	router.Use(CORSMiddleware)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the API!",
		})
	})

	log.Println(os.Getenv("DBUSER"))
	// router.GET("/ping", func(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, gin.H{
	// "message": "pong",
	// })
	// })
	// iini pakai getheader
	router.GET("/ping", func(ctx *gin.Context) {
		requestId := ctx.GetHeader("X-Request-ID")
		contentType := ctx.GetHeader("Content-Type")
		ctx.JSON(http.StatusOK, gin.H{
			"message":     "pong",
			"requestId":   requestId,
			"contentType": contentType,
		})
	})

	// BELAJAR PATH PARAM DAN QUERY PARAM
	// UNUTK PENAMAAN PERAN DAN QUERY DI GOLANG DAN POSTMAN BOLEH/BISA BERBEDA ASALKAN
	// BSIA JUGA MENGGUNAKAN MIDDLEWARE DIDALAM METHOD, URUTANYA,
	// ENDPOINT -> MIDDLEWARE -> CONTROLLER/HANDLE
	router.GET("/ping/:id/:param2", func(ctx *gin.Context) {
		pingID := ctx.Param("id")
		param2 := ctx.Param("param2")
		q := ctx.Query("q")
		ctx.JSON(http.StatusOK, gin.H{
			"param":  pingID,
			"param2": param2,
			"q":      q,
		})
	})

	// mencoba 2
	// router.GET("/data/:userID/:objectID", func(ctx *gin.Context) {
	// // Ambil nilai dari path params
	// userID := ctx.Param("userID")
	// objectID := ctx.Param("objectID")

	// // Ambil nilai dari query param (opsional)
	// queryParam := ctx.Query("q")

	// // Kirim respons JSON yang berisi semua nilai yang diambil
	// ctx.JSON(http.StatusOK, gin.H{
	// "message": "Request berhasil!",
	// "user_id": userID,
	// "object_id": objectID,
	// "query_string": queryParam,
	// })
	// })

	// jalankan engine gin
	// buat window
	// router.Run("localhost:8080")
	// listen and serve on 0.0.0.0:8080
	// cat all route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "invalid input data",
		})
	})

	router.GET("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "status bad ",
		})
	})

	// router.GET("/user/:name", func(ctx *gin.Context) {
	// 	name := ctx.Param("name")
	// 	ctx.JSON(http.StatusOK, Response{
	// 		Success: true,
	// 		Name:    name,
	// 		Message: "user found successfully " + name,
	// 	})
	// })
	// router.GET("/search", func(ctx *gin.Context) {
	// 	query := ctx.Query("query")

	// 	if query == "" {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Query parameter 'query' is missing",
	// 		})
	// 		return // Penting untuk menghentikan eksekusi
	// 	}

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"search_query": query,
	// 	})
	// })

	router.POST("/ping", func(ctx *gin.Context) {
		body := Body{}
		// data bindiing, memasukkan ke dalam variable golang
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

		if err := ValidateBody(body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"body":    body,
		})

		// router.PATCH("/ping:id", func(ctx *gin.Context) {
		// body := Body{}
		// if err := ctx.ShouldBind(&body); err != nil {
		// ctx.JSON(http.StatusInternalServerError, gin.H{
		// "error": err.Error(),
		// "success": false,
		// })
		// return
		// }
		// if err := ValidateBody(body); err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{
		// "error": err.Error(),
		// })
		// return
		// }
		// ctx.JSON(http.StatusOK, gin.H{
		// "success": true,
		// "body": body,
		// })
		// })

	})

	router.PATCH("/ping/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		body := Body{}
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := ValidateBody(body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Data for ID " + id + " updated successfully",
			"data":    body,
		})
	})

	// connection db
	router.GET("/movie", func(ctx *gin.Context) {
		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil {
			page = 1
		}
		limit := 2
		offset := (page - 1) * limit
		sql := "SELECT id, name FROM director LIMIT $2 OFFSET $1"
		values := []any{offset, limit}
		rows, err := db.Query(ctx.Request.Context(), sql, values...)
		if err != nil {
			log.Println("Internal server error: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"data":    []any{},
			})
			return
		}
		defer rows.Close()
		var directors []Director
		// membaca rows/record
		for rows.Next() {
			var director Director
			if err := rows.Scan(&director.Id, &director.Name); err != nil {
				log.Println("Scan error, ", err.Error())
				return
			}
			directors = append(directors, director)
		}

		if (len(directors)) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"data":    []any{},
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    directors,
		})

	})

	// market
	// router.POST("/products", func(ctx *gin.Context) {
	// var body Product
	// if err := ctx.ShouldBind(&body); err != nil {
	// log.Println("Internal server error: ", err.Error())
	// ctx.JSON(http.StatusInternalServerError, gin.H{
	// "success": false,
	// "data": []any{},
	// })
	// return
	// }
	// sql := "INSERT INTO Product (name, promo_id, price) VALUES ($1,$2,$3) RETURNING id, name"
	// values := []any{body.Name, body.PromoId, body.Price}
	// var newProduct Product
	// if err := db.QueryRow(ctx.Request.Context(), sql, values...).Scan(&newProduct.Id, &newProduct.Name); err != nil {
	// ctx.JSON(http.StatusInternalServerError, gin.H{
	// "success": false,
	// "message": "Failed to create product due to server error",
	// "error": err.Error(),
	// })
	// return
	// }
	// ctx.JSON(http.StatusCreated, gin.H{
	// "success": true,
	// "data": newProduct,
	// })
	// })
	router.POST("/products", func(ctx *gin.Context) {
		var body Product
		if err := ctx.ShouldBindJSON(&body); err != nil {
			log.Printf("Database query failed: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to create product due to server error",
				"error":   err.Error(),
			})
			return
		}

		sql := "INSERT INTO product (name, promo_id, price) VALUES ($1, $2, $3) RETURNING id, name"

		var newProduct Product

		if err := db.QueryRow(ctx.Request.Context(), sql, body.Name, body.PromoId, body.Price).Scan(&newProduct.Id, &newProduct.Name); err != nil {
			log.Printf("Bad request: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid request body",
				"error":   err.Error(),
			})
			return
		}

		// Jika semua berhasil, kembalikan 201 Created
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"data":    newProduct,
		})
	})

	// router.POST("/promos", func(ctx *gin.Context) {
	// var body Product

	// if err := ctx.ShouldBindJSON(&body); err != nil {
	// log.Printf("Bad request: %v", err)
	// ctx.JSON(http.StatusBadRequest, gin.H{
	// "success": false,
	// "message": "Invalid request body",
	// "error": err.Error(),
	// })
	// return
	// }

	// sql := "INSERT INTO promos (name, discount) VALUES ($1, $2) RETURNING id, name, discount"

	// var newProduct Product

	// if err := db.QueryRow(ctx.Request.Context(), sql, body.Name, body.PromoId, body.Price).Scan(&newProduct.Id, &newProduct.Name); err != nil {
	// log.Printf("Database query failed: %v", err)
	// ctx.JSON(http.StatusInternalServerError, gin.H{
	// "success": false,
	// "message": "Failed to create product due to server error",
	// "error": err.Error(),
	// })
	// return
	// }

	// // Jika semua berhasil, kembalikan 201 Created
	// ctx.JSON(http.StatusCreated, gin.H{
	// "success": true,
	// "data": newProduct,
	// })
	// })

	router.Run()
}

func MyLogger(ctx *gin.Context) {
	log.Printf("start")
	start := time.Now()
	// Next digunakan untuk lanjut ke middleware/handler berikutnya
	ctx.Next()
	duration := time.Since(start)
	log.Printf("Durasi Request: %d\n", duration.Microseconds())
}

func CORSMiddleware(ctx *gin.Context) {
	// memasangkan header-header CORS
	// setup whitelist origin

	whitelist := []string{"http://127.0.0.1:5000", "http://127.0.0.1:3001"}
	origin := ctx.GetHeader("Origin")

	if slices.Contains(whitelist, origin) {
		ctx.Header("Access-Control-Allow-Origin", origin)
	}

	// header unukt prflight cors
	ctx.Header("Access-Control-Allow-Methods", "GET")
	ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")

	// tangani apabila bertemu preflight
	if ctx.Request.Method == http.MethodOptions {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Name    string
}

type Body struct {
	Id      int    `json:"id" binding:"required"`
	Message string `json:"msg"`
	Gender  string `json:"gender"`
}

func ValidateBody(body Body) error {
	if body.Id <= 0 {
		return errors.New("id tidak boleh dibawah 0")
	}

	if len(body.Message) < 8 {
		return errors.New("panjang pesan harus diatas 8 karakter")
	}

	re, err := regexp.Compile("^[lLpPmMfF]$")
	if err != nil {
		return err
	}

	if !re.Match([]byte(body.Gender)) {
		return errors.New("gender harus berisikan hruuf l, L, m, M, f, F, p, P")
	}
	return nil
}
