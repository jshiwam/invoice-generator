package invoice

type Address struct {
	address    string
	address2   string
	postalCode string
	city       string
	country    string
}

func NewAddress(address, postalCode, city, country string) Address {
	return Address{address: address, postalCode: postalCode, city: city, country: country}
}

type PhoneNumber struct {
	countryCode string
	number      string
}

func NewPhoneNumber(countryCode, number string) PhoneNumber {
	return PhoneNumber{countryCode: countryCode, number: number}
}

type Contact struct {
	name    string
	address Address
	phone   PhoneNumber
}

func NewContact(name string, address Address, phone PhoneNumber) *Contact {
	return &Contact{name: name, address: address, phone: phone}
}
