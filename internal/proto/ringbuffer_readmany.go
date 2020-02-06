
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
 * Reads a batch of items from the Ringbuffer. If the number of available items after the first read item is smaller
 * than the maxCount, these items are returned. So it could be the number of items read is smaller than the maxCount.
 * If there are less items available than minCount, then this call blacks. Reading a batch of items is likely to
 * perform better because less overhead is involved. A filter can be provided to only select items that need to be read.
 * If the filter is null, all items are read. If the filter is not null, only items where the filter function returns
 * true are returned. Using filters is a good way to prevent getting items that are of no value to the receiver.
 * This reduces the amount of IO and the number of operations being executed, and can result in a significant performance improvement.
 */
//@Generated("84c835aa159babd553f68bcf6c7ba2f9")
const (
    //hex: 0x170900
    RingbufferReadManyRequestMessageType = 1509632
    //hex: 0x170901
    RingbufferReadManyResponseMessageType = 1509633
    RingbufferReadManyRequestStartSequenceFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    RingbufferReadManyRequestMinCountFieldOffset = RingbufferReadManyRequestStartSequenceFieldOffset + LongSizeInBytes
    RingbufferReadManyRequestMaxCountFieldOffset = RingbufferReadManyRequestMinCountFieldOffset + IntSizeInBytes
    RingbufferReadManyRequestInitialFrameSize = RingbufferReadManyRequestMaxCountFieldOffset + IntSizeInBytes
    RingbufferReadManyResponseReadCountFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    RingbufferReadManyResponseNextSeqFieldOffset = RingbufferReadManyResponseReadCountFieldOffset + IntSizeInBytes
    RingbufferReadManyResponseInitialFrameSize = RingbufferReadManyResponseNextSeqFieldOffset + LongSizeInBytes

)

func RingbufferReadManyEncodeRequest(name string, startSequence int64, minCount int32, maxCount int32, /* @Nullable */ filter serialization.Data) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Ringbuffer.ReadMany")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, RingbufferReadManyRequestMessageType)
    EncodeLong(initialFrame.Content, RingbufferReadManyRequestStartSequenceFieldOffset, startSequence)
    EncodeInt(initialFrame.Content, RingbufferReadManyRequestMinCountFieldOffset, minCount)
    EncodeInt(initialFrame.Content, RingbufferReadManyRequestMaxCountFieldOffset, maxCount)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    EncodeNullable(clientMessage, filter, DataCodecEncode)  

    return clientMessage
}


func RingbufferReadManyDecodeResponse(clientMessage *ClientMessage) func() (/*** Number of items that have been read before filtering.*/readCount int32, /*** List of items that have beee read.*/items []serialization.Data, /*** List of sequence numbers for the items that have been read.*//* @Nullable */itemSeqs []int64, /*** Sequence number of the item following the last read item.*/nextSeq int64) {
    return func() (/*** Number of items that have been read before filtering.*/readCount int32, /*** List of items that have beee read.*/items []serialization.Data, /*** List of sequence numbers for the items that have been read.*//* @Nullable */itemSeqs []int64, /*** Sequence number of the item following the last read item.*/nextSeq int64) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        readCount = DecodeInt(initialFrame.Content, RingbufferReadManyResponseReadCountFieldOffset)
        nextSeq = DecodeLong(initialFrame.Content, RingbufferReadManyResponseNextSeqFieldOffset)
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

