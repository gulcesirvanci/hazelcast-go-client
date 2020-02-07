
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
)


/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Returns the number of key-value pairs in the multimap.
 */
//@Generated("50b97aa8bb935ea7ef6e8ef86e9fedc1")
const (
    //hex: 0x020A00
    MultiMapSizeRequestMessageType = 133632
    //hex: 0x020A01
    MultiMapSizeResponseMessageType = 133633
    MultiMapSizeRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    MultiMapSizeResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MultiMapSizeResponseInitialFrameSize = MultiMapSizeResponseResponseFieldOffset + IntSizeInBytes

)

func MultiMapSizeEncodeRequest(name string) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("MultiMap.Size")
	initialFrame := &Frame{Content: make([]byte, MultiMapSizeResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MultiMapSizeRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func MultiMapSizeDecodeResponse(clientMessage *ClientMessage) func() (/*** The number of key-value pairs in the multimap.*/response int32) {
    return func() (/*** The number of key-value pairs in the multimap.*/response int32) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeInt(initialFrame.Content, MultiMapSizeResponseResponseFieldOffset)
        return
    }
}

