---
title: <span class="badge object-type-interface"></span> Playlist
---
# <span class="badge object-type-interface"></span> Playlist

## Definition

```typescript
export interface Playlist {
	// Unique playlist identifier. Generated on creation, either by the
	// creator of the playlist of by the application.
	uid: string;
	// Name of the playlist.
	name: string;
	// Interval sets the time between switching views in a playlist.
	// FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
	interval: string;
	// The ordered list of items that the playlist will iterate over.
	// FIXME! This should not be optional, but changing it makes the godegen awkward
	items?: playlist.PlaylistItem[];
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [PlaylistBuilder](./builder-PlaylistBuilder.md)
