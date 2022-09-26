// Code generated by goa v3.8.5, DO NOT EDIT.
//
// inventory gRPC server types
//
// Command:
// $ goa gen characters/design

package server

import (
	inventorypb "characters/gen/grpc/inventory/pb"
	inventory "characters/gen/inventory"
	inventoryviews "characters/gen/inventory/views"

	goa "goa.design/goa/v3/pkg"
)

// NewListPayload builds the payload of the "list" endpoint of the "inventory"
// service from the gRPC request type.
func NewListPayload(message *inventorypb.ListRequest) *inventory.ListPayload {
	v := &inventory.ListPayload{
		CharacterID: message.CharacterId,
	}
	return v
}

// NewProtoStoredInventoryCollection builds the gRPC response type from the
// result of the "list" endpoint of the "inventory" service.
func NewProtoStoredInventoryCollection(result inventoryviews.StoredInventoryCollectionView) *inventorypb.StoredInventoryCollection {
	message := &inventorypb.StoredInventoryCollection{}
	message.Field = make([]*inventorypb.StoredInventory, len(result))
	for i, val := range result {
		message.Field[i] = &inventorypb.StoredInventory{}
		if val.ID != nil {
			message.Field[i].Id = *val.ID
		}
		if val.CharacterID != nil {
			message.Field[i].CharacterId = *val.CharacterID
		}
		if val.Items != nil {
			message.Field[i].Items = make([]*inventorypb.StoredItem, len(val.Items))
			for j, val := range val.Items {
				message.Field[i].Items[j] = &inventorypb.StoredItem{}
				if val.ID != nil {
					message.Field[i].Items[j].Id = *val.ID
				}
				if val.Name != nil {
					message.Field[i].Items[j].Name = *val.Name
				}
				if val.Description != nil {
					message.Field[i].Items[j].Description = *val.Description
				}
				if val.Damage != nil {
					message.Field[i].Items[j].Damage = *val.Damage
				}
				if val.Healing != nil {
					message.Field[i].Items[j].Healing = *val.Healing
				}
				if val.Protection != nil {
					message.Field[i].Items[j].Protection = *val.Protection
				}
			}
		}
	}
	return message
}

// NewShowPayload builds the payload of the "show" endpoint of the "inventory"
// service from the gRPC request type.
func NewShowPayload(message *inventorypb.ShowRequest, view *string) *inventory.ShowPayload {
	v := &inventory.ShowPayload{
		ID: message.Id,
	}
	v.View = view
	return v
}

// NewProtoShowResponse builds the gRPC response type from the result of the
// "show" endpoint of the "inventory" service.
func NewProtoShowResponse(result *inventoryviews.StoredInventoryView) *inventorypb.ShowResponse {
	message := &inventorypb.ShowResponse{}
	if result.ID != nil {
		message.Id = *result.ID
	}
	if result.CharacterID != nil {
		message.CharacterId = *result.CharacterID
	}
	if result.Items != nil {
		message.Items = make([]*inventorypb.StoredItem, len(result.Items))
		for i, val := range result.Items {
			message.Items[i] = &inventorypb.StoredItem{}
			if val.ID != nil {
				message.Items[i].Id = *val.ID
			}
			if val.Name != nil {
				message.Items[i].Name = *val.Name
			}
			if val.Description != nil {
				message.Items[i].Description = *val.Description
			}
			if val.Damage != nil {
				message.Items[i].Damage = *val.Damage
			}
			if val.Healing != nil {
				message.Items[i].Healing = *val.Healing
			}
			if val.Protection != nil {
				message.Items[i].Protection = *val.Protection
			}
		}
	}
	return message
}

