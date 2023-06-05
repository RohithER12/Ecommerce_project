package deliverymodels

type ViewProductDetail struct {
	Name        string   `json:"name"`
	Price       float64  `json:"price" `
	Description string   `json:"description"`
	Sizes       []string ` json:"sizes"`
	Brand       string   `json:"brand"`
}

type ViewCartDetail struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Size        string  `json:"size"`
	Brand       string  `json:"brand"`
	Quantity    uint    `json:"quantity"`
}

type ViewAddressForUser struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"LastName"`
	PhoneNumber string `json:"phoneNumber"`
	HouseNo     string `json:"houseNo"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postalCode"`
}

type ViewInvoiceForUsers struct {
	Name          string  `json:"name"`
	OrderId       uint    `json:"orderID"`
	TotalAmount   float64 `json:"total_amount"`
	Discount      float64 `json:"discount"`
	BillAmount    float64 `json:"billAmount"`
	PaymentStatus string  `json:"PaymentStatus"`
	FromWallet    float64 `json:"from_wallet"`
}

type ViewOrdersForUsers struct {
	OrderID  uint    `json:"orderID"`
	Product  string  `json:"product"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
	Status   string  `json:"status"`
}

type ViewUserProfile struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Username    string
}
