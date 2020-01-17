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
 * Checks the lock for the specified key.If the lock is acquired then returns true, else returns false.
 */
//@Generated("522977927f66138bb2e11f361ea91359")
const (
    //hex: 0x011200
    MapIsLockedRequestMessageType = 70144
    //hex: 0x011201
    MapIsLockedResponseMessageType = 70145
    MapIsLockedRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapIsLockedResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MapIsLockedResponseInitialFrameSize = MapIsLockedResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapIsLockedRequestParameters struct {

    /**
    * name of map
    */
name string

    /**
    * Key for the map entry to check if it is locked.
    */
key serialization.Data
}

func MapIsLockedEncodeRequest(name string, key serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Map.IsLocked")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapIsLockedRequestMessageType)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, key)
    return clientMessage
}




/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapIsLockedResponseParameters struct {
    /**
    * Returns true if the entry is locked, otherwise returns false
    */
response bool
}



func MapIsLockedDecodeResponse(clientMessage *bufutil.ClientMessagex) *MapIsLockedResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapIsLockedResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, MapIsLockedResponseResponseFieldOffset)
    return response
}

