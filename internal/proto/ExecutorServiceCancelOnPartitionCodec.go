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
 * TODO DOC
 */
//@Generated("bef896e11cc57e4d38c120e1954074bc")
const (
    //hex: 0x080300
    ExecutorServiceCancelOnPartitionRequestMessageType = 525056
    //hex: 0x080301
    ExecutorServiceCancelOnPartitionResponseMessageType = 525057
    ExecutorServiceCancelOnPartitionRequestUuidFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ExecutorServiceCancelOnPartitionRequestInterruptFieldOffset = ExecutorServiceCancelOnPartitionRequestUuidFieldOffset + bufutil.UUIDSizeInBytes
    ExecutorServiceCancelOnPartitionRequestInitialFrameSize = ExecutorServiceCancelOnPartitionRequestInterruptFieldOffset + bufutil.BooleanSizeInBytes
    ExecutorServiceCancelOnPartitionResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    ExecutorServiceCancelOnPartitionResponseInitialFrameSize = ExecutorServiceCancelOnPartitionResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ExecutorServiceCancelOnPartitionRequestParameters struct {

    /**
    * Unique id for the execution.
    */
uuid UUID

    /**
    * If true, then the thread interrupt call can be used to cancel the thread, otherwise interrupt can not be used.
    */
interrupt bool
}

func ExecutorServiceCancelOnPartitionEncodeRequest(uuid UUID, interrupt bool) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("ExecutorService.CancelOnPartition")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ExecutorServiceCancelOnPartitionRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, ExecutorServiceCancelOnPartitionRequestUuidFieldOffset, uuid)
    bufutil.EncodeBoolean(initialFrame.Content, ExecutorServiceCancelOnPartitionRequestInterruptFieldOffset, interrupt)
    clientMessage.Add(initialFrame)
    return clientMessage
}

func ExecutorServiceCancelOnPartitionDecodeRequest(clientMessage *bufutil.ClientMessagex) *ExecutorServiceCancelOnPartitionRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(ExecutorServiceCancelOnPartitionRequestParameters)
    initialFrame := iterator.Next()
    request.uuid = bufutil.DecodeUUID(initialFrame.Content, ExecutorServiceCancelOnPartitionRequestUuidFieldOffset)
    request.interrupt = bufutil.DecodeBoolean(initialFrame.Content, ExecutorServiceCancelOnPartitionRequestInterruptFieldOffset)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ExecutorServiceCancelOnPartitionResponseParameters struct {
    /**
    * True if cancelled successfully, false otherwise.
    */
response bool
}

func ExecutorServiceCancelOnPartitionEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, ExecutorServiceCancelOnPartitionResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ExecutorServiceCancelOnPartitionResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, ExecutorServiceCancelOnPartitionResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func ExecutorServiceCancelOnPartitionDecodeResponse(clientMessage *bufutil.ClientMessagex) *ExecutorServiceCancelOnPartitionResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ExecutorServiceCancelOnPartitionResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, ExecutorServiceCancelOnPartitionResponseResponseFieldOffset)
    return response
}

