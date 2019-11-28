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
 * Returns the number of key-value mappings in this map. If the map contains more than Integer.MAX_VALUE elements,
 * returns Integer.MAX_VALUE.
 */
//@Generated("bbe049b0c515af7de9c1f97b3dbc2f4f")
const (
    //hex: 0x0D0200
    ReplicatedMapSizeRequestMessageType = 852480
    //hex: 0x0D0201
    ReplicatedMapSizeResponseMessageType = 852481
    ReplicatedMapSizeRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ReplicatedMapSizeResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    ReplicatedMapSizeResponseInitialFrameSize = ReplicatedMapSizeResponseResponseFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ReplicatedMapSizeRequestParameters struct {

    /**
    * Name of the ReplicatedMap
    */
name string
}

func ReplicatedMapSizeEncodeRequest(name string) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("ReplicatedMap.Size")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ReplicatedMapSizeRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func ReplicatedMapSizeDecodeRequest(clientMessage *bufutil.ClientMessagex) *ReplicatedMapSizeRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(ReplicatedMapSizeRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ReplicatedMapSizeResponseParameters struct {
    /**
    * the number of key-value mappings in this map.
    */
response int
}

func ReplicatedMapSizeEncodeResponse(response int ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, ReplicatedMapSizeResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ReplicatedMapSizeResponseMessageType)
    bufutil.EncodeInt(initialFrame.Content, ReplicatedMapSizeResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func ReplicatedMapSizeDecodeResponse(clientMessage *bufutil.ClientMessagex) *ReplicatedMapSizeResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ReplicatedMapSizeResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeInt(initialFrame.Content, ReplicatedMapSizeResponseResponseFieldOffset)
    return response
}

