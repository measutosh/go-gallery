package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	// importing the database, lead packages
	"gitlab.com/depot2/go-projects/06_Simple_CRM/database"
	"gitlab.com/depot2/go-projects/06_Simple_CRM/lead"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

)

// setting up the routes here
func setUpRoutes(app *fiber.App) {
	// in the lead package the definition of these function exist, that's why "lead.GetLead" is being used here
	// when someone will hit api/v1/lead endpoint in postman then thses function will be running in certain cases
	app.Get("api/v1/lead", lead.GetLeads)
	app.Get("api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id",lead.DeleteLead)
}

func initDatabase() {
	// defining the database properties
	var err error
	// using gorm to open a connection to the database(leads.db)
	// the connection will be stored inside the database.DBconn
	database.DBconn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to the database")
	}
	fmt.Println("\n======Connection opened to the database\n")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("===::::::===Database migrated===::::::===\n")
}

func main() {
	// creating the instance of fiber
	app := fiber.New()
	// starting the database
	initDatabase()
	setUpRoutes(app)
	// starting the server
	app.Listen(3000)
	// closing the database connection at the end
	defer database.DBconn.Close()

}