---
title: <span class="badge builder"></span> PlaylistBuilder
---
# <span class="badge builder"></span> PlaylistBuilder

## Constructor

```typescript
new PlaylistBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> interval

Interval sets the time between switching views in a playlist.

FIXME: Is this based on a standardized format or what options are available? Can datemath be used?

```typescript
interval(interval: string)
```

### <span class="badge object-method"></span> items

The ordered list of items that the playlist will iterate over.

FIXME! This should not be optional, but changing it makes the godegen awkward

```typescript
items(items: cog.Builder<playlist.PlaylistItem>[])
```

### <span class="badge object-method"></span> name

Name of the playlist.

```typescript
name(name: string)
```

### <span class="badge object-method"></span> uid

Unique playlist identifier. Generated on creation, either by the

creator of the playlist of by the application.

```typescript
uid(uid: string)
```

## See also

 * <span class="badge object-type-interface"></span> [Playlist](./object-Playlist.md)
