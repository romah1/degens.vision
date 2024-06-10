package ipfs_client

import "encoding/json"

type IpfsData struct {
	Name        string      `json:"name"`
	Symbol      string      `json:"symbol"`
	Image       string      `json:"image"`
	Attributes  []Attribute `json:"attributes"`
	Description string      `json:"description"`
}

type Attribute struct {
	TraitType string          `json:"trait_type"`
	Value     json.RawMessage `json:"value"`
}
