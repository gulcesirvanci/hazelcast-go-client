
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
 * Gets the list of distributed objects in the cluster.
 */
//@Generated("9e6a15cbc66465376edd88342f629c2d")
const (
    //hex: 0x000800
    ClientGetDistributedObjectsRequestMessageType = 2048
    //hex: 0x000801
    ClientGetDistributedObjectsResponseMessageType = 2049
    ClientGetDistributedObjectsRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    ClientGetDistributedObjectsResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func ClientGetDistributedObjectsEncodeRequest() *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Client.GetDistributedObjects")
	initialFrame := &Frame{Content: make([]byte, ClientGetDistributedObjectsResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ClientGetDistributedObjectsRequestMessageType)
    clientMessage.Add(initialFrame)
    return clientMessage
}


func ClientGetDistributedObjectsDecodeResponse(clientMessage *ClientMessage) func() (/*** An array of distributed object info in the cluster.*/response []DistributedObjectInfo) {
    return func() (/*** An array of distributed object info in the cluster.*/response []DistributedObjectInfo) {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        var result []DistributedObjectInfo
        //begin frame, list
        iterator.Next()
        for !NextFrameIsDataStructureEndFrame(iterator) {
        result = append(result, DistributedObjectInfoCodecDecode(iterator))
        }
        //end frame, list
        iterator.Next()
        response = result //0.1
        return
    }
}
