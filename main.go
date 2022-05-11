package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/motogoozy/go-weather/forecast"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a zip code: ")
	zip, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatalln("Error reading user input:", readErr)
	}

	fc, err := forecast.GetForecast(strings.TrimSpace(zip))
	if err != nil {
		log.Fatalln("Error getting forecast:", err)
	}

	fmt.Println(forecast.FormatForecast(fc))
}
