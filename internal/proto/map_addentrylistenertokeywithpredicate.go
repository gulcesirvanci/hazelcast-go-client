
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
     "github.com/hazelcast/hazelcast-go-client/serialization"
    "github.com/hazelcast/hazelcast-go-client/core"







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
//@Generated("9f9bc96b7b39d5e978258474d306bd22")
const (
    //hex: 0x011600
    MapAddEntryListenerToKeyWithPredicateRequestMessageType = 71168
    //hex: 0x011601
    MapAddEntryListenerToKeyWithPredicateResponseMessageType = 71169
    MapAddEntryListenerToKeyWithPredicateRequestIncludeValueFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerToKeyWithPredicateRequestListenerFlagsFieldOffset = MapAddEntryListenerToKeyWithPredicateRequestIncludeValueFieldOffset + BooleanSizeInBytes
    MapAddEntryListenerToKeyWithPredicateRequestLocalOnlyFieldOffset = MapAddEntryListenerToKeyWithPredicateRequestListenerFlagsFieldOffset + IntSizeInBytes
    MapAddEntryListenerToKeyWithPredicateRequestInitialFrameSize = MapAddEntryListenerToKeyWithPredicateRequestLocalOnlyFieldOffset + BooleanSizeInBytes
    MapAddEntryListenerToKeyWithPredicateResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerToKeyWithPredicateResponseInitialFrameSize = MapAddEntryListenerToKeyWithPredicateResponseResponseFieldOffset + UUIDSizeInBytes
    MapAddEntryListenerToKeyWithPredicateEventEntryEventTypeFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerToKeyWithPredicateEventEntryUuidFieldOffset = MapAddEntryListenerToKeyWithPredicateEventEntryEventTypeFieldOffset + IntSizeInBytes
    MapAddEntryListenerToKeyWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset = MapAddEntryListenerToKeyWithPredicateEventEntryUuidFieldOffset + UUIDSizeInBytes
    MapAddEntryListenerToKeyWithPredicateEventEntryInitialFrameSize = MapAddEntryListenerToKeyWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset + IntSizeInBytes
    //hex: 0x011602
    MapAddEntryListenerToKeyWithPredicateEventEntryMessageType = 71170


)

func MapAddEntryListenerToKeyWithPredicateEncodeRequest(name string, key serialization.Data, predicate serialization.Data, includeValue bool, listenerFlags int32, localOnly bool) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.AddEntryListenerToKeyWithPredicate")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapAddEntryListenerToKeyWithPredicateRequestMessageType)
    EncodeBoolean(initialFrame.Content, MapAddEntryListenerToKeyWithPredicateRequestIncludeValueFieldOffset, includeValue)
    EncodeInt(initialFrame.Content, MapAddEntryListenerToKeyWithPredicateRequestListenerFlagsFieldOffset, listenerFlags)
    EncodeBoolean(initialFrame.Content, MapAddEntryListenerToKeyWithPredicateRequestLocalOnlyFieldOffset, localOnly)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)


    DataCodecEncode(clientMessage, predicate)

    return clientMessage
}


func MapAddEntryListenerToKeyWithPredicateDecodeResponse(clientMessage *ClientMessage) func() (/*** A unique string which is used as a key to remove the listener.*/response core.Uuid) {
    return func() (/*** A unique string which is used as a key to remove the listener.*/response core.Uuid) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeUUID(initialFrame.Content, MapAddEntryListenerToKeyWithPredicateResponseResponseFieldOffset)
        return
    }
}


type MapAddEntryListenerToKeyWithPredicateHandleEntryFunc func(/* @Nullable */ key serialization.Data, /* @Nullable */ value serialization.Data, /* @Nullable */ oldValue serialization.Data, /* @Nullable */ mergingValue serialization.Data, eventType int32, uuid core.Uuid, numberOfAffectedEntries int32)

func MapAddEntryListenerToKeyWithPredicateHandle(clientMessage *ClientMessage, handleEntry MapAddEntryListenerToKeyWithPredicateHandleEntryFunc){
    messageType := clientMessage.MessageType()
    iterator := clientMessage.FrameIterator()
    if messageType == MapAddEntryListenerToKeyWithPredicateEventEntryMessageType {
        initialFrame := iterator.Next()
        eventType := DecodeInt(initialFrame.Content, MapAddEntryListenerToKeyWithPredicateEventEntryEventTypeFieldOffset)
        uuid := DecodeUUID(initialFrame.Content, MapAddEntryListenerToKeyWithPredicateEventEntryUuidFieldOffset)
        numberOfAffectedEntries := DecodeInt(initialFrame.Content, MapAddEntryListenerToKeyWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset)
        key := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        value := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        oldValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        mergingValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        handleEntry(key, value, oldValue, mergingValue, eventType, uuid, numberOfAffectedEntries)
        return
        }
        // FINEST: Logger.getLogger(super.getClass()).finest("Unknown message type received on event handler :" + messageType)
    }

