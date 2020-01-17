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
 * Retains only the elements in this list that are contained in the specified collection (optional operation).
 * In other words, removes from this list all of its elements that are not contained in the specified collection.
 */
//@Generated("54f47fcfeb7f92bfe1366757d2bc836e")
const (
    //hex: 0x050800
    ListCompareAndRetainAllRequestMessageType = 329728
    //hex: 0x050801
    ListCompareAndRetainAllResponseMessageType = 329729
    ListCompareAndRetainAllRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ListCompareAndRetainAllResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    ListCompareAndRetainAllResponseInitialFrameSize = ListCompareAndRetainAllResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListCompareAndRetainAllRequestParameters struct {

    /**
    * Name of the List
    */
name string

    /**
    * The list of values to compare for retaining.
    */
values []serialization.Data
}

func ListCompareAndRetainAllEncodeRequest(name string, values []serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("List.CompareAndRetainAll")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListCompareAndRetainAllRequestMessageType)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.ListMultiFrameCodecEncode(clientMessage, values, bufutil.DataCodecEncode)
    return clientMessage
}




/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListCompareAndRetainAllResponseParameters struct {
    /**
    * True if this list changed as a result of the call, false otherwise.
    */
response bool
}



func ListCompareAndRetainAllDecodeResponse(clientMessage *bufutil.ClientMessagex) *ListCompareAndRetainAllResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ListCompareAndRetainAllResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, ListCompareAndRetainAllResponseResponseFieldOffset)
    return response
}

