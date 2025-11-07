package app

import (
	"csvtogeojson/internal/geojson"
	"csvtogeojson/internal/voyager"
)

func Execute() {
	var cg geojson.GeoJson = geojson.NewConcreteGeoJson()
	var cv voyager.Voyager = voyager.NewConcreteVoyager(cg)
	cv.Execute()
}
