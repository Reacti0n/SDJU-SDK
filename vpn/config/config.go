package config

// Config config for VPN
type Config struct {
	Identity string `json:"identity"`
	Secret	string `json:"secret"`
}