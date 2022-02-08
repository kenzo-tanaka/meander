package meander

type Place struct {
	*googleGeometry `json:geometry`
	Name            string          `json:"name"`
	Icon            string          `json:"icon"`
	Photos          []*googlePhotos `json:"photos"`
	Vicinity        string          `json:"vicinity"`
}

type googleResponse struct {
	Results []*Place `json:"results"`
}

type googleGeometry struct {
	*googleLocation `json:"location"`
}

type googleLocation struct {
	Lat float64 `josn:"lat"`
	Lng float64 `json:"lng"`
}

type googlePhotos struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

var APIKey string
