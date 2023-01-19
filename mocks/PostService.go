// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	posts "socialmedia/features/posts"
)

// PostService is an autogenerated mock type for the PostService type
type PostService struct {
	mock.Mock
}

// Add provides a mock function with given fields: token, file, newPost
func (_m *PostService) Add(token interface{}, file multipart.FileHeader, newPost posts.Core) (posts.Core, error) {
	ret := _m.Called(token, file, newPost)

	var r0 posts.Core
	if rf, ok := ret.Get(0).(func(interface{}, multipart.FileHeader, posts.Core) posts.Core); ok {
		r0 = rf(token, file, newPost)
	} else {
		r0 = ret.Get(0).(posts.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, multipart.FileHeader, posts.Core) error); ok {
		r1 = rf(token, file, newPost)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: token, postID
func (_m *PostService) Delete(token interface{}, postID int) error {
	ret := _m.Called(token, postID)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, int) error); ok {
		r0 = rf(token, postID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPost provides a mock function with given fields:
func (_m *PostService) GetPost() ([]posts.Core, error) {
	ret := _m.Called()

	var r0 []posts.Core
	if rf, ok := ret.Get(0).(func() []posts.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]posts.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostDetail provides a mock function with given fields: postID
func (_m *PostService) GetPostDetail(postID int) (interface{}, error) {
	ret := _m.Called(postID)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(int) interface{}); ok {
		r0 = rf(postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: token, postID, updateData
func (_m *PostService) Update(token interface{}, postID int, updateData posts.Core) (posts.Core, error) {
	ret := _m.Called(token, postID, updateData)

	var r0 posts.Core
	if rf, ok := ret.Get(0).(func(interface{}, int, posts.Core) posts.Core); ok {
		r0 = rf(token, postID, updateData)
	} else {
		r0 = ret.Get(0).(posts.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, int, posts.Core) error); ok {
		r1 = rf(token, postID, updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPostService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostService creates a new instance of PostService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostService(t mockConstructorTestingTNewPostService) *PostService {
	mock := &PostService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
