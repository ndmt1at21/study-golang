package handlers

import (
	"net/http"
	"strconv"
	"unittest/db/queries"
	"unittest/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	q queries.IUserQueries
}

func NewUserHandler(userQueries queries.IUserQueries) *UserHandler {
	return &UserHandler{q: userQueries}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.CreateUserData
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.q.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": createdUser})
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.q.GetUser(int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := h.q.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
