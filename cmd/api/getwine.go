package main

import (
	"errors"
	"github.com/oli4maes/winedata/internal/data"
	"net/http"
)

func (app *application) getWineHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	wine, err := app.models.Wines.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"wine": wine}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
