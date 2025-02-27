package graph

import "fmt"

var (
	ErrNameDuplicated          = fmt.Errorf("name duplicated")
	ErrAppNotFound             = fmt.Errorf("app not found")
	ErrAppAlreadyExists        = fmt.Errorf("app already exists")
	ErrSecretNotFound          = fmt.Errorf("secret not found")
	ErrSecretAlreadyExists     = fmt.Errorf("secret already exists")
	ErrServiceNotFound         = fmt.Errorf("service not found")
	ErrServiceAlreadyExists    = fmt.Errorf("service already exists")
	ErrStorageCantBeDownsized  = fmt.Errorf("storage can't be downsized")
	ErrDatabaseVersionMismatch = fmt.Errorf(
		"version mismatch, you need to perform a database upgrade",
	)
	ErrDatabaseVersionMustBeGreater = fmt.Errorf(
		"new version must be greater than the current version",
	)
)

type VariableRequiredError struct {
	VariableName string
}

func NewVariableRequiredError(name string) error {
	return &VariableRequiredError{
		VariableName: name,
	}
}

func (e *VariableRequiredError) Error() string {
	return fmt.Sprintf("variable required: %s", e.VariableName)
}
