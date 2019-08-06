package jsonresume

type Basics struct {
	Name string
	Label string
	Picture string
	Email string
	Phone string
	Website string
	Summary string
	Location Location
	Profiles []Profile
}

type Location struct {
	Address string
	PostalCode string
	City string
	CountryCode string
	Region string
}

type Profile struct {
	Network string
	Username string
	Url string
}