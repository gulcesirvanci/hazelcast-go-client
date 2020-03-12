// Copyright (c) 2008-2019, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"github.com/hazelcast/hazelcast-go-client/core"
	"github.com/hazelcast/hazelcast-go-client/core/logger"
	"github.com/hazelcast/hazelcast-go-client/security"
	"sync"
	"sync/atomic"
)

const (
	authenticated = iota
	credentialsFailed
	serializationVersionMismatch
)

const serializationVersion = 1
const clientType = "GO"
var ClientVersion = "0.5-SNAPSHOT" //TODO This should be replace with a build time version variable, BuildInfo etc.
var labels []string //todo: check set := make(map[string]void)


type connectionManager interface {
	//getActiveConnections returns a snapshot of active connections
	getActiveConnections() map[string]*Connection

	//getActiveConnection returns connection if available, nil otherwise
	getActiveConnection(address core.Address) *Connection

	//ConnectionCount returns number of active connections
	ConnectionCount() int32

	//getOrConnect returns associated connection if available, creates new connection otherwise
	getOrConnect(address core.Address, asOwner bool) (*Connection, error)

	//getOrTriggerConnect returns associated connection if available, returns error and triggers new connection creation otherwise
	getOrTriggerConnect(address core.Address) (*Connection, error)

	getOwnerConnection() *Connection

	addListener(listener connectionListener)
	onConnectionClose(connection *Connection, cause error)
	NextConnectionID() int64
	IsAlive() bool
	shutdown()
}

//internal definitions and methods called inside connection manager process

type connectionManagerImpl struct {
	client              *HazelcastClient
	connectionsMutex    sync.RWMutex
	connections         map[string]*Connection
	nextConnectionID    int64
	listenerMutex       sync.Mutex
	connectionListeners atomic.Value
	addressTranslator   AddressTranslator
	isAlive             atomic.Value
	credentials         security.Credentials
	logger              logger.Logger
}


func (cm *connectionManagerImpl) initConnectionTimeoutMillis() int {
	networkConfig := cm.client.Config.NetworkConfig()
	connTimeout := networkConfig.ConnectionTimeout() //const time.Duration todo
	if connTimeout == 0 {
		return intMaxValue
	}else{
		return int(connTimeout)
	}
}

//func (cm *connectionManagerImpl) createExecutorService() ScheduledExecutorService { }

//func (cm *connectionManagerImpl) getOutboundPorts() Collection<Integer> { }


func (cm *connectionManagerImpl) getOrConnect(address core.Address, asOwner bool) (*Connection, error) {
	connection := cm.getConnection(address, asOwner)
	if connection != nil {
		return connection, nil
	}
	return cm.getOrCreateConnectionInternal(address, asOwner)
}


/*func (cm *connectionManagerImpl) Start() {
if !cm.IsAlive() {
	//!isAlive.compareAndSet(f, t)
return
}
startNetworking()

heartbeat.start()
connectToCluster()
if isSmartRoutingEnabled {
executor.scheduleWithFixedDelay(connectToAllClusterMembersTask, 1, 1, TimeUnit.SECONDS)
}
}*/

func startNetworking() {

}

func (cm *connectionManagerImpl) IsAlive() bool {
	return cm.isAlive.Load().(bool)
}

func (cm *connectionManagerImpl) shutdown() {
	if cm.isAlive.Load() == false {
		return
	}
	cm.isAlive.Store(false)
	activeCons := cm.getActiveConnections()
	for _, con := range activeCons {

		con.close(core.NewHazelcastClientNotActiveError("client is shutting down", nil))
	}
}