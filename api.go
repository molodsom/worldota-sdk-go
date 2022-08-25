package worldota

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/molodsom/worldota-sdk-go/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type APIClient struct {
	http.Client
}

type APIRoundTripper struct {
	id    string
	token string
	r     http.RoundTripper
}

func (rt APIRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept-Charset", "utf-8")
	r.Header.Add("Authorization", "Basic "+basicAuth(rt.id, rt.token))
	return rt.r.RoundTrip(r)
}

func Client(id, token string) (c APIClient) {
	c.Transport = APIRoundTripper{id: id, token: token}
	return
}

func (c APIClient) request(method, endpoint string, payload []byte) (body []byte) {
	url := fmt.Sprintf("https://api.worldota.net/api/b2b/v3/%s/", endpoint)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	res, _ := c.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != 200 {
		log.Fatal(fmt.Sprintf("Server returned %d in response to URL: %s\n", res.StatusCode, url))
		return
	}

	body, _ = ioutil.ReadAll(res.Body)
	return
}

func (c APIClient) get(endpoint string, payload []byte) (body []byte) {
	body = c.request("GET", endpoint, payload)
	return
}

func (c APIClient) post(endpoint string, payload []byte) (body []byte) {
	body = c.request("POST", endpoint, payload)
	return
}

func (c APIClient) overview() (res models.Overview) {
	r := c.get("overview", nil)
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c APIClient) hotelInfoDump(req models.LanguageRequest) (res models.HotelDump) {
	payload, _ := json.Marshal(req)
	r := c.post("hotel/info/dump", payload)
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c APIClient) hotelInfoIncrementalDump(req models.LanguageRequest) (res models.HotelDump) {
	payload, _ := json.Marshal(req)
	r := c.post("hotel/info/incremental_dump", payload)
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c APIClient) hotelReviewsDump(req models.LanguageRequest) (res models.HotelDump) {
	payload, _ := json.Marshal(req)
	r := c.post("hotel/reviews/dump", payload)
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c APIClient) hotelRegionDump(req models.LanguageRequest) (res models.HotelDump) {
	payload, _ := json.Marshal(req)
	r := c.post("hotel/region/dump", payload)
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c APIClient) hotelInfo(req models.HotelInfoRequest) (res models.HotelInfo) {
	payload, _ := json.Marshal(req)
	r := c.post("hotel/info", payload)
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c APIClient) searchMulticomplete(req models.SearchMulticompleteRequest) (res models.SearchMulticomplete) {
	payload, _ := json.Marshal(req)
	r := c.post("search/multicomplete", payload)
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c APIClient) searchSerpRegion(req models.SearchRequest) (res models.SearchResult) {
	payload, _ := json.Marshal(req)
	r := c.post("search/serp/region", payload)
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c APIClient) searchSerpGeo(req models.SearchGeoRequest) (res models.SearchResult) {
	payload, _ := json.Marshal(req)
	r := c.post("search/serp/geo", payload)
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c APIClient) searchSerpHotels(req models.SearchHotelsRequest) (res models.SearchResult) {
	payload, _ := json.Marshal(req)
	r := c.post("search/serp/hotels", payload)
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c APIClient) searchHotelPage(req models.SearchHotelPage) (res models.SearchResult) {
	payload, _ := json.Marshal(req)
	r := c.post("search/hp", payload)
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Fatal(err)
	}
	return
}
