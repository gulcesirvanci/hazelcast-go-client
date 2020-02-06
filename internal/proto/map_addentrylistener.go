
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
//@Generated("40af85308ddb90631ea32c9c50482b06")
const (
    //hex: 0x011900
    MapAddEntryListenerRequestMessageType = 71936
    //hex: 0x011901
    MapAddEntryListenerResponseMessageType = 71937
    MapAddEntryListenerRequestIncludeValueFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerRequestListenerFlagsFieldOffset = MapAddEntryListenerRequestIncludeValueFieldOffset + BooleanSizeInBytes
    MapAddEntryListenerRequestLocalOnlyFieldOffset = MapAddEntryListenerRequestListenerFlagsFieldOffset + IntSizeInBytes
    MapAddEntryListenerRequestInitialFrameSize = MapAddEntryListenerRequestLocalOnlyFieldOffset + BooleanSizeInBytes
    MapAddEntryListenerResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerResponseInitialFrameSize = MapAddEntryListenerResponseResponseFieldOffset + UUIDSizeInBytes
    MapAddEntryListenerEventEntryEventTypeFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerEventEntryUuidFieldOffset = MapAddEntryListenerEventEntryEventTypeFieldOffset + IntSizeInBytes
    MapAddEntryListenerEventEntryNumberOfAffectedEntriesFieldOffset = MapAddEntryListenerEventEntryUuidFieldOffset + UUIDSizeInBytes
    MapAddEntryListenerEventEntryInitialFrameSize = MapAddEntryListenerEventEntryNumberOfAffectedEntriesFieldOffset + IntSizeInBytes
    //hex: 0x011902
    MapAddEntryListenerEventEntryMessageType = 71938


)

func MapAddEntryListenerEncodeRequest(name string, includeValue bool, listenerFlags int32, localOnly bool) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.AddEntryListener")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapAddEntryListenerRequestMessageType)
    EncodeBoolean(initialFrame.Content, MapAddEntryListenerRequestIncludeValueFieldOffset, includeValue)
    EncodeInt(initialFrame.Content, MapAddEntryListenerRequestListenerFlagsFieldOffset, listenerFlags)
    EncodeBoolean(initialFrame.Content, MapAddEntryListenerRequestLocalOnlyFieldOffset, localOnly)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func MapAddEntryListenerDecodeResponse(clientMessage *ClientMessage) func() (/*** A unique string which is used as a key to remove the listener.*/response core.Uuid) {
    return func() (/*** A unique string which is used as a key to remove the listener.*/response core.Uuid) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeUUID(initialFrame.Content, MapAddEntryListenerResponseResponseFieldOffset)
        return
    }
}


type MapAddEntryListenerHandleEntryFunc func(/* @Nullable */ key serialization.Data, /* @Nullable */ value serialization.Data, /* @Nullable */ oldValue serialization.Data, /* @Nullable */ mergingValue serialization.Data, eventType int32, uuid core.Uuid, numberOfAffectedEntries int32)

func MapAddEntryListenerHandle(clientMessage *ClientMessage, handleEntry MapAddEntryListenerHandleEntryFunc){
    messageType := clientMessage.MessageType()
    iterator := clientMessage.FrameIterator()
    if messageType == MapAddEntryListenerEventEntryMessageType {
        initialFrame := iterator.Next()
        eventType := DecodeInt(initialFrame.Content, MapAddEntryListenerEventEntryEventTypeFieldOffset)
        uuid := DecodeUUID(initialFrame.Content, MapAddEntryListenerEventEntryUuidFieldOffset)
        numberOfAffectedEntries := DecodeInt(initialFrame.Content, MapAddEntryListenerEventEntryNumberOfAffectedEntriesFieldOffset)
        key := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        value := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        oldValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        mergingValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        handleEntry(key, value, oldValue, mergingValue, eventType, uuid, numberOfAffectedEntries)
        return
        }
        // FINEST: Logger.getLogger(super.getClass()).finest("Unknown message type received on event handler :" + messageType)
    }

