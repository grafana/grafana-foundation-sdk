---
title: <span class="badge builder"></span> TraceqlFilterBuilder
---
# <span class="badge builder"></span> TraceqlFilterBuilder

## Constructor

```php
new TraceqlFilterBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> id

Uniquely identify the filter, will not be used in the query generation

```php
id(string $id)
```

### <span class="badge object-method"></span> operator

The operator that connects the tag to the value, for example: =, >, !=, =~

```php
operator(string $operator)
```

### <span class="badge object-method"></span> scope

The scope of the filter, can either be unscoped/all scopes, resource or span

```php
scope(\Grafana\Foundation\Tempo\TraceqlSearchScope $scope)
```

### <span class="badge object-method"></span> tag

The tag for the search filter, for example: .http.status_code, .service.name, status

```php
tag(string $tag)
```

### <span class="badge object-method"></span> value

The value for the search filter

@param string|array<string> $value

```php
value($value)
```

### <span class="badge object-method"></span> valueType

The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query

```php
valueType(string $valueType)
```

## See also

 * <span class="badge object-type-class"></span> [TraceqlFilter](./object-TraceqlFilter.md)
