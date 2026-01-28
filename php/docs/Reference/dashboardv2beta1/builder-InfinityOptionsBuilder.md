---
title: <span class="badge builder"></span> InfinityOptionsBuilder
---
# <span class="badge builder"></span> InfinityOptionsBuilder

## Constructor

```php
new InfinityOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> body

```php
body(string $body)
```

### <span class="badge object-method"></span> datasourceUid

```php
datasourceUid(string $datasourceUid)
```

### <span class="badge object-method"></span> headers

@param array<array<string>> $headers

```php
headers(array $headers)
```

### <span class="badge object-method"></span> method

```php
method(\Grafana\Foundation\Dashboardv2beta1\HttpRequestMethod $method)
```

### <span class="badge object-method"></span> queryParams

These are 2D arrays of strings, each representing a key-value pair

We are defining them this way because we can't generate a go struct that

that would have exactly two strings in each sub-array

@param array<array<string>> $queryParams

```php
queryParams(array $queryParams)
```

### <span class="badge object-method"></span> url

```php
url(string $url)
```

## See also

 * <span class="badge object-type-class"></span> [InfinityOptions](./object-InfinityOptions.md)
