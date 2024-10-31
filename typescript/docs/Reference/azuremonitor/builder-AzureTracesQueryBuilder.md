---
title: <span class="badge builder"></span> AzureTracesQueryBuilder
---
# <span class="badge builder"></span> AzureTracesQueryBuilder

## Constructor

```typescript
new AzureTracesQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> filters

Filters for property values.

```typescript
filters(filters: cog.Builder<azuremonitor.AzureTracesFilter>[])
```

### <span class="badge object-method"></span> operationId

Operation ID. Used only for Traces queries.

```typescript
operationId(operationId: string)
```

### <span class="badge object-method"></span> query

KQL query to be executed.

```typescript
query(query: string)
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

```typescript
resources(resources: string[])
```

### <span class="badge object-method"></span> resultFormat

Specifies the format results should be returned as.

```typescript
resultFormat(resultFormat: azuremonitor.ResultFormat)
```

### <span class="badge object-method"></span> traceTypes

Types of events to filter by.

```typescript
traceTypes(traceTypes: string[])
```

## See also

 * <span class="badge object-type-interface"></span> [AzureTracesQuery](./object-AzureTracesQuery.md)
