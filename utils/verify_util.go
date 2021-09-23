package utils

import (
	"regexp"
)

// Verify the legality of the data

// VerifyEmailFormat verify email format
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyChinaPhoneNumberFormat verify china phone number format
func VerifyChinaPhoneNumberFormat(phoneNumber string) bool {
	pattern := `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[0-35-9]\d{2}|4(?:0\d|1[0-2]|9\d))|9[0-35-9]\d{2}|6[2567]\d{2}|4(?:(?:10|4[01])\d{3}|[68]\d{4}|[579]\d{2}))\d{6}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phoneNumber)
}
