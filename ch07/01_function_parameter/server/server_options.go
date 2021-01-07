package server

import (
	"github.com/pkg/errors"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch07/01_function_parameter/utils"
)

type ServerOption func(*options) error

func MaxConcurrentConnections(n int) ServerOption {
	return func(o *options) error {
		if n > Config.MaxConcurrentConnections {
			return errors.New("error setting MaxConcurrentConnections")
		}
		o.maxConcurrentConnections = n
		return nil
	}
}

func MaxNumber(n int) ServerOption {
	return func(o *options) error {
		o.maxNumber = n
		return nil
	}
}

func FormatNumber(fn convert) ServerOption {
	return func(o *options) error {
		o.convertFn = fn
		return nil
	}
}

func UseNumberHandler(b bool) ServerOption {
	return func(o *options) error {
		o.useNumberHandler = b
		return nil
	}
}
