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
 * Returns the number of key-value mappings in this map.  If the map contains more than Integer.MAX_VALUE elements,
 * returns Integer.MAX_VALUE
 */
//@Generated("2726d6ee0ae13c6b27017f4a46dc053c")
const (
    //hex: 0x012A00
    MapSizeRequestMessageType = 76288
    //hex: 0x012A01
    MapSizeResponseMessageType = 76289
    MapSizeRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapSizeResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MapSizeResponseInitialFrameSize = MapSizeResponseResponseFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapSizeRequestParameters struct {

    /**
    * of map
    */
name string
}

func MapSizeEncodeRequest(name string) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Map.Size")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapSizeRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func MapSizeDecodeRequest(clientMessage *bufutil.ClientMessagex) *MapSizeRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(MapSizeRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapSizeResponseParameters struct {
    /**
    * the number of key-value mappings in this map
    */
response int
}

func MapSizeEncodeResponse(response int ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, MapSizeResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapSizeResponseMessageType)
    bufutil.EncodeInt(initialFrame.Content, MapSizeResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func MapSizeDecodeResponse(clientMessage *bufutil.ClientMessagex) *MapSizeResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapSizeResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeInt(initialFrame.Content, MapSizeResponseResponseFieldOffset)
    return response
}

