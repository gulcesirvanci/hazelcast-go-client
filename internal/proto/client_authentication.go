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
//@Generated("971ffb12c81547ce4375c9acc7e2cbfa")
const (
    //hex: 0x000100
    ClientAuthenticationRequestMessageType = 256
    //hex: 0x000101
    ClientAuthenticationResponseMessageType = 257
    ClientAuthenticationRequestUuidFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ClientAuthenticationRequestSerializationVersionFieldOffset = ClientAuthenticationRequestUuidFieldOffset + bufutil.UUIDSizeInBytes
    ClientAuthenticationRequestPartitionCountFieldOffset = ClientAuthenticationRequestSerializationVersionFieldOffset + bufutil.ByteSizeInBytes
    ClientAuthenticationRequestClusterIdFieldOffset = ClientAuthenticationRequestPartitionCountFieldOffset + bufutil.IntSizeInBytes
    ClientAuthenticationRequestInitialFrameSize = ClientAuthenticationRequestClusterIdFieldOffset + bufutil.UUIDSizeInBytes
    ClientAuthenticationResponseStatusFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    ClientAuthenticationResponseUuidFieldOffset = ClientAuthenticationResponseStatusFieldOffset + bufutil.ByteSizeInBytes
    ClientAuthenticationResponseSerializationVersionFieldOffset = ClientAuthenticationResponseUuidFieldOffset + bufutil.UUIDSizeInBytes
    ClientAuthenticationResponsePartitionCountFieldOffset = ClientAuthenticationResponseSerializationVersionFieldOffset + bufutil.ByteSizeInBytes
    ClientAuthenticationResponseClusterIdFieldOffset = ClientAuthenticationResponsePartitionCountFieldOffset + bufutil.IntSizeInBytes
    ClientAuthenticationResponseInitialFrameSize = ClientAuthenticationResponseClusterIdFieldOffset + bufutil.UUIDSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ClientAuthenticationRequestParameters struct {

    /**
    * Cluster name that client will connect to.
    */
clusterName string

    /**
    * Name of the user for authentication.
    * Used in case Client Identity Config, otherwise it should be passed null.
    */
/* @Nullable */ username string

    /**
    * Password for the user.
    * Used in case Client Identity Config, otherwise it should be passed null.
    */
/* @Nullable */ password string

    /**
    * Unique string identifying the connected client uniquely.
    */
/* @Nullable */ uuid string

    /**
    * The type of the client. E.g. JAVA, CPP, CSHARP, etc.
    */
clientType string

    /**
    * client side supported version to inform server side
    */
serializationVersion byte

    /**
    * The Hazelcast version of the client. (e.g. 3.7.2)
    */
clientHazelcastVersion string

    /**
    * the name of the client instance
    */
clientName string

    /**
    * User defined labels of the client instance
    */
labels []string

    /**
    * the expected partition count of the cluster. Checked on the server side when provided.
    * Authentication fails and 3:NOT_ALLOWED_IN_CLUSTER returned, in case of mismatch
    */
partitionCount int32

    /**
    * the expected id of the cluster. Checked on the server side when provided.
    * Authentication fails and 3:NOT_ALLOWED_IN_CLUSTER returned, in case of mismatch
    */
/* @Nullable */ clusterId string
}

func ClientAuthenticationEncodeRequest(clusterName string, /* @Nullable */ username string, /* @Nullable */ password string, /* @Nullable */ uuid string, clientType string, serializationVersion byte, clientHazelcastVersion string, clientName string, labels []string, partitionCount int32, /* @Nullable */ clusterId string) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Client.Authentication")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ClientAuthenticationRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, ClientAuthenticationRequestUuidFieldOffset, uuid)
    bufutil.EncodeByte(initialFrame.Content, ClientAuthenticationRequestSerializationVersionFieldOffset, serializationVersion)
    bufutil.EncodeInt(initialFrame.Content, ClientAuthenticationRequestPartitionCountFieldOffset, partitionCount)
    bufutil.EncodeUUID(initialFrame.Content, ClientAuthenticationRequestClusterIdFieldOffset, clusterId)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, clusterName)
    bufutil.EncodeNullable(clientMessage, username, bufutil.StringCodecEncode)
    bufutil.EncodeNullable(clientMessage, password, bufutil.StringCodecEncode)
    bufutil.StringCodecEncode(clientMessage, clientType)
    bufutil.StringCodecEncode(clientMessage, clientHazelcastVersion)
    bufutil.StringCodecEncode(clientMessage, clientName)
    bufutil.ListMultiFrameCodecEncode(clientMessage, labels, bufutil.StringCodecEncode)
    return clientMessage
}




/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ClientAuthenticationResponseParameters struct {
    /**
    * TODO DOC
    */
status byte
    /**
    * TODO DOC
    */
/* @Nullable */ address Address
    /**
    * Unique string identifying the connected client uniquely.
    */
/* @Nullable */ uuid string
    /**
    * client side supported version to inform server side
    */
serializationVersion byte
    /**
    * TODO DOC
    */
serverHazelcastVersion string
    /**
    * the expected partition count of the cluster. Checked on the server side when provided.
    * Authentication fails and 3:NOT_ALLOWED_IN_CLUSTER returned, in case of mismatch
    */
partitionCount int32
    /**
    * the expected id of the cluster. Checked on the server side when provided.
    * Authentication fails and 3:NOT_ALLOWED_IN_CLUSTER returned, in case of mismatch
    */
clusterId string
}



func ClientAuthenticationDecodeResponse(clientMessage *bufutil.ClientMessagex) *ClientAuthenticationResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ClientAuthenticationResponseParameters)
    initialFrame := iterator.Next()
    response.status = bufutil.DecodeByte(initialFrame.Content, ClientAuthenticationResponseStatusFieldOffset)
    response.uuid = bufutil.DecodeUUID(initialFrame.Content, ClientAuthenticationResponseUuidFieldOffset)
    response.serializationVersion = bufutil.DecodeByte(initialFrame.Content, ClientAuthenticationResponseSerializationVersionFieldOffset)
    response.partitionCount = bufutil.DecodeInt(initialFrame.Content, ClientAuthenticationResponsePartitionCountFieldOffset)
    response.clusterId = bufutil.DecodeUUID(initialFrame.Content, ClientAuthenticationResponseClusterIdFieldOffset)
    response.address = bufutil.DecodeNullable(iterator, bufutil.AddressCodecDecode) 
    response.serverHazelcastVersion = bufutil.StringCodecDecode(iterator)
    return response
}

