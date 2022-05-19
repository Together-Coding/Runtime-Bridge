package runtimes

import (
	"github.com/gin-gonic/gin"
	"github.com/together-coding/runtime-bridge/db"
	"gorm.io/gorm/clause"
	"net/http"
)

func Register(router *gin.RouterGroup) {
	router.GET("/available", SupportedLang)
}

type SupportedLangResp struct {
	Image    []RuntimeImage      `json:"image"`
	Language []SupportedLanguage `json:"language"`
}

// SupportedLang godoc
// @Summary  Return supported languages available
// @Description Return supported languages available
// @Tags     runtimes
// @Produce  json
// @Success  200 {object} SupportedLangResp
// @Router   /runtimes/available [get]
func SupportedLang(c *gin.Context) {
	var languages []SupportedLanguage
	db.DB.Preload(clause.Associations).Order("`order` ASC").Find(&languages)

	// Flatten runtime images. See #2
	images := make([]RuntimeImage, 0, 5)
	for _, _lang := range languages {
		for _, _image := range _lang.RuntimeImages {
			images = append(images, _image)
		}
	}

	c.JSON(http.StatusOK, SupportedLangResp{
		Image:    images,
		Language: languages,
	})
}
