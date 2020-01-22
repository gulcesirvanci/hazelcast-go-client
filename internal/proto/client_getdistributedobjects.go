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
    "github.com/hazelcast/hazelcast-go-client/internal/proto/bufutil"
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
//@Generated("f2cd657c71d046fe3a48f268166acea0")
const (
    //hex: 0x000800
    ClientGetDistributedObjectsRequestMessageType = 2048
    //hex: 0x000801
    ClientGetDistributedObjectsResponseMessageType = 2049
    ClientGetDistributedObjectsRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ClientGetDistributedObjectsResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ClientGetDistributedObjectsRequestParameters struct {
}

func ClientGetDistributedObjectsEncodeRequest() *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetOperationName("Client.GetDistributedObjects")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ClientGetDistributedObjectsRequestMessageType)
    clientMessage.Add(initialFrame)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ClientGetDistributedObjectsResponseParameters struct {
    /**
    * An array of distributed object info in the cluster.
    */
response []DistributedObjectInfo
}



func ClientGetDistributedObjectsDecodeResponse(clientMessage *bufutil.ClientMessage) *ClientGetDistributedObjectsResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ClientGetDistributedObjectsResponseParameters)
    //empty initial frame
    iterator.Next()
    var result []DistributedObjectInfo
        //begin frame, list
        iterator.Next()
        for !bufutil.NextFrameIsDataStructureEndFrame(iterator) {
        result = append(result, bufutil.DistributedObjectInfoCodecDecode(iterator))
        }
        //end frame, list
        iterator.Next()
        response.response = result
    return response
}

