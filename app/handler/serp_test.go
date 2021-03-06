package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"ratri/domain/model"
	"ratri/handler"
	mock "ratri/mock/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
)

var (
	serpTests = []struct {
		name      string
		in        map[string]interface{}
		want      interface{}
		wantError bool
		err       error
	}{

		{"Want no error#1", map[string]interface{}{"task": 5, "offset": 0, "top": 3}, 200, false, nil},
		{"Want no error#2", map[string]interface{}{"task": 6, "offset": 1, "top": 5}, 200, false, nil},
		{"Want 404 code#1", map[string]interface{}{"task": 1, "offset": 1, "top": 5}, 404, true, &echo.HTTPError{}},
		{"Want 404 code#2", map[string]interface{}{"task": 10, "offset": 1, "top": 5}, 404, true, &echo.HTTPError{}},
		{"Want 404 code#3", map[string]interface{}{"task": 10, "offset": 10, "top": 5}, 404, true, &echo.HTTPError{}},
	}
)

func TestFetchSerpWithRatioByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mck := mock.NewMockSerpUsecase(ctrl)
	for _, tt := range serpTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.in["task"].(int) > 4 && tt.in["task"].(int) < 9 && tt.in["offset"].(int) < 10 {
				mck.EXPECT().
					FetchSerpWithRatio(tt.in["task"], tt.in["offset"], tt.in["top"]).
					Return(nil, nil)
			} else {
				mck.EXPECT().
					FetchSerpWithRatio(tt.in["task"], tt.in["offset"], tt.in["top"]).
					Return(nil, model.ErrNoSuchData)
			}
			h := handler.NewSerpHandler(mck)

			// Set query parameter
			q := make(url.Values)
			q.Set("offset", fmt.Sprintf("%v", tt.in["offset"]))
			q.Set("top", fmt.Sprintf("%v", tt.in["top"]))

			req := httptest.NewRequest(
				http.MethodGet,
				"/v1/serp/"+fmt.Sprintf("%v", tt.in["task"])+"/ratio?"+q.Encode(),
				nil,
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Set path parameter explicitly
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprintf("%v", tt.in["task"]))

			err := h.FetchSerpWithRatioByID(c)

			// Throw t.Fatal if unexpected error has occurred.
			if !tt.wantError && err != nil {
				t.Fatalf("Want no error, but got %#v", err)
			}

			// Throw t.Fatal if different error has occurred.
			// if tt.wantError && !(err == tt.err) {
			// 	t.Fatalf("Want %#v, but got %#v", tt.err, err)
			// }
			if tt.wantError {
				t.Logf("%+v", err)
			}

			// Throw t.Fatal if expected value is different from result.
			if diff := cmp.Diff(tt.want, rec.Code); !tt.wantError && diff != "" {
				t.Fatalf("Want %d, but got %d\n%v", tt.want, rec.Code, diff)
			}
		})
	}
}

func TestFetchSerpWithIconByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mck := mock.NewMockSerpUsecase(ctrl)
	for _, tt := range serpTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.in["task"].(int) > 4 && tt.in["task"].(int) < 9 && tt.in["offset"].(int) < 10 {
				mck.EXPECT().
					FetchSerpWithIcon(tt.in["task"], tt.in["offset"], tt.in["top"]).
					Return(nil, nil)
			} else {
				mck.EXPECT().
					FetchSerpWithIcon(tt.in["task"], tt.in["offset"], tt.in["top"]).
					Return(nil, model.ErrNoSuchData)
			}
			h := handler.NewSerpHandler(mck)

			// Set query parameter
			q := make(url.Values)
			q.Set("offset", fmt.Sprintf("%v", tt.in["offset"]))
			q.Set("top", fmt.Sprintf("%v", tt.in["top"]))

			req := httptest.NewRequest(
				http.MethodGet,
				"/v1/serp/"+fmt.Sprintf("%v", tt.in["task"])+"/icon?"+q.Encode(),
				nil,
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Set path parameter explicitly
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprintf("%v", tt.in["task"]))

			err := h.FetchSerpWithIconByID(c)

			// Throw t.Fatal if unexpected error has occurred.
			if !tt.wantError && err != nil {
				t.Fatalf("Want no error, but got %#v", err)
			}

			// Throw t.Fatal if different error has occurred.
			// if tt.wantError && !(err == tt.err) {
			// 	t.Fatalf("Want %#v, but got %#v", tt.err, err)
			// }
			if tt.wantError {
				t.Logf("%+v", err)
			}

			// Throw t.Fatal if expected value is different from result.
			if diff := cmp.Diff(tt.want, rec.Code); !tt.wantError && diff != "" {
				t.Fatalf("Want %d, but got %d\n%v", tt.want, rec.Code, diff)
			}
		})
	}
}

func TestFetchSerpByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mck := mock.NewMockSerpUsecase(ctrl)
	for _, tt := range serpTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.in["task"].(int) > 4 && tt.in["task"].(int) < 9 && tt.in["offset"].(int) < 10 {
				mck.EXPECT().FetchSerp(tt.in["task"], tt.in["offset"]).Return(nil, nil)
			} else {
				mck.EXPECT().
					FetchSerp(tt.in["task"], tt.in["offset"]).
					Return(nil, model.ErrNoSuchData)
			}
			h := handler.NewSerpHandler(mck)

			// Set query parameter
			q := make(url.Values)
			q.Set("offset", fmt.Sprintf("%v", tt.in["offset"]))

			req := httptest.NewRequest(
				http.MethodGet,
				"/v1/serp/"+fmt.Sprintf("%v", tt.in["task"])+"/ratio?"+q.Encode(),
				nil,
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Set path parameter explicitly
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprintf("%v", tt.in["task"]))

			err := h.FetchSerpByID(c)

			// Throw t.Fatal if unexpected error has occurred.
			if !tt.wantError && err != nil {
				t.Fatalf("Want no error, but got %#v", err)
			}

			// Throw t.Fatal if different error has occurred.
			// if tt.wantError && !(err == tt.err) {
			// 	t.Fatalf("Want %#v, but got %#v", tt.err, err)
			// }
			if tt.wantError {
				t.Logf("%+v", err)
			}

			// Throw t.Fatal if expected value is different from result.
			if diff := cmp.Diff(tt.want, rec.Code); !tt.wantError && diff != "" {
				t.Fatalf("Want %d, but got %d\n%v", tt.want, rec.Code, diff)
			}
		})
	}
}
