package natsyaml

import (
	"fmt"


)

type NatsConfig struct {
	Server Serv_struct `yaml:"Server"`
	Logging Log_struct `yaml:"Logging"`
	JetStream Jets_struct `yaml:"JetStream"`
	Authorisation Auth_struct `yaml:"Authorisation"`
	TLS Tls_struct `yaml:"TLS"`
	Cluster Cluster_struct `yaml:"Cluster"`
	Profiling Profil_struct `yaml:"Profilling"`
}

type Serv_struct struct {
  Addr string `yaml:"addr"`                  // set server address (default: 0.0.0.0)
  Net  string `yaml:"net"`                   // Bind to host address 
  Port string `yaml:"port"`                  // port for clients (default: 4222)
  Name string `yaml:"name"`
  Server_name string `yaml:"server_name"`    // Server name (default: auto)
  Pid string `yaml:"pid"`                    // File to store PID
  Http_port string `yaml:"http_port"`        // Use port for http monitoring
  Https_port string `yaml:"https_port"`      // Use port for https monitoring
  Config string `yaml:"config"`              // Configuration file
  Test bool `yaml:"test"`                    // Test configuration and exit
  Signal string `yaml:"signal"`              // <signal>[=<pid>]    Send signal to nats-server process (ldm, stop, quit, term, reopen, reload)
                                             //    <pid> can be either a PID (e.g. 1) or the path to a PID file (e.g. /var/run/nats-server.pid)
  CL_Adv string `yaml:"client_advertise"`    // Client URL to advertise to other servers
  Ports_FilDir string `yaml:"ports_file_dir"` // Creates a ports file in the specified directory (<executable_name>_<pid>.ports).
}

type Log_struct struct {
  LogFil string `yaml:"log"`                 // File to redirect log output
  LogTime bool `yaml:"logtime"`              // Timestamp log entries (default: true)
  LogTimUtc bool `yaml:"logtime_utc"`        // Timestamps in UTC instead of local timezone (default: false)
  Syslog bool `yaml:"syslog"`                // Log to syslog or windows event log
  RemSysLog string `yaml:"remote_syslog"`    // Syslog server addr (udp://localhost:514)
  Debug bool `yaml:"debug"`                  // Enable debugging output
  Trace string `yaml:"trace"`                // Trace the raw protocol
//    -VV                              Verbose trace (traces system account as well)
//    -DV                              Debug and trace
//    -DVV                             Debug and verbose trace (traces system account as well)
  LogMaxSize string `yaml:"log_size_limit"`   // Logfile size limit (default: auto)
  TraceMaxMsg int `yaml:"max_traced_msg_len"` // Maximum printable length for traced messages (default: unlimited)
}

type Jets_struct struct {
  JetStr bool `yaml:"jetstream"`            // Enable JetStream functionality
  StoreDir string `yaml:"store_dir"`        // Set the storage directory
}

type Auth_struct struct {
  User string `yaml:"user"`                  // User required for connections
  Pass string `yaml:"pass"`                  // Password required for connections
  Token string`yaml:"auth"`                  // Authorization token required for connections
}

type Tls_struct struct {
  Tls bool `yaml:"tls"`                      // Enable TLS, do not verify clients (default: false)
  TlsCert string `yaml:"tlscert"`            // Server certificate file
  Tlskey string `yaml:"tlskey"`              // Private key for server certificate
  Tlsverify string `yaml:"tlsverify"`        // Enable TLS, verify client certificates
  TlsCaCert string `yaml:"tlscacert"`        // Client certificate CA for verification
}

type Cluster_struct struct {
  Routes string `yaml:"routes"`              // Routes to solicit and connect <rurl-1, rurl-2> 
  Cluster string `yaml:"cluster"`            // Cluster URL for solicited routes <cluster-url>
  ClusterName string `yaml:"cluster_name"`   // Cluster Name, if not set one will be dynamically generated
  NoAdvert bool `yaml:"no_advertise"`        // Do not advertise known cluster information to clients
  ClusterAdvert string `yaml:"cluster_advertise"`  // Cluster URL to advertise to other servers
  ClRetries int `yaml:"connect_retries"`     // For implicit routes, number of connect retries
  ClListen string `yaml:"cluster_listen"`    // Cluster url from which members can solicit routes
}

type Profil_struct struct {
   Profile int `yaml:"profile"`    // Profiling HTTP port
}

//Common:
//    -h, --help                       Show this message
//    -v, --version                    Show version
//        --help_tls                   TLS help


func RdYaml(yamlFilnam string) (config NatsConfig, err error) {



	return config, nil
}

func PrintNatsConfig (config NatsConfig) {
	fmt.Println("********** Nats Config *************")


	fmt.Println("******** End Nats Config ***********")

}
