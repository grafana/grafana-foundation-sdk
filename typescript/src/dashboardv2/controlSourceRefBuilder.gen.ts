// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class ControlSourceRefBuilder implements cog.Builder<dashboardv2.ControlSourceRef> {
    protected readonly internal: dashboardv2.ControlSourceRef;

    constructor() {
        this.internal = dashboardv2.defaultControlSourceRef();
        this.internal.type = "datasource";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.ControlSourceRef {
        return this.internal;
    }

    // The plugin type-id
    group(group: string): this {
        this.internal.group = group;
        return this;
    }
}

