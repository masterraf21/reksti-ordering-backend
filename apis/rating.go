package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masterraf21/reksti-ordering-backend/models"
	httpUtils "github.com/masterraf21/reksti-ordering-backend/utils/http"
)

type ratingAPI struct {
	ratingRepo models.RatingRepository
}

// NewRatingAPI will initiate rating API
func NewRatingAPI(r *mux.Router, rr models.RatingRepository) {
	ratingAPI := &ratingAPI{
		ratingRepo: rr,
	}

	r.HandleFunc("/rating", ratingAPI.getAll).Methods("GET")
	r.HandleFunc("/rating", ratingAPI.createRating).Methods("POST")
	r.HandleFunc("/rating/menu/{id}", ratingAPI.getByMenu).Methods("GET")
	r.HandleFunc("/rating/menu/score/{id}", ratingAPI.getMenuScore).Methods("GET")
	r.HandleFunc("/rating/{id}", ratingAPI.deleteRating).Methods("DELETE")
	r.HandleFunc("/rating/{id}", ratingAPI.getByID).Methods("GET")
}

func (t *ratingAPI) getAll(w http.ResponseWriter, r *http.Request) {
	result, err := t.ratingRepo.GetAll()
	if err != nil {
		httpUtils.HandleError(
			w,
			r,
			err,
			"failed to get rating data",
			http.StatusInternalServerError,
		)
		return
	}

	var data struct {
		Data []models.Rating `json:"data"`
	}

	data.Data = result
	httpUtils.HandleJSONResponse(w, r, data)
}

func (t *ratingAPI) createRating(w http.ResponseWriter, r *http.Request) {
	m := models.Rating{}
	json.NewDecoder(r.Body).Decode(&m)
	if m.Remarks == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := t.ratingRepo.Store(&m)
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to create rating", http.StatusInternalServerError)
		return
	}
	var response struct {
		ID uint32 `json:"rating_id"`
	}
	response.ID = id
	httpUtils.HandleJSONResponse(w, r, response)
}

func (t *ratingAPI) deleteRating(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	_, err = t.ratingRepo.GetByID(uint32(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = t.ratingRepo.DeleteByID(uint32(id))
	if err != nil {
		httpUtils.HandleError(w, r, err, "failed to delete rating", http.StatusInternalServerError)
		return
	}
	httpUtils.HandleNoJSONResponse(w)
}

func (t *ratingAPI) getByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	m, err := t.ratingRepo.GetByID(uint32(id))
	if err != nil {
		httpUtils.HandleError(w, r, err, "failted to get rating by id", http.StatusInternalServerError)
		return
	}

	var data struct {
		Data *models.Rating `json:"data"`
	}
	data.Data = m
	httpUtils.HandleJSONResponse(w, r, data)
}

func (t *ratingAPI) getByMenu(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	m, err := t.ratingRepo.GetByMenu(uint32(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var data struct {
		Data []models.Rating `json:"data"`
	}
	data.Data = m
	httpUtils.HandleJSONResponse(w, r, data)
}

func (t *ratingAPI) getMenuScore(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	m, err := t.ratingRepo.GetMenuScore(uint32(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	type Score struct {
		Score float32 `json:"score"`
	}
	var data struct {
		Data Score `json:"data"`
	}
	data.Data = Score{Score: m}
	httpUtils.HandleJSONResponse(w, r, data)
}
