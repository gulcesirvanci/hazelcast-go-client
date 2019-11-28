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
 * This method can be used to recover from a possible event loss situation.
 * This method tries to make consistent the data in this `QueryCache` with the data in the underlying `IMap`
 * by replaying the events after last consistently received ones. As a result of this replaying logic, same event may
 * appear more than once to the `QueryCache` listeners.
 * This method returns `false` if the event is not in the buffer of event publisher side. That means recovery is not
 * possible.
 */
//@Generated("9d9c8998e1c1ad4789a0ad9674b4e91d")
const (
    //hex: 0x160500
    ContinuousQuerySetReadCursorRequestMessageType = 1443072
    //hex: 0x160501
    ContinuousQuerySetReadCursorResponseMessageType = 1443073
    ContinuousQuerySetReadCursorRequestSequenceFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ContinuousQuerySetReadCursorRequestInitialFrameSize = ContinuousQuerySetReadCursorRequestSequenceFieldOffset + bufutil.LongSizeInBytes
    ContinuousQuerySetReadCursorResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    ContinuousQuerySetReadCursorResponseInitialFrameSize = ContinuousQuerySetReadCursorResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ContinuousQuerySetReadCursorRequestParameters struct {

    /**
    * Name of the map.
    */
mapName string

    /**
    * Name of query cache.
    */
cacheName string

    /**
    * The cursor position of the accumulator to be set.
    */
sequence int64
}

func ContinuousQuerySetReadCursorEncodeRequest(mapName string, cacheName string, sequence int64) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("ContinuousQuery.SetReadCursor")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ContinuousQuerySetReadCursorRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, ContinuousQuerySetReadCursorRequestSequenceFieldOffset, sequence)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, mapName)
    StringCodec.encode(clientMessage, cacheName)
    return clientMessage
}

func ContinuousQuerySetReadCursorDecodeRequest(clientMessage *bufutil.ClientMessagex) *ContinuousQuerySetReadCursorRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(ContinuousQuerySetReadCursorRequestParameters)
    initialFrame := iterator.Next()
    request.sequence = bufutil.DecodeLong(initialFrame.Content, ContinuousQuerySetReadCursorRequestSequenceFieldOffset)
    request.mapName = StringCodec.decode(iterator)
    request.cacheName = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ContinuousQuerySetReadCursorResponseParameters struct {
    /**
    * True if the cursor position could be set, false otherwise.
    */
response bool
}

func ContinuousQuerySetReadCursorEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, ContinuousQuerySetReadCursorResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ContinuousQuerySetReadCursorResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, ContinuousQuerySetReadCursorResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func ContinuousQuerySetReadCursorDecodeResponse(clientMessage *bufutil.ClientMessagex) *ContinuousQuerySetReadCursorResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ContinuousQuerySetReadCursorResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, ContinuousQuerySetReadCursorResponseResponseFieldOffset)
    return response
}

