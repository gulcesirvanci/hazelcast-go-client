
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
 * Removes all the entries with the given key. The collection is NOT backed by the map, so changes to the map are
 * NOT reflected in the collection, and vice-versa.
 */
//@Generated("5d1abc260da85216ec7e4ee8623aacf2")
const (
    //hex: 0x021500
    MultiMapRemoveEntryRequestMessageType = 136448
    //hex: 0x021501
    MultiMapRemoveEntryResponseMessageType = 136449
    MultiMapRemoveEntryRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MultiMapRemoveEntryRequestInitialFrameSize = MultiMapRemoveEntryRequestThreadIdFieldOffset + LongSizeInBytes
    MultiMapRemoveEntryResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MultiMapRemoveEntryResponseInitialFrameSize = MultiMapRemoveEntryResponseResponseFieldOffset + BooleanSizeInBytes

)

func MultiMapRemoveEntryEncodeRequest(name string, key serialization.Data, value serialization.Data, threadId int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("MultiMap.RemoveEntry")
	initialFrame := &Frame{Content: make([]byte, MultiMapRemoveEntryResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MultiMapRemoveEntryRequestMessageType)
    EncodeLong(initialFrame.Content, MultiMapRemoveEntryRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)


    DataCodecEncode(clientMessage, value)

    return clientMessage
}


func MultiMapRemoveEntryDecodeResponse(clientMessage *ClientMessage) func() (/*** True if the size of the multimap changed after the remove operation, false otherwise.*/response bool) {
    return func() (/*** True if the size of the multimap changed after the remove operation, false otherwise.*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, MultiMapRemoveEntryResponseResponseFieldOffset)
        return
    }
}

