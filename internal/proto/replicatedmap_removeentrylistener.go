
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
    "github.com/hazelcast/hazelcast-go-client/core"
)


/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Removes the specified entry listener. If there is no such listener added before, this call does no change in the
 * cluster and returns false.
 */
//@Generated("e6d19c936660ba50ed534eacc00892c7")
const (
    //hex: 0x0D0E00
    ReplicatedMapRemoveEntryListenerRequestMessageType = 855552
    //hex: 0x0D0E01
    ReplicatedMapRemoveEntryListenerResponseMessageType = 855553
    ReplicatedMapRemoveEntryListenerRequestRegistrationIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ReplicatedMapRemoveEntryListenerRequestInitialFrameSize = ReplicatedMapRemoveEntryListenerRequestRegistrationIdFieldOffset + UUIDSizeInBytes
    ReplicatedMapRemoveEntryListenerResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    ReplicatedMapRemoveEntryListenerResponseInitialFrameSize = ReplicatedMapRemoveEntryListenerResponseResponseFieldOffset + BooleanSizeInBytes

)

func ReplicatedMapRemoveEntryListenerEncodeRequest(name string, registrationId core.Uuid) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("ReplicatedMap.RemoveEntryListener")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ReplicatedMapRemoveEntryListenerRequestMessageType)
    EncodeUUID(initialFrame.Content, ReplicatedMapRemoveEntryListenerRequestRegistrationIdFieldOffset, registrationId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func ReplicatedMapRemoveEntryListenerDecodeResponse(clientMessage *ClientMessage) func() (/*** True if registration is removed, false otherwise.*/response bool) {
    return func() (/*** True if registration is removed, false otherwise.*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, ReplicatedMapRemoveEntryListenerResponseResponseFieldOffset)
        return
    }
}

