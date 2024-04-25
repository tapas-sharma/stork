package pureutils

import (
	"fmt"
	"github.com/devans10/pugo/flasharray"
	"strings"

	"github.com/portworx/torpedo/pkg/units"
)

// GetAppDataDir checks the pod namespace prefix, and returns a path that we can
// write data to on a volume. Because the mountpath varies so heavily between applications
// (some have multiple PVCs, some have configmap vols, etc. etc.) this is the easiest way
// at the moment. As more apps get added to our test suite, we should add them here.
func GetAppDataDir(namespace string) (string, int) {
	if strings.HasPrefix(namespace, "nginx-without-enc") {
		return "/usr/share/nginx/html", units.GiB / 2
	}
	if strings.HasPrefix(namespace, "wordpress") {
		return "/var/www/html", units.GiB / 2
	}
	if strings.HasPrefix(namespace, "elasticsearch") {
		return "/usr/share/elasticsearch/data", units.GiB * 2
	}
	if strings.HasPrefix(namespace, "mysql-without-enc") {
		return "/var/lib/mysql", units.GiB
	}
	if strings.HasPrefix(namespace, "nginx-fa-davol") {
		return "/data", units.GiB
	}
	if strings.HasPrefix(namespace, "nginx-fa-darawvol") {
		return "/dev/xvda", units.GiB
	}
	return "", 0
}

// PureCreateClientAndConnect Create FA Client and Connect
func PureCreateClientAndConnect(faMgmtEndpoint string, apiToken string) (*flasharray.Client, error) {
	faClient, err := flasharray.NewClient(faMgmtEndpoint, "", "", apiToken,
		"", false, false, "", nil)
	if err != nil {
		return nil, err
	}
	return faClient, nil
}

// GetFAClientMapFromPXPureSecret takes a PXPureSecret and returns a map of mgmt endpoints to FA clients
func GetFAClientMapFromPXPureSecret(secret PXPureSecret) (map[string]*flasharray.Client, error) {
	clientMap := make(map[string]*flasharray.Client)
	for _, fa := range secret.Arrays {
		faClient, err := PureCreateClientAndConnect(fa.MgmtEndPoint, fa.APIToken)
		if err != nil {
			return nil, fmt.Errorf("failed to create FA client for [%s]. Err: [%v]", fa.MgmtEndPoint, err)
		}
		clientMap[fa.MgmtEndPoint] = faClient
	}
	return clientMap, nil
}

// GetFAMgmtEndPoints , Get Lists of all management Endpoints from FA Secrets
func GetFAMgmtEndPoints(secret PXPureSecret) []string {
	mgmtEndpoints := []string{}
	for _, faDetails := range secret.Arrays {
		mgmtEndpoints = append(mgmtEndpoints, faDetails.MgmtEndPoint)
	}
	return mgmtEndpoints
}

// GetApiTokenForMgmtEndpoints Returns API token for Mgmt Endpoints
func GetApiTokenForMgmtEndpoints(secret PXPureSecret, mgmtEndPoint string) string {
	for _, faDetails := range secret.Arrays {
		if faDetails.MgmtEndPoint == mgmtEndPoint {
			return faDetails.APIToken
		}
	}
	return ""
}

// CreateVolumeOnFABackend Creates Volume on FA Backend
func CreateVolumeOnFABackend(faClient *flasharray.Client, volName string, size int) (*flasharray.Volume, error) {
	volume, err := faClient.Volumes.CreateVolume(volName, size)
	if err != nil {
		return nil, err
	}
	return volume, nil
}

// ListAllTheVolumesFromSpecificFA returns list of all the volumes from the FA Specified
func ListAllTheVolumesFromSpecificFA(faClient *flasharray.Client) ([]flasharray.Volume, error) {
	volumes, err := faClient.Volumes.ListVolumes(nil)
	if err != nil {
		return nil, err
	}
	return volumes, nil
}

// GetAllHostGroups Get all Available Host Groups from array
func GetAllHostGroups(faClient *flasharray.Client) ([]flasharray.Hostgroup, error) {
	hostGroup, err := faClient.Hostgroups.ListHostgroups(nil)
	if err != nil {
		return nil, err
	}
	return hostGroup, nil
}

// ListAllHostGroupConnections Returns list of all host group connections from array
func ListAllHostGroupConnections(faClient *flasharray.Client, hostGroupName string) ([]flasharray.HostgroupConnection, error) {
	hostGroupConnections := []flasharray.HostgroupConnection{}
	hostGroupConnect, err := faClient.Hostgroups.ListHostgroupConnections(hostGroupName)
	if err != nil {
		return nil, err
	}
	for _, eachHostConnection := range hostGroupConnect {
		hostGroupConnections = append(hostGroupConnections, eachHostConnection)
	}
	return hostGroupConnections, nil
}

