// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

package ajni

/*
#include <jni.h>
#include "jni_call.h"
#include <stdlib.h>
*/
import "C"

const (
	// JNIFalse as defined in include/jni.h:1158
	JNIFalse = 0
	// JNITrue as defined in include/jni.h:1159
	JNITrue = 1
	// JNIVersion11 as defined in include/jni.h:1161
	JNIVersion11 = 0x00010001
	// JNIVersion12 as defined in include/jni.h:1162
	JNIVersion12 = 0x00010002
	// JNIVersion14 as defined in include/jni.h:1163
	JNIVersion14 = 0x00010004
	// JNIVersion16 as defined in include/jni.h:1164
	JNIVersion16 = 0x00010006
	// JNIOk as defined in include/jni.h:1166
	JNIOk = (0)
	// JNIErr as defined in include/jni.h:1167
	JNIErr = (-1)
	// JNIEdetached as defined in include/jni.h:1168
	JNIEdetached = (-2)
	// JNIEversion as defined in include/jni.h:1169
	JNIEversion = (-3)
	// JNICommit as defined in include/jni.h:1171
	JNICommit = 1
	// JNIAbort as defined in include/jni.h:1172
	JNIAbort = 2
	// MotionEventActionPointerIndexShift as defined in android/input.h:236
	MotionEventActionPointerIndexShift = 8
	// SensorStandardGravity as defined in android/sensor.h:92
	SensorStandardGravity = (9.80665)
	// SensorMagneticFieldEarthMax as defined in android/sensor.h:94
	SensorMagneticFieldEarthMax = (60.0)
	// SensorMagneticFieldEarthMin as defined in android/sensor.h:96
	SensorMagneticFieldEarthMin = (30.0)
	// TtsEnginePropertyConfig as defined in android/tts.h:33
	TtsEnginePropertyConfig = "engineConfig"
	// TtsEnginePropertyPitch as defined in android/tts.h:34
	TtsEnginePropertyPitch = "pitch"
	// TtsEnginePropertyRate as defined in android/tts.h:35
	TtsEnginePropertyRate = "rate"
	// TtsEnginePropertyVolume as defined in android/tts.h:36
	TtsEnginePropertyVolume = "volume"
)

// JobjectRefType as declared in include/jni.h:141
type JobjectRefType int32

// JobjectRefType enumeration from include/jni.h:141
const (
	JNIinvalidreftype    JobjectRefType = iota
	JNIlocalreftype      JobjectRefType = 1
	JNIglobalreftype     JobjectRefType = 2
	JNIweakglobalreftype JobjectRefType = 3
)

// LogPriority as declared in android/log.h:89
type LogPriority int32

// LogPriority enumeration from android/log.h:89
const (
	LogUnknown LogPriority = iota
	LogDefault LogPriority = 1
	LogVerbose LogPriority = 2
	LogDebug   LogPriority = 3
	LogInfo    LogPriority = 4
	LogWarn    LogPriority = 5
	LogError   LogPriority = 6
	LogFatal   LogPriority = 7
	LogSilent  LogPriority = 8
)
