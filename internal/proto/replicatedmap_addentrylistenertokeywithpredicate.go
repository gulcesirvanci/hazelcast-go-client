
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
 * Adds an continuous entry listener for this map. The listener will be notified for map add/remove/update/evict
 * events filtered by the given predicate.
 */
//@Generated("bd20125c26566b86c05561b24576c93e")
const (
    //hex: 0x0D0A00
    ReplicatedMapAddEntryListenerToKeyWithPredicateRequestMessageType = 854528
    //hex: 0x0D0A01
    ReplicatedMapAddEntryListenerToKeyWithPredicateResponseMessageType = 854529
    ReplicatedMapAddEntryListenerToKeyWithPredicateRequestLocalOnlyFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ReplicatedMapAddEntryListenerToKeyWithPredicateRequestInitialFrameSize = ReplicatedMapAddEntryListenerToKeyWithPredicateRequestLocalOnlyFieldOffset + BooleanSizeInBytes
    ReplicatedMapAddEntryListenerToKeyWithPredicateResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    ReplicatedMapAddEntryListenerToKeyWithPredicateResponseInitialFrameSize = ReplicatedMapAddEntryListenerToKeyWithPredicateResponseResponseFieldOffset + UUIDSizeInBytes
    ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryEventTypeFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryUuidFieldOffset = ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryEventTypeFieldOffset + IntSizeInBytes
    ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset = ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryUuidFieldOffset + UUIDSizeInBytes
    ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryInitialFrameSize = ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset + IntSizeInBytes
    //hex: 0x0D0A02
    ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryMessageType = 854530


)

func ReplicatedMapAddEntryListenerToKeyWithPredicateEncodeRequest(name string, key serialization.Data, predicate serialization.Data, localOnly bool) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("ReplicatedMap.AddEntryListenerToKeyWithPredicate")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ReplicatedMapAddEntryListenerToKeyWithPredicateRequestMessageType)
    EncodeBoolean(initialFrame.Content, ReplicatedMapAddEntryListenerToKeyWithPredicateRequestLocalOnlyFieldOffset, localOnly)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)


    DataCodecEncode(clientMessage, predicate)

    return clientMessage
}


func ReplicatedMapAddEntryListenerToKeyWithPredicateDecodeResponse(clientMessage *ClientMessage) func() (/*** A unique string  which is used as a key to remove the listener.*/response core.Uuid) {
    return func() (/*** A unique string  which is used as a key to remove the listener.*/response core.Uuid) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeUUID(initialFrame.Content, ReplicatedMapAddEntryListenerToKeyWithPredicateResponseResponseFieldOffset)
        return
    }
}


type ReplicatedMapAddEntryListenerToKeyWithPredicateHandleEntryFunc func(/* @Nullable */ key serialization.Data, /* @Nullable */ value serialization.Data, /* @Nullable */ oldValue serialization.Data, /* @Nullable */ mergingValue serialization.Data, eventType int32, uuid core.Uuid, numberOfAffectedEntries int32)

func ReplicatedMapAddEntryListenerToKeyWithPredicateHandle(clientMessage *ClientMessage, handleEntry ReplicatedMapAddEntryListenerToKeyWithPredicateHandleEntryFunc){
    messageType := clientMessage.MessageType()
    iterator := clientMessage.FrameIterator()
    if messageType == ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryMessageType {
        initialFrame := iterator.Next()
        eventType := DecodeInt(initialFrame.Content, ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryEventTypeFieldOffset)
        uuid := DecodeUUID(initialFrame.Content, ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryUuidFieldOffset)
        numberOfAffectedEntries := DecodeInt(initialFrame.Content, ReplicatedMapAddEntryListenerToKeyWithPredicateEventEntryNumberOfAffectedEntriesFieldOffset)
        key := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        value := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        oldValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        mergingValue := DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  

        handleEntry(key, value, oldValue, mergingValue, eventType, uuid, numberOfAffectedEntries)
        return
        }
        // FINEST: Logger.getLogger(super.getClass()).finest("Unknown message type received on event handler :" + messageType)
    }

