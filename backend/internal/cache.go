package internal

import (
	"context"

	rs "github.com/brxyxn/go_gpr_nclouds/backend/internal/database/cache/handlers"
	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
	"github.com/go-redis/redis/v8"
)

func (a *App) cacheRoutes() {
	// Cache Routes
	cacheHandler := rs.NewHandlers(a.Cache)
	a.Router.HandleFunc("/api/v1/cache/users", cacheHandler.CreateUser).Methods("POST")
	a.Router.HandleFunc("/api/v1/cache/users", cacheHandler.GetCounter).Methods("GET")
}

func (a *App) initializeCache(bindAddr, password string, dbname int) {
	u.Log.Info("Initializing redis cache...")
	a.Ctx = context.Background()
	a.Cache = redis.NewClient(&redis.Options{
		Addr:     bindAddr,
		Password: password,
		DB:       dbname,
	})

	ping := a.Cache.Ping(a.Ctx)
	if ping.Val() != "PONG" {
		u.Log.Error("Error opening a new connection to the Redis Cache.")
		return
	}
	u.Log.Info("Connected to redis cache db", dbname, "at", bindAddr)
}
