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
    "github.com/hazelcast/hazelcast-go-client/serialization"
)

/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Inserts the specified element at the specified position in this list (optional operation). Shifts the element
 * currently at that position (if any) and any subsequent elements to the right (adds one to their indices).
 */
//@Generated("401910d8234530f7e63daa5caf89f157")
const (
    //hex: 0x051100
    ListAddWithIndexRequestMessageType = 332032
    //hex: 0x051101
    ListAddWithIndexResponseMessageType = 332033
    ListAddWithIndexRequestIndexFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ListAddWithIndexRequestInitialFrameSize = ListAddWithIndexRequestIndexFieldOffset + bufutil.IntSizeInBytes
    ListAddWithIndexResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListAddWithIndexRequestParameters struct {

    /**
    * Name of the List
    */
name string

    /**
    * index at which the specified element is to be inserted
    */
index int32

    /**
    * Value to be inserted.
    */
value serialization.Data
}

func ListAddWithIndexEncodeRequest(name string, index int32, value serialization.Data) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetOperationName("List.AddWithIndex")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListAddWithIndexRequestMessageType)
    bufutil.EncodeInt(initialFrame.Content, ListAddWithIndexRequestIndexFieldOffset, index)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, value)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListAddWithIndexResponseParameters struct {
}



func ListAddWithIndexDecodeResponse(clientMessage *bufutil.ClientMessage) *ListAddWithIndexResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ListAddWithIndexResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

