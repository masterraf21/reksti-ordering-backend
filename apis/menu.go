package apis

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/reksti-ordering-backend/models"
	httpUtils "github.com/masterraf21/reksti-ordering-backend/utils/http"
)

type menuAPI struct {
	menuUsecase models.MenuUsecase
}

// NewMenuAPI will initiate menu API
func NewMenuAPI(r *mux.Router, mu models.MenuUsecase) {
	menuAPI := &menuAPI{
		menuUsecase: mu,
	}

	r.HandleFunc("/menu", menuAPI.getAll).Methods("GET")
	r.HandleFunc("/menu", menuAPI.createMenu).Methods("POST")
	r.HandleFunc("/menu/type", menuAPI.getAllType).Methods("GET")
	r.HandleFunc("/menu/type", menuAPI.createType).Methods("POST")
	r.HandleFunc("/menu/type/{id}", menuAPI.deleteType).Methods("DELETE")
	r.HandleFunc("/menu/{id}", menuAPI.deleteMenu).Methods("DELETE")
	r.HandleFunc("/menu/{id}", menuAPI.getByID).Methods("GET")
}

func (t *menuAPI) getAll(w http.ResponseWriter, r *http.Request) {
	result, err := t.menuUsecase.GetAll()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get menu data",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.MenuComp `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (t *menuAPI) createMenu(w http.ResponseWriter, r *http.Request) {
	m := models.Menu{}
	json.NewDecoder(r.Body).Decode(&m)
	if m.Name == "" {
		httpUtils.HandleError(w, r, errors.New("No Name provided"), "please provide name", http.StatusBadRequest)
		return
	}
	id, err := t.menuUsecase.CreateMenu(context.Background(), &m)
	if err != nil {
		httpUtils.HandleError(w, r, err, "cannot create menu", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"menu_id"`
	}
	response.ID = id

	httpUtils.HandleJSONResponse(w, r, response)
}

func (t *menuAPI) deleteMenu(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	_, err = t.menuUsecase.DeleteMenu(context.Background(), uint32(id))
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to delete menu", http.StatusInternalServerError)
		return
	}

	httpUtils.HandleNoJSONResponse(w)
}

func (t *menuAPI) getByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		httpUtils.HandleError(w, r, err, "id not integer", http.StatusBadRequest)
	}
	res, err := t.menuUsecase.GetByID(uint32(id))
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to get menu by id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data *models.MenuComp `json:"data"`
	}
	data.Data = res
	httpUtils.HandleJSONResponse(w, r, data)
}

func (t *menuAPI) getAllType(w http.ResponseWriter, r *http.Request) {
	result, err := t.menuUsecase.GetAllType()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get menu data",
			http.StatusInternalServerError,
		)
		return
	}

	var dataa struct {
		Data []models.MenuType `json:"data"`
	}

	dataa.Data = result
	httpUtils.HandleJSONResponse(w, r, dataa)
}

func (t *menuAPI) createType(w http.ResponseWriter, r *http.Request) {
	m := models.MenuType{}
	json.NewDecoder(r.Body).Decode(&m)
	if m.TypeName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := t.menuUsecase.CreateType(context.Background(), &m)
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to create menu tyoe", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID uint32 `json:"menu_type_id"`
	}

	response.ID = id

	httpUtils.HandleJSONResponse(w, r, response)
}

func (t *menuAPI) deleteType(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	_, err = t.menuUsecase.DeleteType(context.Background(), uint32(id))
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to delete menu type", http.StatusInternalServerError)
		return
	}
	httpUtils.HandleNoJSONResponse(w)
}
