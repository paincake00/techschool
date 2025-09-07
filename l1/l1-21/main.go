package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

/*
	Реализация адаптера XmlToJson.
	Например, моему сервису нужны данные только в JSON формате,
	а какая-нибудь внешняя библиотека выдает их в XML.
*/

// DataSource имеет методы, с которыми может работать мой сервис
type DataSource interface {
	GetData() ([]byte, error)
}

// XMLService поставляет данные в XML
type XMLService struct{}

func (x *XMLService) GetXMLData() ([]byte, error) {
	// какие-то XML данные, полученные из библиотеки условно
	someData := struct {
		XMLName xml.Name `xml:"user"`
		Name    string   `xml:"name"`
		Age     int      `xml:"age"`
	}{
		Name: "Alice",
		Age:  30,
	}

	return xml.Marshal(someData)
}

type XmlToJsonAdapter struct {
	// оборачиваем несовместимый XML сервис
	service *XMLService
}

// GetData один из методов для реализации интерфейса DataSource
func (x *XmlToJsonAdapter) GetData() ([]byte, error) {
	xmlData, err := x.service.GetXMLData()
	if err != nil {
		return nil, err
	}

	// структура с аннотациями и для XML, и для JSON
	var tmp struct {
		Name string `xml:"name" json:"name"`
		Age  int    `xml:"age" json:"age"`
	}
	if err = xml.Unmarshal(xmlData, &tmp); err != nil {
		return nil, err
	}

	return json.Marshal(tmp)
}

func main() {
	var dataSource DataSource = &XmlToJsonAdapter{service: &XMLService{}}
	data, _ := dataSource.GetData()
	fmt.Println(string(data)) // {"name":"Alice","age":30}
}
