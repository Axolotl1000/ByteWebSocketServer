package main

var (
	CONNECTED    = []byte{0xF0, 0x00}
	AUTH         = []byte{0xF0, 0x01}
	AUTH_SUCCESS = []byte{0xF0, 0x02}
)

var (
	ILLEGA_REQUEST = []byte{0xF1, 0x00}
	CLIENT_EXISTS  = []byte{0xF1, 0x01}
	AUTH_FAILED    = []byte{0xF1, 0x02}
	USER_NOT_FOUND = []byte{0xF1, 0x03}
	INTERNAL_ERROR = []byte{0xF1, 0x04}
)
