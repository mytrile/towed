package towed

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	apiURL           = "http://m.sofiatraffic.bg/tow-away/"
	notLiftedMessage = "Няма информация този автомобил да е принудително преместен"
)

func makeRequest(licensePlate string) string {
	data := url.Values{}
	data.Set("carnumber", licensePlate)

	client := &http.Client{}
	response, _ := client.PostForm(apiURL, data)

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	return string(contents)
}

// Lifted проверява дали автомобил с подадения като аргумент
// регистрационен номер е регистриран като принудително преместен
func Lifted(licensePlate string) bool {

	result := makeRequest(licensePlate)

	return !strings.Contains(result, notLiftedMessage)

}
