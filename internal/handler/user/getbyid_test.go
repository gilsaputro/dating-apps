package user

import (
	"context"
	"fmt"
	user_service "gilsaputro/dating-apps/internal/service/user"
	mock_service "gilsaputro/dating-apps/internal/service/user/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
)

func TestUserHandler_GetByIDUserHandler(t *testing.T) {
	type args struct {
		token   string
		body    string
		timeout int
	}
	type want struct {
		body string
		code int
	}
	tests := []struct {
		name        string
		args        args
		mockFunc    func(m mock_service.MockUserServiceMethod)
		mockContext func() (context.Context, func())
		want        want
	}{
		{
			name: "success flow",
			args: args{
				body:    `1`,
				timeout: 5,
				token:   "token_baru",
			},
			mockFunc: func(m mock_service.MockUserServiceMethod) {
				m.EXPECT().GetUserByID(user_service.GetByIDServiceRequest{
					TokenRequest: "token_baru",
					UserId:       1,
				}).Return(user_service.UserServiceInfo{
					UserId:      1,
					Username:    "1",
					Fullname:    "1",
					Email:       "1",
					CreatedDate: "1",
				}, nil)
			},
			mockContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			want: want{
				code: 200,
				body: `{"data":{"user":{"id":1,"username":"1","email":"1","fullname":"1","created_date":"1"}},"code":200,"message":"success"}`,
			},
		},
		{
			name: "error on service flow",
			args: args{
				body:    `1`,
				timeout: 5,
				token:   "token_baru",
			},
			mockFunc: func(m mock_service.MockUserServiceMethod) {
				m.EXPECT().GetUserByID(user_service.GetByIDServiceRequest{
					TokenRequest: "token_baru",
					UserId:       1,
				}).Return(user_service.UserServiceInfo{}, fmt.Errorf("some error"))
			},
			mockContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			want: want{
				code: 500,
				body: `{"code":500,"message":"some error"}`,
			},
		},
		{
			name: "error on service flow user already exists",
			args: args{
				body:    `1`,
				timeout: 5,
				token:   "token_baru",
			},
			mockFunc: func(m mock_service.MockUserServiceMethod) {
				m.EXPECT().GetUserByID(user_service.GetByIDServiceRequest{
					TokenRequest: "token_baru",
					UserId:       1,
				}).Return(user_service.UserServiceInfo{}, user_service.ErrUnauthorized)
			},
			mockContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			want: want{
				code: 401,
				body: `{"code":401,"message":"unauthorized"}`,
			},
		},
		{
			name: "error on service data not found",
			args: args{
				body:    `1`,
				timeout: 5,
				token:   "token_baru",
			},
			mockFunc: func(m mock_service.MockUserServiceMethod) {
				m.EXPECT().GetUserByID(user_service.GetByIDServiceRequest{
					TokenRequest: "token_baru",
					UserId:       1,
				}).Return(user_service.UserServiceInfo{}, user_service.ErrDataNotFound)
			},
			mockContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			want: want{
				code: 404,
				body: `{"code":404,"message":"data not found"}`,
			},
		},
		{
			name: "error on invalid token",
			args: args{
				body:    `1`,
				timeout: 5,
				token:   "",
			},
			mockFunc: func(m mock_service.MockUserServiceMethod) {
			},
			mockContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			want: want{
				code: 500,
				body: `{"code":500,"message":"Internal Server Error"}`,
			},
		},
		{
			name: "error on invalid body value",
			args: args{
				body:    ``,
				timeout: 5,
				token:   "token_baru",
			},
			mockFunc: func(m mock_service.MockUserServiceMethod) {
			},
			mockContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			want: want{
				code: 400,
				body: `{"code":400,"message":"Invalid Parameter Request"}`,
			},
		},
		{
			name: "got timeout flow",
			args: args{
				body:    `1`,
				timeout: 0,
				token:   "token_baru",
			},
			mockFunc: func(m mock_service.MockUserServiceMethod) {
				m.EXPECT().GetUserByID(user_service.GetByIDServiceRequest{
					TokenRequest: "token_baru",
					UserId:       1,
				}).Return(user_service.UserServiceInfo{}, nil).AnyTimes()
			},
			mockContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			want: want{
				code: 504,
				body: `{"code":504,"message":"Timeout"}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mService := mock_service.NewMockUserServiceMethod(mockCtrl)
			tt.mockFunc(*mService)
			defer mockCtrl.Finish()
			handler := UserHandler{
				service:      mService,
				timeoutInSec: tt.args.timeout,
			}
			r := httptest.NewRequest(http.MethodGet, "/user/{id}", strings.NewReader(tt.args.body))
			ctx, cancel := tt.mockContext()
			defer cancel()
			r = r.WithContext(ctx)
			if len(tt.args.token) > 0 {
				r = r.WithContext(context.WithValue(r.Context(), "token", tt.args.token))
			}
			w := httptest.NewRecorder()

			vars := map[string]string{
				"id": tt.args.body,
			}
			r = mux.SetURLVars(r, vars)

			handler.GetByIDUserHandler(w, r)
			result := w.Result()
			resBody, err := ioutil.ReadAll(result.Body)

			if err != nil {
				t.Fatalf("Error read body err = %v\n", err)
			}

			if string(resBody) != tt.want.body {
				t.Fatalf("GetStatHandler body got =%s, want %s \n", string(resBody), tt.want.body)
			}

			if result.StatusCode != tt.want.code {
				t.Fatalf("GetStatHandler status code got =%d, want %d \n", result.StatusCode, tt.want.code)
			}
		})
	}
}
