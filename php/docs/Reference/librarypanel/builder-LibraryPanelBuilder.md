---
title: <span class="badge builder"></span> LibraryPanelBuilder
---
# <span class="badge builder"></span> LibraryPanelBuilder

## Constructor

```php
new LibraryPanelBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> description

Panel description

```php
description(string $description)
```

### <span class="badge object-method"></span> folderUid

Folder UID

```php
folderUid(string $folderUid)
```

### <span class="badge object-method"></span> meta

Object storage metadata

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Librarypanel\LibraryElementDTOMeta> $meta

```php
meta(\Grafana\Foundation\Cog\Builder $meta)
```

### <span class="badge object-method"></span> model

TODO: should be the same panel schema defined in dashboard

Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel> $model

```php
model(\Grafana\Foundation\Cog\Builder $model)
```

### <span class="badge object-method"></span> name

Panel name (also saved in the model)

```php
name(string $name)
```

### <span class="badge object-method"></span> schemaVersion

Dashboard version when this was saved (zero if unknown)

```php
schemaVersion(int $schemaVersion)
```

### <span class="badge object-method"></span> type

The panel type (from inside the model)

```php
type(string $type)
```

### <span class="badge object-method"></span> uid

Library element UID

```php
uid(string $uid)
```

### <span class="badge object-method"></span> version

panel version, incremented each time the dashboard is updated.

```php
version(int $version)
```

## See also

 * <span class="badge object-type-class"></span> [LibraryPanel](./object-LibraryPanel.md)
