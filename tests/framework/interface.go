// Copyright 2022 The etcd Authors
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

package framework

import (
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/tests/v3/framework/config"
)

type testRunner interface {
	TestMain(m *testing.M)
	BeforeTest(testing.TB)
	NewCluster(testing.TB, config.ClusterConfig) Cluster
}

type Cluster interface {
	Close() error
	Client() Client
}

type Client interface {
	Put(key, value string) error
	Get(key string, opts config.GetOptions) (*clientv3.GetResponse, error)
	Delete(key string, opts config.DeleteOptions) (*clientv3.DeleteResponse, error)
	Compact(rev int64, opts config.CompactOption) (*clientv3.CompactResponse, error)
}
