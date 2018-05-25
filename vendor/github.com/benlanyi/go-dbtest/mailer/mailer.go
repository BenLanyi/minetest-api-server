package mailer

import "fmt"

type sendMessage interface {
	generateMessage() string
}

type RealEmail struct {
	Token string
	Email string
}

type ConsoleEmail struct {
	Token string
}

func (c ConsoleEmail) generateMessage() string {
	message := fmt.Sprintf("To activate your account follow the below link\nhttp://localhost:3000/verify/%s", c.Token)
	return message
}

func (c RealEmail) generateMessage() string {
	return "Not Implemented"
}

// SendMail : sends mail
func SendMail(s sendMessage) {
	fmt.Println(s.generateMessage())
}