// NewShowNotFoundError builds the gRPC error response type from the error of
// the "show" endpoint of the "inventory" service.
func NewShowNotFoundError(er *inventory.NotFound) *inventorypb.ShowNotFoundError {
	message := &inventorypb.ShowNotFoundError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// NewShowItemPayload builds the payload of the "showItem" endpoint of the
// "inventory" service from the gRPC request type.
func NewShowItemPayload(message *inventorypb.ShowItemRequest, view *string) *inventory.ShowItemPayload {
	v := &inventory.ShowItemPayload{
		ID: message.Id,
	}
	v.View = view
	return v
}

// NewProtoStoredItemCollection builds the gRPC response type from the result
// of the "showItem" endpoint of the "inventory" service.
func NewProtoStoredItemCollection(result inventoryviews.StoredItemCollectionView) *inventorypb.StoredItemCollection {
	message := &inventorypb.StoredItemCollection{}
	message.Field = make([]*inventorypb.StoredItem, len(result))
	for i, val := range result {
		message.Field[i] = &inventorypb.StoredItem{}
		if val.ID != nil {
			message.Field[i].Id = *val.ID
		}
		if val.Name != nil {
			message.Field[i].Name = *val.Name
		}
		if val.Description != nil {
			message.Field[i].Description = *val.Description
		}
		if val.Damage != nil {
			message.Field[i].Damage = *val.Damage
		}
		if val.Healing != nil {
			message.Field[i].Healing = *val.Healing
		}
		if val.Protection != nil {
			message.Field[i].Protection = *val.Protection
		}
	}
	return message
}

// NewShowItemNotFoundError builds the gRPC error response type from the error
// of the "showItem" endpoint of the "inventory" service.
func NewShowItemNotFoundError(er *inventory.NotFound) *inventorypb.ShowItemNotFoundError {
	message := &inventorypb.ShowItemNotFoundError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// NewAddPayload builds the payload of the "add" endpoint of the "inventory"
// service from the gRPC request type.
func NewAddPayload(message *inventorypb.AddRequest) *inventory.AddPayload {
	v := &inventory.AddPayload{
		CharacterID: message.CharacterId,
	}
	return v
}

// NewProtoAddResponse builds the gRPC response type from the result of the
// "add" endpoint of the "inventory" service.
func NewProtoAddResponse(result string) *inventorypb.AddResponse {
	message := &inventorypb.AddResponse{}
	message.Field = result
	return message
}

// NewAddItemPayload builds the payload of the "addItem" endpoint of the
// "inventory" service from the gRPC request type.
func NewAddItemPayload(message *inventorypb.AddItemRequest) *inventory.AddItemPayload {
	v := &inventory.AddItemPayload{
		ID:     message.Id,
		ItemID: message.ItemId,
	}
	if message.View != "" {
		v.View = &message.View
	}
	return v
}

// NewProtoAddItemResponse builds the gRPC response type from the result of the
// "addItem" endpoint of the "inventory" service.
func NewProtoAddItemResponse(result string) *inventorypb.AddItemResponse {
	message := &inventorypb.AddItemResponse{}
	message.Field = result
	return message
}

// NewRemoveItemPayload builds the payload of the "removeItem" endpoint of the
// "inventory" service from the gRPC request type.
func NewRemoveItemPayload(message *inventorypb.RemoveItemRequest) *inventory.RemoveItemPayload {
	v := &inventory.RemoveItemPayload{
		ID:     message.Id,
		ItemID: message.ItemId,
	}
	return v
}

// NewProtoRemoveItemResponse builds the gRPC response type from the result of
// the "removeItem" endpoint of the "inventory" service.
func NewProtoRemoveItemResponse() *inventorypb.RemoveItemResponse {
	message := &inventorypb.RemoveItemResponse{}
	return message
}

// NewRemovePayload builds the payload of the "remove" endpoint of the
// "inventory" service from the gRPC request type.
func NewRemovePayload(message *inventorypb.RemoveRequest) *inventory.RemovePayload {
	v := &inventory.RemovePayload{
		ID: message.Id,
	}
	return v
}

// NewProtoRemoveResponse builds the gRPC response type from the result of the
// "remove" endpoint of the "inventory" service.
func NewProtoRemoveResponse() *inventorypb.RemoveResponse {
	message := &inventorypb.RemoveResponse{}
	return message
}

// NewUpdatePayload builds the payload of the "update" endpoint of the
// "inventory" service from the gRPC request type.
func NewUpdatePayload(message *inventorypb.UpdateRequest) *inventory.UpdatePayload {
	v := &inventory.UpdatePayload{
		ID: message.Id,
	}
	if message.Inventory != nil {
		v.Inventory = protobufInventorypbInventory2ToInventoryInventory(message.Inventory)
	}
	return v
}

// NewProtoUpdateResponse builds the gRPC response type from the result of the
// "update" endpoint of the "inventory" service.
func NewProtoUpdateResponse() *inventorypb.UpdateResponse {
	message := &inventorypb.UpdateResponse{}
	return message
}

// ValidateStoredItem runs the validations defined on StoredItem.
func ValidateStoredItem(elem *inventorypb.StoredItem) (err error) {

	return
}

// ValidateAddItemRequest runs the validations defined on AddItemRequest.
func ValidateAddItemRequest(message *inventorypb.AddItemRequest) (err error) {
	if message.View != "" {
		if !(message.View == "default" || message.View == "tiny") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.view", message.View, []interface{}{"default", "tiny"}))
		}
	}
	return
}

// ValidateUpdateRequest runs the validations defined on UpdateRequest.
func ValidateUpdateRequest(message *inventorypb.UpdateRequest) (err error) {
	if message.Inventory == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("inventory", "message"))
	}
	if message.Inventory != nil {
		if err2 := ValidateInventory2(message.Inventory); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateInventory2 runs the validations defined on Inventory2.
func ValidateInventory2(inventory *inventorypb.Inventory2) (err error) {
	if inventory.Character == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("character", "inventory"))
	}
	if inventory.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "inventory"))
	}
	return
}

