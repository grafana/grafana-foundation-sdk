---
title: <span class="badge builder"></span> AlertRuleInstanceTotalsBuilder
---
# <span class="badge builder"></span> AlertRuleInstanceTotalsBuilder

Count of alert instances per state. error also counts instances whose evaluation

errored but were mapped to another state via execErrState, so it can overlap.

## Constructor

```java
new AlertRuleInstanceTotalsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AlertRuleInstanceTotals build()
```

### <span class="badge object-method"></span> error

```java
public AlertRuleInstanceTotalsBuilder error(Long error)
```

### <span class="badge object-method"></span> firing

```java
public AlertRuleInstanceTotalsBuilder firing(Long firing)
```

### <span class="badge object-method"></span> healthy

```java
public AlertRuleInstanceTotalsBuilder healthy(Long healthy)
```

### <span class="badge object-method"></span> nodata

```java
public AlertRuleInstanceTotalsBuilder nodata(Long nodata)
```

### <span class="badge object-method"></span> pending

```java
public AlertRuleInstanceTotalsBuilder pending(Long pending)
```

### <span class="badge object-method"></span> recovering

```java
public AlertRuleInstanceTotalsBuilder recovering(Long recovering)
```

## See also

 * <span class="badge object-type-class"></span> [AlertRuleInstanceTotals](./object-AlertRuleInstanceTotals.md)
