// package repositories

// import (
// 	"context"
// 	"log"

// 	"github.com/habibmrizki/gin/internal/models"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type directorRepository struct {
// 	db *pgxpool.Pool
// }

// func NewdirectorRepository(db *pgxpool.Pool) *directorRepository {
// 	return &directorRepository{
// 		db: db,
// 	}
// }

// func (s *directorRepository) GetdirectorData(reqContext context.Context, offset, limit int) ([]models.Director, error) {
// 	sql := "SELECT id, name FROM director LIMIT $2 OFFSET $1"
// 	values := []any{offset, limit}
// 	rows, err := s.db.Query(reqContext, sql, values...)
// 	if err != nil {
// 		log.Println("Internal server error: ", err.Error())

// 		return []models.Director{}, err
// 	}
// 	defer rows.Close()
// 	var directors models.Director
// 	// membaca rows/record
// 	for rows.Next() {
// 		var director models.Director
// 		if err := rows.Scan(&director.Id, &director.Name); err != nil {
// 			log.Println("Scan error, ", err.Error())
// 			return []models.Director{}, err
// 		}
// 		directors = append(directors, director)
// 	}

//		return directors, nil
//	}
package repositories

import (
	"context"
	"log"

	"github.com/habibmrizki/gin/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DirectorRepository struct {
	db *pgxpool.Pool
}

func NewdirectorRepository(db *pgxpool.Pool) *DirectorRepository {
	return &DirectorRepository{
		db: db,
	}
}

func (s *DirectorRepository) GetdirectorData(reqContext context.Context, offset, limit int) ([]models.Director, error) {
	sql := "SELECT id, name FROM director LIMIT $2 OFFSET $1"
	values := []any{offset, limit}
	rows, err := s.db.Query(reqContext, sql, values...)
	if err != nil {
		log.Println("Internal server error: ", err.Error())
		return []models.Director{}, err
	}
	defer rows.Close()

	var directors []models.Director

	// membaca rows/record
	for rows.Next() {
		var director models.Director
		if err := rows.Scan(&director.Id, &director.Name); err != nil {
			log.Println("Scan error, ", err.Error())
			return []models.Director{}, err
		}
		directors = append(directors, director)
	}

	return directors, nil
}
