// Code generated by goa v3.8.5, DO NOT EDIT.
//
// character gRPC client encoders and decoders
//
// Command:
// $ goa gen characters/design

package client

import (
	character "characters/gen/character"
	characterviews "characters/gen/character/views"
	characterpb "characters/gen/grpc/character/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildListFunc builds the remote method to invoke for "character" service
// "list" endpoint.
func BuildListFunc(grpccli characterpb.CharacterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.List(ctx, reqpb.(*characterpb.ListRequest), opts...)
		}
		return grpccli.List(ctx, &characterpb.ListRequest{}, opts...)
	}
}

// DecodeListResponse decodes responses from the character list endpoint.
func DecodeListResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*characterpb.StoredCharacterCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("character", "list", "*characterpb.StoredCharacterCollection", v)
	}
	res := NewListResult(message)
	vres := characterviews.StoredCharacterCollection{Projected: res, View: view}
	if err := characterviews.ValidateStoredCharacterCollection(vres); err != nil {
		return nil, err
	}
	return character.NewStoredCharacterCollection(vres), nil
}

// BuildShowFunc builds the remote method to invoke for "character" service
// "show" endpoint.
func BuildShowFunc(grpccli characterpb.CharacterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Show(ctx, reqpb.(*characterpb.ShowRequest), opts...)
		}
		return grpccli.Show(ctx, &characterpb.ShowRequest{}, opts...)
	}
}

// EncodeShowRequest encodes requests sent to character show endpoint.
func EncodeShowRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*character.ShowPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("character", "show", "*character.ShowPayload", v)
	}
	if payload.View != nil {
		(*md).Append("view", *payload.View)
	}
	return NewProtoShowRequest(payload), nil
}

// DecodeShowResponse decodes responses from the character show endpoint.
func DecodeShowResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*characterpb.ShowResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("character", "show", "*characterpb.ShowResponse", v)
	}
	res := NewShowResult(message)
	vres := &characterviews.StoredCharacter{Projected: res, View: view}
	if err := characterviews.ValidateStoredCharacter(vres); err != nil {
		return nil, err
	}
	return character.NewStoredCharacter(vres), nil
}

// BuildAddFunc builds the remote method to invoke for "character" service
// "add" endpoint.
func BuildAddFunc(grpccli characterpb.CharacterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Add(ctx, reqpb.(*characterpb.AddRequest), opts...)
		}
		return grpccli.Add(ctx, &characterpb.AddRequest{}, opts...)
	}
}

// EncodeAddRequest encodes requests sent to character add endpoint.
func EncodeAddRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*character.Character)
	if !ok {
		return nil, goagrpc.ErrInvalidType("character", "add", "*character.Character", v)
	}
	return NewProtoAddRequest(payload), nil
}

// DecodeAddResponse decodes responses from the character add endpoint.
func DecodeAddResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*characterpb.AddResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("character", "add", "*characterpb.AddResponse", v)
	}
	res := NewAddResult(message)
	return res, nil
}

// BuildRemoveFunc builds the remote method to invoke for "character" service
// "remove" endpoint.
func BuildRemoveFunc(grpccli characterpb.CharacterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Remove(ctx, reqpb.(*characterpb.RemoveRequest), opts...)
		}
		return grpccli.Remove(ctx, &characterpb.RemoveRequest{}, opts...)
	}
}

// EncodeRemoveRequest encodes requests sent to character remove endpoint.
func EncodeRemoveRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*character.RemovePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("character", "remove", "*character.RemovePayload", v)
	}
	return NewProtoRemoveRequest(payload), nil
}

// BuildUpdateFunc builds the remote method to invoke for "character" service
// "update" endpoint.
func BuildUpdateFunc(grpccli characterpb.CharacterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Update(ctx, reqpb.(*characterpb.UpdateRequest), opts...)
		}
		return grpccli.Update(ctx, &characterpb.UpdateRequest{}, opts...)
	}
}

// EncodeUpdateRequest encodes requests sent to character update endpoint.
func EncodeUpdateRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*character.UpdatePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("character", "update", "*character.UpdatePayload", v)
	}
	return NewProtoUpdateRequest(payload), nil
}