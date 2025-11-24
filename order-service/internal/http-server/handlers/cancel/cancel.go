package cancel

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	resp "order-service/internal/lib/api/response"
	"order-service/internal/lib/logger/slWrap"
	"order-service/internal/storage"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	OrderId int64 `json:"order_id" validate:"required"`
}

type OrderDeleter interface {
	DeleteOrder(
		orderId int64,
	) error
}

func New(log *slog.Logger, orderDeleter OrderDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.cancel.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		if err := render.DecodeJSON(r.Body, &req); errors.Is(err, io.EOF) { //TODO: put this in function
			log.Error("request body is empty")

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("request body is empty"))
			return			
		}

		if err := validator.New().Struct(&req); err != nil {
			validationErr := err.(validator.ValidationErrors)

			log.Error("Invalid request", slWrap.Err(err))

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.ValidateErr(validationErr))
			return
		}

		if req.OrderId < 0 {
			log.Error(fmt.Sprintf("invalid order id: %d", req.OrderId))

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("invalid order id"))
			return
		}

		if err := orderDeleter.DeleteOrder(req.OrderId); err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				log.Error("order not found", slWrap.Err(err))

				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, resp.Error("order not found"))
				return
			}

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("failed to delete order"))
		}

		render.JSON(w, r, resp.OK())
	}
}