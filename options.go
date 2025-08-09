package webserver

type ServerOptions struct {
	Addr       string
	Handlers   map[string]Handlers
	StaticDirs map[string]string
}

func (o *ServerOptions) populateDefaults() {
	if o.Addr == "" {
		o.Addr = DefaultAddr
	}
}
