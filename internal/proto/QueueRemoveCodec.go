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
 * Retrieves and removes the head of this queue.  This method differs from poll only in that it throws an exception
 * if this queue is empty.
 */
//@Generated("92d8d3d4b8c396ecb55830b58e8570e0")
const (
    //hex: 0x030400
    QueueRemoveRequestMessageType = 197632
    //hex: 0x030401
    QueueRemoveResponseMessageType = 197633
    QueueRemoveRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    QueueRemoveResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    QueueRemoveResponseInitialFrameSize = QueueRemoveResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type QueueRemoveRequestParameters struct {

    /**
    * Name of the Queue
    */
name string

    /**
    * Element to be removed from this queue, if present
    */
value serialization.Data
}

func QueueRemoveEncodeRequest(name string, value serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Queue.Remove")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, QueueRemoveRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, value)
    return clientMessage
}

func QueueRemoveDecodeRequest(clientMessage *bufutil.ClientMessagex) *QueueRemoveRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(QueueRemoveRequestParameters)
    //empty initial frame
    iterator.Next()
    request.name = StringCodec.decode(iterator)
    request.value = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type QueueRemoveResponseParameters struct {
    /**
    * <tt>true</tt> if this queue changed as a result of the call
    */
response bool
}

func QueueRemoveEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, QueueRemoveResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, QueueRemoveResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, QueueRemoveResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func QueueRemoveDecodeResponse(clientMessage *bufutil.ClientMessagex) *QueueRemoveResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(QueueRemoveResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, QueueRemoveResponseResponseFieldOffset)
    return response
}

