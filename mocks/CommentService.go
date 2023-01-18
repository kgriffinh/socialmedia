// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	comments "socialmedia/features/comments"

	mock "github.com/stretchr/testify/mock"
)

// CommentService is an autogenerated mock type for the CommentService type
type CommentService struct {
	mock.Mock
}

// Add provides a mock function with given fields: token, newComment
func (_m *CommentService) Add(token interface{}, newComment comments.Core) (comments.Core, error) {
	ret := _m.Called(token, newComment)

	var r0 comments.Core
	if rf, ok := ret.Get(0).(func(interface{}, comments.Core) comments.Core); ok {
		r0 = rf(token, newComment)
	} else {
		r0 = ret.Get(0).(comments.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, comments.Core) error); ok {
		r1 = rf(token, newComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: token, commentID
func (_m *CommentService) Delete(token interface{}, commentID int) error {
	ret := _m.Called(token, commentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int) error); ok {
		r0 = rf(token, commentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: token, commentID, updateData
func (_m *CommentService) Update(token interface{}, commentID int, updateData comments.Core) (comments.Core, error) {
	ret := _m.Called(token, commentID, updateData)

	var r0 comments.Core
	if rf, ok := ret.Get(0).(func(interface{}, int, comments.Core) comments.Core); ok {
		r0 = rf(token, commentID, updateData)
	} else {
		r0 = ret.Get(0).(comments.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, int, comments.Core) error); ok {
		r1 = rf(token, commentID, updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentService creates a new instance of CommentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentService(t mockConstructorTestingTNewCommentService) *CommentService {
	mock := &CommentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}