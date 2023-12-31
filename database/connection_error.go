package database

type connectionError struct{}

var ConnectionError connectionError

func (connectionError) Error() string {
	return "connection error"
}
