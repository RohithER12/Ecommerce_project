package deliverymodels

type UserLoginInputs struct {
	Username string `bson:"username" json:"username" binding:"required"`
	Password string `bson:"password" json:"password" binding:"required"`
}

type User struct {
	FirstName   string `json:"firstname" binding:"required" `
	LastName    string `json:"lastname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phonenumber" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type UpdateUser struct {
	FirstName   string `json:"firstname" `
	LastName    string `json:"lastname" `
	Email       string `json:"email" `
	PhoneNumber string `json:"phonenumber"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type Address struct {
	HouseNo    string `json:"houseNo" binding:"required"`
	Street     string `json:"street" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	PostalCode string `json:"postalCode" binding:"required"`
}
