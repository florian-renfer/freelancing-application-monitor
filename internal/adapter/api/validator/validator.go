package validator

type Validator interface {
	Validate(any) error
	Messages() []string
}
