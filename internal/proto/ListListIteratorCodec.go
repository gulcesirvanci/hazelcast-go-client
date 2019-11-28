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
 * Returns a list iterator over the elements in this list (in proper sequence), starting at the specified position
 * in the list. The specified index indicates the first element that would be returned by an initial call to
 * ListIterator#next next. An initial call to ListIterator#previous previous would return the element with the
 * specified index minus one.
 */
//@Generated("089879ec75abcdf5541ead350dff8aba")
const (
    //hex: 0x051700
    ListListIteratorRequestMessageType = 333568
    //hex: 0x051701
    ListListIteratorResponseMessageType = 333569
    ListListIteratorRequestIndexFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ListListIteratorRequestInitialFrameSize = ListListIteratorRequestIndexFieldOffset + bufutil.IntSizeInBytes
    ListListIteratorResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListListIteratorRequestParameters struct {

    /**
    * Name of the List
    */
name string

    /**
    * index of the first element to be returned from the list iterator next
    */
index int
}

func ListListIteratorEncodeRequest(name string, index int) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("List.ListIterator")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListListIteratorRequestMessageType)
    bufutil.EncodeInt(initialFrame.Content, ListListIteratorRequestIndexFieldOffset, index)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func ListListIteratorDecodeRequest(clientMessage *bufutil.ClientMessagex) *ListListIteratorRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(ListListIteratorRequestParameters)
    initialFrame := iterator.Next()
    request.index = bufutil.DecodeInt(initialFrame.Content, ListListIteratorRequestIndexFieldOffset)
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListListIteratorResponseParameters struct {
    /**
    * a list iterator over the elements in this list (in proper
    * sequence), starting at the specified position in the list
    */
response []serialization.Data
}

func ListListIteratorEncodeResponse(response []serialization.Data ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, ListListIteratorResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListListIteratorResponseMessageType)
    clientMessage.Add(initialFrame)

    ListMultiFrameCodec.encode(clientMessage, response, DataCodecEncode)
    return clientMessage
}

func ListListIteratorDecodeResponse(clientMessage *bufutil.ClientMessagex) *ListListIteratorResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ListListIteratorResponseParameters)
    //empty initial frame
    iterator.Next()
    response.response = ListMultiFrameCodec.decode(iterator, DataCodecDecode)
    return response
}

