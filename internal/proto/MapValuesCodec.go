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
 * Returns a collection clone of the values contained in this map.
 * The collection is NOT backed by the map, so changes to the map are NOT reflected in the collection, and vice-versa.
 * This method is always executed by a distributed query, so it may throw a QueryResultSizeExceededException
 * if query result size limit is configured.
 */
//@Generated("ec0233003056f76e9bf819fc02c6ff44")
const (
    //hex: 0x012400
    MapValuesRequestMessageType = 74752
    //hex: 0x012401
    MapValuesResponseMessageType = 74753
    MapValuesRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapValuesResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapValuesRequestParameters struct {

    /**
    * name of map
    */
name string
}

func MapValuesEncodeRequest(name string) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Map.Values")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapValuesRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func MapValuesDecodeRequest(clientMessage *bufutil.ClientMessagex) *MapValuesRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(MapValuesRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapValuesResponseParameters struct {
    /**
    * All values in the map
    */
response []serialization.Data
}

func MapValuesEncodeResponse(response []serialization.Data ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, MapValuesResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapValuesResponseMessageType)
    clientMessage.Add(initialFrame)

    ListMultiFrameCodec.encode(clientMessage, response, DataCodecEncode)
    return clientMessage
}

func MapValuesDecodeResponse(clientMessage *bufutil.ClientMessagex) *MapValuesResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapValuesResponseParameters)
    //empty initial frame
    iterator.Next()
    response.response = ListMultiFrameCodec.decode(iterator, DataCodecDecode)
    return response
}

