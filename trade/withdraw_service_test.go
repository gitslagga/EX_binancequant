package trade

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type withdrawServiceTestSuite struct {
	baseTestSuite
}

func TestWithdrawService(t *testing.T) {
	suite.Run(t, new(withdrawServiceTestSuite))
}

func (s *withdrawServiceTestSuite) TestCreateWithdraw() {
	data := []byte(`{
        "msg": "success",
        "success": true
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "ETH"
	address := "myaddress"
	amount := "0.01"
	name := "eth"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":    coin,
			"address": address,
			"amount":  amount,
			"name":    name,
		})
		s.assertRequestEqual(e, r)
	})

	err := s.client.NewCreateWithdrawService().Coin(coin).
		Address(address).Amount(amount).Name(name).Do(newContext())
	s.r().NoError(err)
}

func (s *withdrawServiceTestSuite) TestListWithdraws() {
	data := []byte(`{
        "withdrawList": [
            {
                "amount": 1,
                "address": "0x6915f16f8791d0a1cc2bf47c13a6b2a92000504b",
                "coin": "ETH",
                "applyTime": 1508198532000,
                "status": 4
            },
            {
                "amount": 0.005,
                "address": "0x6915f16f8791d0a1cc2bf47c13a6b2a92000504b",
                "txId": "0x80aaabed54bdab3f6de5868f89929a2371ad21d666f20f7393d1a3389fad95a1",
                "coin": "ETH",
                "applyTime": 1508198532000,
                "status": 4
            }
        ],
        "success": true
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "ETH"
	status := 0
	startTime := int64(1508198532000)
	endTime := int64(1508198532001)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":      coin,
			"status":    status,
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	withdraws, err := s.client.NewListWithdrawsService().Coin(coin).
		Status(status).StartTime(startTime).EndTime(endTime).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	s.Len(withdraws, 2)
	e1 := &Withdraw{
		Amount:    1,
		Address:   "0x6915f16f8791d0a1cc2bf47c13a6b2a92000504b",
		Coin:      "ETH",
		ApplyTime: 1508198532000,
		Status:    4,
	}
	e2 := &Withdraw{
		Amount:    0.005,
		Address:   "0x6915f16f8791d0a1cc2bf47c13a6b2a92000504b",
		TxID:      "0x80aaabed54bdab3f6de5868f89929a2371ad21d666f20f7393d1a3389fad95a1",
		Coin:      "ETH",
		ApplyTime: 1508198532000,
		Status:    4,
	}
	s.assertWithdrawEqual(e1, &(*withdraws)[0])
	s.assertWithdrawEqual(e2, &(*withdraws)[1])
}

func (s *withdrawServiceTestSuite) assertWithdrawEqual(e, a *Withdraw) {
	r := s.r()
	r.InDelta(e.Amount, a.Amount, 0.0000001, "Amount")
	r.Equal(e.Address, a.Address, "Address")
	r.Equal(e.Coin, a.Coin, "Coin")
	r.Equal(e.ApplyTime, a.ApplyTime, "ApplyTime")
	r.Equal(e.Status, a.Status, "Status")
}

func (s *withdrawServiceTestSuite) TestGetWithdrawFee() {
	data := []byte(`{"success": true,"withdrawFee": 0.00050}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "BTC"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("coin", coin)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetWithdrawFeeService().Coin(coin).Do(newContext())
	s.r().NoError(err)
	s.r().Equal(res.Fee, 0.0005, "Fee")
}
