// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	service "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/service"
)

// AuthMetadataServiceClient is an autogenerated mock type for the AuthMetadataServiceClient type
type AuthMetadataServiceClient struct {
	mock.Mock
}

type AuthMetadataServiceClient_FlyteClient struct {
	*mock.Call
}

func (_m AuthMetadataServiceClient_FlyteClient) Return(_a0 *service.FlyteClientResponse, _a1 error) *AuthMetadataServiceClient_FlyteClient {
	return &AuthMetadataServiceClient_FlyteClient{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AuthMetadataServiceClient) OnFlyteClient(ctx context.Context, in *service.FlyteClientRequest, opts ...grpc.CallOption) *AuthMetadataServiceClient_FlyteClient {
	c := _m.On("FlyteClient", ctx, in, opts)
	return &AuthMetadataServiceClient_FlyteClient{Call: c}
}

func (_m *AuthMetadataServiceClient) OnFlyteClientMatch(matchers ...interface{}) *AuthMetadataServiceClient_FlyteClient {
	c := _m.On("FlyteClient", matchers...)
	return &AuthMetadataServiceClient_FlyteClient{Call: c}
}

// FlyteClient provides a mock function with given fields: ctx, in, opts
func (_m *AuthMetadataServiceClient) FlyteClient(ctx context.Context, in *service.FlyteClientRequest, opts ...grpc.CallOption) (*service.FlyteClientResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *service.FlyteClientResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.FlyteClientRequest, ...grpc.CallOption) *service.FlyteClientResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.FlyteClientResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.FlyteClientRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AuthMetadataServiceClient_OAuth2Metadata struct {
	*mock.Call
}

func (_m AuthMetadataServiceClient_OAuth2Metadata) Return(_a0 *service.OAuth2MetadataResponse, _a1 error) *AuthMetadataServiceClient_OAuth2Metadata {
	return &AuthMetadataServiceClient_OAuth2Metadata{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AuthMetadataServiceClient) OnOAuth2Metadata(ctx context.Context, in *service.OAuth2MetadataRequest, opts ...grpc.CallOption) *AuthMetadataServiceClient_OAuth2Metadata {
	c := _m.On("OAuth2Metadata", ctx, in, opts)
	return &AuthMetadataServiceClient_OAuth2Metadata{Call: c}
}

func (_m *AuthMetadataServiceClient) OnOAuth2MetadataMatch(matchers ...interface{}) *AuthMetadataServiceClient_OAuth2Metadata {
	c := _m.On("OAuth2Metadata", matchers...)
	return &AuthMetadataServiceClient_OAuth2Metadata{Call: c}
}

// OAuth2Metadata provides a mock function with given fields: ctx, in, opts
func (_m *AuthMetadataServiceClient) OAuth2Metadata(ctx context.Context, in *service.OAuth2MetadataRequest, opts ...grpc.CallOption) (*service.OAuth2MetadataResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *service.OAuth2MetadataResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.OAuth2MetadataRequest, ...grpc.CallOption) *service.OAuth2MetadataResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.OAuth2MetadataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.OAuth2MetadataRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}