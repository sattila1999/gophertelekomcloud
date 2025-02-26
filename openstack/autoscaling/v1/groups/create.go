package groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the AS group name. The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	Name string `json:"scaling_group_name" required:"true"`
	// Specifies the AS configuration ID, which can be obtained using the API for querying AS configurations.
	ConfigurationID string `json:"scaling_configuration_id" required:"true"`
	// Specifies the expected number of instances. The default value is the minimum number of instances.
	// The value ranges from the minimum number of instances to the maximum number of instances.
	DesireInstanceNumber int `json:"desire_instance_number,omitempty"`
	// Specifies the minimum number of instances. The default value is 0.
	MinInstanceNumber int `json:"min_instance_number,omitempty"`
	// Specifies the maximum number of instances. The default value is 0.
	MaxInstanceNumber int `json:"max_instance_number,omitempty"`
	// Specifies the cooldown period (in seconds). The value ranges from 0 to 86400 and is 300 by default.
	// After a scaling action is triggered, the system starts the cooldown period. During the cooldown period,
	// scaling actions triggered by alarms will be denied. Scheduled, periodic, and manual scaling actions are not affected..
	CoolDownTime int `json:"cool_down_time,omitempty"`
	// Specifies the ID of a classic load balancer listener. The system supports the binding of up to six
	// load balancer listeners, the IDs of which are separated using a comma (,).
	// This parameter is alternative to lbaas_listeners.
	LBListenerID string `json:"lb_listener_id,omitempty"`
	// Specifies information about an enhanced load balancer. The system supports the binding of up to six load balancers.
	// This parameter is in list data structure.This parameter is alternative to lb_listener_id.
	LBaaSListeners []LBaaSListener `json:"lbaas_listeners,omitempty"`
	// Specifies the AZ information. The instances added in a scaling action will be created in a specified AZ.
	// If you do not specify an AZ, the system automatically specifies one.
	AvailableZones []string `json:"available_zones,omitempty"`
	// Specifies network information. The system supports up to five subnets. The first subnet transferred
	// serves as the primary NIC of the ECS by default. This parameter is in data structure.
	Networks []ID `json:"networks" required:"true"`
	// Specifies the security group information. A maximum of one security group can be selected. This parameter is in data structure.
	// If the security group is specified both in the AS configuration and AS group,
	// scaled ECS instances will be added to the security group specified in the AS configuration.
	// If the security group is not specified in either of them, scaled ECS instances will be added to the default security group.
	// For your convenience, you are advised to specify the security group in the AS configuration.
	SecurityGroup []ID `json:"security_groups,omitempty"`
	// Specifies the VPC ID, which can be obtained using the API for querying VPCs. For details,
	// see "Querying VPCs" in Virtual Private Network API Reference.
	VpcID string `json:"vpc_id" required:"true"`
	// Specifies the health check method for instances in the AS group. The health check methods include ELB_AUDIT and NOVA_AUDIT.
	// When load balancing is configured for an AS group, the default value is ELB_AUDIT. Otherwise, the default value is NOVA_AUDIT.
	// ELB_AUDIT: indicates the ELB health check, which takes effect in an AS group with a listener.
	// NOVA_AUDIT: indicates the ECS instance health check, which is the health check method delivered with AS.
	HealthPeriodicAuditMethod string `json:"health_periodic_audit_method,omitempty"`
	// Specifies the instance health check period. The value can be 1, 5, 15, 60, or 180 in the unit of minutes.
	// If this parameter is not specified, the default value is 5.
	// If the value is set to 0, health check is performed every 10 seconds.
	HealthPeriodicAuditTime int `json:"health_periodic_audit_time,omitempty"`
	// Specifies the grace period for instance health check. The unit is second and value range is 0-86400. The default value is 600.
	// The health check grace period starts after an instance is added to an AS group and is enabled.
	// The AS group will start checking the instance status only after the grace period ends.
	// This parameter is valid only when the instance health check method of the AS group is ELB_AUDIT.
	HealthPeriodicAuditGrace int `json:"health_periodic_audit_grace_period,omitempty"`
	// Specifies the instance removal policy.
	// OLD_CONFIG_OLD_INSTANCE (default): The earlier-created instances based on the earlier-created AS configurations are removed first.
	// OLD_CONFIG_NEW_INSTANCE: The later-created instances based on the earlier-created AS configurations are removed first.
	// OLD_INSTANCE: The earlier-created instances are removed first.
	// NEW_INSTANCE: The later-created instances are removed first.
	InstanceTerminatePolicy string `json:"instance_terminate_policy,omitempty"`
	// Specifies the notification mode.
	// EMAIL refers to notification by email.
	Notifications []string `json:"notifications,omitempty"`
	// Specifies whether to delete the EIP bound to the ECS when deleting the ECS.
	// The value can be true or false. The default value is false.
	// true: deletes the EIP bound to the ECS when deleting the ECS.
	// false: only unbinds the EIP bound to the ECS when deleting the ECS.
	IsDeletePublicip *bool `json:"delete_publicip,omitempty"`
	// Specifies whether to delete the data disks attached to the ECS when deleting the ECS.
	// The value can be true or false. The default value is false.
	// true: deletes the data disks attached to the ECS when deleting the ECS.
	// false: only detaches the data disks attached to the ECS when deleting the ECS.
	IsDeleteVolume *bool `json:"delete_volume,omitempty"`
	// Specifies the enterprise project ID, which is used to specify the enterprise project to which the AS group belongs.
	// If the value is 0 or left blank, the AS group belongs to the default enterprise project.
	// If the value is a UUID, the AS group belongs to the enterprise project corresponding to the UUID.
	// If an enterprise project is configured for an AS group, ECSs created in this AS group also belong to this enterprise project.
	// Otherwise, the default enterprise project will be used.
	EnterpriseProjectID string `json:"enterprise_project_id,omitempty"`
	// Specifies the priority policy used to select target AZs when adjusting the number of instances in an AS group.
	// EQUILIBRIUM_DISTRIBUTE (default): When adjusting the number of instances,
	// ensure that instances in each AZ in the available_zones list is evenly distributed.
	// If instances cannot be added in the target AZ, select another AZ based on the PICK_FIRST policy.
	// PICK_FIRST: When adjusting the number of instances, target AZs are determined in the order in the available_zones list.
	MultiAZPriorityPolicy string `json:"multi_az_priority_policy,omitempty"`
	// Specifies the description of the AS group. The value can contain 1 to 256 characters.
	Description string `json:"description,omitempty"`
}

type LBaaSListener struct {
	ListenerID string `json:"listener_id"`
	// Specifies the backend ECS group ID.
	PoolID string `json:"pool_id" required:"true"`
	// Specifies the backend protocol ID, which is the port on which a backend ECS listens for traffic. The port ID ranges from 1 to 65535.
	ProtocolPort int `json:"protocol_port" required:"true"`
	// Specifies the weight, which determines the portion of requests a backend ECS processes
	// when being compared to other backend ECSs added to the same listener. The value of this parameter ranges from 0 to 100.
	Weight int `json:"weight" required:"true"`
}

type ID struct {
	ID string `json:"id" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("scaling_group"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		ID string `json:"scaling_group_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
