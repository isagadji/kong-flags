package flags

type HttpServerFlags struct {
	HTTPAddr        string `kong:"required,name=http-addr,env=HTTP_ADDR,group='HTTP server'"`
	HTTPMetricsAddr string `kong:"required,name=http-metrics-addr,env=HTTP_METRICS_ADDR,group='HTTP metrics server'"`
}
