
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
 * Adds a MapListener for this map. To receive an event, you should implement a corresponding MapListener
 * sub-interface for that event.
 */
//@Generated("7a82474f00d2edba24a2a2bc28d4cada")
const (
    //hex: 0x011800
    MapAddEntryListenerToKeyRequestMessageType = 71680
    //hex: 0x011801
    MapAddEntryListenerToKeyResponseMessageType = 71681
    MapAddEntryListenerToKeyRequestIncludeValueFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerToKeyRequestListenerFlagsFieldOffset = MapAddEntryListenerToKeyRequestIncludeValueFieldOffset + BooleanSizeInBytes
    MapAddEntryListenerToKeyRequestLocalOnlyFieldOffset = MapAddEntryListenerToKeyRequestListenerFlagsFieldOffset + IntSizeInBytes
    MapAddEntryListenerToKeyRequestInitialFrameSize = MapAddEntryListenerToKeyRequestLocalOnlyFieldOffset + BooleanSizeInBytes
    MapAddEntryListenerToKeyResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerToKeyResponseInitialFrameSize = MapAddEntryListenerToKeyResponseResponseFieldOffset + UUIDSizeInBytes
    MapAddEntryListenerToKeyEventEntryEventTypeFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerToKeyEventEntryUuidFieldOffset = MapAddEntryListenerToKeyEventEntryEventTypeFieldOffset + IntSizeInBytes
    MapAddEntryListenerToKeyEventEntryNumberOfAffectedEntriesFieldOffset = MapAddEntryListenerToKeyEventEntryUuidFieldOffset + UUIDSizeInBytes
    MapAddEntryListenerToKeyEventEntryInitialFrameSize = MapAddEntryListenerToKeyEventEntryNumberOfAffectedEntriesFieldOffset + IntSizeInBytes
    //hex: 0x011802
    MapAddEntryListenerToKeyEventEntryMessageType = 71682


)

func MapAddEntryListenerToKeyEncodeRequest(name string, key serialization.Data, includeValue bool, listenerFlags int32, localOnly bool) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.AddEntryListenerToKey")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapAddEntryListenerToKeyRequestMessageType)
    EncodeBoolean(initialFrame.Content, MapAddEntryListenerToKeyRequestIncludeValueFieldOffset, includeValue)
    EncodeInt(initialFrame.Content, MapAddEntryListenerToKeyRequestListenerFlagsFieldOffset, listenerFlags)
    EncodeBoolean(initialFrame.Content, MapAddEntryListenerToKeyRequestLocalOnlyFieldOffset, localOnly)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)

    return clientMessage
}


func MapAddEntryListenerToKeyDecodeResponse(clientMessage *ClientMessage) func() ( /*** A unique string which is used as a key to remove the listener.*/response core.Uuid) {
    return func() (/*** A unique string which is used as a key to remove the listener.*/response core.Uuid) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeUUID(initialFrame.Content, MapAddEntryListenerToKeyResponseResponseFieldOffset)
        return
    }
}


type MapAddEntryListenerToKeyHandleEntryFunc func(/* @Nullable */ key serialization.Data, /* @Nullable */ value serialization.Data, /* @Nullable */ oldValue serialization.Data, /* @Nullable */ mergingValue serialization.Data, eventType int32, uuid core.Uuid, numberOfAffectedEntries int32)

func MapAddEntryListenerToKeyHandle(clientMessage *ClientMessage, handleEntry MapAddEntryListenerToKeyHandleEntryFunc){
    messageType := clientMessage.MessageType()
    iterator := clientMessage.FrameIterator()
    if messageType == MapAddEntryListenerToKeyEventEntryMessageType {
        initialFrame := iterator.Next()
        eventType := DecodeInt(initialFrame.Content, MapAddEntryListenerToKeyEventEntryEventTypeFieldOffset)
        uuid := DecodeUUID(initialFrame.Content, MapAddEntryListenerToKeyEventEntryUuidFieldOffset)
        numberOfAffectedEntries := DecodeInt(initialFrame.Content, MapAddEntryListenerToKeyEventEntryNumberOfAffectedEntriesFieldOffset)
        key := DecodeNullable(iterator, DataCodecDecode).(serialization.Data)

        value := DecodeNullable(iterator, DataCodecDecode).(serialization.Data)

        oldValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data)

        mergingValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data)

        handleEntry(key, value, oldValue, mergingValue, eventType, uuid, numberOfAffectedEntries)
        return
        }
        // FINEST: Logger.getLogger(super.getClass()).finest("Unknown message type received on event handler :" + messageType)
    }

