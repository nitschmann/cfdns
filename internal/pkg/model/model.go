package model

// Model is the base interface of the model package and should be implemented by all structs
type Model interface {
	Validate() error
}
