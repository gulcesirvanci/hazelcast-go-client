
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
 * Tries to remove the entry with the given key from this map within the specified timeout value.
 * If the key is already locked by another thread and/or member, then this operation will wait the timeout
 * amount for acquiring the lock.
 */
//@Generated("6a966fac02b57d7df30600f705347972")
const (
    //hex: 0x010B00
    MapTryRemoveRequestMessageType = 68352
    //hex: 0x010B01
    MapTryRemoveResponseMessageType = 68353
    MapTryRemoveRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapTryRemoveRequestTimeoutFieldOffset = MapTryRemoveRequestThreadIdFieldOffset + LongSizeInBytes
    MapTryRemoveRequestInitialFrameSize = MapTryRemoveRequestTimeoutFieldOffset + LongSizeInBytes
    MapTryRemoveResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapTryRemoveResponseInitialFrameSize = MapTryRemoveResponseResponseFieldOffset + BooleanSizeInBytes

)

func MapTryRemoveEncodeRequest(name string, key serialization.Data, threadId int64, timeout int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.TryRemove")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapTryRemoveRequestMessageType)
    EncodeLong(initialFrame.Content, MapTryRemoveRequestThreadIdFieldOffset, threadId)
    EncodeLong(initialFrame.Content, MapTryRemoveRequestTimeoutFieldOffset, timeout)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)

    return clientMessage
}


func MapTryRemoveDecodeResponse(clientMessage *ClientMessage) func() ( /*** Returns true if successful, otherwise returns false*/response bool) {
    return func() (/*** Returns true if successful, otherwise returns false*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, MapTryRemoveResponseResponseFieldOffset)
        return
    }
}

