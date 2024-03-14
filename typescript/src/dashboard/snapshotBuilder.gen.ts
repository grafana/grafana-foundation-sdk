// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// A dashboard snapshot shares an interactive dashboard publicly.
// It is a read-only version of a dashboard, and is not editable.
// It is possible to create a snapshot of a snapshot.
// Grafana strips away all sensitive information from the dashboard.
// Sensitive information stripped: queries (metric, template,annotation) and panel links.
export class SnapshotBuilder implements cog.Builder<dashboard.Snapshot> {
    private readonly internal: dashboard.Snapshot;

    constructor() {
        this.internal = dashboard.defaultSnapshot();
    }

    build(): dashboard.Snapshot {
        return this.internal;
    }

    // Time when the snapshot was created
    created(created: string): this {
        this.internal.created = created;
        return this;
    }

    // Time when the snapshot expires, default is never to expire
    expires(expires: string): this {
        this.internal.expires = expires;
        return this;
    }

    // Is the snapshot saved in an external grafana instance
    external(external: boolean): this {
        this.internal.external = external;
        return this;
    }

    // external url, if snapshot was shared in external grafana instance
    externalUrl(externalUrl: string): this {
        this.internal.externalUrl = externalUrl;
        return this;
    }

    // original url, url of the dashboard that was snapshotted
    originalUrl(originalUrl: string): this {
        this.internal.originalUrl = originalUrl;
        return this;
    }

    // Unique identifier of the snapshot
    id(id: number): this {
        this.internal.id = id;
        return this;
    }

    // Optional, defined the unique key of the snapshot, required if external is true
    key(key: string): this {
        this.internal.key = key;
        return this;
    }

    // Optional, name of the snapshot
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    // org id of the snapshot
    orgId(orgId: number): this {
        this.internal.orgId = orgId;
        return this;
    }

    // last time when the snapshot was updated
    updated(updated: string): this {
        this.internal.updated = updated;
        return this;
    }

    // url of the snapshot, if snapshot was shared internally
    url(url: string): this {
        this.internal.url = url;
        return this;
    }

    // user id of the snapshot creator
    userId(userId: number): this {
        this.internal.userId = userId;
        return this;
    }
}
