---
title: <span class="badge builder"></span> ExprTypeSqlResultAssertionsBuilder
---
# <span class="badge builder"></span> ExprTypeSqlResultAssertionsBuilder

## Constructor

```php
new ExprTypeSqlResultAssertionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> maxFrames

Maximum frame count

```php
maxFrames(int $maxFrames)
```

### <span class="badge object-method"></span> type

Type asserts that the frame matches a known type structure.

Possible enum values:

 - `""` 

 - `"timeseries-wide"` 

 - `"timeseries-long"` 

 - `"timeseries-many"` 

 - `"timeseries-multi"` 

 - `"directory-listing"` 

 - `"table"` 

 - `"numeric-wide"` 

 - `"numeric-multi"` 

 - `"numeric-long"` 

 - `"log-lines"` 

```php
type(\Grafana\Foundation\Expr\TypeSqlType $type)
```

### <span class="badge object-method"></span> typeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

@param array<int> $typeVersion

```php
typeVersion(array $typeVersion)
```

## See also

 * <span class="badge object-type-class"></span> [ExprTypeSqlResultAssertions](./object-ExprTypeSqlResultAssertions.md)
