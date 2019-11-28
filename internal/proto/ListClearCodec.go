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
 * Removes all of the elements from this list (optional operation). The list will be empty after this call returns.
 */
//@Generated("8b38dee143023087742986ac6dd61720")
const (
    //hex: 0x050900
    ListClearRequestMessageType = 329984
    //hex: 0x050901
    ListClearResponseMessageType = 329985
    ListClearRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ListClearResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListClearRequestParameters struct {

    /**
    * Name of the List
    */
name string
}

func ListClearEncodeRequest(name string) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("List.Clear")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListClearRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func ListClearDecodeRequest(clientMessage *bufutil.ClientMessagex) *ListClearRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(ListClearRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListClearResponseParameters struct {
}

func ListClearEncodeResponse() *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, ListClearResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListClearResponseMessageType)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func ListClearDecodeResponse(clientMessage *bufutil.ClientMessagex) *ListClearResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ListClearResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

