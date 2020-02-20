
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
 * Adds listener for backup acks
 */
//@Generated("980058bb53e7edb2704cef1eb742ffe1")
const (
    //hex: 0x000F00
    ClientLocalBackupListenerRequestMessageType = 3840
    //hex: 0x000F01
    ClientLocalBackupListenerResponseMessageType = 3841
    ClientLocalBackupListenerRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    ClientLocalBackupListenerResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    ClientLocalBackupListenerResponseInitialFrameSize = ClientLocalBackupListenerResponseResponseFieldOffset + UUIDSizeInBytes
    ClientLocalBackupListenerEventBackupSourceInvocationCorrelationIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ClientLocalBackupListenerEventBackupInitialFrameSize = ClientLocalBackupListenerEventBackupSourceInvocationCorrelationIdFieldOffset + LongSizeInBytes
    //hex: 0x000F02
    ClientLocalBackupListenerEventBackupMessageType = 3842


)

func ClientLocalBackupListenerEncodeRequest() *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Client.LocalBackupListener")
	initialFrame := &Frame{Content: make([]byte, ClientLocalBackupListenerResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ClientLocalBackupListenerRequestMessageType)
    clientMessage.Add(initialFrame)
    return clientMessage
}


func ClientLocalBackupListenerDecodeResponse(clientMessage *ClientMessage) func() (/*** Returns the registration id for the listener.*/response *core.Uuid) {
    return func() (/*** Returns the registration id for the listener.*/response *core.Uuid) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeUUID(initialFrame.Content, ClientLocalBackupListenerResponseResponseFieldOffset)
        return
    }
}


type ClientLocalBackupListenerHandleBackupFunc func(sourceInvocationCorrelationId int64)

func ClientLocalBackupListenerHandle(clientMessage *ClientMessage, handlebackup ClientLocalBackupListenerHandleBackupFunc){
    messageType := clientMessage.MessageType()
    iterator := clientMessage.FrameIterator()
    if messageType == ClientLocalBackupListenerEventBackupMessageType {
        initialFrame := iterator.Next()
        sourceInvocationCorrelationId := DecodeLong(initialFrame.Content, ClientLocalBackupListenerEventBackupSourceInvocationCorrelationIdFieldOffset)
        handlebackup(sourceInvocationCorrelationId)
        return
        }
        // FINEST: Logger.getLogger(super.getClass()).finest("Unknown message type received on event handler :" + messageType)
    }
