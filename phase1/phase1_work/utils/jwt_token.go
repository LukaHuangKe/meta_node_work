package utils

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserInfo struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func ParseToken(tokenString string) (*UserInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserInfo{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if userInfo, ok := token.Claims.(*UserInfo); ok && token.Valid {
		return userInfo, nil
	}

	return nil, errors.New("invalid token")
}

// ========== 生成 Token ==========
func GenerateToken(userID int64) (string, error) {
	claims := &UserInfo{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JwtSecret))
}

func GetUserID(c *gin.Context) int64 {
	userID, _ := c.Get("userID")
	return userID.(int64)
}
