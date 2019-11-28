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
_"bytes"
"github.com/hazelcast/hazelcast-go-client/serialization"
_ "github.com/hazelcast/hazelcast-go-client"
)

/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Returns the capacity of this Ringbuffer.
 */
//@Generated("46ea9bb07e953ea70a65412f28feb0ea")
const (
    //hex: 0x170400
    RingbufferCapacityRequestMessageType = 1508352
    //hex: 0x170401
    RingbufferCapacityResponseMessageType = 1508353
    RingbufferCapacityRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    RingbufferCapacityResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    RingbufferCapacityResponseInitialFrameSize = RingbufferCapacityResponseResponseFieldOffset + bufutil.LongSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type RingbufferCapacityRequestParameters struct {

    /**
    * Name of the Ringbuffer
    */
name string
}

func RingbufferCapacityEncodeRequest(name string) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Ringbuffer.Capacity")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, RingbufferCapacityRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func RingbufferCapacityDecodeRequest(clientMessage *bufutil.ClientMessagex) *RingbufferCapacityRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(RingbufferCapacityRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type RingbufferCapacityResponseParameters struct {
    /**
    * the capacity
    */
response int64
}

func RingbufferCapacityEncodeResponse(response int64 ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, RingbufferCapacityResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, RingbufferCapacityResponseMessageType)
    bufutil.EncodeLong(initialFrame.Content, RingbufferCapacityResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func RingbufferCapacityDecodeResponse(clientMessage *bufutil.ClientMessagex) *RingbufferCapacityResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(RingbufferCapacityResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeLong(initialFrame.Content, RingbufferCapacityResponseResponseFieldOffset)
    return response
}

