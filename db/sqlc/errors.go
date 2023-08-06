package db

type StoreError struct {
	Code int
	Err  error
}

func (ce StoreError) Error() string {
	return ce.Err.Error()
}
