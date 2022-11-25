package repository

import (
	"database/sql"

	"exercises_go/3_exercise_go_viperNlog/logging"

	"exercises_go/3_exercise_go_viperNlog/config"
	"exercises_go/3_exercise_go_viperNlog/model"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DRIVERNAME = config.GetConfig("database.driverName")
	USER       = config.GetConfig("database.user")
	PASS       = config.GetConfig("database.password")
	PROTOCOL   = config.GetConfig("database.protocol")
	HOST       = config.GetConfig("database.host")
	PORT       = config.GetConfig("database.port")
	DBNAME     = config.GetConfig("database.dbName")
)

type Repository struct {
	db *sql.DB
}

func NewRepository() *Repository {
	db, err := sql.Open(DRIVERNAME, USER+":"+PASS+"@"+PROTOCOL+"("+HOST+":"+PORT+")/"+DBNAME)
	if err != nil {
		logging.LogERROR("Can not connect to db" + DBNAME + "! " + err.Error())
	}
	logging.LogINFO("Connected to db " + DBNAME + " successfully!")
	return &Repository{
		db: db,
	}
}

// Create add car to DB
func (r *Repository) Create(Make string, Price string) {
	query, err := r.db.Query("INSERT INTO `carsdb`.`cars` (`make`, `price`) VALUES (?, ?)", Make, Price)
	if err != nil {
		panic(err.Error())
	}
	logging.LogINFO("User add: " + Make + ", " + Price)
	query.Close()
}

// Get gets Cars from DB
func (r *Repository) Get() []model.Car {
	var cars []model.Car

	rows, err := r.db.Query("select * from carsdb.cars")
	if err != nil {
		logging.LogERROR("Error during get data from " + DBNAME)
	}
	for rows.Next() {
		var car model.Car
		rows.Scan(&car.Id, &car.Make, &car.Price)
		cars = append(cars, car)
	}

	logging.LogINFO("User sent request to DB - " + DBNAME + " and got results successfully")

	return cars
}

// GetById get car By Id from DB
func (r *Repository) GetById(Id string) model.Car {
	rows, err := r.db.Query("select make from carsdb.cars where id = ?", Id)
	if err != nil {
		logging.LogERROR("Error during get data from " + DBNAME)
	}
	defer rows.Close()

	var car model.Car
	for rows.Next() {
		rows.Scan(&car.Make)
	}

	logging.LogINFO("User request get car by Id - " + Id + " and got results successfully")
	return car
}

// UpdateMake update
func (r *Repository) UpdateMake(Id string, Make string) {
	query, err := r.db.Query("UPDATE `carsdb`.`cars` SET `make` = ? WHERE (`id` = ?)", Make, Id)
	if err != nil {
		logging.LogERROR("Error during get data from " + DBNAME)
	}
	query.Close()

	logging.LogINFO("User request update data by ID - " + Id + " and got results successfully")
}

// Delete delete car from DB
func (r *Repository) Delete(Id string) {
	query, err := r.db.Query("DELETE FROM `carsdb`.`cars` WHERE (`id` = ?)", Id)
	if err != nil {
		logging.LogERROR("Error during get data from " + DBNAME)
	}
	query.Close()

	logging.LogINFO("User requested delete data under ID - " + Id + " and got results successfully")
}
