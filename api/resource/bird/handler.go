package bird

import (
	e "bird-box-go/api/resource/common/err"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type API struct {
	repository *Repository
}

func New(db *gorm.DB) *API {
	return &API{
		repository: NewRepository(db),
	}
}

func (a *API) ListAll(w http.ResponseWriter, r *http.Request) {
	books, err := a.repository.ListAll()
	if err != nil {
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	if len(books) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

func (a *API) GetSingle(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	bird, err := a.repository.GetSingle(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	if err := json.NewEncoder(w).Encode(bird); err != nil {
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	rows, err := a.repository.Delete(id)
	if err != nil {
		e.BadRequest(w, e.RespDBDataRemoveFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
