package main //main.go

import (
	"api/config"
	"api/models"
	"api/routes"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

// @title API de Planos de Faturamento
// @version 1.0
// @description Uma api para gerenciamento de planos de faturamento.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

func main() {

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.LogMode(true)
	config.DB.AutoMigrate(&models.Product{}, &models.CyclePeriod{}, &models.CycleParam{})
	// config.DB.Model(&models.Product{}).AddIndex("idx", "product_id")
	// config.DB.Model(&models.CyclePeriod{}).AddIndex("idx", "cycle_period_id")
	// config.DB.Model(&models.CycleParam{}).AddIndex("idx", "cycle_param_id")
	// config.DB.Model(&models.Product{}).AddForeignKey("product_id", "cycle_periods(cycle_period_id)", "CASCADE", "CASCADE")
	// config.DB.Model(&models.CyclePeriod{}).AddForeignKey("cycle_period_id", "cycle_params(cycle_param_id)", "CASCADE", "CASCADE")
	// config.DB.Model(&models.CycleParam{}).AddForeignKey("cycle_param_id", "cycle_periods(cycle_period_id)", "CASCADE", "CASCADE")

	r := routes.SetupRouter()
	//running
	r.Run()
}
