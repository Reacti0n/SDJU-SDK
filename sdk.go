package SDJU

import (
	"github.com/Reacti0n/SDJU-SDK/vpn"
	vpnConfig "github.com/Reacti0n/SDJU-SDK/vpn/config"
	educationConfig "github.com/Reacti0n/SDJU-SDK/education/config"
	"github.com/Reacti0n/SDJU-SDK/education"
	log "github.com/sirupsen/logrus"
	"os"
)

func init(){
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// SDJUSDK struct
type SDJUSDK struct {

}

// NewSDJUSDK init
func NewSDJUSDK()*SDJUSDK  {
	return &SDJUSDK{}
}

// GetVPN 获取 VPN 实例
func (ss *SDJUSDK)GetVPN(cfg *vpnConfig.Config) *vpn.VPN  {
	return vpn.NewVPN(cfg)
}


// GetEducation 获取教务实例
func (ss *SDJUSDK)GetEducation(cfg *educationConfig.Config) *education.Education  {
	return education.NewEducation(cfg)
}

