
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
    "github.com/hazelcast/hazelcast-go-client/serialization"



)


/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Adds an listener for this collection. Listener will be notified or all collection add/remove events.
 */
//@Generated("cd1aaa7c47eefe949fab45ec02c58b54")
const (
    //hex: 0x031100
    QueueAddListenerRequestMessageType = 200960
    //hex: 0x031101
    QueueAddListenerResponseMessageType = 200961
    QueueAddListenerRequestIncludeValueFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    QueueAddListenerRequestLocalOnlyFieldOffset = QueueAddListenerRequestIncludeValueFieldOffset + BooleanSizeInBytes
    QueueAddListenerRequestInitialFrameSize = QueueAddListenerRequestLocalOnlyFieldOffset + BooleanSizeInBytes
    QueueAddListenerResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    QueueAddListenerResponseInitialFrameSize = QueueAddListenerResponseResponseFieldOffset + UUIDSizeInBytes
    QueueAddListenerEventItemUuidFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    QueueAddListenerEventItemEventTypeFieldOffset = QueueAddListenerEventItemUuidFieldOffset + UUIDSizeInBytes
    QueueAddListenerEventItemInitialFrameSize = QueueAddListenerEventItemEventTypeFieldOffset + IntSizeInBytes
    //hex: 0x031102
    QueueAddListenerEventItemMessageType = 200962


)

func QueueAddListenerEncodeRequest(name string, includeValue bool, localOnly bool) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Queue.AddListener")
	initialFrame := &Frame{Content: make([]byte, QueueAddListenerResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, QueueAddListenerRequestMessageType)
    EncodeBoolean(initialFrame.Content, QueueAddListenerRequestIncludeValueFieldOffset, includeValue)
    EncodeBoolean(initialFrame.Content, QueueAddListenerRequestLocalOnlyFieldOffset, localOnly)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func QueueAddListenerDecodeResponse(clientMessage *ClientMessage) func() (/*** The registration id*/response *core.Uuid) {
    return func() (/*** The registration id*/response *core.Uuid) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeUUID(initialFrame.Content, QueueAddListenerResponseResponseFieldOffset)
        return
    }
}


type QueueAddListenerHandleItemFunc func(/* @Nullable */ item serialization.Data, uuid *core.Uuid, eventType int32)

func QueueAddListenerHandle(clientMessage *ClientMessage, handleItem QueueAddListenerHandleItemFunc){
    messageType := clientMessage.MessageType()
    iterator := clientMessage.FrameIterator()
    if messageType == QueueAddListenerEventItemMessageType {
        initialFrame := iterator.Next()
        uuid := DecodeUUID(initialFrame.Content, QueueAddListenerEventItemUuidFieldOffset)
        eventType := DecodeInt(initialFrame.Content, QueueAddListenerEventItemEventTypeFieldOffset)
        item := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        handleItem(item, uuid, eventType)
        return
        }
        // FINEST: Logger.getLogger(super.getClass()).finest("Unknown message type received on event handler :" + messageType)
    }

