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
	r.HandleFunc("/order/{id_order}", orderAPI.getOrderByID).Methods("GET")
	r.HandleFunc("/order/{id_order}/details", orderAPI.getOrderWithDetails).Methods("GET")
	r.HandleFunc("/order/{id_order}/details", orderAPI.createOrderDetail).Methods("POST")
	r.HandleFunc("/order/detail/{id_order_detail}", orderAPI.getOrderDetailByID).Methods("GET")

	r.HandleFunc("/order/customer/{id_customer}", orderAPI.getOrderByCustomer).Methods("GET")
	r.HandleFunc("/order/customer/{id_customer}/history", orderAPI.getOrderHistoryByCustomer).Methods("GET")
	r.HandleFunc("/order/customer/{id_customer}/ongoing", orderAPI.getOrderOngoingByCustomer).Methods("GET")

	r.HandleFunc("/order/{id_order}/price", orderAPI.updateOrderPrice).Methods("PUT")
	r.HandleFunc("/order/detail/{id_order_detail}/price", orderAPI.updateOrderDetailPrice).Methods("PUT")
	r.HandleFunc("/order/{id_order}/status", orderAPI.updateOrderStatus).Methods("PUT")

	r.HandleFunc("/order/{id_order}", orderAPI.deleteOrderByID).Methods("DELETE")
	r.HandleFunc("/order/detail/{id_order_detail}", orderAPI.deleteOrderDetailByID).Methods("DELETE")
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
	var body struct {
		CustomerID  uint32  `json:"customer_id"`
		TotalPrice  float32 `json:"total_price"`
		OrderStatus int32   `json:"order_status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	order := models.Order{
		CustomerID:  body.CustomerID,
		TotalPrice:  body.TotalPrice,
		OrderStatus: body.OrderStatus,
	}

	id, err := t.orderUsecase.CreateOrder(context.TODO(), &order)
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to create order", http.StatusInternalServerError)
		return
	}

	var response struct {
		OrderID uint32 `json:"order_id"`
	}
	response.OrderID = id

	httpUtils.HandleJSONResponse(w, r, response)
}

func (t *orderAPI) getOrderByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderID, err := strconv.ParseInt(params["id_order"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_order"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := t.orderUsecase.GetOrderByID(uint32(orderID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get user order by id",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data *models.Order `json:"data"`
	}

	data.Data = result

	httpUtils.HandleJSONResponse(w, r, data)
}

func (t *orderAPI) getOrderDetailByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderDetailID, err := strconv.ParseInt(params["id_order_detail"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_order_detail"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := t.orderUsecase.GetOrderDetailByID(uint32(orderDetailID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get user order by id",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data *models.OrderDetails `json:"data"`
	}

	data.Data = result

	httpUtils.HandleJSONResponse(w, r, data)
}

func (t *orderAPI) getOrderWithDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderID, err := strconv.ParseInt(params["id_order"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_order"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	result, err := t.orderUsecase.GetOrderWithDetails(uint32(orderID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get user order by id",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data models.OrderWithDetails `json:"data"`
	}

	data.Data = result

	httpUtils.HandleJSONResponse(w, r, data)
}

func (t *orderAPI) createOrderDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderID, err := strconv.ParseInt(params["id_order"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_order"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	var body struct {
		MenuID     uint32  `json:"menu_id"`
		Quantity   uint32  `json:"quantity"`
		TotalPrice float32 `json:"total_price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	ordDetail := models.OrderDetails{
		OrderID:    uint32(orderID),
		MenuID:     body.MenuID,
		Quantity:   body.Quantity,
		TotalPrice: body.TotalPrice,
	}

	err = t.orderUsecase.CreateOrderDetail(
		context.TODO(),
		&ordDetail,
	)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to create order detail",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}

func (t *orderAPI) updateOrderPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderID, err := strconv.ParseInt(params["id_order"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_order"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	err = t.orderUsecase.UpdateOrderPrice(context.TODO(), uint32(orderID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to update order total price",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}

func (t *orderAPI) updateOrderStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderID, err := strconv.ParseInt(params["id_order"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_order"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	var body struct {
		OrderStatus int32 `json:"order_status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpUtils.HandleError(w, r, err, "bad request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = t.orderUsecase.UpdateOrderStatus(context.TODO(), uint32(orderID), body.OrderStatus)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to update order status",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}

func (t *orderAPI) updateOrderDetailPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderDetailID, err := strconv.ParseInt(params["id_order_detail"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_order_detail"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	err = t.orderUsecase.UpdateOrderDetailPrice(context.TODO(), uint32(orderDetailID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to update order detail total price",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}

func (t *orderAPI) getOrderByCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	custID, err := strconv.ParseInt(params["id_customer"], 10, 64)
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

	result, err := t.orderUsecase.GetOrdersByCustID(uint32(custID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get user order by customer id",
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

func (t *orderAPI) getOrderHistoryByCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	custID, err := strconv.ParseInt(params["id_customer"], 10, 64)
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

	result, err := t.orderUsecase.GetOrdersHistoryByCustID(uint32(custID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get user order by customer id",
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

func (t *orderAPI) getOrderOngoingByCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	custID, err := strconv.ParseInt(params["id_customer"], 10, 64)
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

	result, err := t.orderUsecase.GetOngoingOrdersyByCustID(uint32(custID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get user order by customer id",
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

func (t *orderAPI) deleteOrderByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderID, err := strconv.ParseInt(params["id_order"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_order"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	err = t.orderUsecase.DeleteOrder(context.TODO(), uint32(orderID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to delete order by id",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}

func (t *orderAPI) deleteOrderDetailByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	orderDetailID, err := strconv.ParseInt(params["id_order_detail"], 10, 64)
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			params["id_order_detail"]+" is not integer",
			http.StatusBadRequest,
		)
		return
	}

	err = t.orderUsecase.DeleteOrderDetail(context.TODO(), uint32(orderDetailID))
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to delete order detail by id",
			http.StatusInternalServerError,
		)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}
