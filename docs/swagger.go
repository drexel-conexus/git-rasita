package docs

import "github.com/swaggo/swag"

const docTemplate = `{
	"swagger": "2.0",
	"info": {
		"title": "Git Repository Service API"
	}
}`

// @title           Git Repository Service API
// @version         1.0
// @description     A GitHub-like service API for managing projects and repositories
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func SwaggerInfo() *swag.Spec {
	return &swag.Spec{
		InfoInstanceName: "swagger",
		SwaggerTemplate:  docTemplate,
	}
}
