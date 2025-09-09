package types

type CallbackResponse struct {
	Token string                 `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  map[string]interface{} `json:"user"`
}

// LoginCallbackResp 登录回调返回
type LoginCallbackResp struct {
	AccessToken string   `json:"accessToken"` // 访问令牌，用于身份验证
	Code        string   `json:"code"`        // 临时授权码
	Message     string   `json:"message"`     // 返回的消息或提示信息
	State       string   `json:"state"`       // 状态参数，用于防止 CSRF 攻击
	UserInfo    UserInfo `json:"user_info"`   // 用户信息
}

// UserInfo 表示用户信息数据结构
type UserInfo struct {
	AvatarBig       string `json:"avatar_big"`       // 大头像 URL
	AvatarMiddle    string `json:"avatar_middle"`    // 中头像 URL
	AvatarThumb     string `json:"avatar_thumb"`     // 缩略头像 URL
	AvatarURL       string `json:"avatar_url"`       // 默认头像 URL
	Email           string `json:"email"`            // 邮箱地址
	EmployeeNo      string `json:"employee_no"`      // 员工编号
	EnName          string `json:"en_name"`          // 英文姓名
	EnterpriseEmail string `json:"enterprise_email"` // 企业邮箱
	Mobile          string `json:"mobile"`           // 手机号码
	Name            string `json:"name"`             // 用户姓名
	OpenID          string `json:"open_id"`          // 第三方平台 OpenID
	TenantKey       string `json:"tenant_key"`       // 租户标识
	UnionID         string `json:"union_id"`         // 第三方平台唯一标识
	UserID          string `json:"user_id"`          // 用户唯一 ID
}
