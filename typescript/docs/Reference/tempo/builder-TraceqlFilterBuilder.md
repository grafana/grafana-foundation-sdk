---
title: <span class="badge builder"></span> TraceqlFilterBuilder
---
# <span class="badge builder"></span> TraceqlFilterBuilder

## Constructor

```typescript
new TraceqlFilterBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> id

Uniquely identify the filter, will not be used in the query generation

```typescript
id(id: string)
```

### <span class="badge object-method"></span> operator

The operator that connects the tag to the value, for example: =, >, !=, =~

```typescript
operator(operator: string)
```

### <span class="badge object-method"></span> scope

The scope of the filter, can either be unscoped/all scopes, resource or span

```typescript
scope(scope: tempo.TraceqlSearchScope)
```

### <span class="badge object-method"></span> tag

The tag for the search filter, for example: .http.status_code, .service.name, status

```typescript
tag(tag: string)
```

### <span class="badge object-method"></span> value

The value for the search filter

```typescript
value(value: string | string[])
```

### <span class="badge object-method"></span> valueType

The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query

```typescript
valueType(valueType: string)
```

## See also

 * <span class="badge object-type-interface"></span> [TraceqlFilter](./object-TraceqlFilter.md)
