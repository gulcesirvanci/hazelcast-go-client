
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
 * Removes the specified distributed object listener. If there is no such listener added before, this call does no
 * change in the cluster and returns false.
 */
//@Generated("43d8986365ac19689ba0af80f7d8f5dd")
const (
    //hex: 0x000A00
    ClientRemoveDistributedObjectListenerRequestMessageType = 2560
    //hex: 0x000A01
    ClientRemoveDistributedObjectListenerResponseMessageType = 2561
    ClientRemoveDistributedObjectListenerRequestRegistrationIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ClientRemoveDistributedObjectListenerRequestInitialFrameSize = ClientRemoveDistributedObjectListenerRequestRegistrationIdFieldOffset + UUIDSizeInBytes
    ClientRemoveDistributedObjectListenerResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    ClientRemoveDistributedObjectListenerResponseInitialFrameSize = ClientRemoveDistributedObjectListenerResponseResponseFieldOffset + BooleanSizeInBytes

)

func ClientRemoveDistributedObjectListenerEncodeRequest(registrationId core.Uuid) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Client.RemoveDistributedObjectListener")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ClientRemoveDistributedObjectListenerRequestMessageType)
    EncodeUUID(initialFrame.Content, ClientRemoveDistributedObjectListenerRequestRegistrationIdFieldOffset, registrationId)
    clientMessage.Add(initialFrame)
    return clientMessage
}


func ClientRemoveDistributedObjectListenerDecodeResponse(clientMessage *ClientMessage) func() ( /*** true if the listener existed and removed, false otherwise.*/response bool) {
    return func() (/*** true if the listener existed and removed, false otherwise.*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, ClientRemoveDistributedObjectListenerResponseResponseFieldOffset)
        return
    }
}

