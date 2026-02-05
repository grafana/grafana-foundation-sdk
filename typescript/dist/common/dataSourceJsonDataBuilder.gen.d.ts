import * as cog from '../cog';
import * as common from '../common';
export declare class DataSourceJsonDataBuilder implements cog.Builder<common.DataSourceJsonData> {
    protected readonly internal: common.DataSourceJsonData;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.DataSourceJsonData;
    authType(authType: string): this;
    defaultRegion(defaultRegion: string): this;
    profile(profile: string): this;
    manageAlerts(manageAlerts: boolean): this;
    alertmanagerUid(alertmanagerUid: string): this;
}
