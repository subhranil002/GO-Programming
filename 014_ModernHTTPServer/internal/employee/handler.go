package course

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/subhranil002/GO-CRUD/pkg/response"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	employees, err := h.service.List(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch employees")
		return
	}

	response.JSON(w, http.StatusOK, "Employees retrieved successfully", employees)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	employee, err := h.service.Get(
		r.Context(),
		r.PathValue("id"),
	)
	if err != nil {
		handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, "Employee retrieved successfully", employee)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateEmployeeRequest

	if err := readJSON(w, r, &req); err != nil {
		return
	}

	employee, err := h.service.Create(r.Context(), req)
	if err != nil {
		handleError(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, "Employee created successfully", employee)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var req UpdateEmployeeRequest

	if err := readJSON(w, r, &req); err != nil {
		return
	}

	employee, err := h.service.Update(
		r.Context(),
		r.PathValue("id"),
		req,
	)
	if err != nil {
		handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, "Employee updated successfully", employee)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.service.Delete(
		r.Context(),
		r.PathValue("id"),
	)
	if err != nil {
		handleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, "Employee deleted successfully", nil)
}

func readJSON(
	w http.ResponseWriter,
	r *http.Request,
	dst any,
) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("invalid JSON: %v", err),
		)

		return err
	}

	return nil
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrInvalid):
		response.Error(w, http.StatusBadRequest, err.Error())

	case errors.Is(err, ErrNotFound):
		response.Error(w, http.StatusNotFound, "employee not found")

	default:
		response.Error(w, http.StatusInternalServerError, "internal server error")
	}
}
