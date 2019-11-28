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
 * Gets the current metadata of a cluster.
 */
//@Generated("db3b91284de53053d068b4d702e0e735")
const (
    //hex: 0x200E00
    MCGetClusterMetadataRequestMessageType = 2100736
    //hex: 0x200E01
    MCGetClusterMetadataResponseMessageType = 2100737
    MCGetClusterMetadataRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MCGetClusterMetadataResponseCurrentStateFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MCGetClusterMetadataResponseClusterTimeFieldOffset = MCGetClusterMetadataResponseCurrentStateFieldOffset + bufutil.ByteSizeInBytes
    MCGetClusterMetadataResponseInitialFrameSize = MCGetClusterMetadataResponseClusterTimeFieldOffset + bufutil.LongSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MCGetClusterMetadataRequestParameters struct {
}

func MCGetClusterMetadataEncodeRequest() *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("MC.GetClusterMetadata")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MCGetClusterMetadataRequestMessageType)
    clientMessage.Add(initialFrame)
    return clientMessage
}

func MCGetClusterMetadataDecodeRequest(clientMessage *bufutil.ClientMessagex) *MCGetClusterMetadataRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(MCGetClusterMetadataRequestParameters)
    //empty initial frame
    iterator.Next()
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MCGetClusterMetadataResponseParameters struct {
    /**
    * Current state of the cluster:
    * 0 - ACTIVE
    * 1 - NO_MIGRATION
    * 2 - FROZEN
    * 3 - PASSIVE
    * 4 - IN_TRANSITION (not allowed)
    */
currentState byte
    /**
    * Current version of the member.
    */
memberVersion string
    /**
    * Current Jet version of the member.
    */
/* @Nullable */ jetVersion string
    /**
    * Cluster-wide time in milliseconds.
    */
clusterTime int64
}

func MCGetClusterMetadataEncodeResponse(currentState byte , memberVersion string , /* @Nullable */ jetVersion string , clusterTime int64 ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, MCGetClusterMetadataResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MCGetClusterMetadataResponseMessageType)
    bufutil.EncodeByte(initialFrame.Content, MCGetClusterMetadataResponseCurrentStateFieldOffset, currentState)
    bufutil.EncodeLong(initialFrame.Content, MCGetClusterMetadataResponseClusterTimeFieldOffset, clusterTime)
    clientMessage.Add(initialFrame)

    StringCodec.encode(clientMessage, memberVersion)
    CodecUtil.encodeNullable(clientMessage, jetVersion, StringCodecEncode)
    return clientMessage
}

func MCGetClusterMetadataDecodeResponse(clientMessage *bufutil.ClientMessagex) *MCGetClusterMetadataResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MCGetClusterMetadataResponseParameters)
    initialFrame := iterator.Next()
    response.currentState = bufutil.DecodeByte(initialFrame.Content, MCGetClusterMetadataResponseCurrentStateFieldOffset)
    response.clusterTime = bufutil.DecodeLong(initialFrame.Content, MCGetClusterMetadataResponseClusterTimeFieldOffset)
    response.memberVersion = StringCodec.decode(iterator)
    response.jetVersion = CodecUtil.decodeNullable(iterator, StringCodecDecode) 
    return response
}

