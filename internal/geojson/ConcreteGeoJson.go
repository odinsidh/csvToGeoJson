package geojson

import (
	"encoding/json"
	"log"
	"os"
)

type ConcreteGeoJson struct {
	GeoJsonType string    `json:"type"`
	MetaData    MetaData  `json:"metadata"`
	Features    []Feature `json:"features"`
}

type MetaData struct {
	Name       string `json:"name"`
	Creator    string `json:"creator"`
	Desription string `json:"description"`
}

type Feature struct {
	GeoJsonType string     `json:"type"`
	ID          int        `json:"id"`
	Geometry    Geometry   `json:"geometry"`
	Properties  Properties `json:"properties"`
}

type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
	GeoJsonType string    `json:"type"`
}

type Properties struct {
	Description string `json:"description"`
	IconCaption string `json:"iconCaption"`
	IconContent string `json:"iconContent"`
	MarkerColor string `json:"marker-color"`
}

func NewConcreteGeoJson() GeoJson {
	return &ConcreteGeoJson{}
}

func (self *ConcreteGeoJson) AddHeaders(Name, Creator, Desription string) {
	self.GeoJsonType = "FeatureCollection"
	self.MetaData.Name = Name
	self.MetaData.Creator = Creator
	self.MetaData.Desription = Desription

}

func (self *ConcreteGeoJson) AddPoint(Coordinates []float64, objectID int, Name, Description, Color, Group string) {
	self.Features = append(self.Features, Feature{
		GeoJsonType: "feature",
		ID:          objectID,
		Geometry: Geometry{
			Coordinates: Coordinates,
			GeoJsonType: "Point",
		},
		Properties: Properties{
			Description: Description,
			IconCaption: Name,
			IconContent: Group,
			MarkerColor: Color,
		},
	})
}

func (self *ConcreteGeoJson) Create() {
	self.encode()
}

func (self *ConcreteGeoJson) encode() {
	var fileName string = self.MetaData.Name + "." + self.MetaData.Creator + "." + self.MetaData.Creator + ".geojson"
	var path string = `data\output\` + fileName

	// var output *os.File = os.Stdout
	output, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	jsonEncoder := json.NewEncoder(output)
	jsonEncoder.SetIndent("", "\t")
	jsonEncoder.Encode(self)
}

func (self *ConcreteGeoJson) Clear() {
	self.MetaData.Creator = ""
	self.MetaData.Desription = ""
	self.MetaData.Name = ""
	self.Features = nil
}
