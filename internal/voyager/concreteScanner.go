package voyager

import (
	"bytes"
	"csvtogeojson/internal/geojson"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Concretevoyager struct {
	queue   []os.DirEntry
	geoJson geojson.GeoJson
}

func NewConcreteVoyager(input geojson.GeoJson) Voyager {
	return &Concretevoyager{
		geoJson: input,
	}
}

func (self *Concretevoyager) Execute() {
	self.readDir()
	for i := range self.queue {
		self.proceed(i)
	}

}

func (self *Concretevoyager) readDir() {
	var path string = `data\input`
	container, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	self.queue = container
}

func (self *Concretevoyager) proceed(index int) {
	headers := self.extractHeaders(index)
	body := self.extractBody(index)
	self.geoJson.AddHeaders(headers[0], headers[1], headers[2])
	for _, value := range body {
		container := strings.Split(value, ";")
		coordinateX, err := strconv.ParseFloat(container[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		coordinateY, err := strconv.ParseFloat(container[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		objectID, err := strconv.Atoi(container[2])
		self.geoJson.AddPoint([]float64{coordinateY, coordinateX}, objectID, container[3], container[4], container[5], container[6])
	}
	self.geoJson.Create()
	self.geoJson.Clear()

}

func (self *Concretevoyager) extractHeaders(input int) []string {
	var file string = self.queue[input].Name()
	concrereRegexp := regexp.MustCompile(`^([^.]*)\.+([^.]*)\.+([^.]*)\.+([^.]*)$`)
	regexpContainer := concrereRegexp.FindAllStringSubmatch(file, -1)
	return regexpContainer[0][1:]
}

func (self *Concretevoyager) extractBody(input int) []string {
	var (
		file         string = self.queue[input].Name()
		path         string = `data\input\` + file
		container    []string
		buffer       []byte = make([]byte, 256)
		extendBuffer []byte
		sequence     []byte = []byte{13, 10}
	)

	concreteFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer concreteFile.Close()

cycle:
	for {
		offset, err := concreteFile.Read(buffer)
		extendBuffer = append(extendBuffer, buffer[:offset]...)
		contain := bytes.Contains(extendBuffer, sequence)
		if contain {
			answer := bytes.Split(extendBuffer, sequence)
			for len(answer) > 1 {
				container = append(container, string(answer[0]))
				answer = answer[1:]
			}
			extendBuffer = answer[0]
		}
		if err == io.EOF {
			container = append(container, string(extendBuffer))
			break cycle
		}
	}
	return container[1 : len(container)-1]
}
