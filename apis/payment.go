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
	paymentUsecase     models.PaymentUsecase
	paymentTypeUsecase models.PaymentTypeUsecase
}

// NewPaymentAPI will initiate api
func NewPaymentAPI(r *mux.Router, pac models.PaymentUsecase, ptc models.PaymentTypeUsecase) {
	paymentAPI := &paymentAPI{
		paymentUsecase:     pac,
		paymentTypeUsecase: ptc,
	}

	r.HandleFunc("/payment", paymentAPI.GetAll).Methods("GET")
	r.HandleFunc("/payment/type", paymentAPI.GetAllType).Methods("GET")
	r.HandleFunc("/payment", paymentAPI.Create).Methods("POST")
	r.HandleFunc("/payment/type", paymentAPI.CreateType).Methods("POST")
	r.HandleFunc("/payment/{id_payment}", paymentAPI.GetByID).Methods("GET")
	r.HandleFunc("/payment/type/{id_payment_type}", paymentAPI.GetTypeByID).Methods("GET")
	r.HandleFunc("/payment/{id_payment}", paymentAPI.DeleteByID).Methods("DELETE")
	r.HandleFunc("/payment/type/{id_payment_type}", paymentAPI.DeleteTypeByID).Methods("DELETE")
	r.HandleFunc("/payment/{id_payment}/status", paymentAPI.UpdateStatus).Methods("PUT")
	r.HandleFunc("/payment/list/{id_customer}", paymentAPI.GetListOfPaymentsByCustomerID).Methods("GET")
}

func (p *paymentAPI) UpdateStatus(w http.ResponseWriter, r *http.Request) {
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

	var body struct {
		Status int32 `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = p.paymentUsecase.UpdateStatus(context.TODO(), uint32(paymentID), body.Status)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to update payment status",
			http.StatusBadRequest,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}

func (p *paymentAPI) DeleteTypeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paymentTypeID, err := strconv.ParseInt(params["id_payment_type"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_payment_type"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	err = p.paymentTypeUsecase.DeleteByID(context.TODO(), uint32(paymentTypeID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to delete payment type by id",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}

func (p *paymentAPI) GetTypeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paymentTypeID, err := strconv.ParseInt(params["id_payment_type"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_payment_type"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := p.paymentTypeUsecase.GetByID(uint32(paymentTypeID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get payment type data by id",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data *models.PaymentType `json:"data"`
	}

	data.Data = result

	httpUtils.HandleJSONResponse(w, r, data)
}

func (p *paymentAPI) CreateType(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Method  string `json:"method"`
		Company string `json:"company"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	paymentType := models.PaymentType{
		Method:  body.Method,
		Company: body.Company,
	}

	id, err := p.paymentTypeUsecase.CreatePaymentType(context.TODO(), &paymentType)
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to create payment type", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"payment_type_id"`
	}
	response.ID = id

	httpUtils.HandleJSONResponse(w, r, response)
}

func (p *paymentAPI) GetAllType(w http.ResponseWriter, r *http.Request) {
	result, err := p.paymentTypeUsecase.GetAll()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get payment type data",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.PaymentType `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
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
		OrderID       uint32  `json:"order_id"`
		Amount        float32 `json:"amount"`
		PaymentTypeID uint32  `json:"payment_type_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	payment := models.Payment{
		OrderID:       body.OrderID,
		Amount:        body.Amount,
		PaymentTypeID: body.PaymentTypeID,
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

func (p *paymentAPI) GetListOfPaymentsByCustomerID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	customerID, err := strconv.ParseInt(params["id_customer"], 10, 64)

	result, err := p.paymentUsecase.GetListOfPaymentsByCustomerID(context.TODO(), uint32(customerID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get list of payment by customer id",
			http.StatusInternalServerError,
		)
		return
	}

	var dataa struct {
		Data []models.Payment `json:"data"`
	}

	dataa.Data = result
	httpUtils.HandleJSONResponse(w, r, dataa)
}