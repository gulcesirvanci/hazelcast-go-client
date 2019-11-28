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
 * Inserts the specified element into this queue, waiting up to the specified wait time if necessary for space to
 * become available.
 */
//@Generated("22947584c43d67b3616aab0a8182e08d")
const (
    //hex: 0x030100
    QueueOfferRequestMessageType = 196864
    //hex: 0x030101
    QueueOfferResponseMessageType = 196865
    QueueOfferRequestTimeoutMillisFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    QueueOfferRequestInitialFrameSize = QueueOfferRequestTimeoutMillisFieldOffset + bufutil.LongSizeInBytes
    QueueOfferResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    QueueOfferResponseInitialFrameSize = QueueOfferResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type QueueOfferRequestParameters struct {

    /**
    * Name of the Queue
    */
name string

    /**
    * The element to add
    */
value serialization.Data

    /**
    * Maximum time in milliseconds to wait for acquiring the lock for the key.
    */
timeoutMillis int64
}

func QueueOfferEncodeRequest(name string, value serialization.Data, timeoutMillis int64) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Queue.Offer")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, QueueOfferRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, QueueOfferRequestTimeoutMillisFieldOffset, timeoutMillis)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, value)
    return clientMessage
}

func QueueOfferDecodeRequest(clientMessage *bufutil.ClientMessagex) *QueueOfferRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(QueueOfferRequestParameters)
    initialFrame := iterator.Next()
    request.timeoutMillis = bufutil.DecodeLong(initialFrame.Content, QueueOfferRequestTimeoutMillisFieldOffset)
    request.name = StringCodec.decode(iterator)
    request.value = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type QueueOfferResponseParameters struct {
    /**
    * <tt>True</tt> if the element was added to this queue, else <tt>false</tt>
    */
response bool
}

func QueueOfferEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, QueueOfferResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, QueueOfferResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, QueueOfferResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func QueueOfferDecodeResponse(clientMessage *bufutil.ClientMessagex) *QueueOfferResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(QueueOfferResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, QueueOfferResponseResponseFieldOffset)
    return response
}

