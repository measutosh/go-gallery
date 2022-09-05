package lead

import (
	"gitlab.com/depot2/go-projects/06_Simple_CRM/database"
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	
)

type Lead struct {
	gorm.Model
	Name    string  `json:"name"`
	Company string  `json:"company"`
	Email   string  `json:"email"`
	Phone   int     `json:"phone"`
}

// when api hits this function, it takes Context as input and gets ready to work with the data that comes from the user
// for Getleads, is will be needed which can be accessed usng "c.Params()"
func GetLeads(c *fiber.Ctx) {
	db := database.DBconn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBconn
	var lead Lead
	// it will look for a particular id
	db.Find(&lead, id)
	// respond a JSON using c
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBconn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	// to delete a lead, id is needed as particular lead will be deleted
	id := c.Params("id")
	db := database.DBconn
	// lead is the name of the variable, Lead is the struct type
	var lead Lead
	db.First(&lead, id)
	if lead.Name == ""{
		c.Status(500).Send("No lead found with this ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted")
}