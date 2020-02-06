
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
 * Returns true if this set contains no elements.
 */
//@Generated("19313e21b18e9d52ac9058d5203a4b45")
const (
    //hex: 0x060D00
    SetIsEmptyRequestMessageType = 396544
    //hex: 0x060D01
    SetIsEmptyResponseMessageType = 396545
    SetIsEmptyRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    SetIsEmptyResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    SetIsEmptyResponseInitialFrameSize = SetIsEmptyResponseResponseFieldOffset + BooleanSizeInBytes

)

func SetIsEmptyEncodeRequest(name string) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Set.IsEmpty")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, SetIsEmptyRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func SetIsEmptyDecodeResponse(clientMessage *ClientMessage) func() (/*** True if this set contains no elements*/response bool) {
    return func() (/*** True if this set contains no elements*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, SetIsEmptyResponseResponseFieldOffset)
        return
    }
}

