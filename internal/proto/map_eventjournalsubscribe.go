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
    "github.com/hazelcast/hazelcast-go-client/internal/proto/bufutil"
)

/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Performs the initial subscription to the map event journal.
 * This includes retrieving the event journal sequences of the
 * oldest and newest event in the journal.
 */
//@Generated("f9f5a92260668b4fc86a277d19bbbc19")
const (
    //hex: 0x014100
    MapEventJournalSubscribeRequestMessageType = 82176
    //hex: 0x014101
    MapEventJournalSubscribeResponseMessageType = 82177
    MapEventJournalSubscribeRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapEventJournalSubscribeResponseOldestSequenceFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MapEventJournalSubscribeResponseNewestSequenceFieldOffset = MapEventJournalSubscribeResponseOldestSequenceFieldOffset + bufutil.LongSizeInBytes
    MapEventJournalSubscribeResponseInitialFrameSize = MapEventJournalSubscribeResponseNewestSequenceFieldOffset + bufutil.LongSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapEventJournalSubscribeRequestParameters struct {

    /**
    * name of the map
    */
name string
}

func MapEventJournalSubscribeEncodeRequest(name string) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = true
    clientMessage.SetOperationName("Map.EventJournalSubscribe")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapEventJournalSubscribeRequestMessageType)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapEventJournalSubscribeResponseParameters struct {
    /**
    * Sequence id of the oldest event in the event journal.
    */
oldestSequence int64
    /**
    * Sequence id of the newest event in the event journal.
    */
newestSequence int64
}



func MapEventJournalSubscribeDecodeResponse(clientMessage *bufutil.ClientMessage) *MapEventJournalSubscribeResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapEventJournalSubscribeResponseParameters)
    initialFrame := iterator.Next()
    response.oldestSequence = bufutil.DecodeLong(initialFrame.Content, MapEventJournalSubscribeResponseOldestSequenceFieldOffset)
    response.newestSequence = bufutil.DecodeLong(initialFrame.Content, MapEventJournalSubscribeResponseNewestSequenceFieldOffset)
    return response
}

