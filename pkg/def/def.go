package def

const (
	LogServer     = "server"     // "OpenAPI", "gRPC", "Prometheus metrics", etc.
	LogRemote     = "remote"     // Aligned IPv4:Port "   192.168.0.42:1234 ".
	LogHost       = "host"       // DNS hostname or IPv4/IPv6 address.
	LogPort       = "port"       // TCP/UDP port number.
	LogAddr       = "addr"       // host:port.
	LogHTTPMethod = "httpMethod" // GET, POST, etc.
	LogHTTPStatus = "httpStatus" // Status code: 200, 404, etc.
	LogFunc       = "func"       // RPC/event handler method name, REST resource path.
	LogUserID     = "userID"
	LogHandler    = "handler"
)
