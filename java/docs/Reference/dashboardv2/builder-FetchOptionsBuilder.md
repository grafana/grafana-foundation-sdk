---
title: <span class="badge builder"></span> FetchOptionsBuilder
---
# <span class="badge builder"></span> FetchOptionsBuilder

## Constructor

```java
new FetchOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public FetchOptions build()
```

### <span class="badge object-method"></span> body

```java
public FetchOptionsBuilder body(String body)
```

### <span class="badge object-method"></span> headers

```java
public FetchOptionsBuilder headers(List<List<String>> headers)
```

### <span class="badge object-method"></span> method

```java
public FetchOptionsBuilder method(HttpRequestMethod method)
```

### <span class="badge object-method"></span> queryParams

These are 2D arrays of strings, each representing a key-value pair

We are defining them this way because we can't generate a go struct that

that would have exactly two strings in each sub-array

```java
public FetchOptionsBuilder queryParams(List<List<String>> queryParams)
```

### <span class="badge object-method"></span> url

```java
public FetchOptionsBuilder url(String url)
```

## See also

 * <span class="badge object-type-class"></span> [FetchOptions](./object-FetchOptions.md)
