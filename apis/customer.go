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

type customerAPI struct {
	customerUsecase models.CustomerUsecase
}

// NewCustomerAPI will initiate new api
func NewCustomerAPI(r *mux.Router, cuc models.CustomerUsecase) {
	customerAPI := &customerAPI{
		customerUsecase: cuc,
	}

	r.HandleFunc("/customer", customerAPI.GetAll).Methods("GET")
	r.HandleFunc("/customer", customerAPI.Create).Methods("POST")
	r.HandleFunc("/customer/{id_customer}", customerAPI.GetByID).Methods("GET")
	r.HandleFunc("/customer/{id_customer}", customerAPI.DeleteByID).Methods("DELETE")
}

func (c *customerAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := c.customerUsecase.GetAll()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get customer data",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.Customer `json:"data"`
	}
	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (c *customerAPI) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	customerID, err := strconv.ParseInt(params["id_customer"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_customer"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := c.customerUsecase.GetByID(uint32(customerID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get user customer by id",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data *models.Customer `json:"data"`
	}

	data.Data = result

	httpUtils.HandleJSONResponse(w, r, data)
}

func (c *customerAPI) DeleteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	customerID, err := strconv.ParseInt(params["id_customer"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_customer"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	err = c.customerUsecase.DeleteByID(context.TODO(), uint32(customerID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to delete customer by id",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}

func (c *customerAPI) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		FullName    string `json:"full_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_numer"`
		Username    string `json:"username"`
		Password    string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	cust := models.Customer{
		FullName:    body.FullName,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
		Username:    body.Username,
		Password:    body.Password,
	}

	id, err := c.customerUsecase.CreateCustomer(context.TODO(), &cust)
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to create customer", http.StatusInternalServerError)
		return
	}

	var response struct {
		CustID uint32 `json:"customer_id"`
	}
	response.CustID = id

	httpUtils.HandleJSONResponse(w, r, response)
}
