"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.SnapshotBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
// A dashboard snapshot shares an interactive dashboard publicly.
// It is a read-only version of a dashboard, and is not editable.
// It is possible to create a snapshot of a snapshot.
// Grafana strips away all sensitive information from the dashboard.
// Sensitive information stripped: queries (metric, template,annotation) and panel links.
class SnapshotBuilder {
    constructor() {
        this.internal = dashboard.defaultSnapshot();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Time when the snapshot expires, default is never to expire
    expires(expires) {
        this.internal.expires = expires;
        return this;
    }
    // Is the snapshot saved in an external grafana instance
    external(external) {
        this.internal.external = external;
        return this;
    }
    // external url, if snapshot was shared in external grafana instance
    externalUrl(externalUrl) {
        this.internal.externalUrl = externalUrl;
        return this;
    }
    // original url, url of the dashboard that was snapshotted
    originalUrl(originalUrl) {
        this.internal.originalUrl = originalUrl;
        return this;
    }
    // Unique identifier of the snapshot
    id(id) {
        this.internal.id = id;
        return this;
    }
    // Optional, defined the unique key of the snapshot, required if external is true
    key(key) {
        this.internal.key = key;
        return this;
    }
    // Optional, name of the snapshot
    name(name) {
        this.internal.name = name;
        return this;
    }
    // org id of the snapshot
    orgId(orgId) {
        this.internal.orgId = orgId;
        return this;
    }
    // url of the snapshot, if snapshot was shared internally
    url(url) {
        this.internal.url = url;
        return this;
    }
    dashboard(dashboard) {
        const dashboardResource = dashboard.build();
        this.internal.dashboard = dashboardResource;
        return this;
    }
}
exports.SnapshotBuilder = SnapshotBuilder;
//# sourceMappingURL=snapshotBuilder.gen.js.map