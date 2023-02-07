package main

import "fmt"

// IOtp 抽象的模板类
type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type OTP struct {
	iOtp IOtp
}

func (o *OTP) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	return o.iOtp.sendNotification(message)
}

// SMS 具体的消息发送类
type SMS struct {
	OTP
}

func NewSMS() *SMS {
	sms := &SMS{}
	sms.iOtp = sms
	return sms
}

func (s *SMS) genRandomOTP(length int) string {
	return fmt.Sprintf("生成了%d长度的密码\n", length)
}

func (s *SMS) saveOTPCache(otp string) {
	fmt.Printf("%s password saved\n", otp)
}

func (s *SMS) getMessage(otp string) string {
	return "message + " + otp
}

func (s *SMS) sendNotification(msg string) error {
	fmt.Printf("message %s has been sended\n", msg)
	return nil
}

// 新增加email发送
type Email struct {
	OTP
}

func NewEmail() *Email {
	e := &Email{}
	e.iOtp = e
	return e
}

func (e *Email) genRandomOTP(length int) string {
	return fmt.Sprintf("email方式生成了%d长度的密码", length)
}

func (e *Email) saveOTPCache(otp string) {
	fmt.Println(otp, "saved")
}

func (e *Email) getMessage(opt string) string {
	return "email" + opt
}

func (e *Email) sendNotification(msg string) error {
	fmt.Println(msg, "sended")
	return nil
}

func main() {
	sms := NewSMS()
	sms.genAndSendOTP(10)

	email := NewEmail()
	email.genAndSendOTP(21)
}
