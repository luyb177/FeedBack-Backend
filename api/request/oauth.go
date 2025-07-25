package request

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type GenerateTokenReq struct {
	// Token string `json:"token" binding:"required"`
	// 更改为 TableName
	TableName string `json:"table_name" binding:"required"`
}

type InitTokenReq struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}
