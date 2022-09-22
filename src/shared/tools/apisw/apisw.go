package apisw

import (
	//"bytes"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"net/http"
	"net/url"
	"strings"
)

const (
	BASE_URL = "https://swapi.dev/api/planets/"
)

type responseAPI struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []*planetAPI `json:"results"`
}

type planetAPI struct {
	Name     string   `json:"name"`
	Climate  string   `json:"climate"`
	Terrain  string   `json:"terrain"`
	FilmURLs []string `json:"films"`
}

func newResponse() *responseAPI {
	return &responseAPI{
		Results: make([]*planetAPI, 0),
	}
}

func GetQtdeFilm(planetName string) (int, error) {

	if len(planetName) == 0 {
		return 0, errors.New("planet not found")
	}

	url, err := url.Parse(BASE_URL)
	if err != nil {
		return 0, err
	}

	p := url.Query()
	p.Add("search", planetName)
	url.RawQuery = p.Encode()

	_, qtde, err := getPlanetsSWAPI(url.String(), planetName, 0)

	if err != nil {
		return 0, err
	}
	return qtde, nil
}

func getPlanetsSWAPI(planetsURL, planetName string, qtde int) (string, int, error) {

	req, err := http.NewRequest(http.MethodGet, planetsURL, nil)
	if err != nil {
		return "", 0, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}

	fmt.Println(bytes.NewBuffer(body))

	response := newResponse()
	if err := json.Unmarshal(body, &response); err != nil {
		println("aqui")
		return "", 0, err
	}

	for _, v := range response.Results {
		fmt.Println(v)
		if strings.EqualFold(v.Name, planetName) {
			qtde += len(v.FilmURLs)

		}
	}
	if len(response.Next) == 0 {
		return "", qtde, nil

	}
	return getPlanetsSWAPI(response.Next, "", 0)
}
