/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package compute

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/utils/openstack/clientconfig"

	"sigs.k8s.io/cluster-api-provider-openstack/pkg/cloud/services/networking"
	"sigs.k8s.io/cluster-api-provider-openstack/pkg/cloud/services/provider"
)

type Service struct {
	projectID         string
	computeService    Client
	networkingService *networking.Service
	logger            logr.Logger
}

/*
 NovaMinimumMicroversion is the minimum Nova microversion supported by CAPO
 2.53 corresponds to OpenStack Pike

 For the canonical description of Nova microversions, see
 https://docs.openstack.org/nova/latest/reference/api-microversion-history.html

 CAPO uses server tags, which were added in microversion 2.52.
*/
const NovaMinimumMicroversion = "2.53"

// NewService returns an instance of the compute service.
func NewService(client *gophercloud.ProviderClient, clientOpts *clientconfig.ClientOpts, logger logr.Logger) (*Service, error) {
	computeClient, err := openstack.NewComputeV2(client, gophercloud.EndpointOpts{
		Region: clientOpts.RegionName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create compute service client: %v", err)
	}
	computeClient.Microversion = NovaMinimumMicroversion

	imagesClient, err := openstack.NewImageServiceV2(client, gophercloud.EndpointOpts{
		Region: clientOpts.RegionName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create image service client: %v", err)
	}

	computeService := serviceClient{computeClient, imagesClient}

	if clientOpts.AuthInfo == nil {
		return nil, fmt.Errorf("authInfo must be set")
	}

	projectID, err := provider.GetProjectID(client, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("error retrieveing project id: %v", err)
	}

	networkingService, err := networking.NewService(client, clientOpts, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create networking service: %v", err)
	}

	return &Service{
		projectID:         projectID,
		computeService:    computeService,
		networkingService: networkingService,
		logger:            logger,
	}, nil
}
