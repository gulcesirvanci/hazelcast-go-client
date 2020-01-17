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
 * TODO DOC
 */
//@Generated("208e0e4631150ffbf21ae01640477c8b")
const (
    //hex: 0x000400
    ClientCreateProxyRequestMessageType = 1024
    //hex: 0x000401
    ClientCreateProxyResponseMessageType = 1025
    ClientCreateProxyRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ClientCreateProxyResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ClientCreateProxyRequestParameters struct {

    /**
    * The distributed object name for which the proxy is being created for.
    */
name string

    /**
    * The name of the service. Possible service names are:
    * "hz:impl:listService"
    * "hz:impl:queueService"
    * "hz:impl:setService"
    * "hz:impl:idGeneratorService"
    * "hz:impl:executorService"
    * "hz:impl:mapService"
    * "hz:impl:multiMapService"
    * "hz:impl:splitBrainProtectionService"
    * "hz:impl:replicatedMapService"
    * "hz:impl:ringbufferService"
    * "hz:core:proxyService"
    * "hz:impl:reliableTopicService"
    * "hz:impl:topicService"
    * "hz:core:txManagerService"
    * "hz:impl:xaService"
    */
serviceName string

    /**
    * TODO DOC
    */
target Address
}

func ClientCreateProxyEncodeRequest(name string, serviceName string, target Address) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Client.CreateProxy")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ClientCreateProxyRequestMessageType)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.StringCodecEncode(clientMessage, serviceName)
    bufutil.AddressCodecEncode(clientMessage, target)
    return clientMessage
}




/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ClientCreateProxyResponseParameters struct {
}



func ClientCreateProxyDecodeResponse(clientMessage *bufutil.ClientMessagex) *ClientCreateProxyResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ClientCreateProxyResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

