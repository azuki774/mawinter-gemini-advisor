package mawinter

import (
	"io"
	"net/http"
)

type MawinterClient struct {
	Endpoint string // ex. http://localhost:8080/v2/record
}

func NewMawinterClient(endpoint string) *MawinterClient {
	if endpoint == "mock" {
		endpoint = "http://localhost:8080/v2/record"
	}

	return &MawinterClient{
		Endpoint: endpoint,
	}
}

func (m *MawinterClient) GetMonthlyData(yyyymm string) (string, error) {
	endpoint := m.Endpoint + "?yyyymm=" + yyyymm
	res, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
