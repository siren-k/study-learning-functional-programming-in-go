package workflow

import "encoding/base64"

type Monad func(error) (Data, error)

func Base64ToBytes(d Data) Monad {
	dString := d.(string)
	return func(e error) (Data, error) {
		return base64.StdEncoding.DecodeString(dString)
	}
}

func BytesToString(d Data) Monad {
	dString := d
	return func(e error) (Data, error) {
		return dString, nil
	}
}

func Get(d Data) Monad {
	return func(e error) (Data, error) {
		return d, e
	}
}
