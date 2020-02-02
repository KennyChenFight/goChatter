package handler

import (
	"errors"
	"github.com/KennyChenFight/goChatter/lib/auth"
	"github.com/KennyChenFight/goChatter/lib/constant"
	"github.com/KennyChenFight/goChatter/lib/httputil"
	"github.com/KennyChenFight/goChatter/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"xorm.io/xorm"
)

func UserCreate(c *gin.Context) {
	var user struct {
		model.User `xorm:"extends"`
		Password   string `xorm:"-" json:"password" binding:"required"`
	}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusBadRequest)
		c.Set(constant.Error, err)
		return
	}

	uid, err := uuid.NewRandom()
	if err != nil {
		c.Set(constant.StatusCode, http.StatusInternalServerError)
		c.Set(constant.Error, err)
		return
	}
	user.Id = uid.String()

	if digest, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil {
		c.Set(constant.StatusCode, http.StatusInternalServerError)
		c.Set(constant.Error, err)
		return
	} else {
		user.PasswordDigest = string(digest)
	}
	q := `insert into users(id, email, password_digest, name, self_introduction, picture)
			select ?, ?, ?, ?, ?, ? 
			where not exists (select 1 from users where email = ?)`
	db := c.MustGet(constant.DB).(*xorm.Engine)
	result, err := db.Exec(q, user.Id, user.Email, user.PasswordDigest, user.Name,
		user.SelfIntroduction, user.Picture, user.Email)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusInternalServerError)
		c.Set(constant.Error, err)
		return
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		c.Set(constant.StatusCode, http.StatusForbidden)
		c.Set(constant.Error, errors.New("the email is already used"))
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
		c.Set(constant.StatusCode, http.StatusCreated)
		c.Set(constant.Output, map[string]string{"userId": user.Id})
		return
	}
}

func UserUpdate(c *gin.Context) {
	id := c.Param("id")
	userId := c.MustGet(constant.UserId).(string)
	if id != userId {
		c.Set(constant.StatusCode, http.StatusForbidden)
		c.Set(constant.Error, errors.New("updating others user account is forbidden"))
		return
	}

	var input struct {
		model.User       `xorm:"extends"`
		Password         *string `xorm:"-" json:"password" binding:"omitempty,min=1"`
		OriginalPassword *string `xorm:"-" json:"originalPassword" binding:"omitempty,min=1"`
	}
	dbUpdateFields, err := httputil.BindForUpdate(c, &input)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusBadRequest)
		c.Set(constant.Error, err)
		return
	}

	session := c.MustGet(constant.DbSession).(*xorm.Session)
	if input.Password != nil {
		if input.OriginalPassword == nil {
			c.Set(constant.StatusCode, http.StatusForbidden)
			c.Set(constant.Error, errors.New("please provide the original password"))
			return
		} else {
			var user model.User
			if found, err := session.ID(userId).Get(&user); !found {
				c.Set(constant.StatusCode, http.StatusNotFound)
				c.Set(constant.Error, errNotFound)
				return
			} else {
				if err != nil {
					c.Set(constant.StatusCode, http.StatusInternalServerError)
					c.Set(constant.Error, err)
					return
				}
				if bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(*input.OriginalPassword)) != nil {
					c.Set(constant.StatusCode, http.StatusForbidden)
					c.Set(constant.Error, errors.New("the original password is valid"))
					return
				}
			}

			if digest, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost); err != nil {
				c.Set(constant.StatusCode, http.StatusInternalServerError)
				c.Set(constant.Error, err)
				return
			} else {
				input.PasswordDigest = string(digest)
				dbUpdateFields["password_digest"] = true
			}
		}
	}

	//convert the columnName map into string slice
	var columnNames []string
	for k := range dbUpdateFields {
		columnNames = append(columnNames, k)
	}

	//perform the update to the database
	affected, err := session.ID(userId).Cols(columnNames...).Update(&input)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusInternalServerError)
		c.Set(constant.Error, err)
		return
	} else {
		if affected == 0 {
			c.Set(constant.StatusCode, http.StatusNotFound)
			c.Set(constant.Error, errNotFound)
			return
		} else {
			c.Set(constant.StatusCode, http.StatusNoContent)
			c.Set(constant.Output, nil)
			return
		}
	}
}
