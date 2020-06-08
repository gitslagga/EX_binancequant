package tasks

import (
	"EX_binancequant/data"
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
获取充值地址 (支持多网络) (USER_DATA)
*/
func DepositsAddressService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	coin := c.Query("coin")
	network := c.Query("network")

	mylog.Logger.Info().Msgf("[Task Account] DepositsAddressService request param: %v, %v, %v",
		userID, coin, network)

	if coin == "" {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	depositsAddress := client.NewDepositsAddressService()
	depositsAddress.Coin(coin)
	if network != "" {
		depositsAddress.Network(network)
	}

	list, err := depositsAddress.Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = list

	c.JSON(http.StatusOK, out)
	return
}

/**
获取充值历史（支持多网络） (USER_DATA)
*/
func ListDepositsService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	coin := c.Query("coin")
	status := c.Query("status")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	offset := c.Query("offset")
	limit := c.Query("limit")

	mylog.Logger.Info().Msgf("[Task Account] ListDepositsService request param: %v, %v, %v, %v, %v, %v, %v",
		userID, coin, status, startTime, endTime, offset, limit)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	listDeposits := client.NewListDepositsService()
	if coin != "" {
		listDeposits.Coin(coin)
	}
	if status != "" {
		iStatus, err := strconv.Atoi(status)
		if err == nil {
			listDeposits.Status(iStatus)
		}
	}
	if startTime != "" {
		iStartTime, err := strconv.ParseInt(startTime, 10, 64)
		if err == nil {
			listDeposits.StartTime(iStartTime)
		}
	}
	if endTime != "" {
		iEndTime, err := strconv.ParseInt(endTime, 10, 64)
		if err == nil {
			listDeposits.EndTime(iEndTime)
		}
	}
	if offset != "" {
		iOffset, err := strconv.Atoi(offset)
		if err == nil {
			listDeposits.Offset(iOffset)
		}
	}
	if limit != "" {
		iLimit, err := strconv.Atoi(limit)
		if err == nil {
			listDeposits.Limit(iLimit)
		}
	}

	list, err := listDeposits.Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = list

	c.JSON(http.StatusOK, out)
	return
}

/**
现货账户信息 (USER_DATA)
*/
func SpotAccountService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Account] SpotAccountService request param: %v", userID)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	list, err := client.NewGetAccountService().Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = list

	c.JSON(http.StatusOK, out)
	return
}

/**
划转
*/
func FuturesTransferService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	asset := c.Query("asset")
	amount := c.Query("amount")

	var fType trade.FuturesTransferType
	err := json.Unmarshal([]byte(c.Query("type")), &fType)

	mylog.Logger.Info().Msgf("[Task Account] FuturesTransferService request param: %v, %v, %v, %v",
		userID, asset, amount, fType)

	if asset == "" || amount == "" || err != nil {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	futuresTransfer := client.NewFuturesTransferService()
	futuresTransfer.Asset(asset)
	futuresTransfer.Amount(amount)
	futuresTransfer.Type(fType)

	list, err := futuresTransfer.Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = list

	c.JSON(http.StatusOK, out)
	return
}

/**
获取划转历史
*/
func ListFuturesTransferService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	asset := c.Query("asset")
	startTime, err := strconv.ParseInt(c.Query("startTime"), 10, 64)
	endTime := c.Query("endTime")
	current := c.Query("current")
	size := c.Query("size")

	mylog.Logger.Info().Msgf("[Task Account] ListFuturesTransferService request param: %v, %v, %v, %v, %v, %v",
		userID, asset, startTime, endTime, current, size)

	if asset == "" || err != nil {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	futuresTransfer := client.NewListFuturesTransferService()
	futuresTransfer.Asset(asset)
	futuresTransfer.StartTime(startTime)
	if endTime != "" {
		iEndTime, err := strconv.ParseInt(endTime, 10, 64)
		if err == nil {
			futuresTransfer.EndTime(iEndTime)
		}
	}
	if current != "" {
		iCurrent, err := strconv.ParseInt(current, 10, 64)
		if err == nil {
			futuresTransfer.Current(iCurrent)
		}
	}
	if size != "" {
		iSize, err := strconv.ParseInt(size, 10, 64)
		if err == nil {
			futuresTransfer.Size(iSize)
		}
	}

	list, err := futuresTransfer.Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = list

	c.JSON(http.StatusOK, out)
	return
}

