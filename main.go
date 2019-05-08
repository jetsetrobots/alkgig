package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// root document of a song
type Document struct {
	XMLName    xml.Name  `xml:"document"`
	Name       string    `xml:"name,attr"`
	ID         string    `xml:"id,attr"`
	StartBeats int       `xml:"startBeats,attr"`
	EndBeats   int       `xml:"endBeats,attr"`
	Transport  Transport `xml:"transport"`
	Arrays     []Array   `xml:"array"`
}

type Transport struct {
	XMLName xml.Name `xml:"transport"`
	Name    string   `xml:"name,attr"`
	ID      string   `xml:"id,attr"`
	BPM     int      `xml:"bpm,attr"`
}

type Array struct {
	XMLName xml.Name `xml:"array"`
	Name    string   `xml:"name,attr"`
	Tracks  []Track  `xml:"track"`
}

type Track struct {
	XMLName xml.Name `xml:"track"`
	Name    string   `xml:"name,attr"`
}

func main() {

	wd, err := os.Getwd()
	check(err)
	/*
		if len(os.Args) < 2 {
			log.Panic("need project directory :-(")
		}
		projectPath := filepath.Join(wd, os.Args[1])

		dir, err := ioutil.ReadDir(projectPath)
		check(err)

		var projects []string
		for _, info := range dir {
			path := filepath.Join(projectPath, info.Name())
			extension := filepath.Ext(path)
			if info.IsDir() && extension == ".alk" {
				name := path[0 : len(path)-len(extension)]
				projects = append(projects, name)
			}
		}

		fmt.Println(projects)
	*/
	path := filepath.Join(wd, "data", "Song1.alk", "Song1 project.alk")
	fmt.Println(path)
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	check(err)

	var document Document
	xml.Unmarshal(byteValue, &document)
	fmt.Println(document)

}
