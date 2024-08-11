package https

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/5822791760/hr/pkg/errs"
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
	ToHttp() errs.HttpErr
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

func GetParam(r *http.Request, key string) (string, errs.Err) {
	data := chi.URLParam(r, key)
	if data == "" {
		return "", errs.NewQueryNotExistErr(key)
	}

	return data, nil
}

func GetParamInt(r *http.Request, key string) (int, errs.Err) {
	query, err := GetParam(r, key)
	if err != nil {
		return 0, err
	}

	data, xerr := strconv.Atoi(query)
	if xerr != nil {
		return 0, errs.NewInternalServerErr(xerr)
	}

	return data, nil
}

func GetQuery(r *http.Request, key string) (string, errs.Err) {
	query := r.URL.Query()
	data := query.Get(key)

	if data == "" {
		return "", errs.NewQueryNotExistErr(key)
	}

	return data, nil
}

func GetQueryInt(r *http.Request, key string) (int, errs.Err) {
	query, err := GetQuery(r, key)
	if err != nil {
		return 0, err
	}

	data, xerr := strconv.Atoi(query)
	if xerr != nil {
		return 0, errs.NewInternalServerErr(xerr)
	}

	return data, nil
}

func ParseBody(r *http.Request, dest interface{}) errs.Err {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(dest); err != nil {
		return errs.NewInternalServerErr(err)
	}

	return nil
}
