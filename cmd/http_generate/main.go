package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/kawamou/jptime"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

//go:generate go run main.go
func main() {
	url := "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
	importer := jptime.NewHTTPCSVImporter(url)
	resp, err := importer.Import()
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	reader := csv.NewReader(transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder()))
	holidays, err := jptime.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}
	dst, err := os.Create("../../holidays2.go")
	if err != nil {
		log.Fatal(err)
	}
	exporter := jptime.NewGoFileExporter(dst)
	if err := exporter.Export(holidays); err != nil {
		log.Fatal(err)
	}
}
