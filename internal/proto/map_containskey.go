
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
 * Returns true if this map contains a mapping for the specified key.
 */
//@Generated("7bb7db526597c1a40948fa5e3b2dd73d")
const (
    //hex: 0x010600
    MapContainsKeyRequestMessageType = 67072
    //hex: 0x010601
    MapContainsKeyResponseMessageType = 67073
    MapContainsKeyRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapContainsKeyRequestInitialFrameSize = MapContainsKeyRequestThreadIdFieldOffset + LongSizeInBytes
    MapContainsKeyResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapContainsKeyResponseInitialFrameSize = MapContainsKeyResponseResponseFieldOffset + BooleanSizeInBytes

)

func MapContainsKeyEncodeRequest(name string, key serialization.Data, threadId int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Map.ContainsKey")
	initialFrame := &Frame{Content: make([]byte, MapContainsKeyResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapContainsKeyRequestMessageType)
    EncodeLong(initialFrame.Content, MapContainsKeyRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)

    return clientMessage
}


func MapContainsKeyDecodeResponse(clientMessage *ClientMessage) func() (/*** Returns true if the key exists, otherwise returns false.*/response bool) {
    return func() (/*** Returns true if the key exists, otherwise returns false.*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, MapContainsKeyResponseResponseFieldOffset)
        return
    }
}

