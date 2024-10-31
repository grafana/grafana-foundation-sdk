---
title: <span class="badge builder"></span> LineConfigBuilder
---
# <span class="badge builder"></span> LineConfigBuilder

## Constructor

```typescript
new LineConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> lineColor

```typescript
lineColor(lineColor: string)
```

### <span class="badge object-method"></span> lineInterpolation

```typescript
lineInterpolation(lineInterpolation: common.LineInterpolation)
```

### <span class="badge object-method"></span> lineStyle

```typescript
lineStyle(lineStyle: cog.Builder<common.LineStyle>)
```

### <span class="badge object-method"></span> lineWidth

```typescript
lineWidth(lineWidth: number)
```

### <span class="badge object-method"></span> spanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```typescript
spanNulls(spanNulls: boolean | number)
```

## See also

 * <span class="badge object-type-interface"></span> [LineConfig](./object-LineConfig.md)
