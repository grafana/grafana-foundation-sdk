// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// A dashboard snapshot shares an interactive dashboard publicly.
// It is a read-only version of a dashboard, and is not editable.
// It is possible to create a snapshot of a snapshot.
// Grafana strips away all sensitive information from the dashboard.
// Sensitive information stripped: queries (metric, template,annotation) and panel links.
export class SnapshotBuilder implements cog.Builder<dashboard.Snapshot> {
    protected readonly internal: dashboard.Snapshot;

    constructor() {
        this.internal = dashboard.defaultSnapshot();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.Snapshot {
        return this.internal;
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

    // url of the snapshot, if snapshot was shared internally
    url(url: string): this {
        this.internal.url = url;
        return this;
    }

    dashboard(dashboard: cog.Builder<dashboard.Dashboard>): this {
        const dashboardResource = dashboard.build();
        this.internal.dashboard = dashboardResource;
        return this;
    }
}
