package consts

type EnvMode string

func (e EnvMode) String() string {
	return string(e)
}

const (
	EnvModeDevelopment EnvMode = "development"
	EnvModeProduction  EnvMode = "production"
	EnvModeTest        EnvMode = "test"
)
