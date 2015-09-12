package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Unknwon/macaron"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/Unknwon/fws/modules/middleware"
	"github.com/Unknwon/fws/routers/v1"
)

func Test_Fibonacci(t *testing.T) {
	Convey("Test Fibonacci calculator", t, func() {
		// Setup dummy server.
		m := macaron.New()
		m.Use(macaron.Renderer())
		m.Use(middleware.Contexter())

		m.Get("/fibonacci", v1.Fibonacci)

		Convey("Invalid parameters", func() {
			Convey("Give a negative number", func() {
				resp := httptest.NewRecorder()
				req, err := http.NewRequest("GET", "/fibonacci?n=-1", nil)
				So(err, ShouldBeNil)
				m.ServeHTTP(resp, req)

				So(resp.Code, ShouldEqual, 422)
				So(resp.Body.String(), ShouldEqual, `{"error":"Unprocessable number: -1"}`)
			})

			Convey("Give a super large number", func() {
				resp := httptest.NewRecorder()
				req, err := http.NewRequest("GET", "/fibonacci?n=150", nil)
				So(err, ShouldBeNil)
				m.ServeHTTP(resp, req)

				So(resp.Code, ShouldEqual, 422)
				So(resp.Body.String(), ShouldEqual, `{"error":"Unprocessable number: 150"}`)
			})
		})

		Convey("Valid parameters", func() {
			// 0
			resp := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/fibonacci?n=0", nil)
			So(err, ShouldBeNil)
			m.ServeHTTP(resp, req)

			So(resp.Code, ShouldEqual, 200)
			So(resp.Body.String(), ShouldEqual, `{"sequence":[]}`)

			// 1
			resp = httptest.NewRecorder()
			req, err = http.NewRequest("GET", "/fibonacci?n=1", nil)
			So(err, ShouldBeNil)
			m.ServeHTTP(resp, req)

			So(resp.Code, ShouldEqual, 200)
			So(resp.Body.String(), ShouldEqual, `{"sequence":[0]}`)

			// 2
			resp = httptest.NewRecorder()
			req, err = http.NewRequest("GET", "/fibonacci?n=2", nil)
			So(err, ShouldBeNil)
			m.ServeHTTP(resp, req)

			So(resp.Code, ShouldEqual, 200)
			So(resp.Body.String(), ShouldEqual, `{"sequence":[0,1]}`)

			// 5
			resp = httptest.NewRecorder()
			req, err = http.NewRequest("GET", "/fibonacci?n=5", nil)
			So(err, ShouldBeNil)
			m.ServeHTTP(resp, req)

			So(resp.Code, ShouldEqual, 200)
			So(resp.Body.String(), ShouldEqual, `{"sequence":[0,1,1,2,3]}`)

			// 10
			resp = httptest.NewRecorder()
			req, err = http.NewRequest("GET", "/fibonacci?n=10", nil)
			So(err, ShouldBeNil)
			m.ServeHTTP(resp, req)

			So(resp.Code, ShouldEqual, 200)
			So(resp.Body.String(), ShouldEqual, `{"sequence":[0,1,1,2,3,5,8,13,21,34]}`)
		})
	})
}