// ValidateStoredCharacter runs the validations defined on StoredCharacter.
func ValidateStoredCharacter(character *inventorypb.StoredCharacter) (err error) {

	return
}

// protobufInventorypbInventory2ToInventoryInventory builds a value of type
// *inventory.Inventory from a value of type *inventorypb.Inventory2.
func protobufInventorypbInventory2ToInventoryInventory(v *inventorypb.Inventory2) *inventory.Inventory {
	res := &inventory.Inventory{}
	if v.Character != nil {
		res.Character = protobufInventorypbStoredCharacterToInventoryStoredCharacter(v.Character)
	}
	if v.Items != nil {
		res.Items = make([]*inventory.StoredItem, len(v.Items))
		for i, val := range v.Items {
			res.Items[i] = &inventory.StoredItem{
				ID:         val.Id,
				Name:       val.Name,
				Damage:     val.Damage,
				Healing:    val.Healing,
				Protection: val.Protection,
			}
			if val.Description != "" {
				res.Items[i].Description = &val.Description
			}
		}
	}

	return res
}

// protobufInventorypbStoredCharacterToInventoryStoredCharacter builds a value
// of type *inventory.StoredCharacter from a value of type
// *inventorypb.StoredCharacter.
func protobufInventorypbStoredCharacterToInventoryStoredCharacter(v *inventorypb.StoredCharacter) *inventory.StoredCharacter {
	res := &inventory.StoredCharacter{
		ID:         v.Id,
		Name:       v.Name,
		Health:     v.Health,
		Experience: v.Experience,
	}
	if v.Description != "" {
		res.Description = &v.Description
	}

	return res
}

// svcInventoryInventoryToInventorypbInventory2 builds a value of type
// *inventorypb.Inventory2 from a value of type *inventory.Inventory.
func svcInventoryInventoryToInventorypbInventory2(v *inventory.Inventory) *inventorypb.Inventory2 {
	res := &inventorypb.Inventory2{}
	if v.Character != nil {
		res.Character = svcInventoryStoredCharacterToInventorypbStoredCharacter(v.Character)
	}
	if v.Items != nil {
		res.Items = make([]*inventorypb.StoredItem, len(v.Items))
		for i, val := range v.Items {
			res.Items[i] = &inventorypb.StoredItem{
				Id:         val.ID,
				Name:       val.Name,
				Damage:     val.Damage,
				Healing:    val.Healing,
				Protection: val.Protection,
			}
			if val.Description != nil {
				res.Items[i].Description = *val.Description
			}
		}
	}

	return res
}

// svcInventoryStoredCharacterToInventorypbStoredCharacter builds a value of
// type *inventorypb.StoredCharacter from a value of type
// *inventory.StoredCharacter.
func svcInventoryStoredCharacterToInventorypbStoredCharacter(v *inventory.StoredCharacter) *inventorypb.StoredCharacter {
	res := &inventorypb.StoredCharacter{
		Id:         v.ID,
		Name:       v.Name,
		Health:     v.Health,
		Experience: v.Experience,
	}
	if v.Description != nil {
		res.Description = *v.Description
	}

	return res
}