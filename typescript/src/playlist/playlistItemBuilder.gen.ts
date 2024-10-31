// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as playlist from '../playlist';

export class PlaylistItemBuilder implements cog.Builder<playlist.PlaylistItem> {
    protected readonly internal: playlist.PlaylistItem;

    constructor() {
        this.internal = playlist.defaultPlaylistItem();
    }

    /**
     * Builds the object.
     */
    build(): playlist.PlaylistItem {
        return this.internal;
    }

    // Type of the item.
    type(type: "dashboard_by_uid" | "dashboard_by_id" | "dashboard_by_tag"): this {
        this.internal.type = type;
        return this;
    }

    // Value depends on type and describes the playlist item.
    // 
    //  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
    //  is not portable as the numerical identifier is non-deterministic between different instances.
    //  Will be replaced by dashboard_by_uid in the future. (deprecated)
    //  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
    //  dashboards behind the tag will be added to the playlist.
    //  - dashboard_by_uid: The value is the dashboard UID
    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    // Title is an unused property -- it will be removed in the future
    title(title: string): this {
        this.internal.title = title;
        return this;
    }
}
