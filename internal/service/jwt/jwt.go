package jwt

import (
	"baby-daily-api/internal/model"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

const (
	IssUser = "baby.lyranxi.com"
	Secret  = "nyxXczTNQV3sx5td"
	Expire  = 24 * 15 * 3600
)

type CustomClaims struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	OpenId   string `json:"openId"`
	jwt.StandardClaims
}

type JwtService struct {
	signKey        []byte
	issuer         string
	expireDuration int64
}

func NewJwtService() *JwtService {
	return &JwtService{
		signKey:        []byte(Secret),
		issuer:         IssUser,
		expireDuration: int64(24 * 15 * time.Hour.Seconds()),
	}
}

// CreateToken 生成token
func (s *JwtService) CreateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			Id:       user.ID,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			OpenId:   user.OpenId,
			StandardClaims: jwt.StandardClaims{
				Issuer:    s.issuer,
				ExpiresAt: time.Now().Unix() + s.expireDuration,
				NotBefore: time.Now().Unix() - 1000,
				Id:        strconv.Itoa(user.ID),
			},
		},
	)
	return token.SignedString(s.signKey)
}

// ParseToken 解析token
func (s *JwtService) ParseToken(tokenString string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return s.signKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invaild token")
	}

	return &model.User{
		ID:       claims.Id,
		Nickname: claims.Nickname,
		OpenId:   claims.OpenId,
		Avatar:   claims.Avatar,
	}, nil
}
