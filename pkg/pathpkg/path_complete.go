package pathpkg

type PathService struct {
	Root      string
	Service   string
	Env       string
	GlobalEnv string
}

func PathServiceGetComplete(paramPathRoot string, paramEnvMode string, paramServiceName string) PathService {
	// Init Path
	//thisPathRoot := PathGetRoot()

	thisPathService := PathCreateService(paramPathRoot, paramServiceName)
	thisPathEnv := PathCreateEnv(paramEnvMode, thisPathService)
	thisPathGlobalEnv := PathCreateGlobalEnv(paramPathRoot, paramEnvMode)

	return PathService{
		Root:      paramPathRoot,
		Service:   thisPathService,
		Env:       thisPathEnv,
		GlobalEnv: thisPathGlobalEnv,
	}
}
