// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class DataSourceJsonDataBuilder implements cog.Builder<common.DataSourceJsonData> {
    protected readonly internal: common.DataSourceJsonData;

    constructor() {
        this.internal = common.defaultDataSourceJsonData();
    }

    /**
     * Builds the object.
     */
    build(): common.DataSourceJsonData {
        return this.internal;
    }

    authType(authType: string): this {
        this.internal.authType = authType;
        return this;
    }

    defaultRegion(defaultRegion: string): this {
        this.internal.defaultRegion = defaultRegion;
        return this;
    }

    profile(profile: string): this {
        this.internal.profile = profile;
        return this;
    }

    manageAlerts(manageAlerts: boolean): this {
        this.internal.manageAlerts = manageAlerts;
        return this;
    }

    alertmanagerUid(alertmanagerUid: string): this {
        this.internal.alertmanagerUid = alertmanagerUid;
        return this;
    }
}
