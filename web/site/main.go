package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Environment struct {
	Log    *log.Logger

}

var env Environment

var (
	SiteHTML =
		`
    <!DOCTYPE html>
    <html>
      <head>
        <meta http-equiv="content-type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">

        <title>{{.Title}}</title>

        <style>@import url('{{.GoogleFontsURL}}');</style>
        <style>{{.Style}}</style>
      </head>
      <body>
        <section id="top-container">
<header></header>
          <div class="top-content"></div>
        </section>
        <section id="main-container">
          <div class="main-content">
            <h1>{{.Content}}</h1>
          </div>
        </section>
        <section id="bottom-container">
          <div class="bottom-content"></div>
</section> 
      <script>
      </script>
      </body>
    </html>
    `
    SiteCSS string =
    	`
html, body {
  margin: 0;
  padding: 0;
  background-color: #f5f5f5;
  font-family: 'Helvetica Neue', Arial, Helvetica, sans-serif;
}

header {
  margin: 0 auto;
  height: 5vh;
  background-color: #2d2d29;
}

h1 {
  text-align: center;
}

.main-content {
  margin: 0 auto;
  height: 90vh;
  min-height: 360px;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  flex-wrap: wrap;
  background-color: #f5f5f5;
}

.bottom-content {
  margin: 0 auto;
  height: 20vh;
  background-color: #2d2d29;
}

nav {
  display: flex;
  flex-direction: row;
  justify-content: center;
  position: absolute;
  margin: 0 auto;
  width: 100vw;
  bottom: 5vh;
}

.contact {
  display: flex;
  flex-direction: row;
  justify-content: center;
  position: absolute;
  margin: 0 auto;
  width: 100vw;
  bottom: -15vh;
}

nav ul,
.contact ul {
  list-style: none;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  padding: 0;
}

nav ul li,
.contact ul li {
  display: inline;
  padding: 0px 10vw;
}

nav ul li a {
  text-decoration: none;
  color: #ffcfcf;
  font-family: 'Open Sans', Helvetica, sans-serif;
  transition: color .35s ease;
  font-size: 3.5vw;
}

nav ul li a:hover {
  color: #2d2d29;
  transition: color .35s ease;
}
.contact ul li {
  color: #cfcfff;
  font-family: 'Open Sans', Helvetica, sans-serif;
  transition: color .35s ease;
  font-size: 3.5vw;
}
`
)

func init() {
	// Create environment
	env = Environment{
		nil,
	}
	// Set up logger
	env.Log = log.New(os.Stdout, "CUESITE *** ", 0)
	env.Log.Println("Booting up Cue website")

	// Setup shared templates

}

func cueSiteHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("site")
	te, _ := t.Parse(SiteHTML)
	type site struct {
		Title          string
		GoogleFontsURL string
		Content        string
		Style          string
	}
	cueSite := &site{
		"Cue Labs",
		"https://fonts.googleapis.com/css?family=Open+Sans",
		"Cue Labs",
		SiteCSS,
	}
	if err := te.Execute(w, cueSite); err != nil {
		env.Log.Println("Error while templating in main()")
		env.Log.Println(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", cueSiteHandler)
	http.ListenAndServe(":8000", r)
}
