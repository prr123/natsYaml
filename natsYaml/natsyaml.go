package natsyaml

import (
	"fmt"
	"os"
	"reflect"
//	"encoding/json/v2"
	"strconv"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/goccy/go-yaml"
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
  Ports_Dir string `yaml:"ports_file_dir"`   // Creates a ports file in the specified directory (<executable_name>_<pid>.ports).
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
  TlsKey string `yaml:"tlskey"`              // Private key for server certificate
  TlsVerify string `yaml:"tlsverify"`        // Enable TLS, verify client certificates
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


func RdYaml(filnam string) (config NatsConfig, err error) {

    yamlData, err := os.ReadFile(filnam)
    if err != nil {return config, fmt.Errorf("cannot read yaml file [%s]: %v", filnam, err)}

    err = yaml.Unmarshal([]byte(yamlData), &config)
    if err != nil {return config, fmt.Errorf("cannot unmarshal yamldata: %v", err)}

	return config, nil
}

func RdYamlOpt(filnam string) (opt server.Options, err error) {

	var config NatsConfig
    yamlData, err := os.ReadFile(filnam)
    if err != nil {return opt, fmt.Errorf("cannot read yaml file [%s]: %v", filnam, err)}

    err = yaml.Unmarshal([]byte(yamlData), &config)
    if err != nil {return opt, fmt.Errorf("cannot unmarshal yamldata: %v", err)}

	// config file
	// todo

	// server
	scfg := config.Server
	if len(scfg.Server_name) > 0 {
		opt.ServerName = scfg.Server_name
	} else {
		if len(scfg.Name) > 0 {
			opt.ServerName = scfg.Name
		} else {
			opt.ServerName = "NatsServ"
		}
	}

	// server port
	port := scfg.Port
	portNum, err := strconv.Atoi(port)
	if err != nil {return opt, fmt.Errorf("yaml file: port number is not an integer")}
	if portNum>65000 {return opt, fmt.Errorf("yaml file: port number is too high")}
	if portNum<1000 {return opt, fmt.Errorf("yaml file: port number is too low (<1000)")}
	opt.Port = portNum

	// server address
	if len(scfg.Addr) > 0 {
		opt.Host = scfg.Addr
	} else {
		if len(scfg.Net) > 0 {
			opt.Host = scfg.Net
		} else {
			opt.Host = "0.0.0.0"
		}
	}

	// DontListen
	opt.DontListen = false

	// pid file
	opt.PidFile = scfg.Pid

	//

	opt.HTTPPort, err = strconv.Atoi(scfg.Http_port)
	if err != nil {return opt, fmt.Errorf("yaml file: http port number is not an integer")}

	opt.HTTPSPort, err = strconv.Atoi(scfg.Https_port)
	if err != nil {return opt, fmt.Errorf("yaml file: https port number is not an integer")}

	opt.PortsFileDir = scfg.Ports_Dir

	//client advertise (client url)
	opt.ClientAdvertise = scfg.CL_Adv


	return opt, nil
}


func PrintNatsConfig (config NatsConfig) {
	fmt.Println("********** Nats Config *************")
	scfg := config.Server
	fmt.Println("  Server:")
	fmt.Printf("        Addr:              %s\n", scfg.Addr) // server address
	fmt.Printf("        Net:               %s\n", scfg.Net)
	fmt.Printf("        Port:              %s\n", scfg.Port)
	fmt.Printf("        Name:              %s\n", scfg.Name) // same as Server_name
	fmt.Printf("        Server Name:       %s\n", scfg.Server_name)
	fmt.Printf("        Pid:               %s\n", scfg.Pid) // File where PID is stored
	fmt.Printf("        Http_port:         %s\n", scfg.Http_port)
	fmt.Printf("        Https_port:        %s\n", scfg.Https_port)
	fmt.Printf("        Config:            %s\n", scfg.Config)  // Configuration file
	fmt.Printf("        Test:              %t\n", scfg.Test)
	fmt.Printf("        Signal:            %s\n", scfg.Signal)  // <signal>[=<pid>]    Send signal to nats-server process (ldm, stop, quit, term, reopen, reload)
	fmt.Printf("        Client_Adv:        %s\n", scfg.CL_Adv)
	fmt.Printf("        Ports_Dir:         %s\n", scfg.Ports_Dir)    // Creates a ports file in the specified directory (<executable_name>_<pid>.ports).


	lcfg := config.Logging
	fmt.Println("\n  Logging:")
	fmt.Printf("        Log File:          %s\n", lcfg.LogFil)    // File to redirect log output
	fmt.Printf("        Log Time Stamp:    %t\n", lcfg.LogTime)   // Timestamp log entries (default: true)
	fmt.Printf("        UTC Log Time:      %t\n", lcfg.LogTimUtc) // use UTC in time stamps
	fmt.Printf("        Syslog:            %t\n", lcfg.Syslog)    // Log to syslog or windows event log
	fmt.Printf("        Remote Adr:        %s\n", lcfg.RemSysLog) // Syslog server addr (udp://localhost:514)
	fmt.Printf("        Debug:             %t\n", lcfg.Debug)
	fmt.Printf("        Trace:             %s\n", lcfg.Trace)    //
				                                              //    -VV  Verbose trace (traces system account as well)
                                                              //    -DV  Debug and trace
                                                              //    -DVV  Debug and verbose trace (traces system account as well)
	fmt.Printf("        Max Log Size:      %s\n", lcfg.LogMaxSize)  // should be a number but default is 'auto'
	fmt.Printf("        Max Trace Msg Len: %d\n", lcfg.TraceMaxMsg) // Maximum printable length for traced messages (default: unlimited)


	jcfg := config.JetStream
	fmt.Println("\n  JetStream:")
	fmt.Printf("        JetStream:         %t\n", jcfg.JetStr)   // Enables JetStream functionality
	fmt.Printf("        Store Dir:         %s\n", jcfg.StoreDir)

	acfg := config.Authorisation
	fmt.Println("\n  Authorisation:")
	fmt.Printf("        User:              %s\n", acfg.User)
	fmt.Printf("        Passwd:            %s\n", acfg.Pass)
	fmt.Printf("        Token:             %s\n", acfg.Token)

	tcfg := config.TLS
	fmt.Println("\n  TLS:")
	fmt.Printf("        TLS:               %t\n", tcfg.Tls)  // true, if tls in use
	fmt.Printf("        TLS Cert:          %s\n", tcfg.TlsCert)
	fmt.Printf("        TLS Key:           %s\n", tcfg.TlsKey)
	fmt.Printf("        TLS Verify:        %s\n", tcfg.TlsVerify) // Enable TLS, verify client certificates
	fmt.Printf("        TLS Ca Cert:       %s\n", tcfg.TlsCaCert) // Certificate Authority Cert

	clcfg := config.Cluster
	fmt.Println("\n  Cluster:")
	fmt.Printf("        Routes:            %s\n", clcfg.Routes)
	fmt.Printf("        Name:              %s\n", clcfg.ClusterName)
	fmt.Printf("        NoAdvert:          %t\n", clcfg.NoAdvert)
	fmt.Printf("        ClAdvert:          %s\n", clcfg.ClusterAdvert)
	fmt.Printf("        ClRetries:         %d\n", clcfg.ClRetries)
	fmt.Printf("        ClListen:          %s\n", clcfg.ClListen)

	pcfg := config.Profiling
	fmt.Println("\n  Profiling:")
	fmt.Printf("        Profile:           %d\n", pcfg.Profile)


	fmt.Println("******** End Nats Config ***********")

}

func PrintNatsServerOpt (opt server.Options)(err error) {
	fmt.Println("********** Nats Server Options *************")
	val := reflect.ValueOf(opt)

	// 2. Loop through the number of fields
	for i := 0; i < val.NumField(); i++ {
    	field := val.Field(i)
    	fieldName := val.Type().Field(i).Name
    	// Safely verify if you are allowed to read the interface value
    	if field.CanInterface() {
        	fieldValue := field.Interface()
			fmt.Printf("%-30s: %v\n", fieldName, fieldValue)
    	} else {
			fmt.Printf("%-30s: no field velu [%d]\n", fieldName, i)
    	}
	}
	fmt.Println("******** End Nats Server Options ***********")

	return nil
}
