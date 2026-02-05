import * as cog from '../cog';
import * as common from '../common';
export declare class TableSortByFieldStateBuilder implements cog.Builder<common.TableSortByFieldState> {
    protected readonly internal: common.TableSortByFieldState;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.TableSortByFieldState;
    displayName(displayName: string): this;
    desc(desc: boolean): this;
}
