// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// Sort by field state
export class TableSortByFieldStateBuilder implements cog.Builder<common.TableSortByFieldState> {
    protected readonly internal: common.TableSortByFieldState;

    constructor() {
        this.internal = common.defaultTableSortByFieldState();
    }

    /**
     * Builds the object.
     */
    build(): common.TableSortByFieldState {
        return this.internal;
    }

    // Sets the display name of the field to sort by
    displayName(displayName: string): this {
        this.internal.displayName = displayName;
        return this;
    }

    // Flag used to indicate descending sort order
    desc(desc: boolean): this {
        this.internal.desc = desc;
        return this;
    }
}
