// Code generated by mockery v2.7.4. DO NOT EDIT.

package webhookmock

import (
	context "context"

	model "github.com/slok/kubewebhook/v2/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// Webhook is an autogenerated mock type for the Webhook type
type Webhook struct {
	mock.Mock
}

// ID provides a mock function with given fields:
func (_m *Webhook) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Kind provides a mock function with given fields:
func (_m *Webhook) Kind() model.WebhookKind {
	ret := _m.Called()

	var r0 model.WebhookKind
	if rf, ok := ret.Get(0).(func() model.WebhookKind); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.WebhookKind)
	}

	return r0
}

// Review provides a mock function with given fields: ctx, ar
func (_m *Webhook) Review(ctx context.Context, ar model.AdmissionReview) (model.AdmissionResponse, error) {
	ret := _m.Called(ctx, ar)

	var r0 model.AdmissionResponse
	if rf, ok := ret.Get(0).(func(context.Context, model.AdmissionReview) model.AdmissionResponse); ok {
		r0 = rf(ctx, ar)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.AdmissionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.AdmissionReview) error); ok {
		r1 = rf(ctx, ar)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
