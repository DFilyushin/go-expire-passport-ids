package controllers

import (
	"encoding/json"
	"expired-passport-checker/internal/api/responses"
	"expired-passport-checker/internal/service"
	"net/http"
)

type ApiController struct{}

// CheckPassport godoc
// @Summary Check passport
// @Description Check passport by series and number
// @Tags api
// @Accept  json
// @Produce  json
// @Param series query string true "Passport series"
// @Param number query string true "Passport number"
// @Success 200 {object} PassportCheckerResponse
// @Router /checkPassport [get]
func (c ApiController) CheckPassport(service *service.PassportIdService) http.HandlerFunc {
	/*Проверка паспорта*/
	return func(writer http.ResponseWriter, request *http.Request) {
		params := request.URL.Query()
		var series = params.Get("series")
		var number = params.Get("number")
		if (series == "") || (number == "") {
			http.Error(writer, "Number and series important parameters!", http.StatusBadRequest)
			return
		}

		isPassportCorrect := service.CheckPassport(series, number)
		response := responses.PassportCheckerResponse{
			PassportSeries: series,
			PassportNumber: number,
			Result:         isPassportCorrect,
		}

		json.NewEncoder(writer).Encode(response)
	}
}
