package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaiquecaires/hexagon-finances-server/adapters/driven/repository"
	"github.com/kaiquecaires/hexagon-finances-server/domain/entities"
	"github.com/kaiquecaires/hexagon-finances-server/domain/services"
)

type SignupProps struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	SocialName string `json:"social_name"`
}

func Start() {
	r := gin.Default()

	r.POST("/signup", func(ctx *gin.Context) {
		var data SignupProps

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error on binding json"})
			return
		}

		user, err := entities.CreateUser(data.Email, data.Name, data.SocialName)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		signup := services.Signup{
			UserRepository: repository.UserRepository{},
		}

		createdUser, err := signup.Signup(user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, createdUser)
	})

	r.Run(":5001")
}
