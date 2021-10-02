// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

type Teacher struct {
	User
}

type TeacherDto struct {
	UserDto
}

func (t Teacher) ToDto(dto TeacherDto) {
	dto.ID = t.ID
	dto.Account = t.Account
	dto.Password = t.Password
	dto.ProfileID = t.Profile.ID
	dto.Name = t.Name
	dto.Phone = t.Phone
	dto.Email = t.Email
}
