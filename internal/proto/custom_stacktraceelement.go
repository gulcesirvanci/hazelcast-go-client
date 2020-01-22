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
    "github.com/hazelcast/hazelcast-go-client/internal/proto/bufutil"
)

type StackTraceElement struct {
className string
methodName string
fileName string
lineNumber int32
}

//@Generated("ebcc9867a14e608b8fd5b4969a3d2de0")
const (
    StackTraceElementLineNumberFieldOffset = 0
    StackTraceElementInitialFrameSize = StackTraceElementLineNumberFieldOffset + bufutil.IntSizeInBytes
)

func StackTraceElementCodecEncode(clientMessage *bufutil.ClientMessage, stackTraceElement StackTraceElement) {
        clientMessage.Add(bufutil.BeginFrame)
        initialFrame := &bufutil.Frame{Content: make([]byte, StackTraceElementInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
        bufutil.EncodeInt(initialFrame.Content, StackTraceElementLineNumberFieldOffset, stackTraceElement.lineNumber)
        clientMessage.Add(initialFrame)
        bufutil.StringCodecEncode(clientMessage, stackTraceElement.className)
        bufutil.StringCodecEncode(clientMessage, stackTraceElement.methodName)
        bufutil.EncodeNullable(clientMessage, stackTraceElement.fileName, bufutil.StringCodecEncode)

        clientMessage.Add(bufutil.EndFrame)
    }

func StackTraceElementCodecDecode(iterator *bufutil.ForwardFrameIterator)  StackTraceElement  {
        // begin frame
        iterator.Next()
        initialFrame := iterator.Next()
        lineNumber := bufutil.DecodeInt(initialFrame.Content, StackTraceElementLineNumberFieldOffset)
        className := bufutil.StringCodecDecode(iterator)
        methodName := bufutil.StringCodecDecode(iterator)
        fileName := bufutil.DecodeNullable(iterator, bufutil.StringCodecDecode).(string)
        bufutil.FastForwardToEndFrame(iterator)
        return StackTraceElement { className, methodName, fileName, lineNumber }
    }