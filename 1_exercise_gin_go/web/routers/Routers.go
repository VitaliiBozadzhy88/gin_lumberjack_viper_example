package routers

import (
	"exercises_go/1_exercise_gin_go/web/handlers"
	"github.com/gin-gonic/gin"
)

const HOST = "localhost"
const PORT = "8080"

func NewRouter() {
	router := gin.Default()

	router.GET("/cars", handlers.GetCars)
	router.POST("/newCar", handlers.PostCar)
	router.GET("/getCarById/:id", handlers.GetCarById)

	router.Run(HOST + ":" + PORT)
}
