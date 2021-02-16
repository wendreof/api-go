package handlers

import "text/template"

//Models store the html page template
var Models = template.Must(template.ParseFiles("../html/basic.html"))
