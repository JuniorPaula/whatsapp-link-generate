package main

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

// views is the jet template set
var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./views"),
	jet.InDevelopmentMode(),
)

func RenderPage(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) error {
	var vars jet.VarMap

	if data != nil {
		vars = jet.VarMap{}
		vars.Set("data", data)
	}

	view, err := views.GetTemplate(templateName)
	if err != nil {
		return err
	}

	err = view.Execute(w, vars, nil)
	if err != nil {
		return err
	}

	return nil
}
