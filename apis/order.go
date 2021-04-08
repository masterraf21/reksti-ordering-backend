package apis

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masterraf21/reksti-ordering-backend/models"
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
}

func (t *orderAPI) createOrder(w http.ResponseWriter, r *http.Request) {
}
