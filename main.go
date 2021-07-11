package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

const (
	apiKey = "7fd53f036ee073010a46519a358b643f"
	apiUri = "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=" + apiKey
)

func main() {
	apiResult, err := getWeatherDetails("Sydney")
	if err != nil {
		fmt.Println("Something went wrong! Please try again!", err.Error())
	}
	fmt.Println(apiResult)

}

func getWeatherDetails(city string) (string, error) {
	requestUri := fmt.Sprintf(apiUri, city)
	result, err := http.Get(requestUri)

	if err != nil {
		fmt.Println("There was an error when calling the Weather API. Please try again.")
		fmt.Println("The error is ", err.Error())
		os.Exit(1)
	}

	println(reflect.TypeOf(result))
	println(result)

	processedResponse, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println("There was an error processing the API response. Please try again!", err.Error())
	}

	println(reflect.TypeOf(processedResponse))
	println(processedResponse)

	return string(processedResponse), err

}
