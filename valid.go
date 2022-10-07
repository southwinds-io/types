package types

type Valid interface {
	Validate() error
}
