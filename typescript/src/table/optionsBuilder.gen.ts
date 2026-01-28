// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as table from '../table';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<table.Options> {
    protected readonly internal: table.Options;

    constructor() {
        this.internal = table.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): table.Options {
        return this.internal;
    }

    // Represents the index of the selected frame
    frameIndex(frameIndex: number): this {
        this.internal.frameIndex = frameIndex;
        return this;
    }

    // Controls whether the panel should show the header
    showHeader(showHeader: boolean): this {
        this.internal.showHeader = showHeader;
        return this;
    }

    // Controls whether the header should show icons for the column types
    showTypeIcons(showTypeIcons: boolean): this {
        this.internal.showTypeIcons = showTypeIcons;
        return this;
    }

    // Used to control row sorting
    sortBy(sortBy: cog.Builder<common.TableSortByFieldState>[]): this {
        const sortByResources = sortBy.map(builder1 => builder1.build());
        this.internal.sortBy = sortByResources;
        return this;
    }

    // Controls footer options
    footer(footer: cog.Builder<common.TableFooterOptions>): this {
        const footerResource = footer.build();
        this.internal.footer = footerResource;
        return this;
    }

    // Controls the height of the rows
    cellHeight(cellHeight: common.TableCellHeight): this {
        this.internal.cellHeight = cellHeight;
        return this;
    }
}

