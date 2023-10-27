package webapi

import (
	"client-info-service/internal/config"
	"client-info-service/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type ClientInfoEnricher struct {
	cfg config.ClientInfoEnricher
}

func NewClientInfoEnricher(cfg config.ClientInfoEnricher) *ClientInfoEnricher {
	return &ClientInfoEnricher{cfg: cfg}
}

type ageResponse struct {
	Age int `json:"age"`
}

type genderResponse struct {
	Gender string `json:"gender"`
}

type countryResponse struct {
	Countries []country `json:"country"`
}

type country struct {
	Id          string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

func (enricher *ClientInfoEnricher) Enrich(client model.Client) (model.ClientInfo, error) {
	info := model.ClientInfo{Client: client}
	params := url.Values{}
	params.Add("name", info.Name)
	// 1. Получить возраст
	_ = enricher.cfg.GetAgeURL
	getAgeURL := enricher.cfg.GetAgeURL + "?" + params.Encode()
	resp, err := http.Get(getAgeURL)
	if err != nil {
		fmt.Println(err)
		return info, err
	}
	if resp.StatusCode != http.StatusOK {
		return info, errors.New("something went wrong for knowing age")
	}
	defer resp.Body.Close()

	var ageResp ageResponse
	err = json.NewDecoder(resp.Body).Decode(&ageResp)
	if err != nil {
		return info, err
	}
	info.Age = ageResp.Age

	// 2. Получить пол
	getGenderUrl := enricher.cfg.GetGenderURL + "?" + params.Encode()
	resp, err = http.Get(getGenderUrl)
	if err != nil {
		fmt.Println(err)
		return info, err
	}
	if resp.StatusCode != http.StatusOK {
		return info, errors.New("something went wrong for knowing sex")
	}
	defer resp.Body.Close()

	var genderResp genderResponse
	err = json.NewDecoder(resp.Body).Decode(&genderResp)
	if err != nil {
		return info, err
	}
	info.Gender = genderResp.Gender

	// 3. Получить страну
	getCountryUrl := enricher.cfg.GetCountryURL + "?" + params.Encode()
	resp, err = http.Get(getCountryUrl)
	if err != nil {
		fmt.Println(err)
		return info, err
	}
	defer resp.Body.Close()

	var countryResp countryResponse
	err = json.NewDecoder(resp.Body).Decode(&countryResp)
	if err != nil {
		return info, err
	}
	if len(countryResp.Countries) == 0 {
		return info, errors.New("No country for this user")
	}
	info.CountryId = countryResp.Countries[0].Id

	return info, nil
}
