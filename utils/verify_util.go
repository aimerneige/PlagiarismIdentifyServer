// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package utils

import (
	"regexp"
)

const (
	EMIAL_VERIFY_PATTERN = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	PHONE_VERIFY_PATTERN = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[0-35-9]\d{2}|4(?:0\d|1[0-2]|9\d))|9[0-35-9]\d{2}|6[2567]\d{2}|4(?:(?:10|4[01])\d{3}|[68]\d{4}|[579]\d{2}))\d{6}$`
)

// VerifyEmailFormat verify email format
func VerifyEmailFormat(email string) bool {
	if email == "" {
		return false
	}
	reg := regexp.MustCompile(EMIAL_VERIFY_PATTERN)
	return reg.MatchString(email)
}

// VerifyChinaPhoneNumberFormat verify china phone number format
func VerifyChinaPhoneNumberFormat(phoneNumber string) bool {
	if phoneNumber == "" {
		return false
	}
	reg := regexp.MustCompile(PHONE_VERIFY_PATTERN)
	return reg.MatchString(phoneNumber)
}
