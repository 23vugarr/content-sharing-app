package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostLogin(r *gin.Context) {
	fmt.Println("inside post login")
	r.IndentedJSON(http.StatusOK, gin.H{"test": "hello"})
}

func PostRegister(r *gin.Context) {
	fmt.Println("inside post register")
	r.IndentedJSON(http.StatusOK, gin.H{"message": "true"})
}
