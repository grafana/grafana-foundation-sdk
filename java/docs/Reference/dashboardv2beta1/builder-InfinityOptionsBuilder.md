---
title: <span class="badge builder"></span> InfinityOptionsBuilder
---
# <span class="badge builder"></span> InfinityOptionsBuilder

## Constructor

```java
new InfinityOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public InfinityOptions build()
```

### <span class="badge object-method"></span> body

```java
public InfinityOptionsBuilder body(String body)
```

### <span class="badge object-method"></span> datasourceUid

```java
public InfinityOptionsBuilder datasourceUid(String datasourceUid)
```

### <span class="badge object-method"></span> headers

```java
public InfinityOptionsBuilder headers(List<List<String>> headers)
```

### <span class="badge object-method"></span> method

```java
public InfinityOptionsBuilder method(HttpRequestMethod method)
```

### <span class="badge object-method"></span> queryParams

These are 2D arrays of strings, each representing a key-value pair

We are defining them this way because we can't generate a go struct that

that would have exactly two strings in each sub-array

```java
public InfinityOptionsBuilder queryParams(List<List<String>> queryParams)
```

### <span class="badge object-method"></span> url

```java
public InfinityOptionsBuilder url(String url)
```

## See also

 * <span class="badge object-type-class"></span> [InfinityOptions](./object-InfinityOptions.md)
