// GENERATED
// DO NOT EDIT
package errtype

type NotFoundError struct {
    Err error
}

func (e NotFoundError) Error() string {
    return e.Err.Error()
}
func (e NotFoundError) Unwrap() error {
    return e.Err
}

type BadRequestError struct {
    Err error
}

func (e BadRequestError) Error() string {
    return e.Err.Error()
}
func (e BadRequestError) Unwrap() error {
    return e.Err
}

