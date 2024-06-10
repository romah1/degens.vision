package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LmnftClient struct {
	c *http.Client
}

func NewLmnftClient(c *http.Client) *LmnftClient {
	return &LmnftClient{
		c: c,
	}
}

type MultiSearchRequest struct {
	Searches []MultiSearchRequestSearch `json:"searches"`
}

type MultiSearchRequestSearch struct {
	QueryBy              string `json:"query_by"`
	SortBy               string `json:"sort_by"`
	ExcludeFields        string `json:"exclude_fields"`
	HightlightFullFields string `json:"highlight_full_fields"`
	Collection           string `json:"collection"`
	Q                    string `json:"q"`
	FacetBy              string `json:"facet_by"`
}

type MultiSearchResponse struct {
	Results []MultiSearchResponseResult `json:"results"`
}

type MultiSearchResponseResult struct {
	Found int                            `json:"found"`
	Hits  []MultiSearchResponseResultHit `json:"hits"`
}

type MultiSearchResponseResultHit struct {
	Document MultiSearchResponseResultHitDocument `json:"document"`
}

type MultiSearchResponseResultHitDocument struct {
	CollectionName string  `json:"collection_name"`
	Cost           float64 `json:"cost"`
	Deployed       uint64  `json:"deployed"`
	LaunchDate     uint64  `json:"launch_date"`
	Id             string  `json:"id"`
	MaxSupply      int     `json:"max_supply"`
	Owner          string  `json:"owner"`
	TotalMints     int     `json:"total_mints"`
	Type           string  `json:"type"`
}

func (lc *LmnftClient) MultiSearch(request *MultiSearchRequest) (*MultiSearchResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, "https://s.launchmynft.io/multi_search", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("x-typesense-api-key", typesenseApiKey)
	req.URL.RawQuery = query.Encode()
	response, err := lc.c.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad response from multi_search; status_core=%d", response.StatusCode)
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var result MultiSearchResponse
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type DataCollectionRequest struct {
	Owner string
	Id    string
}

type DataCollectionResponse struct {
	PageProps DataCollectionResponsePageProps `json:"pageProps"`
}

type DataCollectionResponsePageProps struct {
	Collection DataCollectionResponsePagePropsCollection `json:"collection"`
}

type DataCollectionResponsePagePropsCollection struct {
	Discord *string `json:"discord"`
	Twitter *string `json:"twitter"`
}

func (lc *LmnftClient) DataCollection(request *DataCollectionRequest) (*DataCollectionResponse, error) {
	url := fmt.Sprintf(
		"https://www.launchmynft.io/_next/data/0hLczpIM12_xuSYWV1Smc/collections/%s/%s.json",
		request.Owner,
		request.Id,
	)

	resp, err := lc.c.Get(url)
	if err != nil {
		return nil, fmt.Errorf("data collection request failed; err=%s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("data collection response is not 200; code=%d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("data collection failed to read body; err=%s", err.Error())
	}

	var response DataCollectionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("data collection failed to unmarshall response; err=%s", err.Error())
	}

	return &response, nil
}

const (
	typesenseApiKey = "UkN4Vnd3V2JMWWVIRlFNcTJ3dng4VGVtMGtvVGxBcmJJTTFFYS9MNXp1WT1Ha3dueyJmaWx0ZXJfYnkiOiJoaWRkZW46ZmFsc2UiLCJleGNsdWRlX2ZpZWxkcyI6ImhpZGRlbiIsInF1ZXJ5X2J5IjoiY29sbGVjdGlvbk5hbWUsb3duZXIiLCJsaW1pdF9oaXRzIjoyMDAsInNuaXBwZXRfdGhyZXNob2xkIjo1MH0="
)
