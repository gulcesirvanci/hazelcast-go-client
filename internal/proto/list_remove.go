
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
 * Removes the first occurrence of the specified element from this list, if it is present (optional operation).
 * If this list does not contain the element, it is unchanged.
 * Returns true if this list contained the specified element (or equivalently, if this list changed as a result of the call).
 */
//@Generated("7ac9c9eecf7f6be955497b886f236be9")
const (
    //hex: 0x050500
    ListRemoveRequestMessageType = 328960
    //hex: 0x050501
    ListRemoveResponseMessageType = 328961
    ListRemoveRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    ListRemoveResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    ListRemoveResponseInitialFrameSize = ListRemoveResponseResponseFieldOffset + BooleanSizeInBytes

)

func ListRemoveEncodeRequest(name string, value serialization.Data) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("List.Remove")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ListRemoveRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, value)

    return clientMessage
}


func ListRemoveDecodeResponse(clientMessage *ClientMessage) func() (/*** True if this list contained the specified element, false otherwise*/response bool) {
    return func() (/*** True if this list contained the specified element, false otherwise*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, ListRemoveResponseResponseFieldOffset)
        return
    }
}

