// Package traefikbodytransform plugin.
package traefikbodytransform

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"fmt"
)

// Config the plugin configuration.
type Config struct {

	TokenTransformQueryParameterFieldName string `json:"tokenTransformQueryParameterFieldName,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		TokenTransformQueryParameterFieldName: "token",
	}
}

// transformer plugin.
type transformer struct {
	next                                  http.Handler
	name                                  string
	tokenTransformQueryParameterFieldName string
}

// New created a new transformer plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	fmt.Println("traefikbodytransform ==> initialized")
	return &transformer{
		next:                                  next,
		name:                                  name,
		tokenTransformQueryParameterFieldName: config.TokenTransformQueryParameterFieldName,
	}, nil
}

func (a *transformer) log(format string) {
	_, writeLogError := os.Stderr.WriteString(a.name + ": " + format)
	if writeLogError != nil {
		panic(writeLogError.Error())
	}
}

func (a *transformer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	transformerOption := make(map[string]bool)
    token := req.URL.Query().Get(a.tokenTransformQueryParameterFieldName)
	fmt.Println("traefikbodytransform:token ==>", token)
	req.Header.Set("Authorization", "Bearer "+token)
	a.next.ServeHTTP(rw, req)
}
