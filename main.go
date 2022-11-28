package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type NationData struct {
	Id         string `json:"ID Nation"`
	Nation     string `json:"Nation"`
	YearOfData string `json:"Year"`
	Population int64  `json:"Population"`
}

type USData struct {
	AllData []NationData `json:"data"`
}

type ZipData struct {
	Country string `json:"country"`
}

func CallApi(url string) ([]byte, error) {
	resp, rerr := http.Get(url)

	if rerr != nil {
		return nil, rerr
	}

	body, berr := io.ReadAll(resp.Body)

	if berr != nil {
		return nil, berr
	}

	return body, nil
}

func main() {
	body, berr := CallApi("https://datausa.io/api/data?drilldowns=Nation&measures=Population")
	if berr != nil {
		fmt.Printf("%v\n", berr)
		os.Exit(1)
	}

	data := USData{}

	jerr := json.Unmarshal(body, &data)
	if jerr != nil {
		fmt.Printf("ERROR: %v\n", jerr)
		os.Exit(3)
	}

	for _, v := range data.AllData {
		if v.YearOfData == "2014" {
			newBody, nberr := CallApi("https://api.zippopotam.us/us/11101")
			if nberr != nil {
				fmt.Printf("%v\n", nberr)
				os.Exit(1)
			}

			ndata := ZipData{}

			jerr := json.Unmarshal(newBody, &ndata)
			if jerr != nil {
				fmt.Printf("ERROR: %v\n", jerr)
				os.Exit(3)
			}

			fmt.Printf("Country: %s\n", ndata.Country)

			break
		}
	}
}
