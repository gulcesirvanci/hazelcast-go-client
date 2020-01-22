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
 * Returns the number of key-value pairs in the multimap.
 */
//@Generated("3334645f0135cb5c019369c15619d70b")
const (
    //hex: 0x020A00
    MultiMapSizeRequestMessageType = 133632
    //hex: 0x020A01
    MultiMapSizeResponseMessageType = 133633
    MultiMapSizeRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MultiMapSizeResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MultiMapSizeResponseInitialFrameSize = MultiMapSizeResponseResponseFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapSizeRequestParameters struct {

    /**
    * Name of the MultiMap
    */
name string
}

func MultiMapSizeEncodeRequest(name string) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = true
    clientMessage.SetOperationName("MultiMap.Size")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MultiMapSizeRequestMessageType)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapSizeResponseParameters struct {
    /**
    * The number of key-value pairs in the multimap.
    */
response int32
}



func MultiMapSizeDecodeResponse(clientMessage *bufutil.ClientMessage) *MultiMapSizeResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MultiMapSizeResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeInt(initialFrame.Content, MultiMapSizeResponseResponseFieldOffset)
    return response
}

