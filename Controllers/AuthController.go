package Controllers
import (
	"net/http"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"goapp/Models"
	"goapp/Utils"
)

type (
	LoginForm struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,max=60,min=6"`
	}

	RegisterForm struct {
		Name     string `json:"name" validate:"required,max=60,min=3"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,max=60,min=6"`
	}

	RefreshForm struct {
		RefreshToken     string `json:"refresh_token" validate:"required"`
	}
)

func Login(c echo.Context) (err error){
	var loginParam LoginForm
	if err = c.Bind(&loginParam); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "Invalid Value",
			})
	}
	if err = c.Validate(loginParam); err != nil {
		return c.JSON(http.StatusInternalServerError,
			Resp{
				Error: err.Error(),
				Message: "Have Error",
		})
	}
	var user Models.User
	err = Models.GetAUserByEmail(&user, loginParam.Email)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Message: "Email or Password Invalid",
			})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginParam.Password))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Message: "Email or Password Invalid",
			})
	}
	token, err := Utils.JwtGeneratorToken(user)
	refreshToken, err := Utils.JwtGeneratorRefreshToken(user)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			Resp{
				Message: "Error create token",
			})
	}
	user.RefreshToken = refreshToken
	err = Models.UpdateAUser(&user)
	return c.JSON(
		http.StatusCreated,
		Resp{
			Result: map[string]string{
				"access_token": token,
				"refresh_token": refreshToken,
			},
			Message: "Login successfully",
		})
}

func Refresh(c echo.Context) (err error){
	var refreshParam RefreshForm
	if err = c.Bind(&refreshParam); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "Invalid Value",
			})
	}
	if err = c.Validate(refreshParam); err != nil {
		return c.JSON(http.StatusInternalServerError,
			Resp{
				Error: err.Error(),
				Message: "Have Error",
		})
	}
	claims, err := Utils.ParseJWT(refreshParam.RefreshToken)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "User Invalid",
			})
	}
	var user Models.User
	err = Models.GetAUserForRefresh(&user, claims.Id, refreshParam.RefreshToken)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Message: "User Invalid",
			})
	}
	token, err := Utils.JwtGeneratorToken(user)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			Resp{
				Message: "Error create token",
			})
	}

	return c.JSON(
		http.StatusCreated,
		Resp{
			Result: map[string]string{
				"access_token": token,
			},
			Message: "Login successfully",
		})
}

func Register(c echo.Context) (err error){
	var registerParam RegisterForm
	if err = c.Bind(&registerParam); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			Resp{
				Error: err.Error(),
				Message: "Invalid Value",
			})
	}
	if err = c.Validate(registerParam); err != nil {
		return c.JSON(http.StatusInternalServerError,
			Resp{
				Error: err.Error(),
				Message: "Have Error",
		})
	}
	count, err := Models.CountByEmail(registerParam.Email)
	if err != nil{
		return c.JSON(http.StatusInternalServerError,
			Resp{
				Error: err.Error(),
				Message: "Have Error",
		})
	}
	if count > 0{
		return c.JSON(http.StatusBadRequest,
			Resp{
				Error: "Email Already Used",
				Message: "Invalid Value",
		})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(registerParam.Password), 5)
	if err != nil{
		return c.JSON(http.StatusInternalServerError,
			Resp{
				Error: err.Error(),
				Message: "Have Error",
		})
	}
	var user Models.User
	user.Email = registerParam.Email
	user.Password = hash
	user.Name = registerParam.Name
	err = Models.CreateAUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
		Resp{
			Error: err.Error(),
			Message: "Have Error",
		})
	}

	return c.JSON(
		http.StatusCreated,
		Resp{
			Result: user,
			Message: "Register successfully",
		})
}
