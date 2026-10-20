package config

type OTP struct {
	Length int    `env:"OTP_CODE_LENGTH"`
	Chars  string `env:"OTP_CODE_CHARS"`
}
