---
title: <span class="badge builder"></span> AlertRuleInstanceTotalsBuilder
---
# <span class="badge builder"></span> AlertRuleInstanceTotalsBuilder

Count of alert instances per state. error also counts instances whose evaluation

errored but were mapped to another state via execErrState, so it can overlap.

## Constructor

```php
new AlertRuleInstanceTotalsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> error

```php
error(int $error)
```

### <span class="badge object-method"></span> firing

```php
firing(int $firing)
```

### <span class="badge object-method"></span> healthy

```php
healthy(int $healthy)
```

### <span class="badge object-method"></span> nodata

```php
nodata(int $nodata)
```

### <span class="badge object-method"></span> pending

```php
pending(int $pending)
```

### <span class="badge object-method"></span> recovering

```php
recovering(int $recovering)
```

## See also

 * <span class="badge object-type-class"></span> [AlertRuleInstanceTotals](./object-AlertRuleInstanceTotals.md)
