package trade

import (
	"context"
	"encoding/json"
)

// GetInfoService query broker account information
type GetInfoService struct {
	c *Client
}

// Do send request
func (s *GetInfoService) Do(ctx context.Context, opts ...RequestOption) (res *GetInfo, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/info",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetInfo)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetInfo define query broker info
type GetInfo struct {
	MaxMakerCommission int `json:"maxMakerCommission"`
	MinMakerCommission int `json:"minMakerCommission"`
	MaxTakerCommission int `json:"maxTakerCommission"`
	MinTakerCommission int `json:"minTakerCommission"`
	SubAccountQty      int `json:"subAccountQty"`
	MaxSubAccountQty   int `json:"maxSubAccountQty"`
}
