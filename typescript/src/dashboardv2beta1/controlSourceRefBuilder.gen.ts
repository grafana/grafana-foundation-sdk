// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class ControlSourceRefBuilder implements cog.Builder<dashboardv2beta1.ControlSourceRef> {
    protected readonly internal: dashboardv2beta1.ControlSourceRef;

    constructor() {
        this.internal = dashboardv2beta1.defaultControlSourceRef();
        this.internal.type = "datasource";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ControlSourceRef {
        return this.internal;
    }

    // The plugin type-id
    group(group: string): this {
        this.internal.group = group;
        return this;
    }
}

