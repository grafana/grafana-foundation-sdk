// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as debug from '../debug';

export class UpdateConfigBuilder implements cog.Builder<debug.UpdateConfig> {
    protected readonly internal: debug.UpdateConfig;

    constructor() {
        this.internal = debug.defaultUpdateConfig();
    }

    /**
     * Builds the object.
     */
    build(): debug.UpdateConfig {
        return this.internal;
    }

    render(render: boolean): this {
        this.internal.render = render;
        return this;
    }

    dataChanged(dataChanged: boolean): this {
        this.internal.dataChanged = dataChanged;
        return this;
    }

    schemaChanged(schemaChanged: boolean): this {
        this.internal.schemaChanged = schemaChanged;
        return this;
    }
}
