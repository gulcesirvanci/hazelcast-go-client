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
 * Removes all of the mappings from this cache. The order that the individual entries are removed is undefined.
 * For every mapping that exists the following are called: any registered CacheEntryRemovedListener if the cache is
 * a write-through cache, the CacheWriter.If the cache is empty, the CacheWriter is not called.
 * This is potentially an expensive operation as listeners are invoked. Use  #clear() to avoid this.
 */
//@Generated("ac5ee66b423a5145b03d2b87a44561dd")
const (
    //hex: 0x130400
    CacheRemoveAllRequestMessageType = 1246208
    //hex: 0x130401
    CacheRemoveAllResponseMessageType = 1246209
    CacheRemoveAllRequestCompletionIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    CacheRemoveAllRequestInitialFrameSize = CacheRemoveAllRequestCompletionIdFieldOffset + bufutil.IntSizeInBytes
    CacheRemoveAllResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type CacheRemoveAllRequestParameters struct {

    /**
    * Name of the cache.
    */
name string

    /**
    * User generated id which shall be received as a field of the cache event upon completion of
    * the request in the cluster.
    */
completionId int
}

func CacheRemoveAllEncodeRequest(name string, completionId int) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Cache.RemoveAll")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, CacheRemoveAllRequestMessageType)
    bufutil.EncodeInt(initialFrame.Content, CacheRemoveAllRequestCompletionIdFieldOffset, completionId)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func CacheRemoveAllDecodeRequest(clientMessage *bufutil.ClientMessagex) *CacheRemoveAllRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(CacheRemoveAllRequestParameters)
    initialFrame := iterator.Next()
    request.completionId = bufutil.DecodeInt(initialFrame.Content, CacheRemoveAllRequestCompletionIdFieldOffset)
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type CacheRemoveAllResponseParameters struct {
}

func CacheRemoveAllEncodeResponse() *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, CacheRemoveAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, CacheRemoveAllResponseMessageType)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func CacheRemoveAllDecodeResponse(clientMessage *bufutil.ClientMessagex) *CacheRemoveAllResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(CacheRemoveAllResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

