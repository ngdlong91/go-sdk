package gosdk

type Service interface {
}

type RunnableService interface {
}

type Application interface {
	RegService(svc Service) error
	// mark a background servie as critical
	// so if Run() method is exit -> app will be stop
	SetCriticalService(svc RunnableService)
	RegMainService(svc RunnableService) error
	OutputEnv()
	Run()
	// reload all services config, log, ... Auto trigger by HUP signal
	Reload()
	Shutdown() <-chan struct{}
	IsShutdown() bool
	GetMainService() RunnableService
	GetServices() []Service
	// Register function to run before exit (include normal exit & fatal)
	RegisterExitHandler(func())
}
