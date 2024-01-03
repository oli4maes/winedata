package main

import (
	"errors"
	"github.com/oli4maes/winedata/internal/data"
	"github.com/oli4maes/winedata/internal/validator"
	"net/http"
)

func (app *application) updateWineHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
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

	var input struct {
		Name *string `json:"name"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		wine.Name = *input.Name
	}

	v := validator.New()

	if data.ValidateWine(v, wine); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Wines.Update(wine)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
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
