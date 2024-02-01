# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [webhooks/adapt/grpc/v1/webhooks.proto](#webhooks_adapt_grpc_v1_webhooks-proto)
    - [App](#webhooks-v1-App)
    - [App.MetadataEntry](#webhooks-v1-App-MetadataEntry)
    - [Attempt](#webhooks-v1-Attempt)
    - [CreateAppsRequest](#webhooks-v1-CreateAppsRequest)
    - [CreateAppsResponse](#webhooks-v1-CreateAppsResponse)
    - [CreateEndpointsRequest](#webhooks-v1-CreateEndpointsRequest)
    - [CreateEndpointsResponse](#webhooks-v1-CreateEndpointsResponse)
    - [CreateMessagesRequest](#webhooks-v1-CreateMessagesRequest)
    - [CreateMessagesResponse](#webhooks-v1-CreateMessagesResponse)
    - [DeleteAppsRequest](#webhooks-v1-DeleteAppsRequest)
    - [DeleteAppsResponse](#webhooks-v1-DeleteAppsResponse)
    - [DeleteEndpointsRequest](#webhooks-v1-DeleteEndpointsRequest)
    - [DeleteEndpointsResponse](#webhooks-v1-DeleteEndpointsResponse)
    - [DeleteMessagesRequest](#webhooks-v1-DeleteMessagesRequest)
    - [DeleteMessagesResponse](#webhooks-v1-DeleteMessagesResponse)
    - [Endpoint](#webhooks-v1-Endpoint)
    - [Endpoint.MetadataEntry](#webhooks-v1-Endpoint-MetadataEntry)
    - [Error](#webhooks-v1-Error)
    - [GetAppsRequest](#webhooks-v1-GetAppsRequest)
    - [GetAppsResponse](#webhooks-v1-GetAppsResponse)
    - [GetAppsResponse.DataEntry](#webhooks-v1-GetAppsResponse-DataEntry)
    - [GetEndpointsRequest](#webhooks-v1-GetEndpointsRequest)
    - [GetEndpointsResponse](#webhooks-v1-GetEndpointsResponse)
    - [GetEndpointsResponse.DataEntry](#webhooks-v1-GetEndpointsResponse-DataEntry)
    - [GetMessagesRequest](#webhooks-v1-GetMessagesRequest)
    - [GetMessagesResponse](#webhooks-v1-GetMessagesResponse)
    - [GetMessagesResponse.DataEntry](#webhooks-v1-GetMessagesResponse-DataEntry)
    - [ListAppsRequest](#webhooks-v1-ListAppsRequest)
    - [ListAppsResponse](#webhooks-v1-ListAppsResponse)
    - [ListEndpointsRequest](#webhooks-v1-ListEndpointsRequest)
    - [ListEndpointsResponse](#webhooks-v1-ListEndpointsResponse)
    - [ListMessagesRequest](#webhooks-v1-ListMessagesRequest)
    - [ListMessagesResponse](#webhooks-v1-ListMessagesResponse)
    - [Message](#webhooks-v1-Message)
    - [PageRequest](#webhooks-v1-PageRequest)
  
    - [WebHookService](#webhooks-v1-WebHookService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="webhooks_adapt_grpc_v1_webhooks-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## webhooks/adapt/grpc/v1/webhooks.proto



<a name="webhooks-v1-App"></a>

### App
Entities


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| uid | [string](#string) |  |  |
| name | [string](#string) |  |  |
| tenantId | [string](#string) |  |  |
| rateLimit | [int32](#int32) |  |  |
| metadata | [App.MetadataEntry](#webhooks-v1-App-MetadataEntry) | repeated |  |






<a name="webhooks-v1-App-MetadataEntry"></a>

### App.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="webhooks-v1-Attempt"></a>

### Attempt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="webhooks-v1-CreateAppsRequest"></a>

### CreateAppsRequest
Req/Res


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [App](#webhooks-v1-App) | repeated |  |






<a name="webhooks-v1-CreateAppsResponse"></a>

### CreateAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [App](#webhooks-v1-App) | repeated |  |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-CreateEndpointsRequest"></a>

### CreateEndpointsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tenant_id | [string](#string) |  |  |
| data | [Endpoint](#webhooks-v1-Endpoint) | repeated |  |






<a name="webhooks-v1-CreateEndpointsResponse"></a>

### CreateEndpointsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Endpoint](#webhooks-v1-Endpoint) | repeated |  |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-CreateMessagesRequest"></a>

### CreateMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tenant_id | [string](#string) |  |  |
| data | [Message](#webhooks-v1-Message) | repeated |  |






<a name="webhooks-v1-CreateMessagesResponse"></a>

### CreateMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Message](#webhooks-v1-Message) | repeated |  |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-DeleteAppsRequest"></a>

### DeleteAppsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [string](#string) | repeated |  |






<a name="webhooks-v1-DeleteAppsResponse"></a>

### DeleteAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-DeleteEndpointsRequest"></a>

### DeleteEndpointsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tenant_id | [string](#string) |  |  |
| ids | [string](#string) | repeated |  |






<a name="webhooks-v1-DeleteEndpointsResponse"></a>

### DeleteEndpointsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-DeleteMessagesRequest"></a>

### DeleteMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tenant_id | [string](#string) |  |  |
| ids | [string](#string) | repeated |  |






<a name="webhooks-v1-DeleteMessagesResponse"></a>

### DeleteMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-Endpoint"></a>

### Endpoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| uid | [string](#string) |  |  |
| url | [string](#string) |  |  |
| secret_id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| disabled | [bool](#bool) |  |  |
| version | [int32](#int32) |  |  |
| rateLimit | [int32](#int32) |  |  |
| metadata | [Endpoint.MetadataEntry](#webhooks-v1-Endpoint-MetadataEntry) | repeated |  |
| filterTypes | [string](#string) | repeated |  |
| channels | [string](#string) | repeated |  |






<a name="webhooks-v1-Endpoint-MetadataEntry"></a>

### Endpoint.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="webhooks-v1-Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int32](#int32) |  |  |
| index | [string](#string) |  |  |
| reason | [string](#string) |  |  |






<a name="webhooks-v1-GetAppsRequest"></a>

### GetAppsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [string](#string) | repeated |  |






<a name="webhooks-v1-GetAppsResponse"></a>

### GetAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [GetAppsResponse.DataEntry](#webhooks-v1-GetAppsResponse-DataEntry) | repeated |  |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-GetAppsResponse-DataEntry"></a>

### GetAppsResponse.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [App](#webhooks-v1-App) |  |  |






<a name="webhooks-v1-GetEndpointsRequest"></a>

### GetEndpointsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tenant_id | [string](#string) |  |  |
| ids | [string](#string) | repeated |  |






<a name="webhooks-v1-GetEndpointsResponse"></a>

### GetEndpointsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [GetEndpointsResponse.DataEntry](#webhooks-v1-GetEndpointsResponse-DataEntry) | repeated |  |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-GetEndpointsResponse-DataEntry"></a>

### GetEndpointsResponse.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Endpoint](#webhooks-v1-Endpoint) |  |  |






<a name="webhooks-v1-GetMessagesRequest"></a>

### GetMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tenant_id | [string](#string) |  |  |
| ids | [string](#string) | repeated |  |






<a name="webhooks-v1-GetMessagesResponse"></a>

### GetMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [GetMessagesResponse.DataEntry](#webhooks-v1-GetMessagesResponse-DataEntry) | repeated |  |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-GetMessagesResponse-DataEntry"></a>

### GetMessagesResponse.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Message](#webhooks-v1-Message) |  |  |






<a name="webhooks-v1-ListAppsRequest"></a>

### ListAppsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [PageRequest](#webhooks-v1-PageRequest) |  |  |






<a name="webhooks-v1-ListAppsResponse"></a>

### ListAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [App](#webhooks-v1-App) | repeated |  |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-ListEndpointsRequest"></a>

### ListEndpointsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tenant_id | [string](#string) |  |  |
| page | [PageRequest](#webhooks-v1-PageRequest) |  |  |






<a name="webhooks-v1-ListEndpointsResponse"></a>

### ListEndpointsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Endpoint](#webhooks-v1-Endpoint) | repeated |  |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-ListMessagesRequest"></a>

### ListMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tenant_id | [string](#string) |  |  |
| page | [PageRequest](#webhooks-v1-PageRequest) |  |  |






<a name="webhooks-v1-ListMessagesResponse"></a>

### ListMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Message](#webhooks-v1-Message) | repeated |  |
| errors | [Error](#webhooks-v1-Error) | repeated |  |






<a name="webhooks-v1-Message"></a>

### Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| event_id | [string](#string) |  |  |
| event_type | [string](#string) |  |  |
| timestamp | [string](#string) |  |  |
| payload | [bytes](#bytes) |  |  |
| tags | [string](#string) | repeated |  |
| channels | [string](#string) | repeated |  |






<a name="webhooks-v1-PageRequest"></a>

### PageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) | optional |  |
| offset | [int32](#int32) |  |  |
| desc | [bool](#bool) |  |  |





 

 

 


<a name="webhooks-v1-WebHookService"></a>

### WebHookService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetApps | [GetAppsRequest](#webhooks-v1-GetAppsRequest) | [GetAppsResponse](#webhooks-v1-GetAppsResponse) |  |
| ListApps | [ListAppsRequest](#webhooks-v1-ListAppsRequest) | [ListAppsResponse](#webhooks-v1-ListAppsResponse) |  |
| CreateApps | [CreateAppsRequest](#webhooks-v1-CreateAppsRequest) | [CreateAppsResponse](#webhooks-v1-CreateAppsResponse) |  |
| DeleteApps | [DeleteAppsRequest](#webhooks-v1-DeleteAppsRequest) | [DeleteAppsResponse](#webhooks-v1-DeleteAppsResponse) |  |
| GetEndpoints | [GetEndpointsRequest](#webhooks-v1-GetEndpointsRequest) | [GetEndpointsResponse](#webhooks-v1-GetEndpointsResponse) |  |
| ListEndpoints | [ListEndpointsRequest](#webhooks-v1-ListEndpointsRequest) | [ListEndpointsResponse](#webhooks-v1-ListEndpointsResponse) |  |
| CreateEndpoints | [CreateEndpointsRequest](#webhooks-v1-CreateEndpointsRequest) | [CreateEndpointsResponse](#webhooks-v1-CreateEndpointsResponse) |  |
| DeleteEndpoints | [DeleteEndpointsRequest](#webhooks-v1-DeleteEndpointsRequest) | [DeleteEndpointsResponse](#webhooks-v1-DeleteEndpointsResponse) |  |
| GetMessages | [GetMessagesRequest](#webhooks-v1-GetMessagesRequest) | [GetMessagesResponse](#webhooks-v1-GetMessagesResponse) |  |
| ListMessages | [ListMessagesRequest](#webhooks-v1-ListMessagesRequest) | [ListMessagesResponse](#webhooks-v1-ListMessagesResponse) |  |
| CreateMessages | [CreateMessagesRequest](#webhooks-v1-CreateMessagesRequest) | [CreateMessagesResponse](#webhooks-v1-CreateMessagesResponse) |  |
| DeleteMessages | [DeleteMessagesRequest](#webhooks-v1-DeleteMessagesRequest) | [DeleteMessagesResponse](#webhooks-v1-DeleteMessagesResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

