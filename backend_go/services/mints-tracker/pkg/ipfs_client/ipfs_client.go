package ipfs_client

import (
	"encoding/json"
	"io"
	"net/http"
)

type IpfsClient struct {
	httpClients   []*http.Client
	curHttpClient int
}

func NewIpfsClient(httpClients []*http.Client) *IpfsClient {
	return &IpfsClient{
		httpClients:   httpClients,
		curHttpClient: 0,
	}
}

func (c *IpfsClient) Get(url string) (*IpfsData, error) {
	httpClient := c.provideHttpClient()
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ipfsData IpfsData
	err = json.Unmarshal(body, &ipfsData)
	if err != nil {
		return nil, err
	}
	if ipfsData.Attributes == nil {
		ipfsData.Attributes = []Attribute{}
	}
	return &ipfsData, nil
}

func (c *IpfsClient) provideHttpClient() *http.Client {
	httpClient := c.httpClients[c.curHttpClient]
	c.curHttpClient = (c.curHttpClient + 1) % len(c.httpClients)
	return httpClient
}
