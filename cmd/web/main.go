package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RajaTobias/GoBookings/pkg/config"
	"github.com/RajaTobias/GoBookings/pkg/handlers"
	"github.com/RajaTobias/GoBookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8082"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InPrduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InPrduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create the cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	app.UseCache = false
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application in port: %s", port))

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
