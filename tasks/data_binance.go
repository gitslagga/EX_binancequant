package tasks

type BinanceCode int

func (c BinanceCode) String() (r string) {
	switch c {
	case -1000:
		r = "处理请求时发生未知错误"
	case -1001:
		r = "内部错误; 无法处理您的请求。 请再试一次"
	case -1002:
		r = "您无权执行此请求参数错误"
	case -1003:
		r = "排队的请求过多"

	case -1106:
		r = "发送了不需要的参数"
	case -1111:
		r = "精度超过为此资产定义的最大值"
	case -1112:
		r = "交易对没有挂单"
	case -1114:
		r = "不需要时发送了TimeInForce参数"
	case -1115:
		r = "无效的timeInForce"
	case -1116:
		r = "无效订单类型"
	case -1117:
		r = "无效买卖方向"
	case -1118:
		r = "新的客户订单ID为空"
	case -1119:
		r = "客户自定义的订单ID为空"
	case -1120:
		r = "无效时间间隔"
	case -1121:
		r = "无效的交易对"
	case -1125:
		r = "该listenKey不存在"
	case -1127:
		r = "查询间隔太大"
	case -1128:
		r = "可选参数组合无效"
	case -1130:
		r = "发送的参数为无效数据"

	case -2010:
		r = "新订单被拒绝"
	case -2011:
		r = "取消订单被拒绝"
	case -2018:
		r = "余额不足"
	case -2019:
		r = "杠杆账户余额不足"
	case -2020:
		r = "无法成交"

	case -3022:
		r = "账号被禁止交易"

	case -4000:
		r = "订单状态不正确"
	case -4001:
		r = "价格小于0"
	case -4002:
		r = "价格超过最大值"
	case -4003:
		r = "数量小于0"
	case -4004:
		r = "数量小于最小值"
	case -4005:
		r = "数量大于最大值"
	case -4006:
		r = "触发价小于最小值"
	case -4007:
		r = "触发价大于最大值"
	case -4008:
		r = "价格精度小于0"
	case -4009:
		r = "最大价格小于最小价格"
	case -4010:
		r = "最大数量小于最小数量"
	case -4024:
		r = "价格低于标价乘数底限"

	default:
	}
	return
}

func BinanceCodeMessage(c int) (r string) {
	return BinanceCode(c).String()
}
