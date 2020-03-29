package main

import (
	"encoding/csv"
	"net/http"
)

func readCSVFromUrl(url string) ([]Record, error) {
	var allRecords []Record
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.LazyQuotes = true
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _,line := range data[1:] {
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
