package apis

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/reksti-ordering-backend/models"
	httpUtils "github.com/masterraf21/reksti-ordering-backend/utils/http"
)

type paymentAPI struct {
	paymentUsecase models.PaymentUsecase
}

// NewPaymentAPI will initiate api
func NewPaymentAPI(r *mux.Router, pac models.PaymentUsecase) {
	paymentAPI := &paymentAPI{
		paymentUsecase: pac,
	}

	r.HandleFunc("/payment", paymentAPI.GetAll).Methods("GET")
	r.HandleFunc("/payment", paymentAPI.Create).Methods("POST")
	r.HandleFunc("/payment/{id_payment}", paymentAPI.GetAll).Methods("GET")
	r.HandleFunc("/payment/{id_payment}", paymentAPI.DeleteByID).Methods("DELETE")
}

func (p *paymentAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := p.paymentUsecase.GetAll()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get payment data",
			http.StatusInternalServerError,
		)
		return
	}
	var data struct {
		Data []models.Payment `json:"data"`
	}
	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (p *paymentAPI) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paymentID, err := strconv.ParseInt(params["id_payment"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_payment"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := p.paymentUsecase.GetByID(uint32(paymentID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get payment data by id",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data *models.Payment `json:"data"`
	}

	data.Data = result

	httpUtils.HandleJSONResponse(w, r, data)
}

func (p *paymentAPI) DeleteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paymentID, err := strconv.ParseInt(params["id_payment"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_payment"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	err = p.paymentUsecase.DeleteByID(context.TODO(), uint32(paymentID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to delete payment by id",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}

func (p *paymentAPI) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		OrderID     uint32  `json:"order_id"`
		Amount      float32 `json:"amount"`
		PaymentType string  `json:"payment_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	payment := models.Payment{
		OrderID:     body.OrderID,
		Amount:      body.Amount,
		PaymentType: body.PaymentType,
	}

	id, err := p.paymentUsecase.CreatePayment(context.TODO(), &payment)
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to create payment", http.StatusInternalServerError)
		return
	}

	var response struct {
		PaymentID uint32 `json:"payment_id"`
	}
	response.PaymentID = id

	httpUtils.HandleJSONResponse(w, r, response)
}
