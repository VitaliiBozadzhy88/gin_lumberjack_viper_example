package routers

import (
	handlers2 "exercises_go/2_exercise_lumberjack_go/app/web/handlers"
	"exercises_go/2_exercise_lumberjack_go/logging"
	"github.com/gin-gonic/gin"
)

const HOST = "localhost"
const PORT = "8080"

func NewRouter() {
	router := gin.Default()

	router.GET("/cars", handlers2.GetCars)
	router.POST("/newCar", handlers2.PostCar)
	router.GET("/getCarById/:id", handlers2.GetCarById)

	err := router.Run(HOST + ":" + PORT)
	if err != nil {
		logging.LogINFO(err.Error())
	}
}
