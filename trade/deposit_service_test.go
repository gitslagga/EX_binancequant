package trade

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type depositServiceTestSuite struct {
	baseTestSuite
}

func TestDepositService(t *testing.T) {
	suite.Run(t, new(depositServiceTestSuite))
}

func (s *depositServiceTestSuite) TestListDeposits() {
	data := []byte(`
    [
		{
			"address": "0xddc66e4313fd6c737b6cae67cad90bb4e0ac7092",
			"addressTag": "",
			"amount": "139.04370000",
			"coin": "USDT",
			"insertTime": 1566791463000,
			"network": "ETH",
			"status": 1,
			"TxID": "0x5759dfe9983a4c7619bce9bc736bb6c26f804091753bf66fa91e7cd5cfeebafd"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":      "USDT",
			"status":    1,
			"startTime": 1508198532000,
			"endTime":   1508198532001,
		})
		s.assertRequestEqual(e, r)
	})
	deposits, err := s.client.NewListDepositsService().Coin("USDT").
		Status(1).StartTime(1508198532000).EndTime(1508198532001).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(deposits, 1)
	e := &Deposit{
		Address:    "0xddc66e4313fd6c737b6cae67cad90bb4e0ac7092",
		AddressTag: "",
		Amount:     "139.04370000",
		Coin:       "USDT",
		InsertTime: 1566791463000,
		Network:    "ETH",
		Status:     1,
		TxID:       "0x5759dfe9983a4c7619bce9bc736bb6c26f804091753bf66fa91e7cd5cfeebafd",
	}
	s.assertDepositEqual(e, deposits[0])
}

func (s *depositServiceTestSuite) assertDepositEqual(e, a *Deposit) {
	r := s.r()
	r.Equal(e.Address, a.Address, "Address")
	r.Equal(e.AddressTag, a.AddressTag, "AddressTag")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Coin, a.Coin, "Coin")
	r.Equal(e.InsertTime, a.InsertTime, "InsertTime")
	r.Equal(e.Coin, a.Coin, "Network")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TxID, a.TxID, "TxID")
}
