
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
 * Removes the mapping for a key from this map if it is present (optional operation).
 * Returns the value to which this map previously associated the key, or null if the map contained no mapping for the key.
 * If this map permits null values, then a return value of null does not necessarily indicate that the map contained no mapping for the key; it's also
 * possible that the map explicitly mapped the key to null. The map will not contain a mapping for the specified key once the
 * call returns.
 */
//@Generated("89abbb5182e26b17af63fe43aba80880")
const (
    //hex: 0x010300
    MapRemoveRequestMessageType = 66304
    //hex: 0x010301
    MapRemoveResponseMessageType = 66305
    MapRemoveRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapRemoveRequestInitialFrameSize = MapRemoveRequestThreadIdFieldOffset + LongSizeInBytes
    MapRemoveResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapRemoveEncodeRequest(name string, key serialization.Data, threadId int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.Remove")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapRemoveRequestMessageType)
    EncodeLong(initialFrame.Content, MapRemoveRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)

    return clientMessage
}


func MapRemoveDecodeResponse(clientMessage *ClientMessage) func() ( /*** Clone of the removed value, not the original (identically equal) value previously put into the map.*//* @Nullable */response serialization.Data) {
    return func() (/*** Clone of the removed value, not the original (identically equal) value previously put into the map.*//* @Nullable */response serialization.Data) {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        response = DecodeNullable(iterator, DataCodecDecode).(serialization.Data)
        return
    }
}

