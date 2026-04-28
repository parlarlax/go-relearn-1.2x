package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate,omitempty"`
}

func main() {
	fmt.Println("=== 1. Marshal XML ===")
	rss := RSS{
		Version: "2.0",
		Channel: Channel{
			Title:       "Go Blog",
			Link:        "https://go.dev/blog",
			Description: "The Go Programming Language Blog",
			Items: []Item{
				{Title: "Go 1.26 Released", Link: "https://go.dev/blog/go1.26", Description: "Release notes"},
				{Title: "Generics Update", Link: "https://go.dev/blog/generics", PubDate: "2026-01-15"},
			},
		},
	}
	data, _ := xml.MarshalIndent(rss, "", "  ")
	fmt.Println(string(data))

	fmt.Println("\n=== 2. Unmarshal XML ===")
	input := `
	<person id="1">
		<name>Alice</name>
		<email>alice@test.com</email>
		<address city="Bangkok" zip="10100"/>
	</person>`
	var person struct {
		XMLName xml.Name `xml:"person"`
		ID      int      `xml:"id,attr"`
		Name    string   `xml:"name"`
		Email   string   `xml:"email"`
		Address struct {
			City string `xml:"city,attr"`
			Zip  string `xml:"zip,attr"`
		} `xml:"address"`
	}
	xml.Unmarshal([]byte(input), &person)
	fmt.Printf("%+v\n", person)
	fmt.Printf("  name=%s email=%s city=%s\n", person.Name, person.Email, person.Address.City)

	fmt.Println("\n=== 3. XML tags reference ===")
	fmt.Println("  xml:\"name\"         → element name")
	fmt.Println("  xml:\"name,attr\"    → attribute")
	fmt.Println("  xml:\"name,chardata\" → character data")
	fmt.Println("  xml:\"name,innerxml\" → raw inner XML")
	fmt.Println("  xml:\"-\"            → skip field")
	fmt.Println("  xml:\"name,omitempty\" → skip if zero")

	fmt.Println("\n=== 4. Encoder (streaming) ===")
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	enc.Encode(map[string]string{"message": "streaming XML"})
	fmt.Println()

	fmt.Println("\n=== 5. Custom MarshalXML ===")
	cdata := CDATA{Content: "<b>bold text</b>"}
	data, _ = xml.Marshal(cdata)
	fmt.Println(string(data))
}

type CDATA struct {
	Content string `xml:",cdata"`
}
