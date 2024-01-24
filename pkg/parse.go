package pkg

import (
	"api-service/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiAge struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type ApiGender struct {
	Count       int    `json:"count"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Probability float64    `json:"probability"`
}

type ApiCountry struct {
	Count   int `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

func Parse(dataUser *types.User) error {
	url := fmt.Sprintf("https://api.agify.io/?name=%v", dataUser.Name) //age
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	dataApiAge := ApiAge{}
	err = json.Unmarshal(data, &dataApiAge)
	if err != nil {
		return err
	}

	url = fmt.Sprintf("https://api.genderize.io/?name=%v", dataUser.Name) //gender
	res, err = http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	dataApiGender := ApiGender{}
	err = json.Unmarshal(data, &dataApiGender)
	if err != nil {
		return err
	}

	url = fmt.Sprintf("https://api.nationalize.io/?name=%v", dataUser.Name) //country
	res, err = http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	dataApiCountry := ApiCountry{}
	err = json.Unmarshal(data, &dataApiCountry)
	if err != nil {
		return err
	}

	//all dataUser
	dataUser.Age = dataApiAge.Age
	dataUser.Gender = dataApiGender.Gender
	dataUser.Nationality = dataApiCountry.Country[0].CountryID
	return nil
}
