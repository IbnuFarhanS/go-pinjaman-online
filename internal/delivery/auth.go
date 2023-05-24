package delivery

import (
	"net/http"
	"os"
	"time"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type authHandler struct {
	borrowerUsecase usecase.BorrowerUsecase
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) { return os.Getenv("jwtKey"), nil })

		if !token.Valid || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("claims", claims)

		c.Next()
	}
}

func LoginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var bor entity.Borrower
	if err := c.ShouldBind(&bor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(bor.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// logic authentication (compare username dan password)
	if user.Username == bor.Username && user.Password == bor.Password {
		// bikin code untuk generate token
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)

		claims["username"] = user.Username
		claims["exp"] = time.Now().Add(time.Minute * 5).Unix() // token akan expired dalam 1 menit

		tokenString, err := token.SignedString(os.Getenv("jwtKey"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString}) // jika login berhasil, dapatkan token string
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
