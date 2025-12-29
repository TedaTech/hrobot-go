package client

import "github.com/TedaTech/hrobot-go/models"

type RobotClient interface {
	SetBaseURL(baseURL string)
	SetUserAgent(userAgent string)
	GetVersion() string
	SetCredentials(username, password string) error
	ValidateCredentials() error

	ServerGetList() ([]models.Server, error)
	ServerGet(id int) (*models.Server, error)
	ServerSetName(id int, input *models.ServerSetNameInput) (*models.Server, error)
	ServerReverse(id int) (*models.Cancellation, error)
	KeyGetList() ([]models.Key, error)
	KeyGet(fingerprint string) (*models.Key, error)
	KeySet(input *models.KeySetInput) (*models.Key, error)
	KeyUpdateName(fingerprint, name string) (*models.Key, error)
	KeyDelete(fingerprint string) error
	IPGetList() ([]models.IP, error)
	RDnsGetList() ([]models.Rdns, error)
	RDnsGet(ip string) (*models.Rdns, error)
	BootLinuxGet(id int) (*models.Linux, error)
	BootLinuxDelete(id int) (*models.Linux, error)
	BootLinuxSet(id int, input *models.LinuxSetInput) (*models.Linux, error)
	BootRescueGet(id int) (*models.Rescue, error)
	BootRescueDelete(id int) (*models.Rescue, error)
	BootRescueSet(id int, input *models.RescueSetInput) (*models.Rescue, error)
	ResetGet(id int) (*models.Reset, error)
	ResetSet(id int, input *models.ResetSetInput) (*models.ResetPost, error)
	FailoverGetList() ([]models.Failover, error)
	FailoverGet(ip string) (*models.Failover, error)
	VSwitchGetList() ([]models.VSwitch, error)
	VSwitchGet(id int) (*models.VSwitch, error)
	VSwitchCreate(input *models.VSwitchCreateInput) (*models.VSwitch, error)
	VSwitchAddServer(vswitchID, serverNumber int) error
	VSwitchRemoveServer(vswitchID, serverNumber int) error
}
