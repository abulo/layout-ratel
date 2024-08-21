package dao

import (
	"github.com/golang-jwt/jwt/v5"
)

// SystemUserToken 用户令牌
type SystemUserToken struct {
	UserId   int64  `json:"userId"`   // 用户ID
	UserName string `json:"userName"` // 用户名
	NickName string `json:"nickName"` // 用户名称
	TenantId int64  `json:"tenantId"` // 租户ID
	jwt.RegisteredClaims
}
