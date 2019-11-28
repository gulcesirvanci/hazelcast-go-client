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
 * Returns true if this list contains no elements
 */
//@Generated("c6e94827f24d54e4726cf8a963bafa78")
const (
    //hex: 0x050D00
    ListIsEmptyRequestMessageType = 331008
    //hex: 0x050D01
    ListIsEmptyResponseMessageType = 331009
    ListIsEmptyRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ListIsEmptyResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    ListIsEmptyResponseInitialFrameSize = ListIsEmptyResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListIsEmptyRequestParameters struct {

    /**
    * Name of the List
    */
name string
}

func ListIsEmptyEncodeRequest(name string) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("List.IsEmpty")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListIsEmptyRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func ListIsEmptyDecodeRequest(clientMessage *bufutil.ClientMessagex) *ListIsEmptyRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(ListIsEmptyRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListIsEmptyResponseParameters struct {
    /**
    * True if this list contains no elements
    */
response bool
}

func ListIsEmptyEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, ListIsEmptyResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListIsEmptyResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, ListIsEmptyResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func ListIsEmptyDecodeResponse(clientMessage *bufutil.ClientMessagex) *ListIsEmptyResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ListIsEmptyResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, ListIsEmptyResponseResponseFieldOffset)
    return response
}

