package services

// holds all services related to retriving a billing code
type DataService struct{
	// holds a pointer to db
	// db *Database
}

// placeholder constructor that initializes the billing code services with any dependencies it requires
func NewDataService(database *Database) *DataService {
	return &DataService{
		db: database
	}
}

func (s *DataService) GetBillingCode() (string, error) {
	code := "placeholder for billing code retrieved from DB"

	return code
}