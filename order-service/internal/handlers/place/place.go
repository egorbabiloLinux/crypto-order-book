package place

import "log/slog"

type OrderSaver interface {
	SaveOrder()
}

func New(log *slog.Logger, )