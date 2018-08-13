package render

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"
)

var globalTplFuncMap = template.FuncMap{
	// don't escape content
	"raw": func(s string) string {
		return s
	},
	"trim": strings.TrimSpace,
	"join": strings.Join,
	// upper first char
	"upFirst": func(s string) string {
		if len(s) != 0 {
			f := s[0]
			// is lower
			if f >= 'a' && f <= 'z' {
				return strings.ToUpper(string(f)) + string(s[1:])
			}
		}

		return s
	},
}

var layoutTplFuncMap = template.FuncMap{
	// include other template file
	"include": func(filePath string) (template.HTML, error) {
		var buf bytes.Buffer
		t := template.Must(template.New("include").ParseFiles(filePath))

		if err := t.Execute(&buf, nil); err != nil {
			panic(err)
			// return "", nil
		}

		// Return safe HTML here since we are rendering our own template.
		return template.HTML(buf.String()), nil
	},
}

func (r *HTTPRenderer) HTMLString(w http.ResponseWriter, status int, html string) error {
	w.Header().Set(ContentType, r.opts.ContentHTML)
	w.WriteHeader(status)

	_, err := w.Write([]byte(html))
	return err
}

func (r *HTTPRenderer) HTML(w http.ResponseWriter, status int, template string, v interface{}) error {
	w.Header().Set(ContentType, r.opts.ContentHTML)
	w.WriteHeader(status)

	return r.Renderer.HTML(w, template, v)
}

// Template
func (r *HTTPRenderer) Template(w http.ResponseWriter, status int, html string) error {
	w.Header().Set(ContentType, r.opts.ContentHTML)
	w.WriteHeader(status)

	_, err := w.Write([]byte(html))
	return err
}

func (r *HTTPRenderer) TplString(w http.ResponseWriter, status int, tplContent string, v interface{}) error {
	w.Header().Set(ContentType, r.opts.ContentHTML)
	w.WriteHeader(status)

	t := template.Must(template.New("").Parse(tplContent))
	if err := t.Execute(w, v); err != nil {
		panic(err)
		return err
	}

	return nil
}
