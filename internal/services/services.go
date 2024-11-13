package services

import "database/sql"

// holds all services related to retriving a billing code
type DataService struct {
	db *sql.DB
}

// placeholder constructor that initializes the billing code services with any dependencies it requires
func NewDataService(database *sql.DB) *DataService {
	return &DataService{
		db: database,
	}
}

func (s *DataService) GetBillingCode() string {
	billingCode := "placeholder for billing code retrieved from DB"

	return billingCode
}
