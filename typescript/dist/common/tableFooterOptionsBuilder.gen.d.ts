import * as cog from '../cog';
import * as common from '../common';
export declare class TableFooterOptionsBuilder implements cog.Builder<common.TableFooterOptions> {
    protected readonly internal: common.TableFooterOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.TableFooterOptions;
    show(show: boolean): this;
    reducer(reducer: string[]): this;
    fields(fields: string[]): this;
    enablePagination(enablePagination: boolean): this;
    countRows(countRows: boolean): this;
}
