package im

const _API_USER_REFRESH_TOKEN_URL = "https://api.netease.im/nimserver/user/refreshToken.action"

type UserRefreshTokenParam struct {
	// 网易云信IM账号，最大长度32字符，必须保证一个
	// APP内唯一
	// 是否必须: true
	Accid string `json:"accid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BD%91%E6%98%93%E4%BA%91%E9%80%9A%E4%BF%A1ID?#%E9%87%8D%E7%BD%AE%E7%BD%91%E6%98%93%E4%BA%91%E4%BF%A1IM%20token
// 重置网易云信IM token
func (y *YunxinIM) ApiUserRefreshToken(accid string) *ImResp {
	return y.PostFrom(_API_USER_REFRESH_TOKEN_URL, UserRefreshTokenParam{accid})
}
