
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
 * Returns whether the multimap contains the given key-value pair.
 */
//@Generated("f603cb78e1a7a80bd2f4da87b88d9dff")
const (
    //hex: 0x020900
    MultiMapContainsEntryRequestMessageType = 133376
    //hex: 0x020901
    MultiMapContainsEntryResponseMessageType = 133377
    MultiMapContainsEntryRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MultiMapContainsEntryRequestInitialFrameSize = MultiMapContainsEntryRequestThreadIdFieldOffset + LongSizeInBytes
    MultiMapContainsEntryResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MultiMapContainsEntryResponseInitialFrameSize = MultiMapContainsEntryResponseResponseFieldOffset + BooleanSizeInBytes

)

func MultiMapContainsEntryEncodeRequest(name string, key serialization.Data, value serialization.Data, threadId int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("MultiMap.ContainsEntry")
	initialFrame := &Frame{Content: make([]byte, MultiMapContainsEntryResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MultiMapContainsEntryRequestMessageType)
    EncodeLong(initialFrame.Content, MultiMapContainsEntryRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)


    DataCodecEncode(clientMessage, value)

    return clientMessage
}


func MultiMapContainsEntryDecodeResponse(clientMessage *ClientMessage) func() (/*** True if the multimap contains the key-value pair, false otherwise.*/response bool) {
    return func() (/*** True if the multimap contains the key-value pair, false otherwise.*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, MultiMapContainsEntryResponseResponseFieldOffset)
        return
    }
}

