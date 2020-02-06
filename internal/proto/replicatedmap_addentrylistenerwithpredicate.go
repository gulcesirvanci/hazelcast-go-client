
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
 * Adds an continuous entry listener for this map. The listener will be notified for map add/remove/update/evict
 * events filtered by the given predicate.
 */
//@Generated("9eb034a3f6757eabf2b8a0dc5a866c0b")
const (
    //hex: 0x0D0B00
    ReplicatedMapAddEntryListenerWithPredicateRequestMessageType = 854784
    //hex: 0x0D0B01
    ReplicatedMapAddEntryListenerWithPredicateResponseMessageType = 854785
    ReplicatedMapAddEntryListenerWithPredicateRequestLocalOnlyFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ReplicatedMapAddEntryListenerWithPredicateRequestInitialFrameSize = ReplicatedMapAddEntryListenerWithPredicateRequestLocalOnlyFieldOffset + BooleanSizeInBytes
    ReplicatedMapAddEntryListenerWithPredicateResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    ReplicatedMapAddEntryListenerWithPredicateResponseInitialFrameSize = ReplicatedMapAddEntryListenerWithPredicateResponseResponseFieldOffset + UUIDSizeInBytes
    ReplicatedMapAddEntryListenerWithPredicateEventEntryEventTypeFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ReplicatedMapAddEntryListenerWithPredicateEventEntryUuidFieldOffset = ReplicatedMapAddEntryListenerWithPredicateEventEntryEventTypeFieldOffset + IntSizeInBytes
    ReplicatedMapAddEntryListenerWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset = ReplicatedMapAddEntryListenerWithPredicateEventEntryUuidFieldOffset + UUIDSizeInBytes
    ReplicatedMapAddEntryListenerWithPredicateEventEntryInitialFrameSize = ReplicatedMapAddEntryListenerWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset + IntSizeInBytes
    //hex: 0x0D0B02
    ReplicatedMapAddEntryListenerWithPredicateEventEntryMessageType = 854786


)

func ReplicatedMapAddEntryListenerWithPredicateEncodeRequest(name string, predicate serialization.Data, localOnly bool) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("ReplicatedMap.AddEntryListenerWithPredicate")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ReplicatedMapAddEntryListenerWithPredicateRequestMessageType)
    EncodeBoolean(initialFrame.Content, ReplicatedMapAddEntryListenerWithPredicateRequestLocalOnlyFieldOffset, localOnly)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, predicate)

    return clientMessage
}


func ReplicatedMapAddEntryListenerWithPredicateDecodeResponse(clientMessage *ClientMessage) func() ( /*** A unique string  which is used as a key to remove the listener.*/response core.Uuid) {
    return func() (/*** A unique string  which is used as a key to remove the listener.*/response core.Uuid) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeUUID(initialFrame.Content, ReplicatedMapAddEntryListenerWithPredicateResponseResponseFieldOffset)
        return
    }
}


type ReplicatedMapAddEntryListenerWithPredicateHandleEntryFunc func(/* @Nullable */ key serialization.Data, /* @Nullable */ value serialization.Data, /* @Nullable */ oldValue serialization.Data, /* @Nullable */ mergingValue serialization.Data, eventType int32, uuid core.Uuid, numberOfAffectedEntries int32)

func ReplicatedMapAddEntryListenerWithPredicateHandle(clientMessage *ClientMessage, handleEntry ReplicatedMapAddEntryListenerWithPredicateHandleEntryFunc){
    messageType := clientMessage.MessageType()
    iterator := clientMessage.FrameIterator()
    if messageType == ReplicatedMapAddEntryListenerWithPredicateEventEntryMessageType {
        initialFrame := iterator.Next()
        eventType := DecodeInt(initialFrame.Content, ReplicatedMapAddEntryListenerWithPredicateEventEntryEventTypeFieldOffset)
        uuid := DecodeUUID(initialFrame.Content, ReplicatedMapAddEntryListenerWithPredicateEventEntryUuidFieldOffset)
        numberOfAffectedEntries := DecodeInt(initialFrame.Content, ReplicatedMapAddEntryListenerWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset)
        key := DecodeNullable(iterator, DataCodecDecode).(serialization.Data)

        value := DecodeNullable(iterator, DataCodecDecode).(serialization.Data)

        oldValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data)

        mergingValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data)

        handleEntry(key, value, oldValue, mergingValue, eventType, uuid, numberOfAffectedEntries)
        return
        }
        // FINEST: Logger.getLogger(super.getClass()).finest("Unknown message type received on event handler :" + messageType)
    }

