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
 * Returns true if this map contains no key-value mappings.
 */
//@Generated("f03dfeb76597ecb47a226964eea2c43f")
const (
    //hex: 0x012B00
    MapIsEmptyRequestMessageType = 76544
    //hex: 0x012B01
    MapIsEmptyResponseMessageType = 76545
    MapIsEmptyRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapIsEmptyResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MapIsEmptyResponseInitialFrameSize = MapIsEmptyResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapIsEmptyRequestParameters struct {

    /**
    * name of map
    */
name string
}

func MapIsEmptyEncodeRequest(name string) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Map.IsEmpty")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapIsEmptyRequestMessageType)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    return clientMessage
}




/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapIsEmptyResponseParameters struct {
    /**
    * true if this map contains no key-value mappings
    */
response bool
}



func MapIsEmptyDecodeResponse(clientMessage *bufutil.ClientMessagex) *MapIsEmptyResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapIsEmptyResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, MapIsEmptyResponseResponseFieldOffset)
    return response
}

