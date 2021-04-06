/*
RabbitMQ Messaging Topology Kubernetes Operator
Copyright 2021 VMware, Inc.

This product is licensed to you under the Mozilla Public License 2.0 license (the "License").  You may not use this product except in compliance with the Mozilla 2.0 License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

package internal

import (
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	topology "github.com/rabbitmq/messaging-topology-operator/api/v1alpha2"
)

func GeneratePermissions(p *topology.Permission) rabbithole.Permissions {
	return rabbithole.Permissions{
		Read:      p.Spec.Permissions.Read,
		Write:     p.Spec.Permissions.Write,
		Configure: p.Spec.Permissions.Configure,
	}
}