<?php

namespace Grafana\Foundation\Trend;

final class PanelConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\Panel $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Trend\PanelBuilder())',
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
            if ($input->transparent !== null && $input->transparent !== false) {
    
        
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
            if ($input->maxPerRow !== null) {
    
        
    $buffer = 'maxPerRow(';
        $arg0 =\var_export($input->maxPerRow, true);
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
            if ($input->transformations !== null && count($input->transformations) >= 1) {
    
        
    $buffer = 'transformations(';
        $tmparg0 = [];
        foreach ($input->transformations as $arg1) {
        $tmptransformationsarg1 ='(new \Grafana\Foundation\Dashboard\DataTransformerConfig(id: '.\var_export($arg1->id, true).','.(($arg1->disabled !== null) ? 'disabled: '.\var_export($arg1->disabled, true).', ' : '').''.(($arg1->filter !== null) ? 'filter: '.'(new \Grafana\Foundation\Dashboard\MatcherConfig(id: '.\var_export($arg1->filter->id, true).','.(($arg1->filter->options !== null) ? 'options: '.\var_export($arg1->filter->options, true).', ' : '').'))'.', ' : '').''.(($arg1->topic !== null) ? 'topic: '.'\Grafana\Foundation\Dashboard\DataTransformerConfigTopic::fromValue("'.$arg1->topic.'")'.', ' : '').'options: '.\var_export($arg1->options, true).',))';
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
            if ($input->hideTimeOverride !== null) {
    
        
    $buffer = 'hideTimeOverride(';
        $arg0 =\var_export($input->hideTimeOverride, true);
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
            if ($input->cacheTimeout !== null && $input->cacheTimeout !== "") {
    
        
    $buffer = 'cacheTimeout(';
        $arg0 =\var_export($input->cacheTimeout, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryCachingTTL !== null) {
    
        
    $buffer = 'queryCachingTTL(';
        $arg0 =\var_export($input->queryCachingTTL, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->displayName !== null && $input->fieldConfig->defaults->displayName !== "") {
    
        
    $buffer = 'displayName(';
        $arg0 =\var_export($input->fieldConfig->defaults->displayName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->unit !== null && $input->fieldConfig->defaults->unit !== "") {
    
        
    $buffer = 'unit(';
        $arg0 =\var_export($input->fieldConfig->defaults->unit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->decimals !== null) {
    
        
    $buffer = 'decimals(';
        $arg0 =\var_export($input->fieldConfig->defaults->decimals, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->min !== null) {
    
        
    $buffer = 'min(';
        $arg0 =\var_export($input->fieldConfig->defaults->min, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->max !== null) {
    
        
    $buffer = 'max(';
        $arg0 =\var_export($input->fieldConfig->defaults->max, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->mappings !== null && count($input->fieldConfig->defaults->mappings) >= 1) {
    
        
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
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->thresholds !== null) {
    
        
    $buffer = 'thresholds(';
        $arg0 = \Grafana\Foundation\Dashboard\ThresholdsConfigConverter::convert($input->fieldConfig->defaults->thresholds);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->color !== null) {
    
        
    $buffer = 'colorScheme(';
        $arg0 = \Grafana\Foundation\Dashboard\FieldColorConverter::convert($input->fieldConfig->defaults->color);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->noValue !== null && $input->fieldConfig->defaults->noValue !== "") {
    
        
    $buffer = 'noValue(';
        $arg0 =\var_export($input->fieldConfig->defaults->noValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && count($input->fieldConfig->overrides) >= 1) {
    
        
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
            if ($input->fieldConfig !== null && count($input->fieldConfig->overrides) >= 1) {
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
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Trend\Options) {
    
        
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Common\VizLegendOptionsConverter::convert($input->options->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Trend\Options) {
    
        
    $buffer = 'tooltip(';
        $arg0 = \Grafana\Foundation\Common\VizTooltipOptionsConverter::convert($input->options->tooltip);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->options !== null && $input->options instanceof \Grafana\Foundation\Trend\Options && $input->options->xField !== null && $input->options->xField !== "") {
    
        
    $buffer = 'xField(';
        $arg0 =\var_export($input->options->xField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->drawStyle !== null) {
    
        
    $buffer = 'drawStyle(';
        $arg0 ='\Grafana\Foundation\Common\GraphDrawStyle::fromValue("'.$input->fieldConfig->defaults->custom->drawStyle.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->gradientMode !== null) {
    
        
    $buffer = 'gradientMode(';
        $arg0 ='\Grafana\Foundation\Common\GraphGradientMode::fromValue("'.$input->fieldConfig->defaults->custom->gradientMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->thresholdsStyle !== null) {
    
        
    $buffer = 'thresholdsStyle(';
        $arg0 = \Grafana\Foundation\Common\GraphThresholdsStyleConfigConverter::convert($input->fieldConfig->defaults->custom->thresholdsStyle);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->transform !== null) {
    
        
    $buffer = 'transform(';
        $arg0 ='\Grafana\Foundation\Common\GraphTransform::fromValue("'.$input->fieldConfig->defaults->custom->transform.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->lineColor !== null && $input->fieldConfig->defaults->custom->lineColor !== "") {
    
        
    $buffer = 'lineColor(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->lineColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->lineWidth !== null) {
    
        
    $buffer = 'lineWidth(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->lineWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->lineInterpolation !== null) {
    
        
    $buffer = 'lineInterpolation(';
        $arg0 ='\Grafana\Foundation\Common\LineInterpolation::fromValue("'.$input->fieldConfig->defaults->custom->lineInterpolation.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->lineStyle !== null) {
    
        
    $buffer = 'lineStyle(';
        $arg0 = \Grafana\Foundation\Common\LineStyleConverter::convert($input->fieldConfig->defaults->custom->lineStyle);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->fillColor !== null && $input->fieldConfig->defaults->custom->fillColor !== "") {
    
        
    $buffer = 'fillColor(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->fillColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->fillOpacity !== null) {
    
        
    $buffer = 'fillOpacity(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->fillOpacity, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->showPoints !== null) {
    
        
    $buffer = 'showPoints(';
        $arg0 ='\Grafana\Foundation\Common\VisibilityMode::fromValue("'.$input->fieldConfig->defaults->custom->showPoints.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->pointSize !== null) {
    
        
    $buffer = 'pointSize(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->pointSize, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->pointColor !== null && $input->fieldConfig->defaults->custom->pointColor !== "") {
    
        
    $buffer = 'pointColor(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->pointColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->axisPlacement !== null) {
    
        
    $buffer = 'axisPlacement(';
        $arg0 ='\Grafana\Foundation\Common\AxisPlacement::fromValue("'.$input->fieldConfig->defaults->custom->axisPlacement.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->axisColorMode !== null) {
    
        
    $buffer = 'axisColorMode(';
        $arg0 ='\Grafana\Foundation\Common\AxisColorMode::fromValue("'.$input->fieldConfig->defaults->custom->axisColorMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->axisLabel !== null && $input->fieldConfig->defaults->custom->axisLabel !== "") {
    
        
    $buffer = 'axisLabel(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->axisLabel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->axisWidth !== null) {
    
        
    $buffer = 'axisWidth(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->axisWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->axisSoftMin !== null) {
    
        
    $buffer = 'axisSoftMin(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->axisSoftMin, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->axisSoftMax !== null) {
    
        
    $buffer = 'axisSoftMax(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->axisSoftMax, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->axisGridShow !== null) {
    
        
    $buffer = 'axisGridShow(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->axisGridShow, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->scaleDistribution !== null) {
    
        
    $buffer = 'scaleDistribution(';
        $arg0 = \Grafana\Foundation\Common\ScaleDistributionConfigConverter::convert($input->fieldConfig->defaults->custom->scaleDistribution);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->axisCenteredZero !== null) {
    
        
    $buffer = 'axisCenteredZero(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->axisCenteredZero, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->barAlignment !== null) {
    
        
    $buffer = 'barAlignment(';
        $arg0 ='\Grafana\Foundation\Common\BarAlignment::fromValue("'.$input->fieldConfig->defaults->custom->barAlignment.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->barWidthFactor !== null) {
    
        
    $buffer = 'barWidthFactor(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->barWidthFactor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->stacking !== null) {
    
        
    $buffer = 'stacking(';
        $arg0 = \Grafana\Foundation\Common\StackingConfigConverter::convert($input->fieldConfig->defaults->custom->stacking);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->hideFrom !== null) {
    
        
    $buffer = 'hideFrom(';
        $arg0 = \Grafana\Foundation\Common\HideSeriesConfigConverter::convert($input->fieldConfig->defaults->custom->hideFrom);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->insertNulls !== null) {
    
        
    $buffer = 'insertNulls(';
        switch (true) {
            case is_bool($input->fieldConfig->defaults->custom->insertNulls):
                $disjunctioninsertNulls =\var_export($input->fieldConfig->defaults->custom->insertNulls, true);
                $arg0 = $disjunctioninsertNulls;
                break;
            case is_float($input->fieldConfig->defaults->custom->insertNulls):
                $disjunctioninsertNulls =\var_export($input->fieldConfig->defaults->custom->insertNulls, true);
                $arg0 = $disjunctioninsertNulls;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->spanNulls !== null) {
    
        
    $buffer = 'spanNulls(';
        switch (true) {
            case is_bool($input->fieldConfig->defaults->custom->spanNulls):
                $disjunctionspanNulls =\var_export($input->fieldConfig->defaults->custom->spanNulls, true);
                $arg0 = $disjunctionspanNulls;
                break;
            case is_float($input->fieldConfig->defaults->custom->spanNulls):
                $disjunctionspanNulls =\var_export($input->fieldConfig->defaults->custom->spanNulls, true);
                $arg0 = $disjunctionspanNulls;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->fillBelowTo !== null && $input->fieldConfig->defaults->custom->fillBelowTo !== "") {
    
        
    $buffer = 'fillBelowTo(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->fillBelowTo, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->pointSymbol !== null && $input->fieldConfig->defaults->custom->pointSymbol !== "") {
    
        
    $buffer = 'pointSymbol(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->pointSymbol, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->axisBorderShow !== null) {
    
        
    $buffer = 'axisBorderShow(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->axisBorderShow, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fieldConfig !== null && $input->fieldConfig->defaults->custom !== null && $input->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Trend\FieldConfig && $input->fieldConfig->defaults->custom->barMaxWidth !== null) {
    
        
    $buffer = 'barMaxWidth(';
        $arg0 =\var_export($input->fieldConfig->defaults->custom->barMaxWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

