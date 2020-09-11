package trade

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type brokerServiceTestSuite struct {
	baseTestSuite
}

func TestBrokerService(t *testing.T) {
	suite.Run(t, new(brokerServiceTestSuite))
}

func (s *brokerServiceTestSuite) TestGetInfo() {
	data := []byte(`{
		"maxMakerCommission":"20",
		"minMakerCommission":"5",
		"maxTakerCommission":"20",
		"minTakerCommission":"5",
		"subAccountQty":400,
		"maxSubAccountQty":1000
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetInfoService().Do(newContext())
	s.r().NoError(err)
	e := &GetInfo{
		MaxMakerCommission: "20",
		MinMakerCommission: "5",
		MaxTakerCommission: "20",
		MinTakerCommission: "5",
		SubAccountQty:      400,
		MaxSubAccountQty:   1000,
	}
	s.assertBrokerEqual(e, res)
}

func (s *brokerServiceTestSuite) assertBrokerEqual(e, a *GetInfo) {
	r := s.r()
	r.Equal(e.MaxMakerCommission, a.MaxMakerCommission, "MaxMakerCommission")
	r.Equal(e.MinMakerCommission, a.MinMakerCommission, "MinMakerCommission")
	r.Equal(e.MaxTakerCommission, a.MaxTakerCommission, "MaxTakerCommission")
	r.Equal(e.MinTakerCommission, a.MinTakerCommission, "MinTakerCommission")
	r.Equal(e.SubAccountQty, a.SubAccountQty, "SubAccountQty")
	r.Equal(e.MaxSubAccountQty, a.MaxSubAccountQty, "MaxSubAccountQty")
}

func (s *brokerServiceTestSuite) TestCreateTransfer() {
	data := []byte(`{
		"txnId":2966662589,
		"clientTranId":"testClientId"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	fromId := "485396905497952257"
	toId := "485396905497953378"
	clientTranId := "testClientId"
	asset := "BTC"
	amount := 9257.0
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"fromId":       fromId,
			"toId":         toId,
			"clientTranId": clientTranId,
			"asset":        asset,
			"amount":       amount,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCreateTransferService().FromId(fromId).ToId(toId).
		FuturesType(1).Asset(asset).Amount(amount).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &CreateTransfer{
		TxnId:        2966662589,
		ClientTranId: "testClientId",
	}
	s.assertCreateTransferEqual(e, res)
}

func (s *brokerServiceTestSuite) assertCreateTransferEqual(e, a *CreateTransfer) {
	r := s.r()
	r.Equal(e.TxnId, a.TxnId, "TxnId")
	r.Equal(e.ClientTranId, a.ClientTranId, "ClientTranId")
}

func (s *brokerServiceTestSuite) TestGetTransfer() {
	data := []byte(`[
        {
            "fromId":"485396905497952257",
            "toId":"485396905497953378",
            "asset":"BTC",
            "qty":"1",
            "time":1544433328000,
            "txnId":"2966662589",
            "clientTranId":"testClientId"
        },
        {
            "fromId":"485396905497952257",
            "toId":"485396905497953378",
            "asset":"ETH",
            "qty":"2",
            "time":1544433328000,
            "txnId":"296666999",
            "clientTranId":"testClientId"
        }
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	subAccountId := "485396905497952257"
	clientTranId := "testClientId"
	startTime := 1544433228000
	endTime := 1544433528000
	page := 1
	limit := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"subAccountId": subAccountId,
			"clientTranId": clientTranId,
			"startTime":    startTime,
			"endTime":      endTime,
			"page":         page,
			"limit":        limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetTransferService().SubAccountId(subAccountId).ClientTranId(clientTranId).
		StartTime(uint64(startTime)).EndTime(uint64(endTime)).Page(page).Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := []*GetTransfer{
		{
			FromId:       "485396905497952257",
			ToId:         "485396905497953378",
			Asset:        "BTC",
			Qty:          "1",
			Time:         1544433328000,
			TxnId:        "2966662589",
			ClientTranId: "testClientId",
		},
		{
			FromId:       "485396905497952257",
			ToId:         "485396905497953378",
			Asset:        "ETH",
			Qty:          "2",
			Time:         1544433328000,
			TxnId:        "296666999",
			ClientTranId: "testClientId",
		},
	}
	s.assertGetTransferEqual(e, res)
}

func (s *brokerServiceTestSuite) assertGetTransferEqual(e, a []*GetTransfer) {
	r := s.r()
	for i := 0; i < len(e); i++ {
		r.Equal(e[i].FromId, a[i].FromId, "FromId")
		r.Equal(e[i].ToId, a[i].ToId, "ToId")
		r.Equal(e[i].Asset, a[i].Asset, "Asset")
		r.Equal(e[i].Qty, a[i].Qty, "Qty")
		r.Equal(e[i].Time, a[i].Time, "Time")
		r.Equal(e[i].TxnId, a[i].TxnId, "TxnId")
		r.Equal(e[i].ClientTranId, a[i].ClientTranId, "ClientTranId")
	}
}

func (s *brokerServiceTestSuite) TestGetSubAccountDepositHist() {
	data := []byte(`[
	  {
		"subaccountId": "485396905497952257",
		"address": "0xddc66e4313fd6c737b6cae67cad90bb4e0ac7092",
		"addressTag": "",
		"amount": "139.04370000",
		"coin": "USDT",
		"insertTime": 1566791463000,
		"network": "ETH",
		"status": 1,
		"txId": "0x5759dfe9983a4c7619bce9bc736bb6c26f804091753bf66fa91e7cd5cfeebafd",     
		"sourceAddress":"xxxxxxxxxxxxxx"
	  },
	  {
		"subaccountId": "485396905497952257",
		"address": "0xddc66e4313fd6c737b6cae67kld90bb4e0ac7092",
		"addressTag": "",
		"amount": "1589.12345678",
		"coin": "BTC",
		"insertTime": 1566791463000,
		"network": "BNB",
		"status": 1,
		"txId": "0x5759dfe9983a4c7619bsdltixngbb6c26f804091753bf66fa91e7cd5cfeebafd",
		"sourceAddress":"xxxxxxxxxxxxxx"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	subAccountId := "485396905497952257"
	coin := "BTC"
	status := 1
	startTime := 1544433228000
	endTime := 1544433528000
	limit := 10
	offset := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"subAccountId": subAccountId,
			"coin":         coin,
			"status":       status,
			"startTime":    startTime,
			"endTime":      endTime,
			"limit":        limit,
			"offset":       offset,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetSubAccountDepositHistService().SubAccountId(subAccountId).Coin(coin).Status(status).
		StartTime(uint64(startTime)).EndTime(uint64(endTime)).Limit(limit).Offset(offset).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := []*GetSubAccountDepositHist{
		{
			SubAccountId:  "485396905497952257",
			Address:       "0xddc66e4313fd6c737b6cae67cad90bb4e0ac7092",
			AddressTag:    "",
			Amount:        "139.04370000",
			Coin:          "USDT",
			InsertTime:    1566791463000,
			Network:       "ETH",
			Status:        1,
			TxId:          "0x5759dfe9983a4c7619bce9bc736bb6c26f804091753bf66fa91e7cd5cfeebafd",
			SourceAddress: "xxxxxxxxxxxxxx",
		}, {
			SubAccountId:  "485396905497952257",
			Address:       "0xddc66e4313fd6c737b6cae67kld90bb4e0ac7092",
			AddressTag:    "",
			Amount:        "1589.12345678",
			Coin:          "BTC",
			InsertTime:    1566791463000,
			Network:       "BNB",
			Status:        1,
			TxId:          "0x5759dfe9983a4c7619bsdltixngbb6c26f804091753bf66fa91e7cd5cfeebafd",
			SourceAddress: "xxxxxxxxxxxxxx",
		},
	}
	s.assertGetSubAccountDepositHistEqual(e, res)
}

func (s *brokerServiceTestSuite) assertGetSubAccountDepositHistEqual(e, a []*GetSubAccountDepositHist) {
	r := s.r()
	for i := 0; i < len(e); i++ {
		r.Equal(e[i].SubAccountId, a[i].SubAccountId, "SubAccountId")
		r.Equal(e[i].Address, a[i].Address, "Address")
		r.Equal(e[i].AddressTag, a[i].AddressTag, "AddressTag")
		r.Equal(e[i].Amount, a[i].Amount, "Amount")
		r.Equal(e[i].Coin, a[i].Coin, "Coin")
		r.Equal(e[i].InsertTime, a[i].InsertTime, "InsertTime")
		r.Equal(e[i].Network, a[i].Network, "Network")
		r.Equal(e[i].Status, a[i].Status, "Status")
		r.Equal(e[i].TxId, a[i].TxId, "TxId")
		r.Equal(e[i].SourceAddress, a[i].SourceAddress, "SourceAddress")
	}
}

func (s *brokerServiceTestSuite) TestGetSubAccountSpotSummary() {
	data := []byte(`{
		"data":[
			{
			   "subAccountId": "485396905497952257", 
				"totalBalanceOfBtc": "0.0355852154360000"
			},
			{ 
				"subAccountId": "485396905497952257", 
				"totalBalanceOfBtc": "0.0233852154360000"
			}
		],
		"timestamp": 1583432900000
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	subAccountId := "485396905497952257"
	page := 1
	size := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"subAccountId": subAccountId,
			"page":         page,
			"size":         size,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetSubAccountSpotSummaryService().SubAccountId(subAccountId).
		Page(uint64(page)).Size(uint64(size)).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &GetSubAccountSpotSummary{
		Data: []*subAccountSpotSummary{
			{
				SubAccountId:      "485396905497952257",
				TotalBalanceOfBtc: "0.0355852154360000",
			},
			{
				SubAccountId:      "485396905497952257",
				TotalBalanceOfBtc: "0.0233852154360000",
			},
		},
		Timestamp: 1583432900000,
	}
	s.assertGetSubAccountSpotSummaryEqual(e, res)
}

func (s *brokerServiceTestSuite) assertGetSubAccountSpotSummaryEqual(e, a *GetSubAccountSpotSummary) {
	r := s.r()
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	for i := 0; i < len(e.Data); i++ {
		r.Equal(e.Data[i].SubAccountId, a.Data[i].SubAccountId, "SubAccountId")
		r.Equal(e.Data[i].TotalBalanceOfBtc, a.Data[i].TotalBalanceOfBtc, "TotalBalanceOfBtc")
	}
}

func (s *brokerServiceTestSuite) TestGetSubAccountFuturesSummary() {
	data := []byte(`{
		"data": [
			{
				"futuresEnable": true,
			   	"subAccountId": "485396905497952257",
				"totalInitialMarginOfUsdt": "0.03558521",
				"totalMaintenanceMarginOfUsdt": "0.02695000", 
				"totalWalletBalanceOfUsdt": "8.23222312",
				"totalUnrealizedProfitOfUsdt": "-0.78628370",
				"totalMarginBalanceOfUsdt": "8.23432343", 
				"totalPositionInitialMarginOfUsdt": "0.33683000", 
				"totalOpenOrderInitialMarginOfUsdt": "0.00000000"
	
			},
			{ 
				"futuresEnable": false,
				"subAccountId": "36753702750323432"
			}
		],
		"timestamp": 1583127900000
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	subAccountId := "485396905497952257"
	page := 1
	size := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"subAccountId": subAccountId,
			"page":         page,
			"size":         size,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetSubAccountFuturesSummaryService().SubAccountId(subAccountId).
		Page(uint64(page)).Size(uint64(size)).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &GetSubAccountFuturesSummary{
		Data: []*subAccountFuturesSummary{
			{
				FuturesEnable:                     true, // if enable futures
				SubAccountId:                      "485396905497952257",
				TotalInitialMarginOfUsdt:          "0.03558521",  //  initial margin
				TotalMaintenanceMarginOfUsdt:      "0.02695000",  // maintenance margin
				TotalWalletBalanceOfUsdt:          "8.23222312",  // wallet balance
				TotalUnrealizedProfitOfUsdt:       "-0.78628370", // unrealized profit
				TotalMarginBalanceOfUsdt:          "8.23432343",  // margin balance
				TotalPositionInitialMarginOfUsdt:  "0.33683000",  // position initial margin
				TotalOpenOrderInitialMarginOfUsdt: "0.00000000",  // open order initial margin
			},
			{
				FuturesEnable: false,
				SubAccountId:  "36753702750323432",
			},
		},
		Timestamp: 1583127900000,
	}
	s.assertGetSubAccountFuturesSummaryEqual(e, res)
}

func (s *brokerServiceTestSuite) assertGetSubAccountFuturesSummaryEqual(e, a *GetSubAccountFuturesSummary) {
	r := s.r()
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	for i := 0; i < len(e.Data); i++ {
		r.Equal(e.Data[i].SubAccountId, a.Data[i].SubAccountId, "SubAccountId")
		r.Equal(e.Data[i].FuturesEnable, a.Data[i].FuturesEnable, "FuturesEnable")
		r.Equal(e.Data[i].TotalInitialMarginOfUsdt, a.Data[i].TotalInitialMarginOfUsdt, "TotalInitialMarginOfUsdt")
		r.Equal(e.Data[i].TotalMaintenanceMarginOfUsdt, a.Data[i].TotalMaintenanceMarginOfUsdt, "TotalMaintenanceMarginOfUsdt")
		r.Equal(e.Data[i].TotalWalletBalanceOfUsdt, a.Data[i].TotalWalletBalanceOfUsdt, "TotalWalletBalanceOfUsdt")
		r.Equal(e.Data[i].TotalUnrealizedProfitOfUsdt, a.Data[i].TotalUnrealizedProfitOfUsdt, "TotalUnrealizedProfitOfUsdt")
		r.Equal(e.Data[i].TotalMarginBalanceOfUsdt, a.Data[i].TotalMarginBalanceOfUsdt, "TotalMarginBalanceOfUsdt")
		r.Equal(e.Data[i].TotalPositionInitialMarginOfUsdt, a.Data[i].TotalPositionInitialMarginOfUsdt, "TotalPositionInitialMarginOfUsdt")
		r.Equal(e.Data[i].TotalOpenOrderInitialMarginOfUsdt, a.Data[i].TotalOpenOrderInitialMarginOfUsdt, "TotalOpenOrderInitialMarginOfUsdt")

	}
}

func (s *brokerServiceTestSuite) TestGetRebateRecentRecord() {
	data := []byte(`[
        {
            "subaccountId":"485396905497952257",
            "income": "0.02063898",
            "asset":"BTC",
            "symbol": "ETHBTC",
            "tradeId": 123456,
            "time":1544433328000
        },
        {
            "subaccountId":"485396905497952257",
            "income": "1.2063898",
            "asset":"USDT",
            "symbol": "BTCUSDT",
            "tradeId": 223456,
            "time":1581580800000
        }
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	subAccountId := "485396905497952257"
	startTime := 1544433228000
	endTime := 1544433528000
	limit := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"subAccountId": subAccountId,
			"startTime":    startTime,
			"endTime":      endTime,
			"limit":        limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetRebateRecentRecordService().SubAccountId(subAccountId).
		StartTime(uint64(startTime)).EndTime(uint64(endTime)).Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := []*GetRebateRecentRecord{
		{
			SubAccountId: "485396905497952257",
			Income:       "0.02063898",
			Asset:        "BTC",
			Symbol:       "ETHBTC",
			TradeId:      123456,
			Time:         1544433328000,
		},
		{
			SubAccountId: "485396905497952257",
			Income:       "1.2063898",
			Asset:        "USDT",
			Symbol:       "BTCUSDT",
			TradeId:      223456,
			Time:         1581580800000,
		},
	}
	s.assertGetRebateRecentRecordEqual(e, res)
}

func (s *brokerServiceTestSuite) assertGetRebateRecentRecordEqual(e, a []*GetRebateRecentRecord) {
	r := s.r()
	for i := 0; i < len(e); i++ {
		r.Equal(e[i].SubAccountId, a[i].SubAccountId, "SubAccountId")
		r.Equal(e[i].Income, a[i].Income, "Income")
		r.Equal(e[i].Asset, a[i].Asset, "Asset")
		r.Equal(e[i].Symbol, a[i].Symbol, "Symbol")
		r.Equal(e[i].TradeId, a[i].TradeId, "TradeId")
		r.Equal(e[i].Time, a[i].Time, "Time")
	}
}

func (s *brokerServiceTestSuite) TestGenerateRebateHistory() {
	data := []byte(`{
		"code": 200,
		"msg": "Historical data is collecting"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	subAccountId := "485396905497952257"
	startTime := 1544433228000
	endTime := 1544433528000
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"subAccountId": subAccountId,
			"startTime":    startTime,
			"endTime":      endTime,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGenerateRebateHistoryService().SubAccountId(subAccountId).StartTime(uint64(startTime)).
		EndTime(uint64(endTime)).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &GenerateRebateHistory{
		Code: 200,
		Msg:  "Historical data is collecting",
	}
	s.assertGenerateRebateHistoryEqual(e, res)
}

func (s *brokerServiceTestSuite) assertGenerateRebateHistoryEqual(e, a *GenerateRebateHistory) {
	r := s.r()
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Msg, a.Msg, "Msg")
}

func (s *brokerServiceTestSuite) TestGetRebateHistory() {
	link := "https://bin-prod-user-rebate-bucket.s3.amazonaws.com/user-rebate/b4b6ca80-bcdc-11ea-8a61-0ad86c4d89f6/part-00000-d67a3f95-97ed-428f-89f0-44929fbd3405-c000.csv?AWSAccessKeyId=ASIAVL364M5ZHDRTC3ES\\u0026Expires=1593833086\\u0026x-amz-security-token=IQoJb3JpZ2luX2VjENP%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaDmFwLW5vcnRoZWFzdC0xIkcwRQIgf3MuxZjiULuxAGXOziJj5%2FjYSCXNyr9ZYRXyT6ZSbasCIQDz45ZUTIMN%2BYpnG%2B%2FVkGFsjfd2ODO0EGfitFmkQcGgyyq%2BAwhsEAAaDDM2OTA5NjQxOTE4NiIMNwJsOnb2HzG85xU4KpsDhgfDP6h5NMC8zNXNJBKHGv%2BtLSCjijh5hrADVnUtWJGylhLrb1YQfweynzwHe4tLm7LIHHvojaT1l62lKy2kUWNaBXTjW0KwQUltuO5EuvrCLHFsuPJbu9493NLI9Bdc9Tg%2BlGgdSIDXxHwt4SGtiPRXhELNvUlKK0HHAL6zDRMMuSsFHivn9NEdm3OoGW3m0XFitK4XRhDhjgxehm8xTJsznjj1UlXq7d%2BcUqrK2rO13%2BVNhOPYQdNE%2FAIy6CA8mVyioNVGDfMmVX9%2BeuGBWFeogIv%2BlkQC%2FiGlsfLT%2FejfYHVREX3NfH5C2MzB8VffKmeUfQbxjGh7GEFeQyraprx5iH6ukVJOCoWHQAnrMlLyitCkxuT7Bc19hRDzKXrw3NnewOuz2CWGD%2Fc8ALSV6xdAhaadq4mGekjHt%2Fyph7fT3Ctjx%2BSR1EXmtzvBe2X%2BR%2FFoTw5ismACvdQdWHoigs9ef66lojfLUWU3CNwITX7nvP%2Fx27ISTduQT2RgynM1XYWMJdC4ZcxLuC71cYxnM%2B7swrbSx8Jyw3jQMLLE%2BvcFOusBQ3q3UmnqDQZ5NzISfzFQRcT1%2Fb5YuWZBttswaf2bWwYy82P%2FeV%2BoFJBjXh3zkc6oTpA0w1FfE2LW3Pz3Rh1E8jOQyGb2IQGNvdByfQVdJztV%2F%2BMtDFd5w0ZkBZwFEPleikx%2Fn5P04VVPX6%2FjY%2BYtM1RDoa8%2FaDFeAxeFQnNnYNNvXJdKF8dax9cbT2tNy1AzmOQDLAywlUwvzbMJdAJMoS04%2FAn%2FjP%2BPpbwgIjvuieHzXOY8gAuqfpod85JlbLPwgpllulGj8wd1J8Ma4wx06UxcYALC%2Bp%2FIAlsW9mVp0jE5kRuwkMe62cCVwQ%3D%3D\\u0026Signature=NXEdVtbPvjY5iv0I6ZTlQGp8baw%3D"
	s.mockDo([]byte(link), nil)
	defer s.assertDo()

	subAccountId := "485396905497952257"
	startTime := 1544433228000
	endTime := 1544433528000
	limit := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"subAccountId": subAccountId,
			"startTime":    startTime,
			"endTime":      endTime,
			"limit":        limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetRebateHistoryService().SubAccountId(subAccountId).
		StartTime(uint64(startTime)).EndTime(uint64(endTime)).Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := []byte(link)
	s.Equal(e, res, "Link")
}
