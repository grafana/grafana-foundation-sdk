<?php

namespace Grafana\Foundation\Heatmap;

final class PanelConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\Panel $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Heatmap\PanelBuilder())',
        ];
            if ($input->id !== null) {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->targets !== null && count($input->targets) >= 1) {
    
        
    $buffer = 'targets(';
        $tmparg0 = [];
        foreach ($input->targets as $arg1) {
        $tmptargetsarg1 = \Grafana\Foundation\Cog\Runtime::get()->convertDataqueryToCode($arg1, );
        $tmparg0[] = $tmptargetsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->title !== null && $input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->description !== null && $input->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->transparent !== false) {
    
        
    $buffer = 'transparent(';
        $arg0 =\var_export($input->transparent, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->gridPos !== null) {
    
        
    $buffer = 'gridPos(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\GridPos(h: '.\var_export($input->gridPos->h, true).',w: '.\var_export($input->gridPos->w, true).',x: '.\var_export($input->gridPos->x, true).',y: '.\var_export($input->gridPos->y, true).','.(($input->gridPos->static !== null) ? 'static: '.\var_export($input->gridPos->static, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->gridPos !== null && $input->gridPos->h !== 9) {
    
        
    $buffer = 'height(';
        $arg0 =\var_export($input->gridPos->h, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->gridPos !== null && $input->gridPos->w !== 12) {
    
        
    $buffer = 'span(';
        $arg0 =\var_export($input->gridPos->w, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->links !== null && count($input->links) >= 1) {
    
        
    $buffer = 'links(';
        $tmparg0 = [];
        foreach ($input->links as $arg1) {
        $tmplinksarg1 = \Grafana\Foundation\Dashboard\DashboardLinkConverter::convert($arg1);
        $tmparg0[] = $tmplinksarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->repeat !== null && $input->repeat !== "") {
    
        
    $buffer = 'repeat(';
        $arg0 =\var_export($input->repeat, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->repeatDirection !== null) {
    
        
    $buffer = 'repeatDirection(';
        $arg0 ='\Grafana\Foundation\Dashboard\PanelRepeatDirection::fromValue("'.$input->repeatDirection.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->maxDataPoints !== null) {
    
        
    $buffer = 'maxDataPoints(';
        $arg0 =\var_export($input->maxDataPoints, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->transformations) >= 1) {
    
        
    $buffer = 'transformations(';
        $tmparg0 = [];
        foreach ($input->transformations as $arg1) {
        $tmptransformationsarg1 ='(new \Grafana\Foundation\Dashboard\DataTransformerConfig(id: '.\var_export($arg1->id, true).','.(($arg1->disabled !== null) ? 'disabled: '.\var_export($arg1->disabled, true).', ' : '').''.(($arg1->filter !== null) ? 'filter: '.'(new \Grafana\Foundation\Dashboard\MatcherConfig(id: '.\var_export($arg1->filter->id, true).','.(($arg1->filter->options !== null) ? 'options: '.\var_export($arg1->filter->options, true).', ' : '').'))'.', ' : '').'options: '.\var_export($arg1->options, true).',))';
        $tmparg0[] = $tmptransformationsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->interval !== null && $input->interval !== "") {
    
        
    $buffer = 'interval(';
        $arg0 =\var_export($input->interval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeFrom !== null && $input->timeFrom !== "") {
    
        
    $buffer = 'timeFrom(';
        $arg0 =\var_export($input->timeFrom, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeShift !== null && $input->timeShift !== "") {
    
        
    $buffer = 'timeShift(';
        $arg0 =\var_export($input->timeShift, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->libraryPanel !== null) {
    
        
    $buffer = 'libraryPanel(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\LibraryPanelRef(name: '.\var_export($input->libraryPanel->name, true).',uid: '.\var_export($input->libraryPanel->uid, true).',))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig->defaults->displayName !== null && $input->fieldConfig->defaults->displayName !== "") {
    
        
    $buffer = 'displayName(';
        $arg0 =\var_export($input->fieldConfig->defaults->displayName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig->defaults->unit !== null && $input->fieldConfig->defaults->unit !== "") {
    
        
    $buffer = 'unit(';
        $arg0 =\var_export($input->fieldConfig->defaults->unit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig->defaults->decimals !== null) {
    
        
    $buffer = 'decimals(';
        $arg0 =\var_export($input->fieldConfig->defaults->decimals, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig->defaults->min !== null) {
    
        
    $buffer = 'min(';
        $arg0 =\var_export($input->fieldConfig->defaults->min, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig->defaults->max !== null) {
    
        
    $buffer = 'max(';
        $arg0 =\var_export($input->fieldConfig->defaults->max, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig->defaults->mappings !== null && count($input->fieldConfig->defaults->mappings) >= 1) {
    
        
    $buffer = 'mappings(';
        $tmparg0 = [];
        foreach ($input->fieldConfig->defaults->mappings as $arg1) {
        switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Dashboard\ValueMap:
                $disjunctionarg1 = \Grafana\Foundation\Dashboard\ValueMapConverter::convert($arg1);
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboard\RangeMap:
                $disjunctionarg1 = \Grafana\Foundation\Dashboard\RangeMapConverter::convert($arg1);
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboard\RegexMap:
                $disjunctionarg1 = \Grafana\Foundation\Dashboard\RegexMapConverter::convert($arg1);
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboard\SpecialValueMap:
                $disjunctionarg1 = \Grafana\Foundation\Dashboard\SpecialValueMapConverter::convert($arg1);
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $tmparg0[] = $tmpmappingsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig->defaults->thresholds !== null) {
    
        
    $buffer = 'thresholds(';
        $arg0 = \Grafana\Foundation\Dashboard\ThresholdsConfigConverter::convert($input->fieldConfig->defaults->thresholds);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig->defaults->color !== null) {
    
        
    $buffer = 'colorScheme(';
        $arg0 = \Grafana\Foundation\Dashboard\FieldColorConverter::convert($input->fieldConfig->defaults->color);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig->defaults->noValue !== null && $input->fieldConfig->defaults->noValue !== "") {
    
        
    $buffer = 'noValue(';
        $arg0 =\var_export($input->fieldConfig->defaults->noValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->fieldConfig->overrides) >= 1) {
    
        
    $buffer = 'overrides(';
        $tmparg0 = [];
        foreach ($input->fieldConfig->overrides as $arg1) {
        $tmpoverridesarg1 = \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverridesConverter::convert($arg1);
        $tmparg0[] = $tmpoverridesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->fieldConfig->overrides) >= 1) {
    foreach ($input->fieldConfig->overrides as $item) {
        
    $buffer = 'withOverride(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\MatcherConfig(id: '.\var_export($item->matcher->id, true).','.(($item->matcher->options !== null) ? 'options: '.\var_export($item->matcher->options, true).', ' : '').'))';
        $buffer .= $arg0;
        $buffer .= ', ';
        $tmparg1 = [];
        foreach ($item->properties as $arg1) {
        $tmppropertiesarg1 ='(new \Grafana\Foundation\Dashboard\DynamicConfigValue(id: '.\var_export($arg1->id, true).','.(($arg1->value !== null) ? 'value: '.\var_export($arg1->value, true).', ' : '').'))';
        $tmparg1[] = $tmppropertiesarg1;
        }
        $arg1 = "[" . implode(", \n", $tmparg1) . "]";
        $buffer .= $arg1;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    }
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->calculate !== null && $input->options->calculate !== false) {
    
        
    $buffer = 'calculate(';
        $arg0 =\var_export($input->options->calculate, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->calculation !== null) {
    
        
    $buffer = 'calculation(';
        $arg0 = \Grafana\Foundation\Common\HeatmapCalculationOptionsConverter::convert($input->options->calculation);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options) {
    
        
    $buffer = 'color(';
        $arg0 = \Grafana\Foundation\Heatmap\HeatmapColorOptionsConverter::convert($input->options->color);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->filterValues !== null) {
    
        
    $buffer = 'filterValues(';
        $arg0 = \Grafana\Foundation\Heatmap\FilterValueRangeConverter::convert($input->options->filterValues);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->rowsFrame !== null) {
    
        
    $buffer = 'rowsFrame(';
        $arg0 = \Grafana\Foundation\Heatmap\RowsHeatmapOptionsConverter::convert($input->options->rowsFrame);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options) {
    
        
    $buffer = 'showValue(';
        $arg0 ='\Grafana\Foundation\Common\VisibilityMode::fromValue("'.$input->options->showValue.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->cellGap !== null && $input->options->cellGap !== 1) {
    
        
    $buffer = 'cellGap(';
        $arg0 =\var_export($input->options->cellGap, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->cellRadius !== null) {
    
        
    $buffer = 'cellRadius(';
        $arg0 =\var_export($input->options->cellRadius, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->cellValues !== null) {
    
        
    $buffer = 'cellValues(';
        $arg0 = \Grafana\Foundation\Heatmap\CellValuesConverter::convert($input->options->cellValues);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options) {
    
        
    $buffer = 'yAxis(';
        $arg0 = \Grafana\Foundation\Heatmap\YAxisConfigConverter::convert($input->options->yAxis);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->legend->show === true) {
    
        
    $buffer = 'showLegend(';
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->legend->show === false) {
    
        
    $buffer = 'hideLegend(';
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->tooltip->show === true) {
    
        
    $buffer = 'showTooltip(';
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->tooltip->show === false) {
    
        
    $buffer = 'hideTooltip(';
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->tooltip->yHistogram !== null && $input->options->tooltip->yHistogram === true) {
    
        
    $buffer = 'showYHistogram(';
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->tooltip->yHistogram !== null && $input->options->tooltip->yHistogram === false) {
    
        
    $buffer = 'hideYHistogram(';
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Heatmap\Options && $input->options->exemplars->color !== "" && $input->options->exemplars->color !== "rgba(255,0,255,0.7)") {
    
        
    $buffer = 'exemplarsColor(';
        $arg0 =\var_export($input->options->exemplars->color, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Heatmap\FieldConfig && $input->fieldConfig->defaults->custom->scaleDistribution !== null) {
    
        
    $buffer = 'scaleDistribution(';
        $arg0 = \Grafana\Foundation\Common\ScaleDistributionConfigConverter::convert($input->fieldConfig->defaults->custom->scaleDistribution);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Heatmap\FieldConfig && $input->fieldConfig->defaults->custom->hideFrom !== null) {
    
        
    $buffer = 'hideFrom(';
        $arg0 = \Grafana\Foundation\Common\HideSeriesConfigConverter::convert($input->fieldConfig->defaults->custom->hideFrom);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

