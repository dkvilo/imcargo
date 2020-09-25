package controller

// Controller data structure
type Controller struct {
	Version string
}

// New - Initialize Controller
func New() (*Controller) {
	return &Controller{}
}
