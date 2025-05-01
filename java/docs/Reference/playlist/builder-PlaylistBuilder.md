---
title: <span class="badge builder"></span> PlaylistBuilder
---
# <span class="badge builder"></span> PlaylistBuilder

## Constructor

```java
new PlaylistBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Playlist build()
```

### <span class="badge object-method"></span> interval

Interval sets the time between switching views in a playlist.

FIXME: Is this based on a standardized format or what options are available? Can datemath be used?

```java
public PlaylistBuilder interval(String interval)
```

### <span class="badge object-method"></span> items

The ordered list of items that the playlist will iterate over.

FIXME! This should not be optional, but changing it makes the godegen awkward

```java
public PlaylistBuilder items(List<PlaylistItem> items)
```

### <span class="badge object-method"></span> name

Name of the playlist.

```java
public PlaylistBuilder name(String name)
```

### <span class="badge object-method"></span> uid

Unique playlist identifier. Generated on creation, either by the

creator of the playlist of by the application.

```java
public PlaylistBuilder uid(String uid)
```

## See also

 * <span class="badge object-type-class"></span> [Playlist](./object-Playlist.md)
