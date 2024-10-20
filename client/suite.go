// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	cwclient "github.com/cloudwego-contrib/cwgo-pkg/config/zookeeper/client"
	"github.com/kitex-contrib/config-zookeeper/utils"
	"github.com/kitex-contrib/config-zookeeper/zookeeper"
)

// ZookeeperClientSuite zookeeper client config suite, configure retry timeout limit and circuitbreak dynamically from zookeeper.
type ZookeeperClientSuite = cwclient.ZookeeperClientSuite

// NewSuite service is the destination service name and client is the local identity.
func NewSuite(service, client string, cli zookeeper.Client, opts ...utils.Option) *ZookeeperClientSuite {
	return cwclient.NewSuite(service, client, cli, opts...)
}
