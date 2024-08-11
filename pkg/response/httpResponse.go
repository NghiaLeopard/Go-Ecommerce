package response

import "net/http"

const (
	INVALID        = 400
	ALREADY_EXIST  = 409
	GET_SUCCESS    = 200
	ACTION_SUCCESS = 201
	UNAUTHORIZED   = 401
	INTERNAL_ERROR = 500
)

type TypeResponse struct {
	Type   string
	Status int
}

var httpResponse = map[int]TypeResponse{
	INVALID: {
		Type:   "INVALID",
		Status: http.StatusBadRequest,
	},
	ALREADY_EXIST: {
		Type:   "ALREADY_EXIST",
		Status: http.StatusConflict,
	},
	GET_SUCCESS: {
		Type:   "SUCCESS",
		Status: http.StatusOK,
	},
	ACTION_SUCCESS: {
		Type:   "SUCCESS",
		Status: http.StatusCreated,
	},
	UNAUTHORIZED: {
		Type:   "UNAUTHORIZED",
		Status: http.StatusUnauthorized,
	},
	INTERNAL_ERROR: {
		Type:   "INTERNAL_SERVER_ERROR",
		Status: http.StatusInternalServerError,
	},
}