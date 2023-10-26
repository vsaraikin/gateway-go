package models

type RequestModel struct{}

type Validator interface {
	// Validate checks whether Requested model is correct before sending it
	Validate() error
}

func (m *RequestModel) Validate() error {
	return nil
}
