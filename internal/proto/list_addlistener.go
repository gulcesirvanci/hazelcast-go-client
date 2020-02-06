
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
 * Adds an item listener for this collection. Listener will be notified for all collection add/remove events.
 */
//@Generated("400aa8017cbc2a561e8bd87a724e150a")
const (
    //hex: 0x050B00
    ListAddListenerRequestMessageType = 330496
    //hex: 0x050B01
    ListAddListenerResponseMessageType = 330497
    ListAddListenerRequestIncludeValueFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ListAddListenerRequestLocalOnlyFieldOffset = ListAddListenerRequestIncludeValueFieldOffset + BooleanSizeInBytes
    ListAddListenerRequestInitialFrameSize = ListAddListenerRequestLocalOnlyFieldOffset + BooleanSizeInBytes
    ListAddListenerResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    ListAddListenerResponseInitialFrameSize = ListAddListenerResponseResponseFieldOffset + UUIDSizeInBytes
    ListAddListenerEventItemUuidFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ListAddListenerEventItemEventTypeFieldOffset = ListAddListenerEventItemUuidFieldOffset + UUIDSizeInBytes
    ListAddListenerEventItemInitialFrameSize = ListAddListenerEventItemEventTypeFieldOffset + IntSizeInBytes
    //hex: 0x050B02
    ListAddListenerEventItemMessageType = 330498


)

func ListAddListenerEncodeRequest(name string, includeValue bool, localOnly bool) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("List.AddListener")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ListAddListenerRequestMessageType)
    EncodeBoolean(initialFrame.Content, ListAddListenerRequestIncludeValueFieldOffset, includeValue)
    EncodeBoolean(initialFrame.Content, ListAddListenerRequestLocalOnlyFieldOffset, localOnly)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func ListAddListenerDecodeResponse(clientMessage *ClientMessage) func() ( /*** Registration id for the listener.*/response core.Uuid) {
    return func() (/*** Registration id for the listener.*/response core.Uuid) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeUUID(initialFrame.Content, ListAddListenerResponseResponseFieldOffset)
        return
    }
}


type ListAddListenerHandleItemFunc func(/* @Nullable */ item serialization.Data, uuid core.Uuid, eventType int32)

func ListAddListenerHandle(clientMessage *ClientMessage, handleItem ListAddListenerHandleItemFunc){
    messageType := clientMessage.MessageType()
    iterator := clientMessage.FrameIterator()
    if messageType == ListAddListenerEventItemMessageType {
        initialFrame := iterator.Next()
        uuid := DecodeUUID(initialFrame.Content, ListAddListenerEventItemUuidFieldOffset)
        eventType := DecodeInt(initialFrame.Content, ListAddListenerEventItemEventTypeFieldOffset)
        item := DecodeNullable(iterator, DataCodecDecode).(serialization.Data)

        handleItem(item, uuid, eventType)
        return
        }
        // FINEST: Logger.getLogger(super.getClass()).finest("Unknown message type received on event handler :" + messageType)
    }

