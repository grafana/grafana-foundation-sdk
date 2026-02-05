import * as cog from '../cog';
import * as dashboard from '../dashboard';
export declare class SnapshotBuilder implements cog.Builder<dashboard.Snapshot> {
    protected readonly internal: dashboard.Snapshot;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboard.Snapshot;
    expires(expires: string): this;
    external(external: boolean): this;
    externalUrl(externalUrl: string): this;
    originalUrl(originalUrl: string): this;
    id(id: number): this;
    key(key: string): this;
    name(name: string): this;
    orgId(orgId: number): this;
    url(url: string): this;
    dashboard(dashboard: cog.Builder<dashboard.Dashboard>): this;
}
