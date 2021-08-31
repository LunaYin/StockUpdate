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
	stocks *crdt.ORMap
}

func NewStock(crdt.EntityID) crdt.EntityHandler {
	return &AllStocks{}
}

func (s *AllStocks) HandleCommand(ctx *crdt.CommandContext, name string, msg proto.Message) (*any.Any, error) {
	switch m := msg.(type) {
	case *GetStockLevel:
		stocks := &AllStockLevels{}
		for _, state := range s.stocks.Entries() {
			var stocklevel StockLevel
			if err := encoding.UnmarshalAny(state.Value.(*crdt.LWWRegister).Value(), &stocklevel); err != nil {
				return nil, fmt.Errorf("failed to unmarshal state: %v", err)
			}
			stocks.Allstocklevels = append(stocks.Allstocklevels, &stocklevel)
		}
		return encoding.MarshalAny(&stocks)
	case *AddOrderInfo:
		addorder, err := encoding.MarshalAny(&OrderInfo{
			UserId: m.GetUserId(),
			ItemId: m.GetItemId(),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal add order info: %v", err)
		}
		key := encoding.String(m.GetUserId())
		reg, err := s.stocks.LWWRegister(key)
		if err != nil {
			return nil, err
		}
		if reg != nil {
			reg.Set(addorder)
		} else {
			reg = crdt.NewLWWRegister(addorder)
		}
		s.stocks.Set(key, reg)
		orderinfos := &AllOrderInfo{}
		var orderinfo OrderInfo
		for _, state := range s.stocks.Entries() {
			if err := encoding.UnmarshalAny(state.Value.(*crdt.LWWRegister).Value(), &orderinfo); err != nil {
				return nil, fmt.Errorf("failed to unmarshal state: %v", err)
			}
			orderinfos.AllorderInfo = append(orderinfos.AllorderInfo, &orderinfo)
		}
		return encoding.MarshalAny(&orderinfo)
	case *AggregateStockLevel:
		if m.GetStockLevel() <= 0 {
			return nil, errors.New("can't add negative quantity")
		}
		addstock, err := encoding.MarshalAny(&StockLevel{
			StoreUid:   m.GetStoreUid(),
			StockLevel: m.GetStockLevel(),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal aggregate stock input: %v", err)
		}
		key := encoding.String(m.GetStoreUid())
		reg, err := s.stocks.LWWRegister(key)
		if err != nil {
			return nil, err
		}
		if reg != nil {
			reg.Set(addstock)
		} else {
			reg = crdt.NewLWWRegister(addstock)
		}
		s.stocks.Set(key, reg)
		stocks := &AllStockLevels{}
		for _, state := range s.stocks.Entries() {
			var stocklevel StockLevel
			if err := encoding.UnmarshalAny(state.Value.(*crdt.LWWRegister).Value(), &stocklevel); err != nil {
				return nil, fmt.Errorf("failed to unmarshal state: %v", err)
			}
			stocks.Allstocklevels = append(stocks.Allstocklevels, &stocklevel)
		}
		return encoding.MarshalAny(stocks)
	}
	return encoding.Empty, nil
}

func (s *AllStocks) Default(ctx *crdt.Context) (crdt.CRDT, error) {
	return crdt.NewORMap(), nil
}

func (s *AllStocks) Set(ctx *crdt.Context, state crdt.CRDT) error {
	s.stocks = state.(*crdt.ORMap)
	return nil
}
