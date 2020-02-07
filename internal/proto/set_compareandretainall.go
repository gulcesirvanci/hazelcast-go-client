
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
 * Retains only the elements in this set that are contained in the specified collection (optional operation).
 * In other words, removes from this set all of its elements that are not contained in the specified collection.
 * If the specified collection is also a set, this operation effectively modifies this set so that its value is the
 * intersection of the two sets.
 */
//@Generated("499a8a0eaabf7b93f263e54e4b708e2a")
const (
    //hex: 0x060800
    SetCompareAndRetainAllRequestMessageType = 395264
    //hex: 0x060801
    SetCompareAndRetainAllResponseMessageType = 395265
    SetCompareAndRetainAllRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    SetCompareAndRetainAllResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    SetCompareAndRetainAllResponseInitialFrameSize = SetCompareAndRetainAllResponseResponseFieldOffset + BooleanSizeInBytes

)

func SetCompareAndRetainAllEncodeRequest(name string, values []serialization.Data) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Set.CompareAndRetainAll")
	initialFrame := &Frame{Content: make([]byte, SetCompareAndRetainAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, SetCompareAndRetainAllRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    clientMessage.Add(BeginFrame)
        for i := 0; i < len(values) ; i++ {
            DataCodecEncode(clientMessage, values[i]) //check line
        }
        clientMessage.Add(EndFrame)



    return clientMessage
}


func SetCompareAndRetainAllDecodeResponse(clientMessage *ClientMessage) func() (/*** true if at least one item in values existed and it is retained, false otherwise. All items not in valueSet but* in the Set are removed.*/response bool) {
    return func() (/*** true if at least one item in values existed and it is retained, false otherwise. All items not in valueSet but* in the Set are removed.*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, SetCompareAndRetainAllResponseResponseFieldOffset)
        return
    }
}

