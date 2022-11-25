package handlers

import (
	"net/http"
	"strconv"

	"exercises_go/1_exercise_gin_go/model"
	"exercises_go/2_exercise_lumberjack_go/logging"
	"github.com/gin-gonic/gin"
)

var cars = []model.Car{
	{"1", "toyota", "12000"},
	{"2", "suzuki", "9700"},
	{"3", "toyota", "18000"},
}

// GetCars show all cars
func GetCars(ctx *gin.Context) {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	logging.LogINFO("User requested: " + "'" + ctx.FullPath() + "'" + " Status code is " + strconv.Itoa(http.StatusOK))
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	ctx.IndentedJSON(http.StatusOK, cars)
}

// PostCar adds new car in Cars
func PostCar(ctx *gin.Context) {
	var newCar model.Car

	if err := ctx.BindJSON(&newCar); err != nil {
		logging.LogERROR("User requested add new data, Error during creation: " + "'" + strconv.ErrRange.Error() + "'")
		return
	}
	cars = append(cars, newCar)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	logging.LogINFO("User requested: " + "'" + ctx.FullPath() + "'" + " Status code is " + strconv.Itoa(http.StatusOK))
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	ctx.IndentedJSON(http.StatusCreated, newCar)
}

// GetCarById shows car by id
func GetCarById(ctx *gin.Context) {
	Id := ctx.Param("id")

	for _, car := range cars {
		if car.Id == Id {
			////////////////////////////////////////////////////////////////////////////////////////////////////////////
			logging.LogINFO("User requested: " + "'" + ctx.FullPath() + "'" + " Status code is " + strconv.Itoa(http.StatusOK))
			////////////////////////////////////////////////////////////////////////////////////////////////////////////

			ctx.IndentedJSON(http.StatusOK, car)
			return
		}
	}
	logging.LogINFO("User requested: " + "'" + ctx.FullPath() + "'" + " Status code is " + strconv.Itoa(http.StatusNotFound))
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}
