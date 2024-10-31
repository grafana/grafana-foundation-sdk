// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class VizTextDisplayOptionsBuilder implements cog.Builder<common.VizTextDisplayOptions> {
    protected readonly internal: common.VizTextDisplayOptions;

    constructor() {
        this.internal = common.defaultVizTextDisplayOptions();
    }

    /**
     * Builds the object.
     */
    build(): common.VizTextDisplayOptions {
        return this.internal;
    }

    // Explicit title text size
    titleSize(titleSize: number): this {
        this.internal.titleSize = titleSize;
        return this;
    }

    // Explicit value text size
    valueSize(valueSize: number): this {
        this.internal.valueSize = valueSize;
        return this;
    }
}
