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
	
	config.DB.AutoMigrate( &models.Products{}, &models.CyclePeriod{}, &models.CycleParam{})
	
	r := routes.SetupRouter()
	//running
	r.Run()
}
