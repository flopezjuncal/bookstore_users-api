package app

import "github.com/gin-gonic/gin"

var (
	//Sera accesible SOLO desde el app package
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
