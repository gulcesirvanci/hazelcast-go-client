
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
 * Adds an continuous entry listener for this map. Listener will get notified for map add/remove/update/evict events
 * filtered by the given predicate.
 */
//@Generated("b12f764fb888ef57a569618b1ffc17b7")
const (
    //hex: 0x011700
    MapAddEntryListenerWithPredicateRequestMessageType = 71424
    //hex: 0x011701
    MapAddEntryListenerWithPredicateResponseMessageType = 71425
    MapAddEntryListenerWithPredicateRequestIncludeValueFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerWithPredicateRequestListenerFlagsFieldOffset = MapAddEntryListenerWithPredicateRequestIncludeValueFieldOffset + BooleanSizeInBytes
    MapAddEntryListenerWithPredicateRequestLocalOnlyFieldOffset = MapAddEntryListenerWithPredicateRequestListenerFlagsFieldOffset + IntSizeInBytes
    MapAddEntryListenerWithPredicateRequestInitialFrameSize = MapAddEntryListenerWithPredicateRequestLocalOnlyFieldOffset + BooleanSizeInBytes
    MapAddEntryListenerWithPredicateResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerWithPredicateResponseInitialFrameSize = MapAddEntryListenerWithPredicateResponseResponseFieldOffset + UUIDSizeInBytes
    MapAddEntryListenerWithPredicateEventEntryEventTypeFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapAddEntryListenerWithPredicateEventEntryUuidFieldOffset = MapAddEntryListenerWithPredicateEventEntryEventTypeFieldOffset + IntSizeInBytes
    MapAddEntryListenerWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset = MapAddEntryListenerWithPredicateEventEntryUuidFieldOffset + UUIDSizeInBytes
    MapAddEntryListenerWithPredicateEventEntryInitialFrameSize = MapAddEntryListenerWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset + IntSizeInBytes
    //hex: 0x011702
    MapAddEntryListenerWithPredicateEventEntryMessageType = 71426


)

func MapAddEntryListenerWithPredicateEncodeRequest(name string, predicate serialization.Data, includeValue bool, listenerFlags int32, localOnly bool) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.AddEntryListenerWithPredicate")
	initialFrame := &Frame{Content: make([]byte, MapAddEntryListenerWithPredicateResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapAddEntryListenerWithPredicateRequestMessageType)
    EncodeBoolean(initialFrame.Content, MapAddEntryListenerWithPredicateRequestIncludeValueFieldOffset, includeValue)
    EncodeInt(initialFrame.Content, MapAddEntryListenerWithPredicateRequestListenerFlagsFieldOffset, listenerFlags)
    EncodeBoolean(initialFrame.Content, MapAddEntryListenerWithPredicateRequestLocalOnlyFieldOffset, localOnly)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, predicate)

    return clientMessage
}


func MapAddEntryListenerWithPredicateDecodeResponse(clientMessage *ClientMessage) func() (/*** A unique string which is used as a key to remove the listener.*/response *core.Uuid) {
    return func() (/*** A unique string which is used as a key to remove the listener.*/response *core.Uuid) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeUUID(initialFrame.Content, MapAddEntryListenerWithPredicateResponseResponseFieldOffset)
        return
    }
}


type MapAddEntryListenerWithPredicateHandleEntryFunc func(/* @Nullable */ key serialization.Data, /* @Nullable */ value serialization.Data, /* @Nullable */ oldValue serialization.Data, /* @Nullable */ mergingValue serialization.Data, eventType int32, uuid *core.Uuid, numberOfAffectedEntries int32)

func MapAddEntryListenerWithPredicateHandle(clientMessage *ClientMessage, handleEntry MapAddEntryListenerWithPredicateHandleEntryFunc){
    messageType := clientMessage.MessageType()
    iterator := clientMessage.FrameIterator()
    if messageType == MapAddEntryListenerWithPredicateEventEntryMessageType {
        initialFrame := iterator.Next()
        eventType := DecodeInt(initialFrame.Content, MapAddEntryListenerWithPredicateEventEntryEventTypeFieldOffset)
        uuid := DecodeUUID(initialFrame.Content, MapAddEntryListenerWithPredicateEventEntryUuidFieldOffset)
        numberOfAffectedEntries := DecodeInt(initialFrame.Content, MapAddEntryListenerWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset)
        key := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        value := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        oldValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        mergingValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        handleEntry(key, value, oldValue, mergingValue, eventType, uuid, numberOfAffectedEntries)
        return
        }
        // FINEST: Logger.getLogger(super.getClass()).finest("Unknown message type received on event handler :" + messageType)
    }

