package errors

type Error struct {
	error map[string][]string
}

func (e *Error) Add(key, err string) {
	if _, ok := e.error[key]; !ok {
		e.error[key] = make([]string, 0, 5)
	}
	e.error[key] = append(e.error[key], err)
}

func (e *Error) Errors() map[string][]string {
	return e.error
}

func (e *Error) ErrorByKey(key string) []string {
	return e.error[key]
}

func (e *Error) HasErrors() bool {
	return len(e.error) > 0
}
func New() *Error {
	return &Error{
		error: make(map[string][]string),
	}
}
