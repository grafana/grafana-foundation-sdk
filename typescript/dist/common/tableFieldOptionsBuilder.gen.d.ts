import * as cog from '../cog';
import * as common from '../common';
export declare class TableFieldOptionsBuilder implements cog.Builder<common.TableFieldOptions> {
    protected readonly internal: common.TableFieldOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.TableFieldOptions;
    width(width: number): this;
    minWidth(minWidth: number): this;
    align(align: common.FieldTextAlignment): this;
    displayMode(displayMode: common.TableCellDisplayMode): this;
    cellOptions(cellOptions: common.TableCellOptions): this;
    hidden(hidden: boolean): this;
    inspect(inspect: boolean): this;
    filterable(filterable: boolean): this;
    hideHeader(hideHeader: boolean): this;
}
