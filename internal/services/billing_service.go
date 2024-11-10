package services

// holds all services related to retriving a billing code
type BillingCodeService struct{}

// constructor that initializes the billing code services with any dependencies it requires
func NewDataService() *BillingCodeService {
	return &BillingCodeService{}
}
