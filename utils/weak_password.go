// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package utils

// IsWeakPassword simply check if password is weak password
func IsWeakPassword(password string) bool {
	if password == "00000000" {
		return true
	}
	if password == "000000000" {
		return true
	}
	if password == "11111111" {
		return true
	}
	if password == "111111111" {
		return true
	}
	if password == "66666666" {
		return true
	}
	if password == "666666666" {
		return true
	}
	if password == "88888888" {
		return true
	}
	if password == "888888888" {
		return true
	}
	if password == "11223344" {
		return true
	}
	if password == "12345678" {
		return true
	}
	if password == "123456789" {
		return true
	}
	if password == "password" {
		return true
	}
	return false
}
