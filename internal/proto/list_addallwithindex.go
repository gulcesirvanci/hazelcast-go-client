
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
 * Inserts all of the elements in the specified collection into this list at the specified position (optional operation).
 * Shifts the element currently at that position (if any) and any subsequent elements to the right (increases their indices).
 * The new elements will appear in this list in the order that they are returned by the specified collection's iterator.
 * The behavior of this operation is undefined if the specified collection is modified while the operation is in progress.
 * (Note that this will occur if the specified collection is this list, and it's nonempty.)
 */
//@Generated("b5f899a7d020514e73705a54085c703a")
const (
    //hex: 0x050E00
    ListAddAllWithIndexRequestMessageType = 331264
    //hex: 0x050E01
    ListAddAllWithIndexResponseMessageType = 331265
    ListAddAllWithIndexRequestIndexFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ListAddAllWithIndexRequestInitialFrameSize = ListAddAllWithIndexRequestIndexFieldOffset + IntSizeInBytes
    ListAddAllWithIndexResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    ListAddAllWithIndexResponseInitialFrameSize = ListAddAllWithIndexResponseResponseFieldOffset + BooleanSizeInBytes

)

func ListAddAllWithIndexEncodeRequest(name string, index int32, valueList []serialization.Data) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("List.AddAllWithIndex")
	initialFrame := &Frame{Content: make([]byte, ListAddAllWithIndexResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ListAddAllWithIndexRequestMessageType)
    EncodeInt(initialFrame.Content, ListAddAllWithIndexRequestIndexFieldOffset, index)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    clientMessage.Add(BeginFrame)
        for i := 0; i < len(valueList) ; i++ {
            DataCodecEncode(clientMessage, valueList[i]) //check line
        }
        clientMessage.Add(EndFrame)



    return clientMessage
}


func ListAddAllWithIndexDecodeResponse(clientMessage *ClientMessage) func() (/*** True if this list changed as a result of the call, false otherwise.*/response bool) {
    return func() (/*** True if this list changed as a result of the call, false otherwise.*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, ListAddAllWithIndexResponseResponseFieldOffset)
        return
    }
}

