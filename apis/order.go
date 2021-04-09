package apis

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masterraf21/reksti-ordering-backend/models"
	httpUtils "github.com/masterraf21/reksti-ordering-backend/utils/http"
)

type orderAPI struct {
	orderUsecase models.OrderUsecase
}

// NewOrderAPI will initiate order API
func NewOrderAPI(r *mux.Router, ouc models.OrderUsecase) {
	orderAPI := &orderAPI{
		orderUsecase: ouc,
	}

	r.HandleFunc("/order", orderAPI.getAllOrders).Methods("GET")
	r.HandleFunc("/order", orderAPI.createOrder).Methods("POST")
}

func (t *orderAPI) getAllOrders(w http.ResponseWriter, r *http.Request) {
	result, err := t.orderUsecase.GetAllOrders()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get order data",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.Order `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (t *orderAPI) createOrder(w http.ResponseWriter, r *http.Request) {
}