// ListAllHosts retuns list of hosts present
func ListAllHosts(faClient *flasharray.Client) ([]flasharray.Host, error) {
	hostGroup, err := faClient.Hosts.ListHosts(nil)
	if err != nil {
		return nil, err
	}
	return hostGroup, nil
}

// CreateNewHostOnFA Creates New Host on FA
func CreateNewHostOnFA(faClient *flasharray.Client, hostName string) (*flasharray.Host, error) {
	host, err := faClient.Hosts.CreateHost(hostName, nil)
	if err != nil {
		return nil, err
	}
	return host, nil
}

// ConnectVolumeToHost Connects Volume to Host
func ConnectVolumeToHost(faClient *flasharray.Client, hostName string, volName string) (*flasharray.ConnectedVolume, error) {
	connectedVol, err := faClient.Hosts.ConnectHost(hostName, volName, nil)
	if err != nil {
		return nil, err
	}
	return connectedVol, nil
}

// UpdateIQNOnSpecificHosts Updates IQN on specific hosts
func UpdateIQNOnSpecificHosts(faClient *flasharray.Client, hostName string, iqnValue string) (*flasharray.Host, error) {
	data1 := make(map[string][]string)
	data1["iqnlist"] = []string{iqnValue}
	host, err := faClient.Hosts.SetHost(hostName, data1)
	if err != nil {
		return nil, err
	}
	return host, nil
}

// GetIqnFromHosts returns list of IQNs associated with the hosts
func GetIqnFromHosts(faClient *flasharray.Client, hostName string) ([]string, error) {
	hosts, err := ListAllHosts(faClient)
	if err != nil {
		return nil, err
	}
	for _, eachHosts := range hosts {
		if eachHosts.Name == hostName {
			return eachHosts.Iqn, nil
		}
	}
	return nil, fmt.Errorf("Unable to fetch iqn details for the host specified [%v]", hostName)
}

// IsIQNExistsOnFA returns True if IQN is already associated to some Node
func IsIQNExistsOnFA(faClient *flasharray.Client, iqnName string) (bool, error) {
	hosts, err := ListAllHosts(faClient)
	if err != nil {
		return false, err
	}
	for _, eachHosts := range hosts {
		for _, eachIqn := range eachHosts.Iqn {
			if eachIqn == iqnName {
				return true, nil
			}
		}
	}
	return false, nil
}

// ListAllNetworkInterfacesOnFA returns all the list of Network Interfaces present
func ListAllNetworkInterfacesOnFA(faClient *flasharray.Client) ([]flasharray.NetworkInterface, error) {
	networkInterface, err := faClient.Networks.ListNetworkInterfaces()
	if err != nil {
		return nil, err
	}
	return networkInterface, nil
}

// FilterSpecificInterfaceBasedOnServiceType returns list of all mgmt interfaces from specific flash array
// interfaceType type can be "manangement", "iscsi", "replication"
func GetSpecificInterfaceBasedOnServiceType(faClient *flasharray.Client, interfaceType string) ([]flasharray.NetworkInterface, error) {
	interfaces := []flasharray.NetworkInterface{}
	allInterface, err := ListAllNetworkInterfacesOnFA(faClient)
	if err != nil {
		return nil, err
	}
	for _, eachInterface := range allInterface {
		for _, eachService := range eachInterface.Services {
			if eachService == interfaceType {
				interfaces = append(interfaces, eachInterface)
			}
		}
	}
	return interfaces, nil
}

// GetNetworkInterfaceDetails Returns list of all network interface details
func GetNetworkInterfaceDetails(faClient *flasharray.Client, iface string) (*flasharray.NetworkInterface, error) {
	networkInterface, err := faClient.Networks.GetNetworkInterface(iface)
	if err != nil {
		return nil, err
	}
	return networkInterface, nil
}

// IsNetworkInterfaceEnabled returns true if network interface is enabled else false
func IsNetworkInterfaceEnabled(faClient *flasharray.Client, iface string) (bool, error) {
	networkInterface, err := GetNetworkInterfaceDetails(faClient, iface)
	if err != nil {
		return false, err
	}
	if networkInterface.Enabled {
		return true, nil
	}
	return false, nil
}

// EnableNetworkInterface enables network interface
func EnableNetworkInterface(faClient *flasharray.Client, iface string) (bool, error) {
	interfaces, err := faClient.Networks.EnableNetworkInterface(iface)
	if err != nil {
		return false, err
	}
	if interfaces.Enabled {
		return true, nil
	}
	return false, fmt.Errorf("Failed to enable network interface [%v]", iface)
}

// DisableNetworkInterface disabled network interface
func DisableNetworkInterface(faClient *flasharray.Client, iface string) (bool, error) {
	interfaces, err := faClient.Networks.DisableNetworkInterface(iface)
	if err != nil {
		return false, err
	}
	if !interfaces.Enabled {
		return true, nil
	}
	return false, fmt.Errorf("Failed to disable network interface [%v]", iface)
}
