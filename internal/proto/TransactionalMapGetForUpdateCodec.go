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
 * Locks the key and then gets and returns the value to which the specified key is mapped. Lock will be released at
 * the end of the transaction (either commit or rollback).
 */
//@Generated("2df7b1d6d1870ad08f176303187092fb")
const (
    //hex: 0x0E0300
    TransactionalMapGetForUpdateRequestMessageType = 918272
    //hex: 0x0E0301
    TransactionalMapGetForUpdateResponseMessageType = 918273
    TransactionalMapGetForUpdateRequestTxnIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalMapGetForUpdateRequestThreadIdFieldOffset = TransactionalMapGetForUpdateRequestTxnIdFieldOffset + bufutil.UUIDSizeInBytes
    TransactionalMapGetForUpdateRequestInitialFrameSize = TransactionalMapGetForUpdateRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    TransactionalMapGetForUpdateResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMapGetForUpdateRequestParameters struct {

    /**
    * Name of the Transactional Map
    */
name string

    /**
    * ID of the this transaction operation
    */
txnId UUID

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation.
    */
threadId int64

    /**
    * The value to which the specified key is mapped
    */
key serialization.Data
}

func TransactionalMapGetForUpdateEncodeRequest(name string, txnId UUID, threadId int64, key serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("TransactionalMap.GetForUpdate")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalMapGetForUpdateRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, TransactionalMapGetForUpdateRequestTxnIdFieldOffset, txnId)
    bufutil.EncodeLong(initialFrame.Content, TransactionalMapGetForUpdateRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, key)
    return clientMessage
}

func TransactionalMapGetForUpdateDecodeRequest(clientMessage *bufutil.ClientMessagex) *TransactionalMapGetForUpdateRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(TransactionalMapGetForUpdateRequestParameters)
    initialFrame := iterator.Next()
    request.txnId = bufutil.DecodeUUID(initialFrame.Content, TransactionalMapGetForUpdateRequestTxnIdFieldOffset)
    request.threadId = bufutil.DecodeLong(initialFrame.Content, TransactionalMapGetForUpdateRequestThreadIdFieldOffset)
    request.name = StringCodec.decode(iterator)
    request.key = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMapGetForUpdateResponseParameters struct {
    /**
    * The value for the specified key
    */
/* @Nullable */ response serialization.Data
}

func TransactionalMapGetForUpdateEncodeResponse(/* @Nullable */ response serialization.Data ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, TransactionalMapGetForUpdateResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalMapGetForUpdateResponseMessageType)
    clientMessage.Add(initialFrame)

    CodecUtil.encodeNullable(clientMessage, response, DataCodecEncode)
    return clientMessage
}

func TransactionalMapGetForUpdateDecodeResponse(clientMessage *bufutil.ClientMessagex) *TransactionalMapGetForUpdateResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(TransactionalMapGetForUpdateResponseParameters)
    //empty initial frame
    iterator.Next()
    response.response = CodecUtil.decodeNullable(iterator, DataCodecDecode) 
    return response
}

