package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (u *UserController) GetUserById(ctx *gin.Context)  {
	idStr :=ctx.Param("id")
	userId, _ := strconv.ParseInt(idStr, 10, 64)
	user, _ :=u.srv.Users().GetUserById(userId)
	ctx.JSON(http.StatusOK, user)
}
