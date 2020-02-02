package handler

import (
	"errors"
	"github.com/KennyChenFight/goChatter/lib/auth"
	"github.com/KennyChenFight/goChatter/lib/constant"
	"github.com/KennyChenFight/goChatter/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"xorm.io/xorm"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusBadRequest)
		c.Set(constant.Error, err)
		return
	}

	var user model.User
	db := c.MustGet(constant.DB).(*xorm.Engine)
	found, err := db.Where("email = ?", input.Email).Get(&user)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusInternalServerError)
		c.Set(constant.Error, err)
		return
	}

	if !found || bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(input.Password)) != nil {
		c.Set(constant.StatusCode, http.StatusUnauthorized)
		c.Set(constant.Error, errors.New("incorrect email or password"))
		return
	}

	if newToken, err := auth.Sign(user.Id); err != nil {
		c.Set(constant.StatusCode, http.StatusInternalServerError)
		c.Set(constant.Error, err)
		return
	} else {
		// update JWT Token
		c.Header("Authorization", newToken)
		// allow CORS
		c.Header("Access-Control-Expose-Headers", "Authorization")

		c.Set(constant.StatusCode, http.StatusOK)
		c.Set(constant.Output, map[string]string{"userId": user.Id})
	}
}
