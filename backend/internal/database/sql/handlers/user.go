package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/brxyxn/go_gpr_nclouds/backend/internal/database/sql/data"
	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
)

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	u.Log.Info("Handling POST Users /sql/users")

	var values *data.User
	var counter data.Counter

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&values); err != nil {
		u.Log.Error("CreateUser handler:", err)
		u.Respond.Error(w, http.StatusBadRequest, u.Consts.InvalidPayload)
		return
	}

	defer r.Body.Close()

	if values.Username == "" || values.Email == "" || values.Password == "" {
		u.Log.Error("CreateUser handler:", u.Consts.RequiredParams)
		u.Respond.Error(w, http.StatusBadRequest, u.Consts.RequiredParams)
		return
	}

	if err := data.CreateUser(h.db, values, &counter); err != nil {
		switch err {
		case sql.ErrNoRows:
			u.Respond.Error(w, http.StatusNotFound, u.Consts.SqlNotFound)
		default:
			u.Log.Error("CreateUser handler:", err)
			u.Respond.Error(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	u.Respond.JSON(w, http.StatusCreated, counter)
}

func (h *Handlers) GetCounter(w http.ResponseWriter, r *http.Request) {
	u.Log.Info("Handling GET Users /sql/users")

	var counter data.Counter
	if err := data.CountUsers(h.db, &counter); err != nil {
		switch err {
		case sql.ErrNoRows:
			u.Respond.Error(w, http.StatusNotFound, u.Consts.SqlNotFound)
		default:
			u.Respond.Error(w, http.StatusInternalServerError, err.Error())
		}
		u.Log.Error("GetCounter handler:", err)
		return
	}

	u.Respond.JSON(w, http.StatusOK, counter)
}
