package tasks

import (
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"EX_binancequant/trade"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
获取充值地址 (支持多网络) (USER_DATA)
*/
func DepositsAddressService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var depositsAddressRequest DepositsAddressRequest
	err := c.ShouldBindQuery(&depositsAddressRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Account] DepositsAddressService request param: %v, %v",
		userID, depositsAddressRequest)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	depositsAddress := client.NewDepositsAddressService()
	depositsAddress.Coin(depositsAddressRequest.Coin)
	if depositsAddressRequest.Network != "" {
		depositsAddress.Network(depositsAddressRequest.Network)
	}

	list, err := depositsAddress.Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	out.RespData = list

	c.Set("responseData", out)
}

/**
获取充值历史（支持多网络） (USER_DATA)
*/
func ListDepositsService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var listDepositsRequest ListDepositsRequest
	err := c.ShouldBindQuery(&listDepositsRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Account] ListDepositsService request param: %v, %v",
		userID, listDepositsRequest)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	listDeposits := client.NewListDepositsService()
	if listDepositsRequest.Coin != "" {
		listDeposits.Coin(listDepositsRequest.Coin)
	}
	if listDepositsRequest.Status != 0 {
		listDeposits.Status(listDepositsRequest.Status)
	}
	if listDepositsRequest.StartTime != 0 {
		listDeposits.StartTime(listDepositsRequest.StartTime)
	}
	if listDepositsRequest.EndTime != 0 {
		listDeposits.EndTime(listDepositsRequest.EndTime)
	}
	if listDepositsRequest.Offset != 0 {
		listDeposits.Offset(listDepositsRequest.Offset)
	}
	if listDepositsRequest.Limit != 0 {
		listDeposits.Limit(listDepositsRequest.Limit)
	}

	list, err := listDeposits.Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	out.RespData = list

	c.Set("responseData", out)
}

/**
现货账户信息 (USER_DATA)
*/
func SpotAccountService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	mylog.Logger.Info().Msgf("[Task Account] SpotAccountService request param: %v", userID)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	list, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	out.RespData = list

	c.Set("responseData", out)
}

/**
划转
*/
func FuturesTransferService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var transferRequest TransferRequest
	err := json.Unmarshal(c.MustGet("requestData").([]byte), &transferRequest)
	if err != nil || transferRequest.Asset == "" || transferRequest.Amount == 0 || transferRequest.Type == 0 {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Account] FuturesTransferService request param: %v, %v",
		userID, transferRequest)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	futuresTransfer := client.NewFuturesTransferService()
	futuresTransfer.Asset(transferRequest.Asset)
	futuresTransfer.Amount(transferRequest.Amount)
	futuresTransfer.Type(trade.FuturesTransferType(transferRequest.Type))

	list, err := futuresTransfer.Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	out.RespData = list

	c.Set("responseData", out)
}

/**
获取划转历史
*/
func ListFuturesTransferService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var listFuturesTransferRequest ListFuturesTransferRequest
	err := c.ShouldBindQuery(&listFuturesTransferRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Account] ListFuturesTransferService request param: %v, %v",
		userID, listFuturesTransferRequest)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	futuresTransfer := client.NewListFuturesTransferService()
	futuresTransfer.Asset(listFuturesTransferRequest.Asset)
	futuresTransfer.StartTime(listFuturesTransferRequest.StartTime)
	if listFuturesTransferRequest.EndTime != 0 {
		futuresTransfer.EndTime(listFuturesTransferRequest.EndTime)
	}
	if listFuturesTransferRequest.Current != 0 {
		futuresTransfer.Current(listFuturesTransferRequest.Current)
	}
	if listFuturesTransferRequest.Size != 0 {
		futuresTransfer.Size(listFuturesTransferRequest.Size)
	}

	list, err := futuresTransfer.Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	out.RespData = list

	c.Set("responseData", out)
}

/**
合约账户信息 (USER_DATA)
*/
func FuturesAccountService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	mylog.Logger.Info().Msgf("[Task Account] FuturesAccountService request param: %v", userID)

	client, err := db.GetFuturesClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	list, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	out.RespData = list

	c.Set("responseData", out)
}

