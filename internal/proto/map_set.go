
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
//@Generated("5e5761ce63e7bb161ff5d947f7f3aa97")
const (
    //hex: 0x010F00
    MapSetRequestMessageType = 69376
    //hex: 0x010F01
    MapSetResponseMessageType = 69377
    MapSetRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapSetRequestTtlFieldOffset = MapSetRequestThreadIdFieldOffset + LongSizeInBytes
    MapSetRequestInitialFrameSize = MapSetRequestTtlFieldOffset + LongSizeInBytes
    MapSetResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapSetEncodeRequest(name string, key serialization.Data, value serialization.Data, threadId int64, ttl int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.Set")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapSetRequestMessageType)
    EncodeLong(initialFrame.Content, MapSetRequestThreadIdFieldOffset, threadId)
    EncodeLong(initialFrame.Content, MapSetRequestTtlFieldOffset, ttl)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)


    DataCodecEncode(clientMessage, value)

    return clientMessage
}


func MapSetDecodeResponse(clientMessage *ClientMessage) func() () {
    return func() () {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        return
    }
}

