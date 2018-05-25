package database

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/benlanyi/go-dbtest/mailer"
	"github.com/benlanyi/go-dbtest/person"

	"golang.org/x/crypto/bcrypt"
)

type register struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), err
}

func (db database) NewUser(body []byte) returnMessage {
	var registerDetails register

	json.Unmarshal(body, &registerDetails)
	fmt.Println("details from auth", registerDetails)
	//generate password hash
	var passwordByte, err = bcrypt.GenerateFromPassword([]byte(registerDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print(err.Error())
	}
	passwordHash := base64.URLEncoding.EncodeToString(passwordByte)
	fmt.Println("Password Hash: ", passwordHash)
	//generate verification token
	verificationToken, err := generateToken()
	if err != nil {
		fmt.Print(err.Error())
	}
	//generate reset token
	resetToken, err := generateToken()
	if err != nil {
		fmt.Print(err.Error())
	}
	//enter details into database
	input := fmt.Sprintf("INSERT INTO users (first_name, last_name, email, password_hash, verification_token, reset_token) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')", registerDetails.FirstName, registerDetails.LastName, registerDetails.Email, passwordHash, verificationToken, resetToken)
	db.database.MustExec(input)

	var mail mailer.ConsoleEmail
	mail.Token = verificationToken
	mailer.SendMail(mail)

	// confirm success
	var returnMessage returnMessage
	returnMessage.Status = "success"
	returnMessage.Message = "User Created"
	return returnMessage

}

func (db database) SendVerification(email string) {

}

func (db database) Verify(token string) returnMessage {
	input := fmt.Sprintf("UPDATE users SET verified = 'true' WHERE verification_token = '%s'", token)
	db.database.MustExec(input)
	var returnMessage returnMessage

	// result, err := db.database.Exec(input)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(result.RowsAffected())
	//
	// if result != "0" {
	// 	returnMessage.Status = "success"
	// 	returnMessage.Message = "User Verified"
	// 	fmt.Println("user verified")
	// } else {
	// 	returnMessage.Status = "fail"
	// 	returnMessage.Message = "Issue Verifying User"
	// 	fmt.Println("failed")
	// }
	returnMessage.Status = "success"
	returnMessage.Message = "User Verified"
	return returnMessage
}

func (db database) LogIn(email string, password string) {
	people := []person.Person{}
	input := fmt.Sprintf("SELECT * FROM person WHERE name = '%s'", name)
	db.database.Select(&people, input)
}
