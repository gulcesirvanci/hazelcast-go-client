
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
 * Evicts the specified key from this map. If a MapStore is defined for this map, then the entry is not deleted
 * from the underlying MapStore, evict only removes the entry from the memory.
 */
//@Generated("5c73f89ec56a9d4b5ae8a1dbfb410a13")
const (
    //hex: 0x011E00
    MapEvictRequestMessageType = 73216
    //hex: 0x011E01
    MapEvictResponseMessageType = 73217
    MapEvictRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapEvictRequestInitialFrameSize = MapEvictRequestThreadIdFieldOffset + LongSizeInBytes
    MapEvictResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapEvictResponseInitialFrameSize = MapEvictResponseResponseFieldOffset + BooleanSizeInBytes

)

func MapEvictEncodeRequest(name string, key serialization.Data, threadId int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.Evict")
	initialFrame := &Frame{Content: make([]byte, MapEvictResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapEvictRequestMessageType)
    EncodeLong(initialFrame.Content, MapEvictRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)

    return clientMessage
}


func MapEvictDecodeResponse(clientMessage *ClientMessage) func() (/*** true if the key is evicted, false otherwise.*/response bool) {
    return func() (/*** true if the key is evicted, false otherwise.*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, MapEvictResponseResponseFieldOffset)
        return
    }
}

