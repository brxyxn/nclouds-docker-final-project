package data

import (
	"context"

	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
	"github.com/go-redis/redis/v8"
)

type User struct {
	UserID   uint64 `json:"userId"`
	Username string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Counter struct {
	Value int `json:"counter"`
}

func CreateUser(rdb *redis.Client, ctx context.Context, user *User) error {
	var err error
	var key, value string

	key = user.Username + "-" + user.Email
	value = user.Password

	err = rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		u.Log.Error("Error creating new key.", err)
		return err
	}

	_, err = rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		u.Log.Error(key, u.Consts.KeyNotExist)
		return err
	}
	if err != nil {
		u.Log.Error(err)
		return err
	}

	return nil
}

func CountUsers(rdb *redis.Client, ctx context.Context, counter *Counter) error {
	count := rdb.DBSize(ctx)
	err := count.Err()
	if err != nil {
		return err
	}

	counter.Value = int(count.Val())

	return nil
}
