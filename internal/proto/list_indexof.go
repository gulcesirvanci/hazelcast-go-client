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
 * Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not
 * contain the element.
 */
//@Generated("3ac9201688dda547b750080df5c20dc4")
const (
    //hex: 0x051400
    ListIndexOfRequestMessageType = 332800
    //hex: 0x051401
    ListIndexOfResponseMessageType = 332801
    ListIndexOfRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ListIndexOfResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    ListIndexOfResponseInitialFrameSize = ListIndexOfResponseResponseFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListIndexOfRequestParameters struct {

    /**
    * Name of the List
    */
name string

    /**
    * Element to search for
    */
value serialization.Data
}

func ListIndexOfEncodeRequest(name string, value serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("List.IndexOf")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ListIndexOfRequestMessageType)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, value)
    return clientMessage
}




/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ListIndexOfResponseParameters struct {
    /**
    * The index of the first occurrence of the specified element in
    * this list, or -1 if this list does not contain the element
    */
response int32
}



func ListIndexOfDecodeResponse(clientMessage *bufutil.ClientMessagex) *ListIndexOfResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ListIndexOfResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeInt(initialFrame.Content, ListIndexOfResponseResponseFieldOffset)
    return response
}

