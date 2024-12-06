---
title: <span class="badge builder"></span> NotificationTemplateBuilder
---
# <span class="badge builder"></span> NotificationTemplateBuilder

## Constructor

```go
func NewNotificationTemplateBuilder() *NotificationTemplateBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *NotificationTemplateBuilder) Build() (NotificationTemplate, error)
```

### <span class="badge object-method"></span> Name

```go
func (builder *NotificationTemplateBuilder) Name(name string) *NotificationTemplateBuilder
```

### <span class="badge object-method"></span> Provenance

```go
func (builder *NotificationTemplateBuilder) Provenance(provenance alerting.Provenance) *NotificationTemplateBuilder
```

### <span class="badge object-method"></span> Template

```go
func (builder *NotificationTemplateBuilder) Template(template string) *NotificationTemplateBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *NotificationTemplateBuilder) Version(version string) *NotificationTemplateBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [NotificationTemplate](./object-NotificationTemplate.md)
