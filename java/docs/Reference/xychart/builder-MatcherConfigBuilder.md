---
title: <span class="badge builder"></span> MatcherConfigBuilder
---
# <span class="badge builder"></span> MatcherConfigBuilder

NOTE: (copied from dashboard_kind.cue, since not exported)

Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.

It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.

## Constructor

```java
new MatcherConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public MatcherConfig build()
```

### <span class="badge object-method"></span> id

The matcher id. This is used to find the matcher implementation from registry.

```java
public MatcherConfigBuilder id(String id)
```

### <span class="badge object-method"></span> options

The matcher options. This is specific to the matcher implementation.

```java
public MatcherConfigBuilder options(Object options)
```

## See also

 * <span class="badge object-type-class"></span> [MatcherConfig](./object-MatcherConfig.md)
