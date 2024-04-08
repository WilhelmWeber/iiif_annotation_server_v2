package libs

/*manifest形式定義ファイル*/

//Presentation API ver2.0
type ManifestV2 struct {
	Context     string                   `json:"@context"`
	Id          string                   `json:"@id"`
	Type        string                   `json:"@type"`
	Label       string                   `json:"label"`
	Description string                   `json:"description"`
	License     string                   `json:"license"`
	Attribution string                   `json:"attribution"`
	Metadata    []map[string]interface{} `json:"metadata"`
	Logo        string                   `json:"logo"`
	Sequences   []SequenceV2             `json:"sequences"`
	Thumbnail   interface{}              `json:"thumbnail"`
}

type SequenceV2 struct {
	Context          string                   `json:"@context"`
	Id               string                   `json:"@id"`
	Type             string                   `json:"@type"`
	Label            string                   `json:"label"`
	ViewingDirection string                   `json:"viewingDirection"`
	ViewingHint      string                   `json:"viewingHint"`
	Thumbnail        interface{}              `json:"thumbnail"`
	Canvases         []map[string]interface{} `json:"canvases"`
	Structures       []map[string]interface{} `json:"structures"`
}

//Presentation API ver3.0
type ManifestV3 struct {
	Context          string                   `json:"@context"`
	Id               string                   `json:"id"`
	Type             string                   `json:"type"`
	Label            interface{}              `json:"label"`
	Metadata         []map[string]interface{} `json:"metadata"`
	Summary          interface{}              `json:"summary"`
	Rights           string                   `json:"rights"`
	Thumbnail        []map[string]interface{} `json:"thumbnail"`
	Provider         []map[string]interface{} `json:"provider"`
	ViewingDirection string                   `json:"viewingDirection"`
	Behavior         interface{}              `json:"behavior"`
	Items            []CanvaseV2              `json:"items"`
}

type CanvaseV2 struct {
	Context     string                   `json:"@context"`
	Id          string                   `json:"id"`
	Type        string                   `json:"type"`
	Label       interface{}              `json:"label"`
	Height      uint                     `json:"height"`
	Width       uint                     `json:"width"`
	Items       []map[string]interface{} `json:"items"`
	Annotations []map[string]interface{} `json:"annotations"`
}
