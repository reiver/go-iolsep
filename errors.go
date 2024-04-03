package iolsep

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilPipeReader = erorr.Error("iolsep: nil pipe-reader")
	errNilPipeWriter = erorr.Error("iolsep: nil pipe-writer")
	errNilReader     = erorr.Error("iolsep: nil reader")
	errNilReceiver   = erorr.Error("iolsep: nil receiver")
	errNilWriter     = erorr.Error("iolsep: nil writer")
)

const (
	errNilReturnEOFFunction   = erorr.Error("iolsep: nil return-eof function")
	errNilReturnErrorFunction = erorr.Error("iolsep: nil return-error function")
	errNilWriteRuneFunction   = erorr.Error("iolsep: nil write-rune function")
)
