import * as common from '../common';
export interface Options {
    frameIndex: number;
    showHeader: boolean;
    showTypeIcons?: boolean;
    sortBy?: common.TableSortByFieldState[];
    footer?: common.TableFooterOptions;
    cellHeight?: common.TableCellHeight;
}
export declare const defaultOptions: () => Options;
export type FieldConfig = common.TableFieldOptions;
export declare const defaultFieldConfig: () => FieldConfig;
