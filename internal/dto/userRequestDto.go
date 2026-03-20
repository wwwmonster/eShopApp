package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	UserLogin
	Phone string `json:"phone"`
}

type VerificationCodeInput struct {
	Code string `json:"code"`
}

type SellerInput struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PhoneNumber       string `json:"phone_number"`
	BankAccountNumber string `json:"bankAccountNumber"`
	SwiftCode         string `json:"swiftCode"`
	PaymentType       string `json:"paymentType"`
}
