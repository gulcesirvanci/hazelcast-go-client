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
 * Removes the specified item listener. Returns silently if the specified listener was not added before.
 */
//@Generated("9260fc2e05a830fc96f973ff87515a9a")
const (
    //hex: 0x050C00
    ListRemoveListenerRequestMessageType = 330752
    //hex: 0x050C01
    ListRemoveListenerResponseMessageType = 330753
    ListRemoveListenerRequestRegistrationIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ListRemoveListenerRequestInitialFrameSize = ListRemoveListenerRequestRegistrationIdFieldOffset + bufutil.UUIDSizeInBytes
    ListRemoveListenerResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    ListRemoveListenerResponseInitialFrameSize = ListRemoveListenerResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListRemoveListenerRequestParameters struct {

    /**
    * Name of the List
    */
name string

    /**
    * The id of the listener which was provided during registration.
    */
registrationId UUID
}

func ListRemoveListenerEncodeRequest(name string, registrationId UUID) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("List.RemoveListener")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListRemoveListenerRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, ListRemoveListenerRequestRegistrationIdFieldOffset, registrationId)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func ListRemoveListenerDecodeRequest(clientMessage *bufutil.ClientMessagex) *ListRemoveListenerRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(ListRemoveListenerRequestParameters)
    initialFrame := iterator.Next()
    request.registrationId = bufutil.DecodeUUID(initialFrame.Content, ListRemoveListenerRequestRegistrationIdFieldOffset)
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListRemoveListenerResponseParameters struct {
    /**
    * True if unregistered, false otherwise.
    */
response bool
}

func ListRemoveListenerEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, ListRemoveListenerResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListRemoveListenerResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, ListRemoveListenerResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func ListRemoveListenerDecodeResponse(clientMessage *bufutil.ClientMessagex) *ListRemoveListenerResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ListRemoveListenerResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, ListRemoveListenerResponseResponseFieldOffset)
    return response
}

