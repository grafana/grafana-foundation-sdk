---
title: <span class="badge builder"></span> TraceqlFilterBuilder
---
# <span class="badge builder"></span> TraceqlFilterBuilder

## Constructor

```java
new TraceqlFilterBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public TraceqlFilter build()
```

### <span class="badge object-method"></span> id

Uniquely identify the filter, will not be used in the query generation

```java
public TraceqlFilterBuilder id(String id)
```

### <span class="badge object-method"></span> operator

The operator that connects the tag to the value, for example: =, >, !=, =~

```java
public TraceqlFilterBuilder operator(String operator)
```

### <span class="badge object-method"></span> scope

The scope of the filter, can either be unscoped/all scopes, resource or span

```java
public TraceqlFilterBuilder scope(TraceqlSearchScope scope)
```

### <span class="badge object-method"></span> tag

The tag for the search filter, for example: .http.status_code, .service.name, status

```java
public TraceqlFilterBuilder tag(String tag)
```

### <span class="badge object-method"></span> value

The value for the search filter

```java
public TraceqlFilterBuilder value(StringOrArrayOfString value)
```

### <span class="badge object-method"></span> valueType

The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query

```java
public TraceqlFilterBuilder valueType(String valueType)
```

## See also

 * <span class="badge object-type-class"></span> [TraceqlFilter](./object-TraceqlFilter.md)
