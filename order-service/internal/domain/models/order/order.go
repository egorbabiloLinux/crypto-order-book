package order

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderType  string
type SideType   string
type StatusType string

const (
	Limit 	 OrderType = "limit"
	Market 	 OrderType = "market"
	Bid 	 SideType = "bid"
	Ask 	 SideType = "ask"
	Active 	 StatusType = "active"
	Partial  StatusType = "partial"
	Filled 	 StatusType = "filled"
	Canceled StatusType = "canceled"
)

func (s SideType) IsValid() bool {
	for _, side := range AllSideTypes() {
		if s == side {
			return true
		}
	}
	return false
}

func (o OrderType) IsValid() bool {
	for _, order := range AllOrderTypes() {
		if o == order {
			return true
		}
	}
	return false
}

func (s StatusType) IsValid() bool {
	for _, status := range AllStatusTypes() {
		if s == status {
			return true
		}
	}
	return false
}

func AllSideTypes() []SideType {
	return []SideType{Bid, Ask}
}

func AllOrderTypes() []OrderType {
	return []OrderType{Limit, Market}
}

func AllStatusTypes() []StatusType {
	return []StatusType{Active, Partial, Filled, Canceled}
}

type Order struct {
	id 	      	int64
	userId    	int64 
	price     	decimal.Decimal
	amount    	decimal.Decimal
	remaining 	decimal.Decimal
	side 	  	SideType
	orderType 	OrderType
	status 		StatusType
	createdAt 	time.Time
	updatedAt 	time.Time
}