
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
 * Removes the specified partition lost listener. If there is no such listener added before, this call does no change
 * in the cluster and returns false.
 */
//@Generated("38fc1d0a29e3ea4dbe3f2968b693f7ee")
const (
    //hex: 0x000700
    ClientRemovePartitionLostListenerRequestMessageType = 1792
    //hex: 0x000701
    ClientRemovePartitionLostListenerResponseMessageType = 1793
    ClientRemovePartitionLostListenerRequestRegistrationIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ClientRemovePartitionLostListenerRequestInitialFrameSize = ClientRemovePartitionLostListenerRequestRegistrationIdFieldOffset + UUIDSizeInBytes
    ClientRemovePartitionLostListenerResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    ClientRemovePartitionLostListenerResponseInitialFrameSize = ClientRemovePartitionLostListenerResponseResponseFieldOffset + BooleanSizeInBytes

)

func ClientRemovePartitionLostListenerEncodeRequest(registrationId core.Uuid) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Client.RemovePartitionLostListener")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ClientRemovePartitionLostListenerRequestMessageType)
    EncodeUUID(initialFrame.Content, ClientRemovePartitionLostListenerRequestRegistrationIdFieldOffset, registrationId)
    clientMessage.Add(initialFrame)
    return clientMessage
}


func ClientRemovePartitionLostListenerDecodeResponse(clientMessage *ClientMessage) func() (/*** true if the listener existed and removed, false otherwise.*/response bool) {
    return func() (/*** true if the listener existed and removed, false otherwise.*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, ClientRemovePartitionLostListenerResponseResponseFieldOffset)
        return
    }
}

