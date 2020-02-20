
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
 * Updates TTL (time to live) value of the entry specified by {@code key} with a new TTL value.
 * New TTL value is valid from this operation is invoked, not from the original creation of the entry.
 * If the entry does not exist or already expired, then this call has no effect.
 * <p>
 * The entry will expire and get evicted after the TTL. If the TTL is 0,
 * then the entry lives forever. If the TTL is negative, then the TTL
 * from the map configuration will be used (default: forever).
 * 
 * If there is no entry with key {@code key}, this call has no effect.
 * 
 * <b>Warning:</b>
 * <p>
 * Time resolution for TTL is seconds. The given TTL value is rounded to the next closest second value.
 */
//@Generated("6b12b9ab22ad15164495b7da94728b98")
const (
    //hex: 0x014300
    MapSetTtlRequestMessageType = 82688
    //hex: 0x014301
    MapSetTtlResponseMessageType = 82689
    MapSetTtlRequestTtlFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapSetTtlRequestInitialFrameSize = MapSetTtlRequestTtlFieldOffset + LongSizeInBytes
    MapSetTtlResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapSetTtlResponseInitialFrameSize = MapSetTtlResponseResponseFieldOffset + BooleanSizeInBytes

)

func MapSetTtlEncodeRequest(name string, key serialization.Data, ttl int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.SetTtl")
	initialFrame := &Frame{Content: make([]byte, MapSetTtlResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapSetTtlRequestMessageType)
    EncodeLong(initialFrame.Content, MapSetTtlRequestTtlFieldOffset, ttl)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)

    return clientMessage
}


func MapSetTtlDecodeResponse(clientMessage *ClientMessage) func() (/*** 'true' if the entry is affected, 'false' otherwise*/response bool) {
    return func() (/*** 'true' if the entry is affected, 'false' otherwise*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, MapSetTtlResponseResponseFieldOffset)
        return
    }
}
