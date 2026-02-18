// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as playlistv0alpha1 from '../playlistv0alpha1';

export class PlaylistBuilder implements cog.Builder<playlistv0alpha1.Playlist> {
    protected readonly internal: playlistv0alpha1.Playlist;

    constructor(title: string) {
        this.internal = playlistv0alpha1.defaultPlaylist();
        this.internal.title = title;
    }

    /**
     * Builds the object.
     */
    build(): playlistv0alpha1.Playlist {
        return this.internal;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    interval(interval: string): this {
        this.internal.interval = interval;
        return this;
    }

    items(items: cog.Builder<playlistv0alpha1.Item>[]): this {
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.items = itemsResources;
        return this;
    }

    item(item: cog.Builder<playlistv0alpha1.Item>): this {
        if (!this.internal.items) {
            this.internal.items = [];
        }
        const itemResource = item.build();
        this.internal.items.push(itemResource);
        return this;
    }
}

