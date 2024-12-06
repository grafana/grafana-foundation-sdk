// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// Auto mode table cell options
export class TableAutoCellOptionsBuilder implements cog.Builder<common.TableAutoCellOptions> {
    protected readonly internal: common.TableAutoCellOptions;

    constructor() {
        this.internal = common.defaultTableAutoCellOptions();
        this.internal.type = "auto";
    }

    /**
     * Builds the object.
     */
    build(): common.TableAutoCellOptions {
        return this.internal;
    }

    wrapText(wrapText: boolean): this {
        this.internal.wrapText = wrapText;
        return this;
    }
}
