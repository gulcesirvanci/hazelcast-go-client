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
 * Destroys the distributed object with the given name on the requested
 * CP group
 */
//@Generated("766403b7af041147dfb953a734f3d599")
const (
    //hex: 0x1E0200
    CPGroupDestroyCPObjectRequestMessageType = 1966592
    //hex: 0x1E0201
    CPGroupDestroyCPObjectResponseMessageType = 1966593
    CPGroupDestroyCPObjectRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    CPGroupDestroyCPObjectResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type CPGroupDestroyCPObjectRequestParameters struct {

    /**
    * CP group id of this distributed object
    */
groupId RaftGroupId

    /**
    * The service of this distributed object
    */
serviceName string

    /**
    * The name of this distributed object
    */
objectName string
}

func CPGroupDestroyCPObjectEncodeRequest(groupId RaftGroupId, serviceName string, objectName string) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("CPGroup.DestroyCPObject")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, CPGroupDestroyCPObjectRequestMessageType)
    clientMessage.Add(initialFrame)
    RaftGroupIdCodec.encode(clientMessage, groupId)
    StringCodec.encode(clientMessage, serviceName)
    StringCodec.encode(clientMessage, objectName)
    return clientMessage
}

func CPGroupDestroyCPObjectDecodeRequest(clientMessage *bufutil.ClientMessagex) *CPGroupDestroyCPObjectRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(CPGroupDestroyCPObjectRequestParameters)
    //empty initial frame
    iterator.Next()
    request.groupId = RaftGroupIdCodec.decode(iterator)
    request.serviceName = StringCodec.decode(iterator)
    request.objectName = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type CPGroupDestroyCPObjectResponseParameters struct {
}

func CPGroupDestroyCPObjectEncodeResponse() *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, CPGroupDestroyCPObjectResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, CPGroupDestroyCPObjectResponseMessageType)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func CPGroupDestroyCPObjectDecodeResponse(clientMessage *bufutil.ClientMessagex) *CPGroupDestroyCPObjectResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(CPGroupDestroyCPObjectResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

