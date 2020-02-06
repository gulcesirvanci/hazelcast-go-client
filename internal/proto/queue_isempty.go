
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

/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Returns true if this collection contains no elements.
 */
//@Generated("47eca59bc6f84211047c8f0b453bf066")
const (
    //hex: 0x031400
    QueueIsEmptyRequestMessageType = 201728
    //hex: 0x031401
    QueueIsEmptyResponseMessageType = 201729
    QueueIsEmptyRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    QueueIsEmptyResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    QueueIsEmptyResponseInitialFrameSize = QueueIsEmptyResponseResponseFieldOffset + BooleanSizeInBytes

)

func QueueIsEmptyEncodeRequest(name string) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Queue.IsEmpty")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, QueueIsEmptyRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func QueueIsEmptyDecodeResponse(clientMessage *ClientMessage) func() ( /*** <tt>True</tt> if this collection contains no elements*/response bool) {
    return func() (/*** <tt>True</tt> if this collection contains no elements*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, QueueIsEmptyResponseResponseFieldOffset)
        return
    }
}

