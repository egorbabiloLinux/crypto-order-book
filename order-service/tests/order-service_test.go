package tests

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

const (
	host = "localhost:8083"
)

func TestPlaceOrder_HappyPath(t *testing.T) {
	u := url.URL {
		Scheme: "http",
		Host: host,
	}

	e := httpexpect.Default(t, u.String())

	e.POST("/order").WithJSON(map[string]interface{}{
		"user_id": 1,
		"price": 1,
		"amount": 1,
		"side": "bid",
		"orderType": "limit",
	}).Expect().Status(http.StatusOK).
	JSON().Object().ContainsKey("order_id")
}