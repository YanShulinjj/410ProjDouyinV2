/* ----------------------------------
*  @author suyame 2022-10-13 17:20:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package jwtx

import (
	"410proj/pkg/xerr"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type MyClaims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

func GetToken(secretKey string, iat, seconds int64, uid uint64) (string, error) {
	claims := MyClaims{
		UserId: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: iat + seconds,
			IssuedAt:  iat,
		},
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// 验证token
func ParseToken(secretKey string, token string) (*MyClaims, xerr.ErrCodeType) {
	// 对token的密钥进行验证
	tokenCaims, err := jwt.ParseWithClaims(token, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, xerr.TokenNotMatchErr
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, xerr.TokenExpiredErr
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, xerr.TokenNotActiveErr
			} else {
				return nil, xerr.TokenNotMatchErr
			}
		}
	}

	// 判断token是否过期
	if claims, ok := tokenCaims.Claims.(*MyClaims); ok && tokenCaims.Valid {
		return claims, 0
	}
	return nil, xerr.TokenNotMatchErr
}

func IsValid(secretKey string, token string) (bool, error) {
	_, errCode := ParseToken(secretKey, token)
	if errCode != 0 {
		return false, errors.Wrapf(xerr.NewErrCode(errCode),
			"Token 验证失败")
	}
	return true, nil
}

func GetUserId(secretKey string, token string) (uint64, error) {
	claim, errCode := ParseToken(secretKey, token)
	if errCode != 0 {
		return 0, errors.Wrapf(xerr.NewErrCode(errCode),
			"Token 验证失败")
	}
	return claim.UserId, nil
}
