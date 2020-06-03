package tasks

import (
	"EX_binancequant/data"
	"EX_binancequant/db"
	"EX_binancequant/mylog"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
获取充值历史 (USER_DATA)
*/
func ListDepositsService(c *gin.Context) {
	out := data.CommonResp{}

	token := c.GetHeader("token")
	userID, err := db.ConvertTokenToUserID(token)
	asset := c.Query("asset")
	status, _ := strconv.Atoi(c.Query("status"))
	startTime, _ := strconv.Atoi(c.Query("startTime"))

	mylog.Logger.Info().Msgf("[Task Service] GetSwapInstrumentPosition request param: %s", userID)

	if err != nil {
		out.ErrorCode = data.EC_PARAMS_ERR
		out.ErrorMessage = data.ErrorCodeMessage(data.EC_PARAMS_ERR)
		c.JSON(http.StatusBadRequest, out)
		return
	}

	client, err := db.GetClientByUserID(userID)
	if err != nil {
		out.ErrorCode = data.EC_NETWORK_ERR
		out.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, out)
		return
	}

	listDeposits := client.NewListDepositsService()
	listDeposits.Asset(asset)
	listDeposits.Status(status)
	listDeposits.StartTime(int64(startTime))

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
