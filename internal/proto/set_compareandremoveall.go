
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
 * Removes from this set all of its elements that are contained in the specified collection (optional operation).
 * If the specified collection is also a set, this operation effectively modifies this set so that its value is the
 * asymmetric set difference of the two sets.
 */
//@Generated("b1e2bafb4cbc3d0e92000db48482d152")
const (
    //hex: 0x060700
    SetCompareAndRemoveAllRequestMessageType = 395008
    //hex: 0x060701
    SetCompareAndRemoveAllResponseMessageType = 395009
    SetCompareAndRemoveAllRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    SetCompareAndRemoveAllResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    SetCompareAndRemoveAllResponseInitialFrameSize = SetCompareAndRemoveAllResponseResponseFieldOffset + BooleanSizeInBytes

)

func SetCompareAndRemoveAllEncodeRequest(name string, values []serialization.Data) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Set.CompareAndRemoveAll")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, SetCompareAndRemoveAllRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    clientMessage.Add(BeginFrame)
        for i := 0; i < len(values) ; i++ {
            DataCodecEncode(clientMessage, values[i]) //check line
        }
        clientMessage.Add(EndFrame)



    return clientMessage
}


func SetCompareAndRemoveAllDecodeResponse(clientMessage *ClientMessage) func() (/*** true if at least one item in values existed and removed, false otherwise.*/response bool) {
    return func() (/*** true if at least one item in values existed and removed, false otherwise.*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, SetCompareAndRemoveAllResponseResponseFieldOffset)
        return
    }
}

