package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type HackerRankData struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		Name    string   `json:"name"`
		Weather string   `json:"weather"`
		Status  []string `json:"status"`
	} `json:"data"`
}

func getWind(str string) string {
	words := strings.Fields(str)
	wind := words[1]
	windSpeed := wind[:len(wind)-4]
	return windSpeed
}

func getHumidity(str string) string{
	words := strings.Fields(str)
	humidity := words[1]
	humidityPercent := humidity[:len(humidity)-1]
	return humidityPercent
}

func getWeather(str string) string{
	words := strings.Fields(str)
	weather := words[0]
	return weather
}

func main() {
	var ans []string
	count := 0
	pageCount := 1;
	for {
		url := "https://jsonmock.hackerrank.com/api/weather/search?name=al&page="
		pc := strconv.Itoa(pageCount)
		resp,err := http.Get(url + pc)
		if err!=nil{
			fmt.Print("Error while making request")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print("Error while converting json to byte array")
		}

		var emp HackerRankData
		err = json.Unmarshal(body, &emp)
		
		if err != nil {
			fmt.Println("Can;t unmarshal the byte array")
			return
		}

		for i :=0 ; i < emp.PerPage && count < emp.Total ;i++{
			name := emp.Data[i].Name
			weather := getWeather(emp.Data[i].Weather)
			wind := getWind(emp.Data[i].Status[0])
			humidity := getHumidity(emp.Data[i].Status[1])
			m := name + "," + weather + "," + wind + ","  + humidity
			ans = append(ans, m)
			count++
		}

		if count == emp.Total{
			break
		}
		pageCount++;
    }

	// sort result on name
	sort.Slice(ans, func(i, j int) bool {
		words1 := strings.Split(ans[i], ",")
		words2 := strings.Split(ans[j], ",")
		return words1[0] < words2[0]
	})

	fmt.Println(ans)
}

// output :- [Altay City,12,25,5 Alwar,15,1,5 Bahawalpur,16,13,36 Dallas,12,2,5 Dallupura,10,9,14 East Jerusalem,13,36,17 Faisalabad,22,4,44 Galați,28,14,11 Jabalpur,36,35,42 Juticalpa,12,24,9 Phaltan,17,14,17 Qapqal,11,29,46 Ualog,15,8,5 Valenca,15,39,80 Vallejo,1,24,56 Waltham,11,32,80]
