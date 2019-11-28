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
 * Retrieves and disposes the result of the execution with the given sequence
 */
//@Generated("3251733b1bcf84e176166fa1d07a779c")
const (
    //hex: 0x180600
    DurableExecutorRetrieveAndDisposeResultRequestMessageType = 1574400
    //hex: 0x180601
    DurableExecutorRetrieveAndDisposeResultResponseMessageType = 1574401
    DurableExecutorRetrieveAndDisposeResultRequestSequenceFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    DurableExecutorRetrieveAndDisposeResultRequestInitialFrameSize = DurableExecutorRetrieveAndDisposeResultRequestSequenceFieldOffset + bufutil.IntSizeInBytes
    DurableExecutorRetrieveAndDisposeResultResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type DurableExecutorRetrieveAndDisposeResultRequestParameters struct {

    /**
    * Name of the executor.
    */
name string

    /**
    * Sequence of the execution.
    */
sequence int
}

func DurableExecutorRetrieveAndDisposeResultEncodeRequest(name string, sequence int) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("DurableExecutor.RetrieveAndDisposeResult")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, DurableExecutorRetrieveAndDisposeResultRequestMessageType)
    bufutil.EncodeInt(initialFrame.Content, DurableExecutorRetrieveAndDisposeResultRequestSequenceFieldOffset, sequence)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func DurableExecutorRetrieveAndDisposeResultDecodeRequest(clientMessage *bufutil.ClientMessagex) *DurableExecutorRetrieveAndDisposeResultRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(DurableExecutorRetrieveAndDisposeResultRequestParameters)
    initialFrame := iterator.Next()
    request.sequence = bufutil.DecodeInt(initialFrame.Content, DurableExecutorRetrieveAndDisposeResultRequestSequenceFieldOffset)
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type DurableExecutorRetrieveAndDisposeResultResponseParameters struct {
    /**
    * The result of the callable execution with the given sequence.
    */
/* @Nullable */ response serialization.Data
}

func DurableExecutorRetrieveAndDisposeResultEncodeResponse(/* @Nullable */ response serialization.Data ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, DurableExecutorRetrieveAndDisposeResultResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, DurableExecutorRetrieveAndDisposeResultResponseMessageType)
    clientMessage.Add(initialFrame)

    CodecUtil.encodeNullable(clientMessage, response, DataCodecEncode)
    return clientMessage
}

func DurableExecutorRetrieveAndDisposeResultDecodeResponse(clientMessage *bufutil.ClientMessagex) *DurableExecutorRetrieveAndDisposeResultResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(DurableExecutorRetrieveAndDisposeResultResponseParameters)
    //empty initial frame
    iterator.Next()
    response.response = CodecUtil.decodeNullable(iterator, DataCodecDecode) 
    return response
}

