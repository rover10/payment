package database

type DBClient interface {
	DBConnect(config *config.Config) error
	CreatePayment(model.Payment) (model.Payment, error)
}
