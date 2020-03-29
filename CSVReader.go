package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func readCSVFromUrl(filename string) ([]Record, error) {
	var allRecords []Record
	var data, _ [][]string
	var qualifiedFileName = fmt.Sprintf("%s.csv", filename)
	if !checkFileExist(qualifiedFileName) {
		writeContentToFile(qualifiedFileName)
	}
	data, e := readContentFromFile(qualifiedFileName)

	if e != nil {
		log.Fatalln(e)
	}

	for _, line := range data[1:] {
		record := Record{
			FIPS:          line[0],
			Admin2:        line[1],
			ProvinceState: line[2],
			CountryRegion: line[3],
			LastUpdate:    line[4],
			Lat:           line[5],
			Long_:         line[6],
			Confirmed:     line[7],
			Deaths:        line[8],
			Recovered:     line[9],
			Active:        line[10],
			CombinedKey:   line[11],
		}
		allRecords = append(allRecords, record)
	}

	return allRecords, nil
}

func checkFileExist(fileName string) bool {
	if _, err := os.Stat(fmt.Sprintf("dataset/%s", fileName)); err == nil {
		return true
	} else {
		return false
	}
}

func writeContentToFile(filename string) {
	//Create the file
	out, _ := os.Create(fmt.Sprintf("dataset/%s", filename))
	defer out.Close()

	//Read from url
	resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/%s", filename))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer resp.Body.Close()

	// write csv content to the file
	if _, err := io.Copy(out, resp.Body); err != nil {
		log.Fatal(err)
	}
}

func readContentFromFile(filename string) ([][]string, error) {
	csvFile, e := os.Open(fmt.Sprintf("dataset/%s", filename))
	if e != nil {
		log.Fatalln(e.Error())
	}

	reader := csv.NewReader(csvFile)
	reader.LazyQuotes = true
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}
