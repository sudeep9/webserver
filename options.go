package webserver

type Certificates struct {
	Cert string
	Key  string
}

type ServerOptions struct {
	Addr       string
	Handlers   map[string]Handlers
	StaticDirs map[string]string
	Certs      *Certificates
}

func (o *ServerOptions) populateDefaults() {
	if o.Addr == "" {
		o.Addr = DefaultAddr
	}
}
