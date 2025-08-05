package jwtinfo

import "github.com/golang-jwt/jwt/v5"

const JwtKey = "oepuqtpowejfvgc;lvmjn290331pq;woerwqje[p]"

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
}
