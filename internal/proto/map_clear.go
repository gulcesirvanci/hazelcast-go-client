
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
 * This method clears the map and invokes MapStore#deleteAll deleteAll on MapStore which, if connected to a database,
 * will delete the records from that database. The MAP_CLEARED event is fired for any registered listeners.
 * To clear a map without calling MapStore#deleteAll, use #evictAll.
 */
//@Generated("4e3d66c96d9a42380749b7d0bbedaa69")
const (
    //hex: 0x012D00
    MapClearRequestMessageType = 77056
    //hex: 0x012D01
    MapClearResponseMessageType = 77057
    MapClearRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    MapClearResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapClearEncodeRequest(name string) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.Clear")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapClearRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func MapClearDecodeResponse(clientMessage *ClientMessage) func() () {
    return func() () {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        return
    }
}

