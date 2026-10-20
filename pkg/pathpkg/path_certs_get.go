package pathpkg

import "path/filepath"

type PathCertByGlobal struct {
	Global  PathCert
	Current PathCert
}

type PathCert struct {
	Cert string
	Key  string
}

func PathCertsGetComplete(paramPathRoot string, paramEnvMode string, paramServiceName string) PathCertByGlobal {
	// Init Path
	//thisPathRoot := PathGetRoot()

	return PathCertByGlobal{
		Global: PathCert{
			Cert: filepath.Join(paramPathRoot, "certs", "root-cert.pem"),
			Key:  filepath.Join(paramPathRoot, "certs", "root-key.pem"),
		},
		Current: PathCert{
			Cert: filepath.Join(paramPathRoot, "internal", paramServiceName, "v1", "certs", paramServiceName+"-service-cert.pem"),
			Key:  filepath.Join(paramPathRoot, "internal", paramServiceName, "v1", "certs", paramServiceName+"-service-key.pem"),
		},
	}
}

//func PathCertsGetGlobalCert(paramEnvMode string) (string, string) {
//	var a = filepath.Join(paramPathRoot, "certs", "root-cert.pem")
//
//}
