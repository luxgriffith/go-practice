package urlshort

import (
	"net/http"
)

type MapHandlerVars struct {
	pathsToUrls map[string]string
	fallback    http.Handler
}

func (m *MapHandlerVars) New(pathsToUrls map[string]string, fallback http.Handler) {
	m.pathsToUrls = pathsToUrls
	m.fallback = fallback
}

var mapVars MapHandlerVars

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	mapVars = MapHandlerVars{pathsToUrls: pathsToUrls, fallback: fallback}
	return http.HandlerFunc(mapHandlerFunc)
}

func mapHandlerFunc(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	mapUrl := mapVars.pathsToUrls[path]
	if mapUrl != "" {
		http.RedirectHandler(mapUrl, 301).ServeHTTP(w, r)
	} else {
		mapVars.fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return nil, nil
}
