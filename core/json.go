package core

import (
	jsoniter "github.com/json-iterator/go"
)

// 重写JSON对象
var json = jsoniter.ConfigCompatibleWithStandardLibrary
var fastJson = jsoniter.ConfigFastest
var FastJsonMarshal = fastJson.Marshal
var FastJsonMarshalIndent = fastJson.MarshalIndent
var FastJsonMarshalToString = fastJson.MarshalToString
var FastJsonUnmarshal = fastJson.Unmarshal
var FastJsonUnmarshalFromString = fastJson.UnmarshalFromString
