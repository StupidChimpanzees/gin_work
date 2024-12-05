package common

import (
	"gin_work/extend/jwt"
	"gin_work/wrap/config"
)

func CheckToken(accessToken, refreshToken, domain, ip string) (*jwt.TokenClaims, bool, bool) {
	if accessToken == "" || refreshToken == "" {
		return nil, false, false
	}
	var claims *jwt.TokenClaims
	var err error
	claims, err = jwt.ParseToken(accessToken, domain)
	// access过期
	if err != nil {
		claims, err = jwt.ParseToken(refreshToken, domain)
		// refresh过期
		if err != nil {
			return nil, false, false
		}
		// ip不一致
		if claims.Ip != ip {
			return nil, false, false
		}
		// access过期,refresh不过期
		return claims, false, true
	} else if claims.Ip != ip {
		return nil, false, false
	}
	// access不过期,refresh不过期
	return claims, true, true
}

func RefreshNewToken(uuid, domain, ip string) (accessToken, refreshToken string, err error) {
	conf := config.Mapping.JWT
	accessToken, err = jwt.GenerateToken(uuid, domain, ip)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = jwt.GenerateToken(uuid, domain, ip, conf.RefreshExpires)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err
}
