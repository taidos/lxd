package drivers

import (
	"github.com/lxc/lxd/lxd/db"
	deviceConfig "github.com/lxc/lxd/lxd/device/config"
	"github.com/lxc/lxd/lxd/instance/instancetype"
	"github.com/lxc/lxd/lxd/state"
	"github.com/lxc/lxd/shared/api"
)

// common provides structure common to all instance types.
type common struct {
	dbType          instancetype.Type
	devPaths        []string
	expandedConfig  map[string]string
	expandedDevices deviceConfig.Devices
	localConfig     map[string]string
	localDevices    deviceConfig.Devices
	profiles        []string
	project         string
	state           *state.State
}

// Project returns instance's project.
func (c *common) Project() string {
	return c.project
}

// Type returns the instance's type.
func (c *common) Type() instancetype.Type {
	return c.dbType
}

// ExpandedConfig returns instance's expanded config.
func (c *common) ExpandedConfig() map[string]string {
	return c.expandedConfig
}

// ExpandedDevices returns instance's expanded device config.
func (c *common) ExpandedDevices() deviceConfig.Devices {
	return c.expandedDevices
}

// LocalConfig returns the instance's local config.
func (c *common) LocalConfig() map[string]string {
	return c.localConfig
}

// LocalDevices returns the instance's local device config.
func (c *common) LocalDevices() deviceConfig.Devices {
	return c.localDevices
}

func (c *common) expandConfig(profiles []api.Profile) error {
	if profiles == nil && len(c.profiles) > 0 {
		var err error
		profiles, err = c.state.Cluster.GetProfiles(c.project, c.profiles)
		if err != nil {
			return err
		}
	}

	c.expandedConfig = db.ExpandInstanceConfig(c.localConfig, profiles)

	return nil
}

func (c *common) expandDevices(profiles []api.Profile) error {
	if profiles == nil && len(c.profiles) > 0 {
		var err error
		profiles, err = c.state.Cluster.GetProfiles(c.project, c.profiles)
		if err != nil {
			return err
		}
	}

	c.expandedDevices = db.ExpandInstanceDevices(c.localDevices, profiles)

	return nil
}

// DevPaths() returns a list of /dev devices which the instance requires access to.
// This is function is only safe to call from within the security
// packages as called during instance startup, the rest of the time this
// will likely return nil.
func (c *common) DevPaths() []string {
	return c.devPaths
}
