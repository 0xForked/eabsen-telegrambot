package domain

type User struct {
	ID				uint 	`gorm:"primarykey"`
	RoleID			uint
	UniqueID		string
	TelegramID		string
	Name			string
	Username		string
	Email			string
	Phone			string
	Status			string
	IntegrationCode	string
}

type UserRepositoryContract interface {
	Profile(TelegramId string) (res int64, err error)
	Connect(Username string, Code string, TelegramId string) (res int64, err error)
}