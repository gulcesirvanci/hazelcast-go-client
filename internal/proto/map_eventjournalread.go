
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
)


/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Reads from the map event journal in batches. You may specify the start sequence,
 * the minumum required number of items in the response, the maximum number of items
 * in the response, a predicate that the events should pass and a projection to
 * apply to the events in the journal.
 * If the event journal currently contains less events than {@code minSize}, the
 * call will wait until it has sufficient items.
 * The predicate, filter and projection may be {@code null} in which case all elements are returned
 * and no projection is applied.
 */
//@Generated("f2f380408ec3009219995f81ed8658ff")
const (
    //hex: 0x014200
    MapEventJournalReadRequestMessageType = 82432
    //hex: 0x014201
    MapEventJournalReadResponseMessageType = 82433
    MapEventJournalReadRequestStartSequenceFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapEventJournalReadRequestMinSizeFieldOffset = MapEventJournalReadRequestStartSequenceFieldOffset + LongSizeInBytes
    MapEventJournalReadRequestMaxSizeFieldOffset = MapEventJournalReadRequestMinSizeFieldOffset + IntSizeInBytes
    MapEventJournalReadRequestInitialFrameSize = MapEventJournalReadRequestMaxSizeFieldOffset + IntSizeInBytes
    MapEventJournalReadResponseReadCountFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapEventJournalReadResponseNextSeqFieldOffset = MapEventJournalReadResponseReadCountFieldOffset + IntSizeInBytes
    MapEventJournalReadResponseInitialFrameSize = MapEventJournalReadResponseNextSeqFieldOffset + LongSizeInBytes

)

func MapEventJournalReadEncodeRequest(name string, startSequence int64, minSize int32, maxSize int32, /* @Nullable */ predicate serialization.Data, /* @Nullable */ projection serialization.Data) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Map.EventJournalRead")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapEventJournalReadRequestMessageType)
    EncodeLong(initialFrame.Content, MapEventJournalReadRequestStartSequenceFieldOffset, startSequence)
    EncodeInt(initialFrame.Content, MapEventJournalReadRequestMinSizeFieldOffset, minSize)
    EncodeInt(initialFrame.Content, MapEventJournalReadRequestMaxSizeFieldOffset, maxSize)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    EncodeNullable(clientMessage, predicate, DataCodecEncode)  


    EncodeNullable(clientMessage, projection, DataCodecEncode)  

    return clientMessage
}


func MapEventJournalReadDecodeResponse(clientMessage *ClientMessage) func() (/*** Number of items that have been read.*/readCount int32, /*** List of items that have been read.*/items []serialization.Data, /*** Sequence numbers of items in the event journal.*//* @Nullable */itemSeqs []int64, /*** Sequence number of the item following the last read item.*/nextSeq int64) {
    return func() (/*** Number of items that have been read.*/readCount int32, /*** List of items that have been read.*/items []serialization.Data, /*** Sequence numbers of items in the event journal.*//* @Nullable */itemSeqs []int64, /*** Sequence number of the item following the last read item.*/nextSeq int64) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        readCount = DecodeInt(initialFrame.Content, MapEventJournalReadResponseReadCountFieldOffset)
        nextSeq = DecodeLong(initialFrame.Content, MapEventJournalReadResponseNextSeqFieldOffset)
        var result []serialization.Data
        //begin frame, list
        iterator.Next()
        for !NextFrameIsDataStructureEndFrame(iterator) {
        result = append(result, DataCodecDecode(iterator))
        }
        //end frame, list
        iterator.Next()
        items = result //0.1
        itemSeqs = DecodeNullable(iterator, LongArrayCodecDecode).([]int64) // 2  
        return
    }
}

