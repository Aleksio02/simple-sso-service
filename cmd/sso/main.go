package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"simple-sso-service/modules/sso/routes"
)

func main() {
	app := gin.Default()

	routes.RegisterSsoRoutes(app)

	app.Static("/assets", "./modules/sso/assets")

	//tmpl := template.Must(template.ParseGlob("modules/sso/templates/*.html"))
	tmpl := template.Must(template.ParseFiles(
		"modules/sso/templates/base.html",
		"modules/sso/templates/register.html",
		"modules/sso/templates/login.html",
	))
	app.SetHTMLTemplate(tmpl)

	//app.LoadHTMLGlob("modules/sso/templates/*.html")
	//app.LoadHTMLGlob("modules/sso/templates/layout/*")
	app.Run(":8080")
}
