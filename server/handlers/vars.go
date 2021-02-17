package handlers

import "text/template"

//ModelBasic store the basic html page template
var ModelBasic = template.Must(template.ParseFiles("../html/basic.html"))

//ModelLocal store the local html page template
var ModelLocal = template.Must(template.ParseFiles("../html/local.html"))
