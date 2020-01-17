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
_"bytes"
"github.com/hazelcast/hazelcast-go-client/serialization"
_ "github.com/hazelcast/hazelcast-go-client"
)

/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Return the all elements of this collection
 */
//@Generated("c257fb0cd3ab247ac97cd26945a28893")
const (
    //hex: 0x050A00
    ListGetAllRequestMessageType = 330240
    //hex: 0x050A01
    ListGetAllResponseMessageType = 330241
    ListGetAllRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ListGetAllResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListGetAllRequestParameters struct {

    /**
    * Name of the List
    */
name string
}

func ListGetAllEncodeRequest(name string) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("List.GetAll")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListGetAllRequestMessageType)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    return clientMessage
}




/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListGetAllResponseParameters struct {
    /**
    * An array of all item values in the list.
    */
response []serialization.Data
}



func ListGetAllDecodeResponse(clientMessage *bufutil.ClientMessagex) *ListGetAllResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ListGetAllResponseParameters)
    //empty initial frame
    iterator.Next()
    response.response = bufutil.ListMultiFrameCodecDecode(iterator, bufutil.DataCodecDecode)
    return response
}

