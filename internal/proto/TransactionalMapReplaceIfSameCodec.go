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
 * Replaces the entry for a key only if currently mapped to a given value. The object to be replaced will be
 * accessible only in the current transaction context until the transaction is committed.
 */
//@Generated("c75152a6d06a132e79de02be3d037850")
const (
    //hex: 0x0E0A00
    TransactionalMapReplaceIfSameRequestMessageType = 920064
    //hex: 0x0E0A01
    TransactionalMapReplaceIfSameResponseMessageType = 920065
    TransactionalMapReplaceIfSameRequestTxnIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalMapReplaceIfSameRequestThreadIdFieldOffset = TransactionalMapReplaceIfSameRequestTxnIdFieldOffset + bufutil.UUIDSizeInBytes
    TransactionalMapReplaceIfSameRequestInitialFrameSize = TransactionalMapReplaceIfSameRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    TransactionalMapReplaceIfSameResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalMapReplaceIfSameResponseInitialFrameSize = TransactionalMapReplaceIfSameResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMapReplaceIfSameRequestParameters struct {

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
    * The specified key.
    */
key serialization.Data

    /**
    * Replace the key value if it is the old value.
    */
oldValue serialization.Data

    /**
    * The new value to replace the old value.
    */
newValue serialization.Data
}

func TransactionalMapReplaceIfSameEncodeRequest(name string, txnId UUID, threadId int64, key serialization.Data, oldValue serialization.Data, newValue serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("TransactionalMap.ReplaceIfSame")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalMapReplaceIfSameRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, TransactionalMapReplaceIfSameRequestTxnIdFieldOffset, txnId)
    bufutil.EncodeLong(initialFrame.Content, TransactionalMapReplaceIfSameRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, key)
    DataCodec.encode(clientMessage, oldValue)
    DataCodec.encode(clientMessage, newValue)
    return clientMessage
}

func TransactionalMapReplaceIfSameDecodeRequest(clientMessage *bufutil.ClientMessagex) *TransactionalMapReplaceIfSameRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(TransactionalMapReplaceIfSameRequestParameters)
    initialFrame := iterator.Next()
    request.txnId = bufutil.DecodeUUID(initialFrame.Content, TransactionalMapReplaceIfSameRequestTxnIdFieldOffset)
    request.threadId = bufutil.DecodeLong(initialFrame.Content, TransactionalMapReplaceIfSameRequestThreadIdFieldOffset)
    request.name = StringCodec.decode(iterator)
    request.key = DataCodec.decode(iterator)
    request.oldValue = DataCodec.decode(iterator)
    request.newValue = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMapReplaceIfSameResponseParameters struct {
    /**
    * true if the value was replaced.
    */
response bool
}

func TransactionalMapReplaceIfSameEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, TransactionalMapReplaceIfSameResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalMapReplaceIfSameResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, TransactionalMapReplaceIfSameResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func TransactionalMapReplaceIfSameDecodeResponse(clientMessage *bufutil.ClientMessagex) *TransactionalMapReplaceIfSameResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(TransactionalMapReplaceIfSameResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, TransactionalMapReplaceIfSameResponseResponseFieldOffset)
    return response
}

