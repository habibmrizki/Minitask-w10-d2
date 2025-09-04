package utils

import (
	"errors"
	"regexp"

	"github.com/habibmrizki/gin/internal/models"
)

func ValidateBody(body models.Body) error {
	if body.Id <= 0 {
		return errors.New("id harus diatas 0")
	}
	if len(body.Message) < 8 {
		return errors.New("panjang pesan harus diatas 8 karakter")
	}
	re, err := regexp.Compile("^[lLpPmMfF]$")
	if err != nil {
		return err
	}
	if !re.Match([]byte(body.Gender)) {
		return errors.New("gender harus berisikan huruf l, L, m, M, f, F, p, P")
	}
	return nil
}
