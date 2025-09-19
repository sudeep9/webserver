package webserver

type Certificates struct {
	Cert string
	Key  string
}

type ServerOptions struct {
	Addr       string
	Handlers   map[string]Handlers
	StaticDirs map[string]string
	certs      *Certificates
}

func (o *ServerOptions) populateDefaults() {
	if o.Addr == "" {
		o.Addr = DefaultAddr
	}
}
