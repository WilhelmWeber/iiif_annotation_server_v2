package controllers

import (
	"fmt"
	"net/http"

	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/model"
	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/repository"
	"github.com/gin-gonic/gin"
)

type ManifestController struct {
	Service *repository.ManifestService
}

func NewManifestController(s *repository.ManifestService) *ManifestController {
	return &ManifestController{s}
}

/*GET service/manifest */
func GetAllManifest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome login user!",
	})
}

/*POST /service/manifest */
func (m *ManifestController) CreateManifest(c *gin.Context) {
	var manifest model.Manifest
	if err := c.ShouldBindJSON(&manifest); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error. Please Inform it of Administrator",
		})
		return
	}
	//同じURLの資料が登録されていればはじく
	if isExist := m.Service.GetByUrl(manifest.Url); isExist.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This IIIF Manifest is already registered. Try another one.",
		})
		return
	}

	if err := m.Service.Create(&manifest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error. Please Inform it of Administrator",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "manifest created with success",
	})
}
