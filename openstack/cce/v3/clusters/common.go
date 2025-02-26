package clusters

import "encoding/json"

type Clusters struct {
	// API type, fixed value Cluster
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiversion" required:"true"`
	// Metadata of a Cluster
	Metadata MetaData `json:"metadata" required:"true"`
	// specifications of a Cluster
	Spec Spec `json:"spec" required:"true"`
	// status of a Cluster
	Status Status `json:"status"`
}

// Metadata required to create a cluster
type MetaData struct {
	// Cluster unique name
	Name string `json:"name"`
	// Cluster unique Id
	Id string `json:"uid"`
	// Cluster tag, key/value pair format
	Labels map[string]string `json:"labels,omitempty"`
	// Cluster annotation, key/value pair format
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Specifications to create a cluster
type Spec struct {
	// Cluster category: CCE, Turbo
	Category string `json:"category,omitempty"`
	// Cluster Type: VirtualMachine, BareMetal, or Windows
	Type string `json:"type" required:"true"`
	// Cluster specifications
	Flavor string `json:"flavor" required:"true"`
	// Cluster's baseline Kubernetes version. The latest version is recommended
	Version string `json:"version,omitempty"`
	// Cluster description
	Description string `json:"description,omitempty"`
	// Public IP ID
	PublicIP string `json:"publicip_id,omitempty"`
	// Node network parameters
	HostNetwork HostNetworkSpec `json:"hostNetwork" required:"true"`
	// Container network parameters
	ContainerNetwork ContainerNetworkSpec `json:"containerNetwork" required:"true"`
	// ENI network parameters
	EniNetwork *EniNetworkSpec `json:"eniNetwork,omitempty"`
	// Authentication parameters
	Authentication AuthenticationSpec `json:"authentication,omitempty"`
	// Charging mode of the cluster, which is 0 (on demand)
	BillingMode int `json:"billingMode,omitempty"`
	// Extended parameter for a cluster
	ExtendParam map[string]string `json:"extendParam,omitempty"`
	// KubernetesSvcIpRange Service CIDR block or the IP address range which the kubernetes clusterIp must fall within.
	// This parameter is available only for clusters of v1.11.7 and later.
	KubernetesSvcIpRange string `json:"kubernetesSvcIpRange,omitempty"`
	// KubeProxyMode Service forwarding mode. One of `iptables`, `ipvs`
	KubeProxyMode string `json:"kubeProxyMode,omitempty"`
	// The system disks and data disks of the master nodes in the cluster are encrypted.
	EnableMasterVolumeEncryption *bool `json:"enableMasterVolumeEncryption,omitempty"`
}

type Status struct {
	// The state of the cluster
	Phase string `json:"phase"`
	// The ID of the Job that is operating asynchronously in the cluster
	JobID string `json:"jobID"`
	// Reasons for the cluster to become current
	Reason string `json:"reason"`
	// The status of each component in the cluster
	Conditions Conditions `json:"conditions"`
	// Kube-apiserver access address in the cluster
	Endpoints []Endpoints `json:"-"`
}

// Node network parameters
type HostNetworkSpec struct {
	// The ID of the VPC used to create the node
	VpcId string `json:"vpc" required:"true"`
	// The ID of the subnet used to create the node
	SubnetId string `json:"subnet" required:"true"`
	// The ID of the high speed network used to create bare metal nodes.
	// This parameter is required when creating a bare metal cluster.
	HighwaySubnet string `json:"highwaySubnet,omitempty"`
	// The ID of the Security Group used to create the node
	SecurityGroupId string `json:"SecurityGroup,omitempty"`
}

// Container network parameters
type ContainerNetworkSpec struct {
	// Container network type: overlay_l2 , underlay_ipvlan or vpc-router
	Mode string `json:"mode" required:"true"`
	// Container network segment: 172.16.0.0/16 ~ 172.31.0.0/16. If there is a network segment conflict, it will be automatically reselected.
	Cidr string `json:"cidr,omitempty"`
}

type EniNetworkSpec struct {
	// Eni network subnet id
	SubnetId string `json:"eniSubnetId" required:"true"`
	// Eni network cidr
	Cidr string `json:"eniSubnetCIDR" required:"true"`
}

// Authentication parameters
type AuthenticationSpec struct {
	// Authentication mode: rbac , x509 or authenticating_proxy
	Mode                string            `json:"mode" required:"true"`
	AuthenticatingProxy map[string]string `json:"authenticatingProxy" required:"true"`
}

type Conditions struct {
	// The type of component
	Type string `json:"type"`
	// The state of the component
	Status string `json:"status"`
	// The reason that the component becomes current
	Reason string `json:"reason"`
}

type Endpoints struct {
	// The address accessed within the user's subnet - OpenTelekomCloud
	Url string `json:"url"`
	// Public network access address - OpenTelekomCloud
	Type string `json:"type"`
	// Internal network address - OTC
	Internal string `json:"internal"`
	// External network address - OTC
	External string `json:"external"`
	// Endpoint of the cluster to be accessed through API Gateway - OTC
	ExternalOTC string `json:"external_otc"`
}

type Certificate struct {
	// API type, fixed value Config
	Kind string `json:"kind"`
	// API version, fixed value v1
	ApiVersion string `json:"apiVersion"`
	// Cluster list
	Clusters []CertClusters `json:"clusters"`
	// User list
	Users []CertUsers `json:"users"`
	// Context list
	Contexts []CertContexts `json:"contexts"`
	// The current context
	CurrentContext string `json:"current-context"`
}

type CertClusters struct {
	// Cluster name
	Name string `json:"name"`
	// Cluster information
	Cluster CertCluster `json:"cluster"`
}

type CertCluster struct {
	// Server IP address
	Server string `json:"server"`
	// Certificate data
	CertAuthorityData string `json:"certificate-authority-data,omitempty"`
}

type CertUsers struct {
	// User name
	Name string `json:"name"`
	// Cluster information
	User CertUser `json:"user"`
}

type CertUser struct {
	// Client certificate
	ClientCertData string `json:"client-certificate-data"`
	// Client key data
	ClientKeyData string `json:"client-key-data"`
}

type CertContexts struct {
	// Context name
	Name string `json:"name"`
	// Context information
	Context CertContext `json:"context"`
}

type CertContext struct {
	// Cluster name
	Cluster string `json:"cluster"`
	// User name
	User string `json:"user"`
}

// UnmarshalJSON helps to unmarshal Status fields into needed values.
// OTC and Huawei have different data types and child fields for `endpoints` field in Cluster Status.
// This function handles the unmarshal for both
func (r *Status) UnmarshalJSON(b []byte) error {
	type tmp Status
	var s struct {
		tmp
		Endpoints []Endpoints `json:"endpoints"`
	}

	err := json.Unmarshal(b, &s)

	if err != nil {
		switch err.(type) {
		case *json.UnmarshalTypeError: // check if type error occurred (handles the different endpoint structure for huawei and otc)
			var s struct {
				tmp
				Endpoints Endpoints `json:"endpoints"`
			}
			err := json.Unmarshal(b, &s)
			if err != nil {
				return err
			}
			*r = Status(s.tmp)
			r.Endpoints = []Endpoints{{Internal: s.Endpoints.Internal,
				External:    s.Endpoints.External,
				ExternalOTC: s.Endpoints.ExternalOTC}}
			return nil
		default:
			return err
		}
	}

	*r = Status(s.tmp)
	r.Endpoints = s.Endpoints

	return err
}
