package main

import (
	"fmt"
	"github.com/oli4maes/winedata/internal/data"
	"github.com/oli4maes/winedata/internal/validator"
	"net/http"
)

func (app *application) createWineHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	wine := &data.Wine{
		Name: input.Name,
	}

	v := validator.New()

	if data.ValidateWine(v, wine); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Wines.Insert(wine)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/wines/%d", wine.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"wine": wine}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
