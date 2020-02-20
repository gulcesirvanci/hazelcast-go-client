
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
)


/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Deploys the list of classes to cluster
 * Each item is a Map.Entry<String, byte[]> in the list.
 * key of entry is full class name, and byte[] is the class definition.
 */
//@Generated("34538bf8bb4f49e04aaa331ff89bc6bd")
const (
    //hex: 0x000D00
    ClientDeployClassesRequestMessageType = 3328
    //hex: 0x000D01
    ClientDeployClassesResponseMessageType = 3329
    ClientDeployClassesRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    ClientDeployClassesResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func ClientDeployClassesEncodeRequest(classDefinitions []*Pair) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Client.DeployClasses")
	initialFrame := &Frame{Content: make([]byte, ClientDeployClassesResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ClientDeployClassesRequestMessageType)
    clientMessage.Add(initialFrame)

            clientMessage.Add(BeginFrame.Copy())

        for _, cp := range classDefinitions{
           StringCodecEncode(clientMessage, (*cp).Key().(string))
           StringCodecEncode(clientMessage, (*cp).Value().([]byte))
        }

        clientMessage.Add(EndFrame.Copy())



    return clientMessage
}


func ClientDeployClassesDecodeResponse(clientMessage *ClientMessage) func() () {
    return func() () {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        return
    }
}
