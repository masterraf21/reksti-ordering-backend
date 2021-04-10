package apis

import (
	"net/http"
	"context"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/masterraf21/reksti-ordering-backend/models"
	httpUtils "github.com/masterraf21/reksti-ordering-backend/utils/http"
)

type menuAPI struct {
	menuUsecase models.MenuUsecase
}

// NewOrderAPI will initiate menu API
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := t.menuUsecase.CreateMenu(context.Background(), &m)
	m.MenuID = id
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (t *menuAPI) deleteMenu(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
	res, err := t.menuUsecase.DeleteMenu(context.Background(), uint32(id))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
    w.Header().Add("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(res); err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
    }
}

func (t *menuAPI) getByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    res, err := t.menuUsecase.GetByID(uint32(id))
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.Header().Add("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(res); err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
    }
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
	m.MenuTypeID = id
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (t *menuAPI) deleteType(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    m, err := t.menuUsecase.DeleteType(context.Background(), uint32(id))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
    w.Header().Add("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(m); err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
    }
}
