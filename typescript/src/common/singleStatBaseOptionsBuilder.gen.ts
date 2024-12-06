// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class SingleStatBaseOptionsBuilder implements cog.Builder<common.SingleStatBaseOptions> {
    protected readonly internal: common.SingleStatBaseOptions;

    constructor() {
        this.internal = common.defaultSingleStatBaseOptions();
    }

    /**
     * Builds the object.
     */
    build(): common.SingleStatBaseOptions {
        return this.internal;
    }

    reduceOptions(reduceOptions: cog.Builder<common.ReduceDataOptions>): this {
        const reduceOptionsResource = reduceOptions.build();
        this.internal.reduceOptions = reduceOptionsResource;
        return this;
    }

    text(text: cog.Builder<common.VizTextDisplayOptions>): this {
        const textResource = text.build();
        this.internal.text = textResource;
        return this;
    }

    orientation(orientation: common.VizOrientation): this {
        this.internal.orientation = orientation;
        return this;
    }
}
