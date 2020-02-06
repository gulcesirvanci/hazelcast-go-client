
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
 * Stores a key-value pair in the multimap.
 */
//@Generated("9c744903a08e78e9510ecfd97ecdc31d")
const (
    //hex: 0x020100
    MultiMapPutRequestMessageType = 131328
    //hex: 0x020101
    MultiMapPutResponseMessageType = 131329
    MultiMapPutRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MultiMapPutRequestInitialFrameSize = MultiMapPutRequestThreadIdFieldOffset + LongSizeInBytes
    MultiMapPutResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MultiMapPutResponseInitialFrameSize = MultiMapPutResponseResponseFieldOffset + BooleanSizeInBytes

)

func MultiMapPutEncodeRequest(name string, key serialization.Data, value serialization.Data, threadId int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("MultiMap.Put")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MultiMapPutRequestMessageType)
    EncodeLong(initialFrame.Content, MultiMapPutRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)


    DataCodecEncode(clientMessage, value)

    return clientMessage
}


func MultiMapPutDecodeResponse(clientMessage *ClientMessage) func() ( /*** True if size of the multimap is increased, false if the multimap already contains the key-value pair.*/response bool) {
    return func() (/*** True if size of the multimap is increased, false if the multimap already contains the key-value pair.*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, MultiMapPutResponseResponseFieldOffset)
        return
    }
}

