/*
* Copyright (c) 2008-2019, Hazelcast, Inc. All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License")
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

package proto

import (
_"bytes"
"github.com/hazelcast/hazelcast-go-client/serialization"
_ "github.com/hazelcast/hazelcast-go-client"
)

/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Shuts down the cluster.
 */
//@Generated("1c9951c38e6ba0ac14be75b15f10d71b")
const (
    //hex: 0x200F00
    MCShutdownClusterRequestMessageType = 2100992
    //hex: 0x200F01
    MCShutdownClusterResponseMessageType = 2100993
    MCShutdownClusterRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MCShutdownClusterResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MCShutdownClusterRequestParameters struct {
}

func MCShutdownClusterEncodeRequest() *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("MC.ShutdownCluster")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MCShutdownClusterRequestMessageType)
    clientMessage.Add(initialFrame)
    return clientMessage
}

func MCShutdownClusterDecodeRequest(clientMessage *bufutil.ClientMessagex) *MCShutdownClusterRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(MCShutdownClusterRequestParameters)
    //empty initial frame
    iterator.Next()
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MCShutdownClusterResponseParameters struct {
}

func MCShutdownClusterEncodeResponse() *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, MCShutdownClusterResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MCShutdownClusterResponseMessageType)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func MCShutdownClusterDecodeResponse(clientMessage *bufutil.ClientMessagex) *MCShutdownClusterResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MCShutdownClusterResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

