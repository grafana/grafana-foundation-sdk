// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

// Source information for controls (e.g. variables or links)
export class DatasourceControlSourceRefBuilder implements cog.Builder<dashboardv2.DatasourceControlSourceRef> {
    protected readonly internal: dashboardv2.DatasourceControlSourceRef;

    constructor() {
        this.internal = dashboardv2.defaultDatasourceControlSourceRef();
        this.internal.type = "datasource";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.DatasourceControlSourceRef {
        return this.internal;
    }

    // The plugin type-id
    group(group: string): this {
        this.internal.group = group;
        return this;
    }
}

