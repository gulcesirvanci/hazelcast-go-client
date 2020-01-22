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
    "github.com/hazelcast/hazelcast-go-client/serialization"
)

/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Adds an item to the tail of the Ringbuffer. If there is space in the ringbuffer, the call
 * will return the sequence of the written item. If there is no space, it depends on the overflow policy what happens:
 * OverflowPolicy OVERWRITE we just overwrite the oldest item in the ringbuffer and we violate the ttl
 * OverflowPolicy FAIL we return -1. The reason that FAIL exist is to give the opportunity to obey the ttl.
 * <p/>
 * This sequence will always be unique for this Ringbuffer instance so it can be used as a unique id generator if you are
 * publishing items on this Ringbuffer. However you need to take care of correctly determining an initial id when any node
 * uses the ringbuffer for the first time. The most reliable way to do that is to write a dummy item into the ringbuffer and
 * use the returned sequence as initial  id. On the reading side, this dummy item should be discard. Please keep in mind that
 * this id is not the sequence of the item you are about to publish but from a previously published item. So it can't be used
 * to find that item.
 */
//@Generated("d87177474087284c50a7527dc12b3c2f")
const (
    //hex: 0x170600
    RingbufferAddRequestMessageType = 1508864
    //hex: 0x170601
    RingbufferAddResponseMessageType = 1508865
    RingbufferAddRequestOverflowPolicyFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    RingbufferAddRequestInitialFrameSize = RingbufferAddRequestOverflowPolicyFieldOffset + bufutil.IntSizeInBytes
    RingbufferAddResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    RingbufferAddResponseInitialFrameSize = RingbufferAddResponseResponseFieldOffset + bufutil.LongSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type RingbufferAddRequestParameters struct {

    /**
    * Name of the Ringbuffer
    */
name string

    /**
    * the OverflowPolicy to use.
    */
overflowPolicy int32

    /**
    * to item to add
    */
value serialization.Data
}

func RingbufferAddEncodeRequest(name string, overflowPolicy int32, value serialization.Data) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetOperationName("Ringbuffer.Add")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, RingbufferAddRequestMessageType)
    bufutil.EncodeInt(initialFrame.Content, RingbufferAddRequestOverflowPolicyFieldOffset, overflowPolicy)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, value)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type RingbufferAddResponseParameters struct {
    /**
    * the sequence of the added item, or -1 if the add failed.
    */
response int64
}



func RingbufferAddDecodeResponse(clientMessage *bufutil.ClientMessage) *RingbufferAddResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(RingbufferAddResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeLong(initialFrame.Content, RingbufferAddResponseResponseFieldOffset)
    return response
}

