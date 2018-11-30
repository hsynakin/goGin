package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	type FirmaArrayStructure struct {
		XMLName           xml.Name  `json:"-"`
		Identifier        string    `json:"identifier"`
		Alias             string    `json:"alias"`
		Title             string    `json:"title"`
		Type              string    `json:"type"`
		FirstCreationTime time.Time `json:"firstCreationTime"`
	}

	type Kullanicilar struct { //xml deki kullanıcıları bir dizi olarak tanımladık.
		XMLName xml.Name              `xml:"UserList"`
		Adanla  []FirmaArrayStructure `xml:"User"`
	}

	var xmlUsers Kullanicilar //Unmarshal öncelikle verilerin saklanacağı bir alan oluşturmamızı gerektirdiği için xmlUsers oluşturduk. https://blog.golang.org/json-and-go (search : decoding)

	xmlFile, err := os.Open("./Users.xml")
	if err != nil {
		log.Println(err)
	}

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &xmlUsers)

}
