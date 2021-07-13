package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/kawamou/jptime"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// 実行方法は`ARG=filePath go generate`です
// 例：`ARG=../../syukujitsu.csv go generate`

//go:generate go run main.go $ARG
func main() {
	arg := os.Args[1]
	importer := jptime.NewLocalCSVImporter(arg)
	r, err := importer.Import()
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(transform.NewReader(r, japanese.ShiftJIS.NewDecoder()))
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
