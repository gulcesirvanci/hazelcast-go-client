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
 * Returns true if this map contains a mapping for the specified key.
 */
//@Generated("a4a227c80d1bf8acc77a0b05da21e75f")
const (
    //hex: 0x010600
    MapContainsKeyRequestMessageType = 67072
    //hex: 0x010601
    MapContainsKeyResponseMessageType = 67073
    MapContainsKeyRequestThreadIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapContainsKeyRequestInitialFrameSize = MapContainsKeyRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    MapContainsKeyResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MapContainsKeyResponseInitialFrameSize = MapContainsKeyResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapContainsKeyRequestParameters struct {

    /**
    * Name of the map.
    */
name string

    /**
    * Key for the map entry.
    */
key serialization.Data

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation.
    */
threadId int64
}

func MapContainsKeyEncodeRequest(name string, key serialization.Data, threadId int64) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = true
    clientMessage.SetOperationName("Map.ContainsKey")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapContainsKeyRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, MapContainsKeyRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, key)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapContainsKeyResponseParameters struct {
    /**
    * Returns true if the key exists, otherwise returns false.
    */
response bool
}



func MapContainsKeyDecodeResponse(clientMessage *bufutil.ClientMessage) *MapContainsKeyResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapContainsKeyResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, MapContainsKeyResponseResponseFieldOffset)
    return response
}

