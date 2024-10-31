// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as playlist from '../playlist';

export class PlaylistBuilder implements cog.Builder<playlist.Playlist> {
    protected readonly internal: playlist.Playlist;

    constructor() {
        this.internal = playlist.defaultPlaylist();
    }

    /**
     * Builds the object.
     */
    build(): playlist.Playlist {
        return this.internal;
    }

    // Unique playlist identifier. Generated on creation, either by the
    // creator of the playlist of by the application.
    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }

    // Name of the playlist.
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    // Interval sets the time between switching views in a playlist.
    // FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
    interval(interval: string): this {
        this.internal.interval = interval;
        return this;
    }

    // The ordered list of items that the playlist will iterate over.
    // FIXME! This should not be optional, but changing it makes the godegen awkward
    items(items: cog.Builder<playlist.PlaylistItem>[]): this {
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.items = itemsResources;
        return this;
    }
}
