package types


type Recipient struct {
	Name string
	Email string
}

type EmailData struct {
	From       string
	To         string
	Name       string
	Subject    string
	Body       string
	ButtonText string
	ButtonLink string
}