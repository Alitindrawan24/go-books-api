package constant

type HeaderKey string

const (
	ClientIPAddress   HeaderKey = "Client-IP-Address"
	Client            HeaderKey = "X-Client"
	HostKey           HeaderKey = "X-Host"
	ClientCFIPAddress HeaderKey = "CF-Connecting-IP"
	RequestIDHeader   HeaderKey = "Request-ID"
)
