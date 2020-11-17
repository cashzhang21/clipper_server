module "clipper_server"

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gin-contrib/static v0.0.0-20200916080430-d45d9a37d28e
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.11
)