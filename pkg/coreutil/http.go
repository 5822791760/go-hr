package coreutil

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/5822791760/hr/pkg/apperr"
	"github.com/go-chi/chi/v5"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

type httpError interface {
	ToHttp() apperr.HttpErr
}

func WriteError(w http.ResponseWriter, err httpError) {
	if err == nil {
		return
	}

	data := err.ToHttp()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(data.Code))
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetParam(r *http.Request, key string) (string, apperr.Err) {
	data := chi.URLParam(r, key)
	if data == "" {
		return "", apperr.NewQueryNotExistErr(key)
	}

	return data, nil
}

func GetParamInt(r *http.Request, key string) (int, apperr.Err) {
	query, err := GetParam(r, key)
	if err != nil {
		return 0, err
	}

	data, xerr := strconv.Atoi(query)
	if xerr != nil {
		return 0, apperr.NewInternalServerErr(xerr)
	}

	return data, nil
}

func GetQuery(r *http.Request, key string) (string, apperr.Err) {
	query := r.URL.Query()
	data := query.Get(key)

	if data == "" {
		return "", apperr.NewQueryNotExistErr(key)
	}

	return data, nil
}

func GetQueryInt(r *http.Request, key string) (int, apperr.Err) {
	query, err := GetQuery(r, key)
	if err != nil {
		return 0, err
	}

	data, xerr := strconv.Atoi(query)
	if xerr != nil {
		return 0, apperr.NewInternalServerErr(xerr)
	}

	return data, nil
}

func ParseBody(r *http.Request, dest interface{}) apperr.Err {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if xerr := decoder.Decode(dest); xerr != nil {
		return apperr.NewInternalServerErr(xerr)
	}

	return nil
}

func GetContext(r *http.Request, db interface{}) context.Context {
	ctx := r.Context()
	return StoreContextDB(ctx, db)
}

func GetTxContext(r *http.Request, tx Transactionable) (context.Context, func(err apperr.Err) apperr.Err, apperr.Err) {
	ctx := r.Context()

	ctx, err := StartTransaction(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

	end := func(err apperr.Err) apperr.Err {
		if err != nil {
			return err
		}

		tx, err := GetContextTx(ctx)
		if err != nil {
			return err
		}

		xerr := tx.Commit()
		if xerr != nil {
			return apperr.NewInternalServerErr(xerr)
		}

		return nil
	}

	return ctx, end, nil
}
