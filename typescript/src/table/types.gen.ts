// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	// Represents the index of the selected frame
	frameIndex: number;
	// Controls whether the panel should show the header
	showHeader: boolean;
	// Controls whether the header should show icons for the column types
	showTypeIcons?: boolean;
	// Used to control row sorting
	sortBy?: common.TableSortByFieldState[];
	// Controls footer options
	footer?: common.TableFooterOptions;
	// Controls the height of the rows
	cellHeight?: common.TableCellHeight;
}

export const defaultOptions = (): Options => ({
	frameIndex: 0,
	showHeader: true,
	showTypeIcons: false,
	footer: { show: false, reducer: [
], countRows: false, },
	cellHeight: common.TableCellHeight.Sm,
});

