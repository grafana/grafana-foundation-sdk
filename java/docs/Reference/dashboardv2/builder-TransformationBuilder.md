---
title: <span class="badge builder"></span> TransformationBuilder
---
# <span class="badge builder"></span> TransformationBuilder

## Constructor

```java
new TransformationBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Transformation build()
```

### <span class="badge object-method"></span> disabled

Disabled transformations are skipped

```java
public TransformationBuilder disabled(Boolean disabled)
```

### <span class="badge object-method"></span> filter

Optional frame matcher. When missing it will be applied to all results

```java
public TransformationBuilder filter(MatcherConfig filter)
```

### <span class="badge object-method"></span> group

The group is the transformation ID

```java
public TransformationBuilder group(String group)
```

### <span class="badge object-method"></span> options

Options to be passed to the transformer

Valid options depend on the transformer id

```java
public TransformationBuilder options(Object options)
```

### <span class="badge object-method"></span> topic

Where to pull DataFrames from as input to transformation

```java
public TransformationBuilder topic(DataTopic topic)
```

## See also

 * <span class="badge object-type-class"></span> [TransformationKind](./object-TransformationKind.md)
