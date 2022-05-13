package Utils

import (
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"goapp/Models"
	"strconv"

)

type JwtCustomClaims struct {
	Id    uint `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	*jwt.StandardClaims
}

func JwtGeneratorToken(user Models.User) (t string, err error){
	claims := &JwtCustomClaims{
		user.Id,
		user.Name,
		user.Email,
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute*20).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err = token.SignedString([]byte(viper.GetString("PRIVATE_KEY")))
	return t, err
}

func JwtGeneratorRefreshToken(user Models.User) (t string, err error){
	claims := &JwtCustomClaims{
		user.Id,
		user.Name,
		user.Email,
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*300).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err = token.SignedString([]byte(viper.GetString("REFRESH_PRIVATE_KEY")))
	return t, err
}


func ParseJWT(tokenString string) (*JwtCustomClaims, error){
    token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(viper.GetString("REFRESH_PRIVATE_KEY")), nil
    })

    if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
        return claims, nil
    } else {
        return nil, err
    }
}

func GetUserID(user *jwt.Token)(uint){
	claims := user.Claims.(*JwtCustomClaims)
	userId := claims.Id
	return userId
}

func UnitToString(u uint)(string){
	i := strconv.FormatUint(uint64(u), 10)
	return i
}