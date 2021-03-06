package rux

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildRequestUrl_Params(t *testing.T) {
	is := assert.New(t)

	b := NewBuildRequestURL()
	b.Path(`/news/{category_id}/{new_id}/detail`)
	b.Params(M{"{category_id}": "100", "{new_id}": "20"})

	is.Equal(b.Build().String(), `/news/100/20/detail`)
}

func TestBuildRequestUrl_Host(t *testing.T) {
	is := assert.New(t)

	b := NewBuildRequestURL()
	b.Scheme("https")
	b.Host("127.0.0.1")
	b.Path(`/news`)

	is.Equal(b.Build().String(), `https://127.0.0.1/news`)
}

func TestBuildRequestURL_User(t *testing.T) {
	is := assert.New(t)

	b := NewBuildRequestURL()
	b.Scheme("https")
	b.User("tom", "123")
	b.Host("127.0.0.1")
	b.Path(`/news`)

	is.Equal(b.Build().String(), `https://tom:123@127.0.0.1/news`)
}

func TestBuildRequestUrl_Queries(t *testing.T) {
	is := assert.New(t)

	var u = make(url.Values)
	u.Add("username", "admin")
	u.Add("password", "12345")

	b := NewBuildRequestURL()
	b.Queries(u)
	b.Path(`/news`)

	is.Equal(b.Build().String(), `/news?password=12345&username=admin`)
}

func TestBuildRequestUrl_Build(t *testing.T) {
	is := assert.New(t)

	r := New()

	homepage := NewNamedRoute("homepage", `/build-test/{name}/{id:\d+}`, emptyHandler, GET)
	homepageFiexdPath := NewNamedRoute("homepage_fiexd_path", `/build-test/fiexd/path`, emptyHandler, GET)

	r.AddRoute(homepage)
	r.AddRoute(homepageFiexdPath)

	b := NewBuildRequestURL()
	b.Params(M{"{name}": "test", "{id}": "20"})

	is.Equal(r.BuildRequestURL("homepage", b).String(), `/build-test/test/20`)
	is.Equal(r.BuildRequestURL("homepage_fiexd_path").String(), `/build-test/fiexd/path`)
}

func TestBuildRequestUrl_With(t *testing.T) {
	r := New()
	is := assert.New(t)

	homepage := NewNamedRoute("homepage", `/build-test/{name}/{id:\d+}`, emptyHandler, GET)

	r.AddRoute(homepage)

	is.Equal(r.BuildRequestURL("homepage", M{
		"{name}":   "test",
		"{id}":     20,
		"username": "demo",
	}).String(), `/build-test/test/20?username=demo`)
}

func TestBuildRequestUrl_WithCustom(t *testing.T) {
	is := assert.New(t)

	b := NewBuildRequestURL()
	b.Path("/build-test/test/{id}")

	is.Equal(b.Build(M{
		"{id}":     20,
		"username": "demo",
	}).String(), `/build-test/test/20?username=demo`)
}

func TestBuildRequestUrl_WithMutilArgs(t *testing.T) {
	r := New()
	is := assert.New(t)

	homepage := NewNamedRoute("homepage", `/build-test/{name}/{id:\d+}`, emptyHandler, GET)

	r.AddRoute(homepage)

	str := r.BuildRequestURL("homepage", "{name}", "test", "{id}", 20, "username", "demo").String()
	is.Equal(`/build-test/test/20?username=demo`, str)
}

func TestBuildRequestUrl_WithMutilArgs2(t *testing.T) {
	r := New()
	is := assert.New(t)

	homepage := NewNamedRoute("homepage", `/build-test`, emptyHandler, GET)

	r.AddRoute(homepage)

	str := r.BuildRequestURL("homepage", "{name}", "test", "{id}", 20, "username", "demo").String()
	is.Equal(`/build-test?username=demo`, str)

	str = r.BuildURL("homepage", "{name}", "test", "{id}", 20).String()
	is.Equal(`/build-test`, str)
}

func TestBuildRequestUrl_WithMutilArgs3(t *testing.T) {
	r := New()
	is := assert.New(t)

	homepage := NewNamedRoute("homepage", `/build-test/{id}`, emptyHandler, GET)

	r.AddRoute(homepage)

	str := r.BuildRequestURL("homepage", "{name}", "test", "{id}", 20, "username", "demo").String()
	is.Equal(`/build-test/20?username=demo`, str)

	str = r.BuildURL("homepage", "{name}", "test", "{id}", 23).String()
	is.Equal(`/build-test/23`, str)
}

func TestBuildRequestUrl_EmptyRoute(t *testing.T) {
	r := New()
	is := assert.New(t)

	homepage := NewNamedRoute("homepage", `/build-test/{name}/{id:\d+}`, emptyHandler, GET)

	r.AddRoute(homepage)

	is.PanicsWithValue("BuildRequestURL get route (name: homepage-empty) is nil", func() {
		r.BuildRequestURL("homepage-empty", "{name}", "test", "{id}", "20", "username", "demo")
	})
}

func TestBuildRequestUrl_ErrorArgs(t *testing.T) {
	r := New()
	is := assert.New(t)

	homepage := NamedRoute("homepage", `/build-test/{name}/{id:\d+}`, emptyHandler, GET)
	r.AddRoute(homepage)

	is.PanicsWithValue("buildRequestURLs odd argument count", func() {
		r.BuildRequestURL("homepage", "one")
	})

	is.PanicsWithValue("buildRequestURLs odd argument count", func() {
		r.BuildRequestURL("homepage", "{name}", "test", "{id}")
	})
}
