// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as stat from '../stat';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<stat.Options> {
    protected readonly internal: stat.Options;

    constructor() {
        this.internal = stat.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): stat.Options {
        return this.internal;
    }

    graphMode(graphMode: common.BigValueGraphMode): this {
        this.internal.graphMode = graphMode;
        return this;
    }

    colorMode(colorMode: common.BigValueColorMode): this {
        this.internal.colorMode = colorMode;
        return this;
    }

    justifyMode(justifyMode: common.BigValueJustifyMode): this {
        this.internal.justifyMode = justifyMode;
        return this;
    }

    textMode(textMode: common.BigValueTextMode): this {
        this.internal.textMode = textMode;
        return this;
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

    wideLayout(wideLayout: boolean): this {
        this.internal.wideLayout = wideLayout;
        return this;
    }

    orientation(orientation: common.VizOrientation): this {
        this.internal.orientation = orientation;
        return this;
    }
}

