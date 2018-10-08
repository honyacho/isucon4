package main

import (
	"net/http"
	"testing"
	"html/template"
	"github.com/martini-contrib/render"
)

type(
	Render struct { }
)

func BenchmarkTest(b *testing.B) {
	req := &http.Request {
		Header: map[string][]string {
			"X-Advertiser-Id": []string{"2"},
		},
	}
	r := &Render{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		routeGetReport(req, r)
	}
}

func (r *Render) JSON(status int, v interface{}) { }
func (r *Render) HTML(status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) { }
func (r *Render) XML(status int, v interface{}) { }
func (r *Render) Data(status int, v []byte) { }
func (r *Render) Text(status int, v string) { }
func (r *Render) Error(status int) { }
func (r *Render) Status(status int) { }
func (r *Render) Redirect(location string, status ...int) { }
func (r *Render) Template() *template.Template { return nil }
func (r *Render) Header() http.Header { return nil }
