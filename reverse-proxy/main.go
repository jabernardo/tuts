package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Endpoint struct {
	Name    string `yaml:name`
	Pattern string `yaml:pattern`
	Path    string `yaml:path`
	Dest    string `yaml:dest`
}

type Config struct {
	Host      string     `yaml:host`
	Port      int        `yaml:port`
	Resources []Endpoint `yaml:endpoints`
}

func main() {
	t := Config{}

	config, err := os.ReadFile("config.yml")

	if err != nil {
		log.Fatalln("Could not found configuration file")
	}

	err = yaml.Unmarshal([]byte(config), &t)

	if err != nil {
		log.Fatalln("Invalid configuration file")
	}

	mux := http.NewServeMux()

	for i := 0; i < len(t.Resources); i++ {
		curr := t.Resources[i]

		log.Printf("Adding Proxy for [%s] Path \"%s\" to \"%s\"", curr.Name, curr.Path, curr.Dest)

		mux.HandleFunc(curr.Pattern, func(w http.ResponseWriter, r *http.Request) {
			proxy := httputil.ReverseProxy{
				Rewrite: func(req *httputil.ProxyRequest) {
					p, _ := url.Parse(curr.Dest)
					req.SetURL(p)

					req.SetXForwarded()
					req.Out.Header.Add("Custom-Header", "Value")
					req.Out.URL.Path = strings.TrimRight(strings.TrimPrefix(r.URL.Path, curr.Path), "/")

					log.Printf("[%s] %s -> %s", curr.Dest, r.URL.Path, p)
				},
			}

			proxy.ServeHTTP(w, r)
		})
	}

	http.ListenAndServe(fmt.Sprintf("%s:%d", t.Host, t.Port), mux)
}
