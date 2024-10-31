// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// Json view cell options
export class TableImageCellOptionsBuilder implements cog.Builder<common.TableImageCellOptions> {
    protected readonly internal: common.TableImageCellOptions;

    constructor() {
        this.internal = common.defaultTableImageCellOptions();
        this.internal.type = "image";
    }

    /**
     * Builds the object.
     */
    build(): common.TableImageCellOptions {
        return this.internal;
    }

    alt(alt: string): this {
        this.internal.alt = alt;
        return this;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }
}