/**
合约账户信息 (USER_DATA)
*/
func FuturesAccountService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)

	mylog.Logger.Info().Msgf("[Task Account] FuturesAccountService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	list, err := client.NewGetAccountService().Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = list

	c.JSON(http.StatusOK, out)
	return
}

/**
提交提现请求。
*/
func CreateWithdrawService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	coin := c.Query("coin")
	withdrawOrderId := c.Query("withdrawOrderId")
	network := c.Query("network")
	address := c.Query("address")
	addressTag := c.Query("addressTag")
	amount, err := strconv.ParseFloat(c.Query("amount"), 64)
	transactionFeeFlag := c.Query("transactionFeeFlag")
	name := c.Query("name")

	if coin == "" || address != "" || err != nil {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Account] CreateWithdrawService request param: %v, %v, %v, %v, %v, %v, %v, %v, %v",
		userID, coin, withdrawOrderId, network, address, addressTag, amount, transactionFeeFlag, name)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	createWithdraw := client.NewCreateWithdrawService()
	createWithdraw.Coin(coin)
	createWithdraw.Address(address)
	createWithdraw.Amount(amount)

	if withdrawOrderId != "" {
		createWithdraw.WithdrawOrderId(withdrawOrderId)
	}
	if network != "" {
		createWithdraw.Network(network)
	}
	if addressTag != "" {
		createWithdraw.AddressTag(addressTag)
	}
	if transactionFeeFlag != "" {
		bTransactionFeeFlag, err := strconv.ParseBool(transactionFeeFlag)
		if err == nil {
			createWithdraw.TransactionFeeFlag(bTransactionFeeFlag)
		}
	}
	if name != "" {
		createWithdraw.Name(name)
	}

	err = createWithdraw.Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = ""

	c.JSON(http.StatusOK, out)
	return
}

/**
获取提币历史 (支持多网络) (USER_DATA)
*/
func ListWithdrawsService(c *gin.Context) {
	out := data.CommonResp{}

	userID := c.MustGet("user_id").(string)
	coin := c.Query("coin")
	status := c.Query("status")
	offset := c.Query("offset")
	limit := c.Query("limit")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	mylog.Logger.Info().Msgf("[Task Account] ListWithdrawsService request param: %v, %v, %v, %v, %v, %v, %v",
		userID, coin, status, startTime, endTime, offset, limit)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	listDeposits := client.NewListWithdrawsService()
	if coin != "" {
		listDeposits.Coin(coin)
	}
	if status != "" {
		iStatus, err := strconv.Atoi(status)
		if err == nil {
			listDeposits.Status(iStatus)
		}
	}
	if offset != "" {
		iOffset, err := strconv.Atoi(offset)
		if err == nil {
			listDeposits.Offset(iOffset)
		}
	}
	if limit != "" {
		iLimit, err := strconv.Atoi(limit)
		if err == nil {
			listDeposits.Limit(iLimit)
		}
	}
	if startTime != "" {
		iStartTime, err := strconv.ParseInt(startTime, 10, 64)
		if err == nil {
			listDeposits.StartTime(iStartTime)
		}
	}
	if endTime != "" {
		iEndTime, err := strconv.ParseInt(endTime, 10, 64)
		if err == nil {
			listDeposits.EndTime(iEndTime)
		}
	}

	list, err := listDeposits.Do(data.NewContext())
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	out.ErrorCode = data.EC_NONE.Code()
	out.ErrorMessage = data.EC_NONE.String()
	out.Data = list

	c.JSON(http.StatusOK, out)
	return
}
