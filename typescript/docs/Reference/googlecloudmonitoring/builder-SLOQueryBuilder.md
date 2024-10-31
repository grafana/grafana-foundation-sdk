---
title: <span class="badge builder"></span> SLOQueryBuilder
---
# <span class="badge builder"></span> SLOQueryBuilder

## Constructor

```typescript
new SLOQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> alignmentPeriod

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```typescript
alignmentPeriod(alignmentPeriod: string)
```

### <span class="badge object-method"></span> goal

SLO goal value.

```typescript
goal(goal: number)
```

### <span class="badge object-method"></span> lookbackPeriod

Specific lookback period for the SLO.

```typescript
lookbackPeriod(lookbackPeriod: string)
```

### <span class="badge object-method"></span> perSeriesAligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```typescript
perSeriesAligner(perSeriesAligner: string)
```

### <span class="badge object-method"></span> projectName

GCP project to execute the query against.

```typescript
projectName(projectName: string)
```

### <span class="badge object-method"></span> selectorName

SLO selector.

```typescript
selectorName(selectorName: string)
```

### <span class="badge object-method"></span> serviceId

ID for the service the SLO is in.

```typescript
serviceId(serviceId: string)
```

### <span class="badge object-method"></span> serviceName

Name for the service the SLO is in.

```typescript
serviceName(serviceName: string)
```

### <span class="badge object-method"></span> sloId

ID for the SLO.

```typescript
sloId(sloId: string)
```

### <span class="badge object-method"></span> sloName

Name of the SLO.

```typescript
sloName(sloName: string)
```

## See also

 * <span class="badge object-type-interface"></span> [SLOQuery](./object-SLOQuery.md)
