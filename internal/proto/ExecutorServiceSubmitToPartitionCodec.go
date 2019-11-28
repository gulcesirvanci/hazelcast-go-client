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
//@Generated("89d84e09a2b54e47d372c7e04687ef0a")
const (
    //hex: 0x080500
    ExecutorServiceSubmitToPartitionRequestMessageType = 525568
    //hex: 0x080501
    ExecutorServiceSubmitToPartitionResponseMessageType = 525569
    ExecutorServiceSubmitToPartitionRequestUuidFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ExecutorServiceSubmitToPartitionRequestInitialFrameSize = ExecutorServiceSubmitToPartitionRequestUuidFieldOffset + bufutil.UUIDSizeInBytes
    ExecutorServiceSubmitToPartitionResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ExecutorServiceSubmitToPartitionRequestParameters struct {

    /**
    * Name of the executor.
    */
name string

    /**
    * Unique id for the execution.
    */
uuid UUID

    /**
    * The callable object to be executed.
    */
callable serialization.Data
}

func ExecutorServiceSubmitToPartitionEncodeRequest(name string, uuid UUID, callable serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("ExecutorService.SubmitToPartition")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ExecutorServiceSubmitToPartitionRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, ExecutorServiceSubmitToPartitionRequestUuidFieldOffset, uuid)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, callable)
    return clientMessage
}

func ExecutorServiceSubmitToPartitionDecodeRequest(clientMessage *bufutil.ClientMessagex) *ExecutorServiceSubmitToPartitionRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(ExecutorServiceSubmitToPartitionRequestParameters)
    initialFrame := iterator.Next()
    request.uuid = bufutil.DecodeUUID(initialFrame.Content, ExecutorServiceSubmitToPartitionRequestUuidFieldOffset)
    request.name = StringCodec.decode(iterator)
    request.callable = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ExecutorServiceSubmitToPartitionResponseParameters struct {
    /**
    * The result of the callable execution.
    */
/* @Nullable */ response serialization.Data
}

func ExecutorServiceSubmitToPartitionEncodeResponse(/* @Nullable */ response serialization.Data ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, ExecutorServiceSubmitToPartitionResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ExecutorServiceSubmitToPartitionResponseMessageType)
    clientMessage.Add(initialFrame)

    CodecUtil.encodeNullable(clientMessage, response, DataCodecEncode)
    return clientMessage
}

func ExecutorServiceSubmitToPartitionDecodeResponse(clientMessage *bufutil.ClientMessagex) *ExecutorServiceSubmitToPartitionResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ExecutorServiceSubmitToPartitionResponseParameters)
    //empty initial frame
    iterator.Next()
    response.response = CodecUtil.decodeNullable(iterator, DataCodecDecode) 
    return response
}

