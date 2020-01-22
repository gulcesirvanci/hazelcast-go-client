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
 * If this map has a MapStore, this method flushes all the local dirty entries by calling MapStore.storeAll()
 * and/or MapStore.deleteAll().
 */
//@Generated("a1331fc66899c4e040aa76b12c39f6bb")
const (
    //hex: 0x010A00
    MapFlushRequestMessageType = 68096
    //hex: 0x010A01
    MapFlushResponseMessageType = 68097
    MapFlushRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapFlushResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapFlushRequestParameters struct {

    /**
    * Name of the map.
    */
name string
}

func MapFlushEncodeRequest(name string) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetOperationName("Map.Flush")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapFlushRequestMessageType)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapFlushResponseParameters struct {
}



func MapFlushDecodeResponse(clientMessage *bufutil.ClientMessage) *MapFlushResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapFlushResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

