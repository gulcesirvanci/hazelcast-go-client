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
 * Evicts all keys from this map except the locked ones. If a MapStore is defined for this map, deleteAll is not
 * called by this method. If you do want to deleteAll to be called use the clear method. The EVICT_ALL event is
 * fired for any registered listeners.
 */
//@Generated("1099a25cb3c3740fe85ee8ce92612e91")
const (
    //hex: 0x011F00
    MapEvictAllRequestMessageType = 73472
    //hex: 0x011F01
    MapEvictAllResponseMessageType = 73473
    MapEvictAllRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapEvictAllResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapEvictAllRequestParameters struct {

    /**
    * name of map
    */
name string
}

func MapEvictAllEncodeRequest(name string) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Map.EvictAll")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapEvictAllRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func MapEvictAllDecodeRequest(clientMessage *bufutil.ClientMessagex) *MapEvictAllRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(MapEvictAllRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapEvictAllResponseParameters struct {
}

func MapEvictAllEncodeResponse() *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, MapEvictAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapEvictAllResponseMessageType)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func MapEvictAllDecodeResponse(clientMessage *bufutil.ClientMessagex) *MapEvictAllResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapEvictAllResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

