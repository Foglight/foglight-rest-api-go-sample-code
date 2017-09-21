package rest

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

//Config config items
type Config struct {
	ServerURL                 string
	AuthToken                 string
	CollectionIntervalMinutes int
	AgentName                 string
}

//NewConfig Parse arguments and create Config
func NewConfig() *Config {
	url := flag.String("server", "http://localhost:8080", "Fogligt server base url.")
	token := flag.String("token", "", "Fogligt REST API token.")
	interval := flag.Int("interval", 5, "Agent data collection interval in minutes, should in 1 to 30 minutes.")
	agentName := flag.String("agent", fmt.Sprintf("%s(pid:%d)", Fqdn(), os.Getpid()), "Agent name.")
	flag.Parse()
	config := &Config{*url, *token, *interval, *agentName}
	if !config.validate() {
		Log("Missing required parameters.\n")
		flag.Usage()
		os.Exit(1)
	}
	Log("Configuration:\n", fmt.Sprintf("ServerURL: %s\nAuth Token: %s\n Interval: %d Minutes\n Agent Name: %s", config.ServerURL, config.AuthToken, config.CollectionIntervalMinutes, config.AgentName))
	return config
}

func (c *Config) validate() bool {
	if c.AuthToken == "" {
		Log("AuthToken is requird.")
		return false
	}
	if c.CollectionIntervalMinutes <= 0 || c.CollectionIntervalMinutes > 30 {
		Log("Collection interval is not valid: ", c.CollectionIntervalMinutes)
		return false
	}
	return true
}

func Fqdn() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknownHost"
	}

	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return hostname
	}

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip, err := ipv4.MarshalText()
			if err != nil {
				return hostname
			}
			hosts, err := net.LookupAddr(string(ip))
			if err != nil || len(hosts) == 0 {
				return hostname
			}
			fqdn := hosts[0]
			return strings.TrimSuffix(fqdn, ".")
		}
	}
	return hostname
}
