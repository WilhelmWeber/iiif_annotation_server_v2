package libs

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*元のマニフェストJSONを取得してくる関数*/
/*IIIFのバージョン(uint)とJSON文字列[]byteを返す(戻り先でunmarshallする)*/
func RequestManifest(url string) ([]byte, uint, error) {
	var iiif_version uint
	var mapped_json map[string]interface{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		resp.Body.Close()
		return nil, 0, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &mapped_json)
	if mapped_json["@context"] == "http://iiif.io/api/presentation/3/context.json" {
		iiif_version = 3
	} else {
		iiif_version = 2
	}

	return body, iiif_version, nil
}

/*Ver2.0 or 2.1のManifest構造体をJSONに変換*/
func V2toJSON(body *ManifestV2) gin.H {
	var sequences []gin.H
	for _, seq := range body.Sequences {
		h := gin.H{
			"@id":              seq.Id,
			"@context":         seq.Context,
			"@type":            seq.Type,
			"label":            seq.Label,
			"viewingDirection": seq.ViewingDirection,
			"viewingHint":      seq.ViewingHint,
			"thumbnail":        seq.Thumbnail,
			"canvases":         seq.Canvases,
			"structures":       seq.Structures,
		}
		sequences = append(sequences, h)
	}
	json := gin.H{
		"@context":    body.Context,
		"@id":         body.Id,
		"@type":       body.Type,
		"label":       body.Label,
		"description": body.Description,
		"license":     body.License,
		"attribution": body.Attribution,
		"metadata":    body.Metadata,
		"logo":        body.Logo,
		"sequences":   sequences,
		"thumbnail":   body.Thumbnail,
	}
	return json
}

func V3toJSON(body *ManifestV3) gin.H {
	var items []gin.H
	for _, item := range body.Items {
		item_json := gin.H{
			"@context":    item.Context,
			"id":          item.Id,
			"type":        item.Type,
			"label":       item.Label,
			"height":      item.Height,
			"width":       item.Width,
			"items":       item.Items,
			"annotations": item.Annotations,
		}
		items = append(items, item_json)
	}
	json := gin.H{
		"@context":         body.Context,
		"id":               body.Id,
		"type":             body.Type,
		"label":            body.Label,
		"metadata":         body.Metadata,
		"summary":          body.Summary,
		"rights":           body.Rights,
		"thumbnail":        body.Thumbnail,
		"provider":         body.Provider,
		"viewingDirection": body.ViewingDirection,
		"behavior":         body.Behavior,
		"items":            items,
	}
	return json
}
