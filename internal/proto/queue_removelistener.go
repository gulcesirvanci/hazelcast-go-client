
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
 * Removes the specified item listener. If there is no such listener added before, this call does no change in the
 * cluster and returns false.
 */
//@Generated("7660c50a41a3e6c8683b4932e9d17749")
const (
    //hex: 0x031200
    QueueRemoveListenerRequestMessageType = 201216
    //hex: 0x031201
    QueueRemoveListenerResponseMessageType = 201217
    QueueRemoveListenerRequestRegistrationIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    QueueRemoveListenerRequestInitialFrameSize = QueueRemoveListenerRequestRegistrationIdFieldOffset + UUIDSizeInBytes
    QueueRemoveListenerResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    QueueRemoveListenerResponseInitialFrameSize = QueueRemoveListenerResponseResponseFieldOffset + BooleanSizeInBytes

)

func QueueRemoveListenerEncodeRequest(name string, registrationId core.Uuid) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Queue.RemoveListener")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, QueueRemoveListenerRequestMessageType)
    EncodeUUID(initialFrame.Content, QueueRemoveListenerRequestRegistrationIdFieldOffset, registrationId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func QueueRemoveListenerDecodeResponse(clientMessage *ClientMessage) func() (/*** True if the item listener is removed, false otherwise*/response bool) {
    return func() (/*** True if the item listener is removed, false otherwise*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, QueueRemoveListenerResponseResponseFieldOffset)
        return
    }
}

