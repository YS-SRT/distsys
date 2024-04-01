package dto

type UserVO struct {
	ID           int    `json:"id" xml:"id"`
	Nickname     string `json:"nickname" xml:"nickname"`
	Portrait     string `json:"portrait" xml:"portrait"`
	Phone        string `json:"phone" xml:"phone"`
	Token        string `json:"token" xml:"token"`
	RefreshToken string `json:"refresh_token" xml:"refresh_token"`
}

type UserLoginByPWDInput struct {
	Phone    string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required,max=20,min=6"`
}

type UserLoginByWechatIDSInput struct {
	//wechat already verified, just trust it
	WechatIDS string `form:"wechat_ids" json:"wechat_ids" xml:"wechat_ids"  binding:"required,max=255"`
}

type UserRegisterInput struct {
	Nickname   string `form:"nickname" json:"nickname" xml:"nickname"  binding:"max=30"`
	WechatUIDS string `form:"wechat_ids" json:"wechat_ids" xml:"wechat_ids"  binding:"max=255"`
	Portrait   string `form:"portrait" json:"portrait" xml:"portrait"  binding:"max=100"`
	Phone      string `form:"phone" json:"phone" xml:"phone"  binding:"max=20"`
	Password   string `form:"password" json:"password" xml:"password"  binding:"max=20,min=6"`
}
