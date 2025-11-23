package place

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"

	"order-service/internal/domain/models/order"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	resp "order-service/internal/lib/api/response"
	"order-service/internal/lib/logger/slWrap"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/shopspring/decimal"
)

type Response struct {
	resp.Response
	OrderId int64 `json:"order_id"`
}

type Request struct {
	UserId 	  int64				`json:"user_id" validate:"required"`
	Price 	  decimal.Decimal	`json:"price" validate:"required"`
	Amount 	  decimal.Decimal	`json:"amount" validate:"required"`
	Side 	  order.SideType	`json:"side" validate:"required"`
	OrderType order.OrderType	`json:"orderType" validate:"required"`
}

type OrderSaver interface {
	SaveOrder(
		userId 	  int64,
		price     decimal.Decimal,
		amount    decimal.Decimal,
		remaining decimal.Decimal,
		side      order.SideType,
		orderType order.OrderType,	
		status 	  order.StatusType,
		createdAt time.Time,
		updatedAt time.Time,
	) (int64, error)
}

func New(log *slog.Logger, orderSaver OrderSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.place.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		if err := render.DecodeJSON(r.Body, &req); errors.Is(err, io.EOF) {
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

		zero := decimal.NewFromInt(0)

		if req.Price.LessThan(zero){
			log.Error("invalid order price")

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("invalid order price"))
			return
		}

		if req.Amount.LessThan(zero) {
			log.Error("invalid order amount")

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("invalid order amount"))
			return			
		}

		if req.Side.IsValid() {
			log.Error("invalid order side type")

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("invalid order side type"))
			return				
		}

		if req.OrderType.IsValid() {
			log.Error("invalid order side type")

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("invalid order side type"))
			return				
		}

		id, err := orderSaver.SaveOrder(
			req.UserId,
			req.Price,
			req.Amount,
			req.Amount,
			req.Side,
			req.OrderType,
			order.Active,
			time.Now(),
			time.Now(),
		)
		if err != nil {
			log.Error("failed to add order", slWrap.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("failed to add order"))
			return
		}

		render.JSON(w, r, Response {
			Response: resp.OK(),
			OrderId: id,
		})
	}
}