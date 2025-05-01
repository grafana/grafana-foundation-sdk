---
title: <span class="badge builder"></span> StreamingQueryBuilder
---
# <span class="badge builder"></span> StreamingQueryBuilder

## Constructor

```java
new StreamingQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public StreamingQuery build()
```

### <span class="badge object-method"></span> bands

```java
public StreamingQueryBuilder bands(Long bands)
```

### <span class="badge object-method"></span> noise

```java
public StreamingQueryBuilder noise(Double noise)
```

### <span class="badge object-method"></span> speed

```java
public StreamingQueryBuilder speed(Double speed)
```

### <span class="badge object-method"></span> spread

```java
public StreamingQueryBuilder spread(Double spread)
```

### <span class="badge object-method"></span> type

Possible enum values:

 - `"fetch"` 

 - `"logs"` 

 - `"signal"` 

 - `"traces"` 

```java
public StreamingQueryBuilder type(StreamingQueryType type)
```

### <span class="badge object-method"></span> url

```java
public StreamingQueryBuilder url(String url)
```

## See also

 * <span class="badge object-type-class"></span> [StreamingQuery](./object-StreamingQuery.md)
