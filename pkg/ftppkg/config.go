package ftppkg

import "time"

type Config struct {
	Timeout time.Duration `env:"FTP_TIMEOUT"`
	//Host     string `env:"FTP_HOST"`
	//Username string `env:"FTP_USERNAME"`
	//Password string `env:"FTP_PASSWORD"`
	//Port     string `env:"FTP_PORT"`
}

// host := "194.147.142.197"
// username := "pz15176"
// password := "MCk6BApm"
// port := "21"
// dial := fmt.Sprintf("%s:%s", host, port)
