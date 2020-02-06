
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
 * If this map has a MapStore, this method flushes all the local dirty entries by calling MapStore.storeAll()
 * and/or MapStore.deleteAll().
 */
//@Generated("5a428aac1ed3097c6936dea5d10cf2df")
const (
    //hex: 0x010A00
    MapFlushRequestMessageType = 68096
    //hex: 0x010A01
    MapFlushResponseMessageType = 68097
    MapFlushRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    MapFlushResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapFlushEncodeRequest(name string) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.Flush")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapFlushRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func MapFlushDecodeResponse(clientMessage *ClientMessage) func() () {
    return func() () {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        return
    }
}

