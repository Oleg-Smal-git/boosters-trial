package controllers

var (
	// These are controller instances.{{ range $controller, $instance := .Controllers }}{{"\n\t"}}{{ $instance }}{{"\t"}}= {{ $controller }}{}{{ end }}
)

var (
	// handlers is a map of routes and functions that control them.
	handlers = map[string]map[string]string { {{ range $method, $routes := .Handlers }}
		"{{ $method }}": {
			{{ range $route, $handler := $routes }}"{{ $route }}": {{ $handler }},{{ end }}
		},{{ end }}
	}
)