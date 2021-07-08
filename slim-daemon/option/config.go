package option

/*
const (
	// ciliumEnvPrefix is the prefix used for environment variables
	ciliumEnvPrefix = "CILIUM_"

	// K8sKubeConfigPath is the absolute path of the kubernetes kubeconfig file
	K8sKubeConfigPath = "k8s-kubeconfig-path"

	// K8sAPIServer is the kubernetes api address server (for https use --k8s-kubeconfig-path instead)
	K8sAPIServer = "k8s-api-server"

	// IdentityChangeGracePeriod is the name of the
	// IdentityChangeGracePeriod option
	IdentityChangeGracePeriod = "identity-change-grace-period"
)

type DaemonConfig struct {
	CreateTime        time.Time
	K8sKubeConfigPath string
	K8sAPIServer      string

	// IdentityChangeGracePeriod is the grace period that needs to pass
	// before an endpoint that has changed its identity will start using
	// that new identity. During the grace period, the new identity has
	// already been allocated and other nodes in the cluster have a chance
	// to whitelist the new upcoming identity of the endpoint.
	IdentityChangeGracePeriod time.Duration
}

var (
	Config = &DaemonConfig{
		CreateTime:                time.Now(),
		IdentityChangeGracePeriod: defaults.IdentityChangeGracePeriod,
	}
)

func (c *DaemonConfig) Populate() {
	c.K8sKubeConfigPath = viper.GetString(K8sKubeConfigPath)
	c.K8sAPIServer = viper.GetString(K8sAPIServer)

}

// K8sAPIDiscoveryEnabled returns true if API discovery of API groups and
// resources is enabled
func (c *DaemonConfig) K8sAPIDiscoveryEnabled() bool {
	return false
}

// K8sLeasesFallbackDiscoveryEnabled returns true if we should fallback to direct API
// probing when checking for support of Leases in case Discovery API fails to discover
// required groups.
func (c *DaemonConfig) K8sLeasesFallbackDiscoveryEnabled() bool {
	return false
}

// BindEnv binds the option name with an deterministic generated environment
// variable which s based on the given optName. If the same optName is bind
// more than 1 time, this function panics.
func BindEnv(optName string) {
	registerOpt(optName)
	viper.BindEnv(optName, getEnvName(optName))
}

// RegisteredOptions maps all options that are bind to viper.
var RegisteredOptions = map[string]struct{}{}

func registerOpt(optName string) {
	_, ok := RegisteredOptions[optName]
	if ok || optName == "" {
		panic(fmt.Errorf("option already registered: %s", optName))
	}
	RegisteredOptions[optName] = struct{}{}
}

// getEnvName returns the environment variable to be used for the given option name.
func getEnvName(option string) string {
	under := strings.Replace(option, "-", "_", -1)
	upper := strings.ToUpper(under)
	return ciliumEnvPrefix + upper
}

*/
