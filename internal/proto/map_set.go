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
 * Puts an entry into this map with a given ttl (time to live) value.Entry will expire and get evicted after the ttl
 * If ttl is 0, then the entry lives forever. Similar to the put operation except that set doesn't
 * return the old value, which is more efficient.
 */
//@Generated("328a297dd5f805adf99438f7e6d18c75")
const (
    //hex: 0x010F00
    MapSetRequestMessageType = 69376
    //hex: 0x010F01
    MapSetResponseMessageType = 69377
    MapSetRequestThreadIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapSetRequestTtlFieldOffset = MapSetRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    MapSetRequestInitialFrameSize = MapSetRequestTtlFieldOffset + bufutil.LongSizeInBytes
    MapSetResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapSetRequestParameters struct {

    /**
    * Name of the map.
    */
name string

    /**
    * Key for the map entry.
    */
key serialization.Data

    /**
    * New value for the map entry.
    */
value serialization.Data

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation.
    */
threadId int64

    /**
    * The duration in milliseconds after which this entry shall be deleted. O means infinite.
    */
ttl int64
}

func MapSetEncodeRequest(name string, key serialization.Data, value serialization.Data, threadId int64, ttl int64) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetOperationName("Map.Set")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapSetRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, MapSetRequestThreadIdFieldOffset, threadId)
    bufutil.EncodeLong(initialFrame.Content, MapSetRequestTtlFieldOffset, ttl)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, key)
    bufutil.DataCodecEncode(clientMessage, value)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapSetResponseParameters struct {
}



func MapSetDecodeResponse(clientMessage *bufutil.ClientMessage) *MapSetResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapSetResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

