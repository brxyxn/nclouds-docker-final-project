package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brxyxn/go_gpr_nclouds/backend/internal/database/cache/data"
	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
)

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	u.Log.Info("Handling POST Users /cache/users")

	w.Header().Set("Content-Type", "application/json")

	var user data.User
	var count data.Counter

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		u.Log.Error("Cache CreateUser handler:", err)
		u.Respond.Error(w, http.StatusBadRequest, u.Consts.InvalidPayload)
		return
	}

	defer r.Body.Close()

	if user.Username == "" || user.Email == "" || user.Password == "" {
		u.Log.Error("Cache CreateUser handler:", u.Consts.RequiredParams)
		u.Respond.Error(w, http.StatusBadRequest, u.Consts.RequiredParams)
	}

	if err := data.CreateUser(h.client, ctx, &user); err != nil {
		u.Log.Error("Cache CreateUser handler:", err)
		u.Respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := data.CountUsers(h.client, ctx, &count); err != nil {
		u.Log.Error("Cache CreateUser handler:", err)
		u.Respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.Respond.JSON(w, http.StatusCreated, count)
}

func (h *Handlers) GetCounter(w http.ResponseWriter, r *http.Request) {
	u.Log.Info("Handling GET Users /cache/users")

	w.Header().Set("Content-Type", "application/json")

	var counter data.Counter
	if err := data.CountUsers(h.client, ctx, &counter); err != nil {
		u.Respond.Error(w, http.StatusInternalServerError, err.Error())
		u.Log.Error("GetCounter handler:", err)
		return
	}

	u.Respond.JSON(w, http.StatusOK, counter)
}
