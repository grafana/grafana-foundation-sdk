---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```typescript
new OptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> alignValue

Controls value alignment on the timelines

```typescript
alignValue(alignValue: common.TimelineValueAlignment)
```

### <span class="badge object-method"></span> legend

```typescript
legend(legend: cog.Builder<common.VizLegendOptions>)
```

### <span class="badge object-method"></span> mergeValues

Merge equal consecutive values

```typescript
mergeValues(mergeValues: boolean)
```

### <span class="badge object-method"></span> perPage

Enables pagination when > 0

```typescript
perPage(perPage: number)
```

### <span class="badge object-method"></span> rowHeight

Controls the row height

```typescript
rowHeight(rowHeight: number)
```

### <span class="badge object-method"></span> showValue

Show timeline values on chart

```typescript
showValue(showValue: common.VisibilityMode)
```

### <span class="badge object-method"></span> timezone

```typescript
timezone(timezone: common.TimeZone[])
```

### <span class="badge object-method"></span> tooltip

```typescript
tooltip(tooltip: cog.Builder<common.VizTooltipOptions>)
```

## See also

 * <span class="badge object-type-interface"></span> [Options](./object-Options.md)
