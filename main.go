package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/darkknight107/weatherApp/types"
)

const (
	apiKey = "7fd53f036ee073010a46519a358b643f"
	apiUri = "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=" + apiKey
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a valid city: ")
	city, _ := inputReader.ReadString('\n')
	city = strings.TrimSuffix(city, "\n")
	city = strings.TrimSuffix(city, "\r")

	apiResult, err := getWeatherDetails(city)
	if err != nil {
		fmt.Println("Something went wrong! Please try again!", err.Error())
	}

	data := types.WeatherResponse{}
	err1 := json.Unmarshal([]byte(string(apiResult)), &data)
	if err1 != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data.Main.Temp)

}

func getWeatherDetails(city string) (string, error) {
	requestUri := fmt.Sprintf(apiUri, city)
	result, err := http.Get(requestUri)

	if err != nil {
		fmt.Println("There was an error when calling the Weather API. Please try again.")
		fmt.Println("The error is ", err.Error())
		os.Exit(1)
	}

	processedResponse, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println("There was an error processing the API response. Please try again!", err.Error())
	}

	return string(processedResponse), err

}
