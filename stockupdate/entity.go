package stockupdate

import (
	"errors"
	"fmt"

	"github.com/cloudstateio/go-support/cloudstate/crdt"
	"github.com/cloudstateio/go-support/cloudstate/encoding"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
)

type AllStocks struct {
	stocks *crdt.ORSet
}

func NewStock(crdt.EntityID) crdt.EntityHandler {
	return &AllStocks{}
}

// func (s *AllStocks) getStocks() (*Stocks, error) {
// 	var stocks Stocks
// 	for _, state := range s.stocks.Value() {
// 		var stock WarehouseStock
// 		if err := encoding.UnmarshalAny(state, &stock); err != nil {
// 			return nil, fmt.Errorf("_________________failed to decodestruct %v", err)
// 		}
// 		stocks.Stocks = append(stocks.Stocks, &stock)
// 	}
// 	return stocks, nil
// }
// func (s *AllStocks) AggregateStock(storeStock *AggregateStoreStock) (*Stocks, error) {
// 	if storeStock.GetQuantity() <= 0 {
// 		return nil, fmt.Errorf("unbale to add negative quantity: %v", storeStock.GetQuantity())
// 	}
// 	newstock, err := encoding.MarshalAny(&WarehouseStock{
// 		WarehouseUid: storeStock.GetWarehouseUid(),
// 		Quantity:     storeStock.GetQuantity(),
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to marshal newstock: %v", err)
// 	}
// 	key := encoding.String(storeStock.GetWarehouseUid())
// 	reg, err := s.stocks.LWWRegister(key)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to register key: %v", err)
// 	}
// 	if reg != nil {
// 		reg.Set(newstock)
// 	} else {
// 		reg = crdt.NewLWWRegister(newstock)
// 	}
// 	s.stocks.Set(key, reg)
// 	return s.getStocks()
// }

func (s *AllStocks) HandleCommand(ctx *crdt.CommandContext, name string, msg proto.Message) (*any.Any, error) {
	switch m := msg.(type) {
	case *GetStoreStock:
		var stocks Stocks
		for _, state := range s.stocks.Value() {
			var warehousestock WarehouseStock
			if err := encoding.UnmarshalAny(state, &warehousestock); err != nil {
				return nil, fmt.Errorf("failed to unmarshal state: %v", err)
			}
			stocks.Stocks = append(stocks.Stocks, &warehousestock)
		}
		return encoding.MarshalAny(&stocks)
	case *AggregateStoreStock:
		if m.GetWarehouseStock().GetQuantity() <= 0 {
			return nil, errors.New("can't add negative quantity")
		}
		addstock, err := encoding.MarshalAny(m.GetWarehouseStock())
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal aggregate stock input: %v", err)
		}
		s.stocks.Add(addstock)
		return encoding.MarshalAny(addstock)
	}
	return encoding.Empty, nil
}

func (s *AllStocks) Default(ctx *crdt.Context) (crdt.CRDT, error) {
	return crdt.NewORSet(), nil
}

func (s *AllStocks) Set(ctx *crdt.Context, state crdt.CRDT) error {
	s.stocks = state.(*crdt.ORSet)
	return nil
}
