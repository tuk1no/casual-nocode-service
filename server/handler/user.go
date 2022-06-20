package handler

import (
	"casual-nocode-service/models"
	"casual-nocode-service/token"
	"net/http"

	"github.com/labstack/echo"
)

// ログインリクエスト
type loginRequest struct {
	Email    string
	Password string
}

// ユーザー登録リクエスト
type signupRequest struct {
	Name     string
	Email    string
	Password string
}

// ユーザーログイン
func (h *Handler) LoginUser(c echo.Context) error {
	param := new(loginRequest)
	if err := c.Bind(param); err != nil {
		return echo.ErrBadRequest
	}

	// 登録済みユーザの検索
	user := new(models.User)
	result := h.DB.Users.
		Where("email == ?", param.Email).
		Limit(1).
		Find(&user)
	if result.Error != nil {
		return result.Error
	}

	// アカウントが無い
	if result.RowsAffected <= 0 {
		return echo.ErrUnauthorized
	}

	// パスワードが異なる
	if !user.Login(param.Password) {
		return echo.ErrUnauthorized
	}

	// トークン生成
	token, err := token.CreateToken(*user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// ユーザ作成
func (h *Handler) SignUpUser(c echo.Context) error {
	param := new(signupRequest)
	if err := c.Bind(param); err != nil {
		return echo.ErrBadRequest
	}

	result := h.DB.Users.
		Where("email == ?", param.Email).
		Limit(1).
		Find(&models.User{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		// emailが登録済み
		return c.NoContent(http.StatusConflict)
	}

	user := models.User{Name: param.Name, Email: param.Email}
	err := user.SetPassword(param.Password)
	if err != nil {
		return err
	}
	result = h.DB.Users.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return c.NoContent(http.StatusOK)
}

// ユーザー情報取得
func (h *Handler) GetUserInfo(c echo.Context) error {
	id := token.GetId(c)
	user := models.User{}
	h.DB.Users.First(&user, id)
	return c.JSON(http.StatusOK, echo.Map{
		"username": user.Name,
		"email":    user.Email,
	})
}
