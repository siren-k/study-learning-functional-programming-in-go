package workflow

type Data interface{}

func Next(m Monad, f func(Data) Monad) Monad {
	return func(e error) (Data, error) {
		newData, newError := m(e)
		if newError != nil {
			return nil, newError
		}
		return f(newData)(newError)
	}
}
