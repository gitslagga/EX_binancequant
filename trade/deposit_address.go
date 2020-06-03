package trade

import (
	"context"
	"encoding/json"
)

// DepositsAddressService deposits address
type DepositsAddressService struct {
	c       *Client
	coin    *string
	network *string
}

// Coin set coin
func (s *DepositsAddressService) Coin(coin string) *DepositsAddressService {
	s.coin = &coin
	return s
}

// Network set network
func (s *DepositsAddressService) Network(network string) *DepositsAddressService {
	s.network = &network
	return s
}

// Do send request
func (s *DepositsAddressService) Do(ctx context.Context, opts ...RequestOption) (address *DepositAddressResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/deposit/address",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.coin != nil {
		m["coin"] = *s.coin
	}
	if s.network != nil {
		m["network"] = *s.network
	}
	r.setParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	var res = new(DepositAddressResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res, nil
}

// DepositAddressResponse define deposit history
type DepositAddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}