/**
提交提现请求。
*/
func CreateWithdrawService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var withdrawRequest WithdrawRequest
	err := json.Unmarshal(c.MustGet("requestData").([]byte), &withdrawRequest)
	if err != nil || withdrawRequest.Coin == "" || withdrawRequest.Address == "" || withdrawRequest.Amount == 0 {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Account] CreateWithdrawService request param: %v, %v",
		userID, withdrawRequest)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	//获取用户现货余额
	account, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	var balance float64
	for _, v := range account.Balances {
		if v.Asset == withdrawRequest.Coin {
			balance, err = strconv.ParseFloat(v.Free, 64)
			if err != nil {
				out.RespCode = EC_FORMAT_ERR
				out.RespDesc = err.Error()
				c.Set("responseData", out)
				return
			}
			break
		}
	}

	if balance < withdrawRequest.Amount {
		out.RespCode = EC_NO_BALANCE
		out.RespDesc = ErrorCodeMessage(EC_NO_BALANCE)
		c.Set("responseData", out)
		return
	}

	//获取用户子账户ID
	subAccountID, err := db.GetSubAccountIdByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	//从子账户现货账户往经纪人现货账户划转
	_, err = trade.BAExClient.NewCreateTransferService().
		Asset(withdrawRequest.Coin).Amount(withdrawRequest.Amount).FromId(subAccountID).Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	//经纪人api key签名发起提币
	createWithdraw := trade.BAExClient.NewCreateWithdrawService()
	createWithdraw.Coin(withdrawRequest.Coin)
	createWithdraw.Address(withdrawRequest.Address)
	createWithdraw.Amount(withdrawRequest.Amount)

	if withdrawRequest.WithdrawOrderId != "" {
		createWithdraw.WithdrawOrderId(withdrawRequest.WithdrawOrderId)
	}
	if withdrawRequest.Network != "" {
		createWithdraw.Network(withdrawRequest.Network)
	}
	if withdrawRequest.AddressTag != "" {
		createWithdraw.AddressTag(withdrawRequest.AddressTag)
	}
	if withdrawRequest.TransactionFeeFlag != false {
		createWithdraw.TransactionFeeFlag(withdrawRequest.TransactionFeeFlag)
	}
	if withdrawRequest.Name != "" {
		createWithdraw.Name(withdrawRequest.Name)
	}

	err = createWithdraw.Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()

	c.Set("responseData", out)
}

/**
获取提币历史 (支持多网络) (USER_DATA)
*/
func ListWithdrawsService(c *gin.Context) {
	out := CommonResp{}

	userID := c.MustGet("user_id").(uint64)

	var listWithdrawsRequest ListWithdrawsRequest
	err := c.ShouldBindQuery(&listWithdrawsRequest)
	if err != nil {
		out.RespCode = EC_PARAMS_ERR
		out.RespDesc = ErrorCodeMessage(EC_PARAMS_ERR)
		c.Set("responseData", out)
		return
	}

	mylog.Logger.Info().Msgf("[Task Account] ListWithdrawsService request param: %v, %v",
		userID, listWithdrawsRequest)

	client, err := db.GetSpotClientByUserID(userID)
	if err != nil {
		out.RespCode = EC_NOT_ACTIVE
		out.RespDesc = ErrorCodeMessage(EC_NOT_ACTIVE)
		c.Set("responseData", out)
		return
	}

	listDeposits := client.NewListWithdrawsService()
	if listWithdrawsRequest.Coin != "" {
		listDeposits.Coin(listWithdrawsRequest.Coin)
	}
	if listWithdrawsRequest.Status != 0 {
		listDeposits.Status(listWithdrawsRequest.Status)
	}
	if listWithdrawsRequest.StartTime != 0 {
		listDeposits.StartTime(listWithdrawsRequest.StartTime)
	}
	if listWithdrawsRequest.EndTime != 0 {
		listDeposits.EndTime(listWithdrawsRequest.EndTime)
	}
	if listWithdrawsRequest.Offset != 0 {
		listDeposits.Offset(listWithdrawsRequest.Offset)
	}
	if listWithdrawsRequest.Limit != 0 {
		listDeposits.Limit(listWithdrawsRequest.Limit)
	}

	list, err := listDeposits.Do(context.Background())
	if err != nil {
		out.RespCode = EC_NETWORK_ERR
		out.RespDesc = err.Error()
		c.Set("responseData", out)
		return
	}

	out.RespCode = EC_NONE.Code()
	out.RespDesc = EC_NONE.String()
	out.RespData = list

	c.Set("responseData", out)
}
