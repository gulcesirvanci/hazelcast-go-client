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
 * Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not
 * contain the element.
 */
//@Generated("51e4a45150d9f988141c27634d0fbcfa")
const (
    //hex: 0x051300
    ListLastIndexOfRequestMessageType = 332544
    //hex: 0x051301
    ListLastIndexOfResponseMessageType = 332545
    ListLastIndexOfRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ListLastIndexOfResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    ListLastIndexOfResponseInitialFrameSize = ListLastIndexOfResponseResponseFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListLastIndexOfRequestParameters struct {

    /**
    * Name of the List
    */
name string

    /**
    * Element to search for
    */
value serialization.Data
}

func ListLastIndexOfEncodeRequest(name string, value serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("List.LastIndexOf")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListLastIndexOfRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, value)
    return clientMessage
}

func ListLastIndexOfDecodeRequest(clientMessage *bufutil.ClientMessagex) *ListLastIndexOfRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(ListLastIndexOfRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    request.value = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListLastIndexOfResponseParameters struct {
    /**
    * the index of the last occurrence of the specified element in
    * this list, or -1 if this list does not contain the element
    */
response int
}

func ListLastIndexOfEncodeResponse(response int ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, ListLastIndexOfResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListLastIndexOfResponseMessageType)
    bufutil.EncodeInt(initialFrame.Content, ListLastIndexOfResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func ListLastIndexOfDecodeResponse(clientMessage *bufutil.ClientMessagex) *ListLastIndexOfResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ListLastIndexOfResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeInt(initialFrame.Content, ListLastIndexOfResponseResponseFieldOffset)
    return response
}

