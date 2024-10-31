package user

import (
	"go-htmx/internal/app"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const groupUser = "user"

type ProfileView struct {
	log            *zap.Logger
	getUserUseCase app.GetUserUseCase
}

func NewProfileView(log *zap.Logger, getUserUseCase app.GetUserUseCase) *ProfileView {
	return &ProfileView{log: log, getUserUseCase: getUserUseCase}
}

func (p *ProfileView) Group() string {
	return groupUser
}

func (p *ProfileView) Method() string {
	return http.MethodGet
}

func (p *ProfileView) Path() string {
	return "/:user_id"
}

func (v *ProfileView) Handle(c *gin.Context) {
	v.log.Info("ProfileView.Handle")
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	user, err := v.getUserUseCase.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.HTML(http.StatusOK, "user/profile", gin.H{
		"User": user,
	})
}
