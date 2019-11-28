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
 * Applies the projection logic on all map entries and returns the result
 */
//@Generated("bf6a5dc1b98ab9195fbe66e495a966b9")
const (
    //hex: 0x013C00
    MapProjectRequestMessageType = 80896
    //hex: 0x013C01
    MapProjectResponseMessageType = 80897
    MapProjectRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapProjectResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapProjectRequestParameters struct {

    /**
    * Name of the map.
    */
name string

    /**
    * projection to transform the entries with. May return null.
    */
projection serialization.Data
}

func MapProjectEncodeRequest(name string, projection serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Map.Project")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapProjectRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, projection)
    return clientMessage
}

func MapProjectDecodeRequest(clientMessage *bufutil.ClientMessagex) *MapProjectRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(MapProjectRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    request.projection = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapProjectResponseParameters struct {
    /**
    * the resulted collection upon transformation to the type of the projection
    */
response []serialization.Data
}

func MapProjectEncodeResponse(response Collection<com.hazelcast.nio.serialization.Data> ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, MapProjectResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapProjectResponseMessageType)
    clientMessage.Add(initialFrame)

    ListMultiFrameCodec.encodeContainsNullable(clientMessage, response, DataCodecEncode)
    return clientMessage
}

func MapProjectDecodeResponse(clientMessage *bufutil.ClientMessagex) *MapProjectResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapProjectResponseParameters)
    //empty initial frame
    iterator.Next()
    response.response = ListMultiFrameCodec.decodeContainsNullable(iterator, DataCodecDecode)
    return response
}

