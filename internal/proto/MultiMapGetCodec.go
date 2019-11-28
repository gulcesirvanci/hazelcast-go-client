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
 * Returns the collection of values associated with the key. The collection is NOT backed by the map, so changes to
 * the map are NOT reflected in the collection, and vice-versa.
 */
//@Generated("8a1e8d0716155643c45763e14ace444e")
const (
    //hex: 0x020200
    MultiMapGetRequestMessageType = 131584
    //hex: 0x020201
    MultiMapGetResponseMessageType = 131585
    MultiMapGetRequestThreadIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MultiMapGetRequestInitialFrameSize = MultiMapGetRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    MultiMapGetResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapGetRequestParameters struct {

    /**
    * Name of the MultiMap
    */
name string

    /**
    * The key whose associated values are to be returned
    */
key serialization.Data

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation.
    */
threadId int64
}

func MultiMapGetEncodeRequest(name string, key serialization.Data, threadId int64) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("MultiMap.Get")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MultiMapGetRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, MultiMapGetRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, key)
    return clientMessage
}

func MultiMapGetDecodeRequest(clientMessage *bufutil.ClientMessagex) *MultiMapGetRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(MultiMapGetRequestParameters)
    initialFrame := iterator.Next()
    request.threadId = bufutil.DecodeLong(initialFrame.Content, MultiMapGetRequestThreadIdFieldOffset)
    request.name = StringCodec.decode(iterator)
    request.key = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapGetResponseParameters struct {
    /**
    * The collection of the values associated with the key.
    */
response []serialization.Data
}

func MultiMapGetEncodeResponse(response []serialization.Data ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, MultiMapGetResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MultiMapGetResponseMessageType)
    clientMessage.Add(initialFrame)

    ListMultiFrameCodec.encode(clientMessage, response, DataCodecEncode)
    return clientMessage
}

func MultiMapGetDecodeResponse(clientMessage *bufutil.ClientMessagex) *MultiMapGetResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MultiMapGetResponseParameters)
    //empty initial frame
    iterator.Next()
    response.response = ListMultiFrameCodec.decode(iterator, DataCodecDecode)
    return response
}

