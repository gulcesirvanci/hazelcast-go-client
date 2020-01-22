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
 * Queries the map based on the specified predicate and returns the values of matching entries.Specified predicate
 * runs on all members in parallel. The collection is NOT backed by the map, so changes to the map are NOT reflected
 * in the collection, and vice-versa. This method is always executed by a distributed query, so it may throw
 * a QueryResultSizeExceededException if query result size limit is configured.
 */
//@Generated("50e1e0e32818b6585603b7e998bed757")
const (
    //hex: 0x0E1100
    TransactionalMapValuesWithPredicateRequestMessageType = 921856
    //hex: 0x0E1101
    TransactionalMapValuesWithPredicateResponseMessageType = 921857
    TransactionalMapValuesWithPredicateRequestTxnIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalMapValuesWithPredicateRequestThreadIdFieldOffset = TransactionalMapValuesWithPredicateRequestTxnIdFieldOffset + bufutil.UUIDSizeInBytes
    TransactionalMapValuesWithPredicateRequestInitialFrameSize = TransactionalMapValuesWithPredicateRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    TransactionalMapValuesWithPredicateResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMapValuesWithPredicateRequestParameters struct {

    /**
    * Name of the Transactional Map
    */
name string

    /**
    * ID of the this transaction operation
    */
txnId string

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation.
    */
threadId int64

    /**
    * Specified query criteria.
    */
predicate serialization.Data
}

func TransactionalMapValuesWithPredicateEncodeRequest(name string, txnId string, threadId int64, predicate serialization.Data) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetOperationName("TransactionalMap.ValuesWithPredicate")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalMapValuesWithPredicateRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, TransactionalMapValuesWithPredicateRequestTxnIdFieldOffset, txnId)
    bufutil.EncodeLong(initialFrame.Content, TransactionalMapValuesWithPredicateRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, predicate)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMapValuesWithPredicateResponseParameters struct {
    /**
    * Result value collection of the query.
    */
response []serialization.Data
}



func TransactionalMapValuesWithPredicateDecodeResponse(clientMessage *bufutil.ClientMessage) *TransactionalMapValuesWithPredicateResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(TransactionalMapValuesWithPredicateResponseParameters)
    //empty initial frame
    iterator.Next()
    var result []serialization.Data
        //begin frame, list
        iterator.Next()
        for !bufutil.NextFrameIsDataStructureEndFrame(iterator) {
        result = append(result, bufutil.DataCodecDecode(iterator))
        }
        //end frame, list
        iterator.Next()
        response.response = result
    return response
}

