package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderType string
type SideType  string
type StatusType string

const (
	Limit 	 OrderType = "limit"
	Market 	 OrderType = "market"
	Bid 	 SideType = "bid"
	Ask 	 SideType = "Ask"
	Active 	 StatusType = "active"
	Partial  StatusType = "partial"
	Filled 	 StatusType = "filled"
	Canceled StatusType = "canceled"
)

type Order struct {
	id 	      	int64
	userId    	int64 
	price     	decimal.Decimal
	amount    	decimal.Decimal
	remaining 	decimal.Decimal
	side 	  	SideType
	orderType 	OrderType
	orderStatus StatusType
	createdAt 	time.Time
	updatedAt 	time.Time
}