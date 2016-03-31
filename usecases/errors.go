package usecases

type InvalidNameError struct{}

func (e *InvalidNameError) Error() string {
	return "A pizza order must have a name so we know whom to deliver to"
}

type InvalidAddressError struct{}

func (e *InvalidAddressError) Error() string {
	return "A pizza order must have a address so we where to deliver your pizza"
}
