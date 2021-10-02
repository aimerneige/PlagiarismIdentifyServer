// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

type ProgramLanguage uint8

const (
	NONE   ProgramLanguage = 0 // if homework type are not programming language
	JAVA   ProgramLanguage = 1
	C_CPP  ProgramLanguage = 2
	PYTHON ProgramLanguage = 3
)
