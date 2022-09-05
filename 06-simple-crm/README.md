
![](06_Simple_CRM_SS.webm)


# Simple CRM

This project is a Simple CRM that uses SQLite database and Gofiber package that picks up the request from the user and operates them on the database.


## Little More
- Uses the given pacakges :- "fmt", "github.com/gofiber/fiber", "github.com/jinzhu/gorm"
- database.go holds the database instance properties, lead.go holds the leads and function of the routes.
- While creating new leads, the sent from the user gets verified by BodyParser, then a new lead gets created.
- SQLite database is handled by fiber, (find, create and delete) are operated
- GORM, the ORM library is responsible for all the data modification


