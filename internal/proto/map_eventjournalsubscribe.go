
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
//@Generated("e4a8a832d16dc02bf38e5fd04f849055")
const (
    //hex: 0x014100
    MapEventJournalSubscribeRequestMessageType = 82176
    //hex: 0x014101
    MapEventJournalSubscribeResponseMessageType = 82177
    MapEventJournalSubscribeRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    MapEventJournalSubscribeResponseOldestSequenceFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapEventJournalSubscribeResponseNewestSequenceFieldOffset = MapEventJournalSubscribeResponseOldestSequenceFieldOffset + LongSizeInBytes
    MapEventJournalSubscribeResponseInitialFrameSize = MapEventJournalSubscribeResponseNewestSequenceFieldOffset + LongSizeInBytes

)

func MapEventJournalSubscribeEncodeRequest(name string) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Map.EventJournalSubscribe")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapEventJournalSubscribeRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func MapEventJournalSubscribeDecodeResponse(clientMessage *ClientMessage) func() ( /*** Sequence id of the oldest event in the event journal.*/oldestSequence int64, /*** Sequence id of the newest event in the event journal.*/newestSequence int64) {
    return func() (/*** Sequence id of the oldest event in the event journal.*/oldestSequence int64, /*** Sequence id of the newest event in the event journal.*/newestSequence int64) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        oldestSequence = DecodeLong(initialFrame.Content, MapEventJournalSubscribeResponseOldestSequenceFieldOffset)
        newestSequence = DecodeLong(initialFrame.Content, MapEventJournalSubscribeResponseNewestSequenceFieldOffset)
        return
    }
}

