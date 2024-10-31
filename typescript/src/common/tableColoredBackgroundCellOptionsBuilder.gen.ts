// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// Colored background cell options
export class TableColoredBackgroundCellOptionsBuilder implements cog.Builder<common.TableColoredBackgroundCellOptions> {
    protected readonly internal: common.TableColoredBackgroundCellOptions;

    constructor() {
        this.internal = common.defaultTableColoredBackgroundCellOptions();
        this.internal.type = "color-background";
    }

    /**
     * Builds the object.
     */
    build(): common.TableColoredBackgroundCellOptions {
        return this.internal;
    }

    mode(mode: common.TableCellBackgroundDisplayMode): this {
        this.internal.mode = mode;
        return this;
    }
}
