# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [message.proto](#message.proto)
    - [Common](#protos.Common) 公共消息
    - [Mail](#protos.Mail) 发送邮件
    - [Message](#protos.Message) 发送消息结构体
    - [Reponse](#protos.Reponse) 响应
    - [Sentry](#protos.Sentry) sentry消息
  
  
  
    - [SendMessageService](#protos.SendMessageService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="message.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## message.proto



<a name="protos.Common"></a>

### Common



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | 消息类型 |
| message | [string](#string) |  | 消息内容 |






<a name="protos.Mail"></a>

### Mail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | 邮件标题 |
| from | [string](#string) |  |  |
| to | [string](#string) |  |  |
| message | [string](#string) |  | 邮件内容 |
| password | [string](#string) |  |  |






<a name="protos.Message"></a>

### Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | 消息类型（sentry, common, mail） |
| common | [Common](#protos.Common) |  | 公共消息体 |
| sentry | [Sentry](#protos.Sentry) |  | sentry消息体 |
| mail | [Mail](#protos.Mail) |  | 邮件消息体 |






<a name="protos.Reponse"></a>

### Reponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int32](#int32) |  | 状态码 |
| msg | [string](#string) |  | 响应信息 |






<a name="protos.Sentry"></a>

### Sentry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| projectName | [string](#string) |  | 项目名称 |
| level | [string](#string) |  | 错误级别 |
| time | [string](#string) |  | 时间 |
| file | [string](#string) |  | 文件路径 |
| message | [string](#string) |  | 错误信息 |
| href | [string](#string) |  | 地址 |
| type | [string](#string) |  | 类型 |





 

 

 


<a name="protos.SendMessageService"></a>

### SendMessageService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendMessage | [Message](#protos.Message) | [Reponse](#protos.Reponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

