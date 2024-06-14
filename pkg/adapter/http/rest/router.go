package rest

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(pokemon PokemonHandler, token string) *gin.Engine {
	r := gin.Default()
	corsCfg := cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"Accept",
			"Origin",
			"Cache-Control",
			"X-Requested-With",
		},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}
	r.Use(cors.New(corsCfg))
	_ = r.Group("/")
	{
		poke := r.Group("pokemon")
		{
			poke.Use(TokenAuthMiddleware(token))
			poke.GET("/:id", pokemon.GetPokemon)
			poke.POST("/", pokemon.PostPokemon)
		}
	}
	return r
}

func TokenAuthMiddleware(requiredtoken string) gin.HandlerFunc {
	if requiredtoken == "" {
		panic("Please set API_TOKEN environment variable")
	}
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"message": "Token required"})
			c.Abort()
			return
		}
		if token != requiredtoken {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid API token"})
			return
		}
		c.Next()
	}
}
