---
title: <span class="badge builder"></span> FetchOptionsBuilder
---
# <span class="badge builder"></span> FetchOptionsBuilder

## Constructor

```typescript
new FetchOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> body

```typescript
body(body: string)
```

### <span class="badge object-method"></span> headers

```typescript
headers(headers: string[][])
```

### <span class="badge object-method"></span> method

```typescript
method(method: dashboard.HttpRequestMethod)
```

### <span class="badge object-method"></span> queryParams

These are 2D arrays of strings, each representing a key-value pair

We are defining this way because we can't generate a go struct that

that would have exactly two strings in each sub-array

```typescript
queryParams(queryParams: string[][])
```

### <span class="badge object-method"></span> url

```typescript
url(url: string)
```

## See also

 * <span class="badge object-type-interface"></span> [FetchOptions](./object-FetchOptions.md)
