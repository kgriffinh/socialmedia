package main

import (
	"log"
	"socialmedia/config"
	"socialmedia/features/users/data"
	"socialmedia/features/users/handler"
	"socialmedia/features/users/services"

	pd "socialmedia/features/posts/data"
	phl "socialmedia/features/posts/handler"
	psrv "socialmedia/features/posts/services"

	cd "socialmedia/features/comments/data"
	chl "socialmedia/features/comments/handler"
	csrv "socialmedia/features/comments/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	// // panggil fungsi Migrate untuk buat table baru di database
	config.Migrate(db)

	userData := data.New(db)
	userSrv := services.New(userData)
	userHdl := handler.New(userSrv)

	postData := pd.New(db)
	postSrv := psrv.New(postData)
	postHdl := phl.New(postSrv)

	commentData := cd.New(db)
	commentSrv := csrv.New(commentData)
	commentHdl := chl.New(commentSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/posts", postHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/posts", postHdl.GetPost(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/posts/:post_id", postHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/posts/:post_id", postHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/posts/:post_id", postHdl.GetPostDetail())

	e.POST("/comments", commentHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/comments/:comment_id", commentHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/comments/:comment_id", commentHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}

}
