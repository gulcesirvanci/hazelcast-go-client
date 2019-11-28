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
 * Removes the first occurrence of the specified element from this list, if it is present (optional operation).
 * If this list does not contain the element, it is unchanged.
 * Returns true if this list contained the specified element (or equivalently, if this list changed as a result of the call).
 */
//@Generated("121562906a296afd900b00b0a0033dda")
const (
    //hex: 0x050500
    ListRemoveRequestMessageType = 328960
    //hex: 0x050501
    ListRemoveResponseMessageType = 328961
    ListRemoveRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ListRemoveResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    ListRemoveResponseInitialFrameSize = ListRemoveResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListRemoveRequestParameters struct {

    /**
    * Name of the List
    */
name string

    /**
    * Element to be removed from this list, if present
    */
value serialization.Data
}

func ListRemoveEncodeRequest(name string, value serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("List.Remove")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListRemoveRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, value)
    return clientMessage
}

func ListRemoveDecodeRequest(clientMessage *bufutil.ClientMessagex) *ListRemoveRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(ListRemoveRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    request.value = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListRemoveResponseParameters struct {
    /**
    * True if this list contained the specified element, false otherwise
    */
response bool
}

func ListRemoveEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, ListRemoveResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListRemoveResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, ListRemoveResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func ListRemoveDecodeResponse(clientMessage *bufutil.ClientMessagex) *ListRemoveResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ListRemoveResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, ListRemoveResponseResponseFieldOffset)
    return response
}

