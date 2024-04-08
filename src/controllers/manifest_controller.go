package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/libs"
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

/*DELETE /service/manifest */
func (m *ManifestController) DeleteManifest(c *gin.Context) {
	var manifest model.Manifest
	if err := c.ShouldBindJSON(&manifest); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error. Please Inform it of Administrator",
		})
		return
	}
	if err := m.Service.Delete(&manifest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error. Please Inform it of Administrator",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "manifest deleted with success",
	})
}

//以下、middlewareを通さなくてもリクエストできる
/*GET /presentation/:manifestId/manifest.json */
func (m *ManifestController) PresentManifest(c *gin.Context) {
	_manifestId := c.Param("manifestId")
	manifestId, _ := strconv.Atoi(_manifestId)

	manifest, err := m.Service.GetById(uint(manifestId))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error. Please Inform it of Administrator",
		})
		return
	}
	if manifest.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}

	//外部にRequestを投げて元のManifestを取得する
	body, ver, err := libs.RequestManifest(manifest.Url)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error. Please Inform it of Administrator",
		})
		return
	}
	switch ver {
	case 2:
		var ManifestBody libs.ManifestV2
		json.Unmarshal(body, &ManifestBody)
		//JSONの書き換え
		ManifestBody.Id = c.Request.Host + c.Request.URL.Path
		for i, canvas := range ManifestBody.Sequences[0].Canvases {
			canvas["otherContent"] = map[string]interface{}{
				// {BASE_URI}/presentation/:manifest_id/:canvas_num/annolist.json
				"@id":   c.Request.Host + "/presentation/" + _manifestId + "/" + strconv.Itoa(i) + "/annnolist.json",
				"@type": "sc:AnnotationList",
			}
		}
		manifest_json := libs.V2toJSON(&ManifestBody)
		c.JSON(http.StatusOK, manifest_json)
	case 3:
		var ManifestBody libs.ManifestV3
		json.Unmarshal(body, &ManifestBody)
		ManifestBody.Id = c.Request.Host + c.Request.URL.Path
		for i, item := range ManifestBody.Items {
			var annotations []map[string]interface{}
			anno := map[string]interface{}{
				"id":   c.Request.Host + "/presentation/" + _manifestId + "/" + strconv.Itoa(i) + "/annolist.json",
				"type": "AnnotationPage",
			}
			annotations = append(annotations, anno)
			item.Annotations = annotations
		}
		manifest_json := libs.V3toJSON(&ManifestBody)
		c.JSON(http.StatusOK, manifest_json)
	default:
		panic("Unreachable")
	}
}
