package internal

type CustomError struct{}

func (e *CustomError) Error() string {
	return "custom error"
}
