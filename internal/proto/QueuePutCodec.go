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
 * Inserts the specified element into this queue, waiting if necessary for space to become available.
 */
//@Generated("680e98ded112c9f2429215a6c7fa7e66")
const (
    //hex: 0x030200
    QueuePutRequestMessageType = 197120
    //hex: 0x030201
    QueuePutResponseMessageType = 197121
    QueuePutRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    QueuePutResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type QueuePutRequestParameters struct {

    /**
    * Name of the Queue
    */
name string

    /**
    * The element to add
    */
value serialization.Data
}

func QueuePutEncodeRequest(name string, value serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Queue.Put")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, QueuePutRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, value)
    return clientMessage
}

func QueuePutDecodeRequest(clientMessage *bufutil.ClientMessagex) *QueuePutRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(QueuePutRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    request.value = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type QueuePutResponseParameters struct {
}

func QueuePutEncodeResponse() *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, QueuePutResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, QueuePutResponseMessageType)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func QueuePutDecodeResponse(clientMessage *bufutil.ClientMessagex) *QueuePutResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(QueuePutResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

