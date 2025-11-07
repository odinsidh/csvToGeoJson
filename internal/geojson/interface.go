package geojson

type GeoJson interface {
	Create()
	AddHeaders(Name, Creator, Desription string)
	AddPoint(Coordinates []float64, objectID int, Name, Description, Color, Group string)
	Clear()
}
