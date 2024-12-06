// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// Colored text cell options
export class TableColorTextCellOptionsBuilder implements cog.Builder<common.TableColorTextCellOptions> {
    protected readonly internal: common.TableColorTextCellOptions;

    constructor() {
        this.internal = common.defaultTableColorTextCellOptions();
        this.internal.type = "color-text";
    }

    /**
     * Builds the object.
     */
    build(): common.TableColorTextCellOptions {
        return this.internal;
    }

    wrapText(wrapText: boolean): this {
        this.internal.wrapText = wrapText;
        return this;
    }
}
