
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
    "github.com/hazelcast/hazelcast-go-client/core"
)


/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Stops receiving messages for the given message listener.If the given listener already removed, this method does nothing.
 */
//@Generated("031ac56fe34c716e780c39d7bb621a18")
const (
    //hex: 0x040300
    TopicRemoveMessageListenerRequestMessageType = 262912
    //hex: 0x040301
    TopicRemoveMessageListenerResponseMessageType = 262913
    TopicRemoveMessageListenerRequestRegistrationIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    TopicRemoveMessageListenerRequestInitialFrameSize = TopicRemoveMessageListenerRequestRegistrationIdFieldOffset + UUIDSizeInBytes
    TopicRemoveMessageListenerResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    TopicRemoveMessageListenerResponseInitialFrameSize = TopicRemoveMessageListenerResponseResponseFieldOffset + BooleanSizeInBytes

)

func TopicRemoveMessageListenerEncodeRequest(name string, registrationId core.Uuid) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Topic.RemoveMessageListener")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, TopicRemoveMessageListenerRequestMessageType)
    EncodeUUID(initialFrame.Content, TopicRemoveMessageListenerRequestRegistrationIdFieldOffset, registrationId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func TopicRemoveMessageListenerDecodeResponse(clientMessage *ClientMessage) func() ( /*** True if registration is removed, false otherwise*/response bool) {
    return func() (/*** True if registration is removed, false otherwise*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, TopicRemoveMessageListenerResponseResponseFieldOffset)
        return
    }
}

