package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

type FundingWalletService struct {
	c                *Client
	asset            *string
	needBtcValuation *string
}

// Asset set asset
func (s *FundingWalletService) Asset(asset string) *FundingWalletService {
	s.asset = &asset
	return s
}

// NeedBtcValuation set needBtcValuation
func (s *FundingWalletService) NeedBtcValuation(needBtcValuation string) *FundingWalletService {
	s.needBtcValuation = &needBtcValuation
	return s
}

// Do send request
func (s *FundingWalletService) Do(ctx context.Context) (res []*FundingWalletResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/asset/get-funding-asset",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.needBtcValuation != nil {
		r.setParam("needBtcValuation", *s.needBtcValuation)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*FundingWalletResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FundingWalletResponse define response of FundingWalletService
type FundingWalletResponse struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	BtcValuation string `json:"btcValuation"`
}
