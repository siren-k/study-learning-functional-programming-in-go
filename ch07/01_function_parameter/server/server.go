package server

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	defaultServerMaxMessageSize     = 1024 * 1024 * 4
	defaultMaxNumber                = 30
	defaultMaxConcurrentConnections = 2
)

var defaultServerOptions = options{
	maxMessageSize:           defaultServerMaxMessageSize,
	maxNumber:                defaultMaxNumber,
	maxConcurrentConnections: defaultMaxConcurrentConnections,
}

type convert func(int) (string, error)

type options struct {
	maxMessageSize           int
	maxNumber                int
	maxConcurrentConnections int
	convertFn                convert
	useNumberHandler         bool
}

type Logger interface {
	Printf(format string, v ...interface{})
}

type Server struct {
	logger  Logger
	opts    options
	handler http.Handler
}

func New(opt ...ServerOption) (*Server, error) {
	opts := defaultServerOptions
	for _, f := range opt {
		err := f(&opts)
		if err != nil {
			return nil, errors.Wrap(err, "error setting option")
		}
	}
	s := &Server{
		opts:   opts,
		logger: log.New(os.Stdout, "", 0),
	}
	s.register()
	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

func (s *Server) register() {
	mux := http.NewServeMux()
	if s.opts.useNumberHandler {
		mux.Handle("/", http.HandlerFunc(s.displayNumber))
	} else {
		mux.Handle("/", http.FileServer(http.Dir("./")))
	}
	s.handler = mux
}

func (s *Server) displayNumber(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("displayNumber called with number=%s\n", r.URL.Query().Get("number"))
	if numberParam := r.URL.Query().Get("number"); numberParam != "" {
		number, err := strconv.Atoi(numberParam)
		if err != nil {
			writeJson(w, map[string]interface{}{
				"error": fmt.Sprintf("number (%d) too big. Max number: %d", number, s.opts.maxNumber),
			}, http.StatusBadRequest)
		} else {
			var convertedNumber string
			if s.opts.convertFn == nil {
				convertedNumber = numberParam
			} else {
				convertedNumber, err = s.opts.convertFn(number)
			}
			if err != nil {
				writeJson(w, map[string]interface{}{
					"error": "error running convertFn number",
				}, http.StatusBadRequest)
			} else {
				writeJson(w, map[string]interface{}{
					"displayNumber": convertedNumber,
				})
			}
		}
	} else {
		writeJson(w, map[string]interface{}{
			"error": "missing number",
		}, http.StatusBadRequest)
	}
}

func writeJson(w http.ResponseWriter, v interface{}, statuses ...int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if len(statuses) > 0 {
		w.WriteHeader(statuses[0])
	}
	json.NewEncoder(w).Encode(v)
}
