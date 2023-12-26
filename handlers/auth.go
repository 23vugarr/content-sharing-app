package handlers

import (
	"fmt"
	"net/http"

	_ "github.com/23vugarr/content-sharing-app/cmd/docs"
	"github.com/23vugarr/content-sharing-app/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *models.DB
}

func PostLogin(r *gin.Context) {
	fmt.Println("inside post login")
	r.IndentedJSON(http.StatusOK, gin.H{"test": "hello"})
}

func PostRegister(r *gin.Context) {
	fmt.Println("inside post register")
	r.IndentedJSON(http.StatusOK, gin.H{"message": "true"})
}

// GetUser       godoc
// @Summary      Get books array
// @Description  Responds with the list of all books as JSON.
// @Tags         books
// @Produce      json
// @Success      200  {array}  models.Book
// @Router       /books [get]
func (h *Handler) GetUser(r *gin.Context) {
	username := r.Param("username")
	fmt.Println(username)

	user, err := h.DB.FetchUserByUsername(username)
	if err != nil {
		r.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	r.JSON(http.StatusOK, gin.H{"user": user})

}
