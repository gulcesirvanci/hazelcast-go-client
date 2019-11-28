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
 * Loads all keys into the store. This is a batch load operation so that an implementation can optimize the multiple loads.
 */
//@Generated("fcb7077f72200c3fc9c2e2fc979ad319")
const (
    //hex: 0x012000
    MapLoadAllRequestMessageType = 73728
    //hex: 0x012001
    MapLoadAllResponseMessageType = 73729
    MapLoadAllRequestReplaceExistingValuesFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapLoadAllRequestInitialFrameSize = MapLoadAllRequestReplaceExistingValuesFieldOffset + bufutil.BooleanSizeInBytes
    MapLoadAllResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapLoadAllRequestParameters struct {

    /**
    * name of map
    */
name string

    /**
    * when <code>true</code>, existing values in the Map will
    * be replaced by those loaded from the MapLoader
    */
replaceExistingValues bool
}

func MapLoadAllEncodeRequest(name string, replaceExistingValues bool) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Map.LoadAll")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapLoadAllRequestMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, MapLoadAllRequestReplaceExistingValuesFieldOffset, replaceExistingValues)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func MapLoadAllDecodeRequest(clientMessage *bufutil.ClientMessagex) *MapLoadAllRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(MapLoadAllRequestParameters)
    initialFrame := iterator.Next()
    request.replaceExistingValues = bufutil.DecodeBoolean(initialFrame.Content, MapLoadAllRequestReplaceExistingValuesFieldOffset)
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapLoadAllResponseParameters struct {
}

func MapLoadAllEncodeResponse() *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, MapLoadAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapLoadAllResponseMessageType)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func MapLoadAllDecodeResponse(clientMessage *bufutil.ClientMessagex) *MapLoadAllResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapLoadAllResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

