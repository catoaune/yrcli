package main

import (
	"fmt"
	"testing"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
  )

func TestYrAPI(t *testing.T) {
	assert.Equal(t, "Gvarv", readLocationFromXml())
	
	//assert.Equal(t, 123, 123, "they should be equal")
}

func readLocationFromXml() string {

	v := Weatherdata{}

	data := `
		<weatherdata>
			<location>
				<name>Gvarv</name>
			</location>
		</weatherdata>
	`


	err := xml.Unmarshal([]byte(data), &v)

	if err != nil {
		fmt.Printf("error: %v", err)
		return ""
	}

	return v.Location.Name
	//return "Gvarv"
}

type Weatherdata struct {
	XMLName xml.Name `xml:"weatherdata"`
	Location struct {
		Name string `xml:"name"`
	} `xml:"location"`//L//ocation `xml:""`
}

// type Location struct {
// 	// XMLName xml.Name `xml:"location"`
// 	// Name string `xml:"name"`
	
// 	//Name  string `xml:"location>name"`
// 	 `xml:"location"`
// }

// type Name struct {
// 	Name string `xml:"name"`
// }