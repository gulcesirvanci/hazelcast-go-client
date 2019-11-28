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
 * Fetches specified number of entries from the specified partition starting from specified table index.
 */
//@Generated("6bd9871939440b30f12efc0a3865cdb6")
const (
    //hex: 0x013900
    MapFetchEntriesRequestMessageType = 80128
    //hex: 0x013901
    MapFetchEntriesResponseMessageType = 80129
    MapFetchEntriesRequestTableIndexFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapFetchEntriesRequestBatchFieldOffset = MapFetchEntriesRequestTableIndexFieldOffset + bufutil.IntSizeInBytes
    MapFetchEntriesRequestInitialFrameSize = MapFetchEntriesRequestBatchFieldOffset + bufutil.IntSizeInBytes
    MapFetchEntriesResponseTableIndexFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MapFetchEntriesResponseInitialFrameSize = MapFetchEntriesResponseTableIndexFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapFetchEntriesRequestParameters struct {

    /**
    * Name of the map.
    */
name string

    /**
    * The slot number (or index) to start the iterator
    */
tableIndex int

    /**
    * The number of items to be batched
    */
batch int
}

func MapFetchEntriesEncodeRequest(name string, tableIndex int, batch int) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Map.FetchEntries")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapFetchEntriesRequestMessageType)
    bufutil.EncodeInt(initialFrame.Content, MapFetchEntriesRequestTableIndexFieldOffset, tableIndex)
    bufutil.EncodeInt(initialFrame.Content, MapFetchEntriesRequestBatchFieldOffset, batch)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func MapFetchEntriesDecodeRequest(clientMessage *bufutil.ClientMessagex) *MapFetchEntriesRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(MapFetchEntriesRequestParameters)
    initialFrame := iterator.Next()
    request.tableIndex = bufutil.DecodeInt(initialFrame.Content, MapFetchEntriesRequestTableIndexFieldOffset)
    request.batch = bufutil.DecodeInt(initialFrame.Content, MapFetchEntriesRequestBatchFieldOffset)
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapFetchEntriesResponseParameters struct {
    /**
    * The slot number (or index) to start the iterator
    */
tableIndex int
    /**
    * TODO DOC
    */
entries java.util.List<[]serialization.Data, []serialization.Data>
}

func MapFetchEntriesEncodeResponse(tableIndex int , entries []serialization.Data, []serialization.Data ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, MapFetchEntriesResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapFetchEntriesResponseMessageType)
    bufutil.EncodeInt(initialFrame.Content, MapFetchEntriesResponseTableIndexFieldOffset, tableIndex)
    clientMessage.Add(initialFrame)

    EntryListCodec.encode(clientMessage, entries, DataCodecEncode, DataCodecEncode)
    return clientMessage
}

func MapFetchEntriesDecodeResponse(clientMessage *bufutil.ClientMessagex) *MapFetchEntriesResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapFetchEntriesResponseParameters)
    initialFrame := iterator.Next()
    response.tableIndex = bufutil.DecodeInt(initialFrame.Content, MapFetchEntriesResponseTableIndexFieldOffset)
    response.entries = EntryListCodec.decode(iterator, DataCodecDecode, DataCodecDecode)
    return response
}

