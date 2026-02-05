---
title: <span class="badge builder"></span> ContactPointBuilder
---
# <span class="badge builder"></span> ContactPointBuilder

## Constructor

```typescript
new ContactPointBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> disableResolveMessage

```typescript
disableResolveMessage(disableResolveMessage: boolean)
```

### <span class="badge object-method"></span> name

Name is used as grouping key in the UI. Contact points with the

same name will be grouped in the UI.

```typescript
name(name: string)
```

### <span class="badge object-method"></span> provenance

```typescript
provenance(provenance: string)
```

### <span class="badge object-method"></span> settings

```typescript
settings(settings: alerting.Json)
```

### <span class="badge object-method"></span> type

```typescript
type(type: "alertmanager" | "dingding" | "discord" | "email" | "googlechat" | "kafka" | "line" | "opsgenie" | "pagerduty" | "pushover" | "sensugo" | "slack" | "teams" | "telegram" | "threema" | "victorops" | "webhook" | "wecom")
```

### <span class="badge object-method"></span> uid

UID is the unique identifier of the contact point. The UID can be

set by the user.

```typescript
uid(uid: string)
```

## See also

 * <span class="badge object-type-interface"></span> [ContactPoint](./object-ContactPoint.md)
