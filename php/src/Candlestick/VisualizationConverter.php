<?php

namespace Grafana\Foundation\Candlestick;

final class VisualizationConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\VizConfigKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Candlestick\VisualizationBuilder())',
        ];
            if ($input->spec->fieldConfig->defaults->displayName !== null && $input->spec->fieldConfig->defaults->displayName !== "") {
    
        
    $buffer = 'displayName(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->displayName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->displayNameFromDS !== null && $input->spec->fieldConfig->defaults->displayNameFromDS !== "") {
    
        
    $buffer = 'displayNameFromDS(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->displayNameFromDS, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->description !== null && $input->spec->fieldConfig->defaults->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->path !== null && $input->spec->fieldConfig->defaults->path !== "") {
    
        
    $buffer = 'path(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->path, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->unit !== null && $input->spec->fieldConfig->defaults->unit !== "") {
    
        
    $buffer = 'unit(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->unit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->decimals !== null) {
    
        
    $buffer = 'decimals(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->decimals, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->min !== null) {
    
        
    $buffer = 'min(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->min, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->max !== null) {
    
        
    $buffer = 'max(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->max, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->mappings !== null && count($input->spec->fieldConfig->defaults->mappings) >= 1) {
    
        
    $buffer = 'mappings(';
        $tmparg0 = [];
        foreach ($input->spec->fieldConfig->defaults->mappings as $arg1) {
        switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\ValueMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\ValueMap(type: '.\var_export($arg1->type, true).',options: '.\var_export($arg1->options, true).',))';
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\RangeMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\RangeMap(type: '.\var_export($arg1->type, true).',options: '.'(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RangeMapOptions('.(($arg1->options->from !== null) ? 'from: '.\var_export($arg1->options->from, true).', ' : '').''.(($arg1->options->to !== null) ? 'to: '.\var_export($arg1->options->to, true).', ' : '').'result: '.'(new \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult('.(($arg1->options->result->text !== null) ? 'text: '.\var_export($arg1->options->result->text, true).', ' : '').''.(($arg1->options->result->color !== null) ? 'color: '.\var_export($arg1->options->result->color, true).', ' : '').''.(($arg1->options->result->icon !== null) ? 'icon: '.\var_export($arg1->options->result->icon, true).', ' : '').''.(($arg1->options->result->index !== null) ? 'index: '.\var_export($arg1->options->result->index, true).', ' : '').'))'.',))'.',))';
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\RegexMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\RegexMap(type: '.\var_export($arg1->type, true).',options: '.'(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions(pattern: '.\var_export($arg1->options->pattern, true).',result: '.'(new \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult('.(($arg1->options->result->text !== null) ? 'text: '.\var_export($arg1->options->result->text, true).', ' : '').''.(($arg1->options->result->color !== null) ? 'color: '.\var_export($arg1->options->result->color, true).', ' : '').''.(($arg1->options->result->icon !== null) ? 'icon: '.\var_export($arg1->options->result->icon, true).', ' : '').''.(($arg1->options->result->index !== null) ? 'index: '.\var_export($arg1->options->result->index, true).', ' : '').'))'.',))'.',))';
                $tmpmappingsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\SpecialValueMap:
                $disjunctionarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\SpecialValueMap(type: '.\var_export($arg1->type, true).',options: '.'(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1SpecialValueMapOptions(match: '.'\Grafana\Foundation\Dashboardv2beta1\SpecialValueMatch::fromValue("'.$arg1->options->match.'")'.',result: '.'(new \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult('.(($arg1->options->result->text !== null) ? 'text: '.\var_export($arg1->options->result->text, true).', ' : '').''.(($arg1->options->result->color !== null) ? 'color: '.\var_export($arg1->options->result->color, true).', ' : '').''.(($arg1->options->result->icon !== null) ? 'icon: '.\var_export($arg1->options->result->icon, true).', ' : '').''.(($arg1->options->result->index !== null) ? 'index: '.\var_export($arg1->options->result->index, true).', ' : '').'))'.',))'.',))';
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
            if ($input->spec->fieldConfig->defaults->thresholds !== null) {
    
        
    $buffer = 'thresholds(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ThresholdsConfigConverter::convert($input->spec->fieldConfig->defaults->thresholds);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->color !== null) {
    
        
    $buffer = 'colorScheme(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\FieldColorConverter::convert($input->spec->fieldConfig->defaults->color);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->links !== null && count($input->spec->fieldConfig->defaults->links) >= 1) {
    
        
    $buffer = 'dataLinks(';
        $tmparg0 = [];
        foreach ($input->spec->fieldConfig->defaults->links as $arg1) {
        $tmplinksarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmplinksarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->actions !== null && count($input->spec->fieldConfig->defaults->actions) >= 1) {
    
        
    $buffer = 'actions(';
        $tmparg0 = [];
        foreach ($input->spec->fieldConfig->defaults->actions as $arg1) {
        $tmpactionsarg1 = \Grafana\Foundation\Dashboardv2beta1\ActionConverter::convert($arg1);
        $tmparg0[] = $tmpactionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->noValue !== null && $input->spec->fieldConfig->defaults->noValue !== "") {
    
        
    $buffer = 'noValue(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->noValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->spec->fieldConfig->overrides) >= 1) {
    
        
    $buffer = 'overrides(';
        $tmparg0 = [];
        foreach ($input->spec->fieldConfig->overrides as $arg1) {
        $tmpoverridesarg1 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverridesConverter::convert($arg1);
        $tmparg0[] = $tmpoverridesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->spec->fieldConfig->overrides) >= 1) {
    foreach ($input->spec->fieldConfig->overrides as $item) {
        
    $buffer = 'overrideByName(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\MatcherConfig(id: '.\var_export($item->matcher->id, true).','.(($item->matcher->options !== null) ? 'options: '.\var_export($item->matcher->options, true).', ' : '').'))';
        $buffer .= $arg0;
        $buffer .= ', ';
        $tmparg1 = [];
        foreach ($item->properties as $arg1) {
        $tmppropertiesarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue(id: '.\var_export($arg1->id, true).','.(($arg1->value !== null) ? 'value: '.\var_export($arg1->value, true).', ' : '').'))';
        $tmparg1[] = $tmppropertiesarg1;
        }
        $arg1 = "[" . implode(", \n", $tmparg1) . "]";
        $buffer .= $arg1;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    }
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Candlestick\Options) {
    
        
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Candlestick\VizDisplayMode::fromValue("'.$input->spec->options->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Candlestick\Options) {
    
        
    $buffer = 'candleStyle(';
        $arg0 ='\Grafana\Foundation\Candlestick\CandleStyle::fromValue("'.$input->spec->options->candleStyle.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Candlestick\Options) {
    
        
    $buffer = 'colorStrategy(';
        $arg0 ='\Grafana\Foundation\Candlestick\ColorStrategy::fromValue("'.$input->spec->options->colorStrategy.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Candlestick\Options) {
    
        
    $buffer = 'fields(';
        $arg0 = \Grafana\Foundation\Candlestick\CandlestickFieldMapConverter::convert($input->spec->options->fields);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Candlestick\Options) {
    
        
    $buffer = 'colors(';
        $arg0 = \Grafana\Foundation\Candlestick\CandlestickColorsConverter::convert($input->spec->options->colors);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Candlestick\Options) {
    
        
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Common\VizLegendOptionsConverter::convert($input->spec->options->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Candlestick\Options) {
    
        
    $buffer = 'tooltip(';
        $arg0 = \Grafana\Foundation\Common\VizTooltipOptionsConverter::convert($input->spec->options->tooltip);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->options !== null && $input->spec->options instanceof \Grafana\Foundation\Candlestick\Options && $input->spec->options->includeAllFields !== null && $input->spec->options->includeAllFields !== false) {
    
        
    $buffer = 'includeAllFields(';
        $arg0 =\var_export($input->spec->options->includeAllFields, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->drawStyle !== null) {
    
        
    $buffer = 'drawStyle(';
        $arg0 ='\Grafana\Foundation\Common\GraphDrawStyle::fromValue("'.$input->spec->fieldConfig->defaults->custom->drawStyle.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->gradientMode !== null) {
    
        
    $buffer = 'gradientMode(';
        $arg0 ='\Grafana\Foundation\Common\GraphGradientMode::fromValue("'.$input->spec->fieldConfig->defaults->custom->gradientMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->thresholdsStyle !== null) {
    
        
    $buffer = 'thresholdsStyle(';
        $arg0 = \Grafana\Foundation\Common\GraphThresholdsStyleConfigConverter::convert($input->spec->fieldConfig->defaults->custom->thresholdsStyle);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->transform !== null) {
    
        
    $buffer = 'transform(';
        $arg0 ='\Grafana\Foundation\Common\GraphTransform::fromValue("'.$input->spec->fieldConfig->defaults->custom->transform.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->lineColor !== null && $input->spec->fieldConfig->defaults->custom->lineColor !== "") {
    
        
    $buffer = 'lineColor(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->lineColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->lineWidth !== null) {
    
        
    $buffer = 'lineWidth(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->lineWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->lineInterpolation !== null) {
    
        
    $buffer = 'lineInterpolation(';
        $arg0 ='\Grafana\Foundation\Common\LineInterpolation::fromValue("'.$input->spec->fieldConfig->defaults->custom->lineInterpolation.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->lineStyle !== null) {
    
        
    $buffer = 'lineStyle(';
        $arg0 = \Grafana\Foundation\Common\LineStyleConverter::convert($input->spec->fieldConfig->defaults->custom->lineStyle);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->fillColor !== null && $input->spec->fieldConfig->defaults->custom->fillColor !== "") {
    
        
    $buffer = 'fillColor(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->fillColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->fillOpacity !== null) {
    
        
    $buffer = 'fillOpacity(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->fillOpacity, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->showPoints !== null) {
    
        
    $buffer = 'showPoints(';
        $arg0 ='\Grafana\Foundation\Common\VisibilityMode::fromValue("'.$input->spec->fieldConfig->defaults->custom->showPoints.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->pointSize !== null) {
    
        
    $buffer = 'pointSize(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->pointSize, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->pointColor !== null && $input->spec->fieldConfig->defaults->custom->pointColor !== "") {
    
        
    $buffer = 'pointColor(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->pointColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->axisPlacement !== null) {
    
        
    $buffer = 'axisPlacement(';
        $arg0 ='\Grafana\Foundation\Common\AxisPlacement::fromValue("'.$input->spec->fieldConfig->defaults->custom->axisPlacement.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->axisColorMode !== null) {
    
        
    $buffer = 'axisColorMode(';
        $arg0 ='\Grafana\Foundation\Common\AxisColorMode::fromValue("'.$input->spec->fieldConfig->defaults->custom->axisColorMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->axisLabel !== null && $input->spec->fieldConfig->defaults->custom->axisLabel !== "") {
    
        
    $buffer = 'axisLabel(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->axisLabel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->axisWidth !== null) {
    
        
    $buffer = 'axisWidth(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->axisWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->axisSoftMin !== null) {
    
        
    $buffer = 'axisSoftMin(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->axisSoftMin, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->axisSoftMax !== null) {
    
        
    $buffer = 'axisSoftMax(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->axisSoftMax, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->axisGridShow !== null) {
    
        
    $buffer = 'axisGridShow(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->axisGridShow, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->scaleDistribution !== null) {
    
        
    $buffer = 'scaleDistribution(';
        $arg0 = \Grafana\Foundation\Common\ScaleDistributionConfigConverter::convert($input->spec->fieldConfig->defaults->custom->scaleDistribution);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->axisCenteredZero !== null) {
    
        
    $buffer = 'axisCenteredZero(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->axisCenteredZero, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->barAlignment !== null) {
    
        
    $buffer = 'barAlignment(';
        $arg0 ='\Grafana\Foundation\Common\BarAlignment::fromValue("'.$input->spec->fieldConfig->defaults->custom->barAlignment.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->barWidthFactor !== null) {
    
        
    $buffer = 'barWidthFactor(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->barWidthFactor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->stacking !== null) {
    
        
    $buffer = 'stacking(';
        $arg0 = \Grafana\Foundation\Common\StackingConfigConverter::convert($input->spec->fieldConfig->defaults->custom->stacking);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->hideFrom !== null) {
    
        
    $buffer = 'hideFrom(';
        $arg0 = \Grafana\Foundation\Common\HideSeriesConfigConverter::convert($input->spec->fieldConfig->defaults->custom->hideFrom);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->insertNulls !== null) {
    
        
    $buffer = 'insertNulls(';
        switch (true) {
            case is_bool($input->spec->fieldConfig->defaults->custom->insertNulls):
                $disjunctioninsertNulls =\var_export($input->spec->fieldConfig->defaults->custom->insertNulls, true);
                $arg0 = $disjunctioninsertNulls;
                break;
            case is_float($input->spec->fieldConfig->defaults->custom->insertNulls):
                $disjunctioninsertNulls =\var_export($input->spec->fieldConfig->defaults->custom->insertNulls, true);
                $arg0 = $disjunctioninsertNulls;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->spanNulls !== null) {
    
        
    $buffer = 'spanNulls(';
        switch (true) {
            case is_bool($input->spec->fieldConfig->defaults->custom->spanNulls):
                $disjunctionspanNulls =\var_export($input->spec->fieldConfig->defaults->custom->spanNulls, true);
                $arg0 = $disjunctionspanNulls;
                break;
            case is_float($input->spec->fieldConfig->defaults->custom->spanNulls):
                $disjunctionspanNulls =\var_export($input->spec->fieldConfig->defaults->custom->spanNulls, true);
                $arg0 = $disjunctionspanNulls;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->fillBelowTo !== null && $input->spec->fieldConfig->defaults->custom->fillBelowTo !== "") {
    
        
    $buffer = 'fillBelowTo(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->fillBelowTo, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->pointSymbol !== null && $input->spec->fieldConfig->defaults->custom->pointSymbol !== "") {
    
        
    $buffer = 'pointSymbol(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->pointSymbol, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->axisBorderShow !== null) {
    
        
    $buffer = 'axisBorderShow(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->axisBorderShow, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fieldConfig->defaults->custom !== null && $input->spec->fieldConfig->defaults->custom instanceof \Grafana\Foundation\Candlestick\FieldConfig && $input->spec->fieldConfig->defaults->custom->barMaxWidth !== null) {
    
        
    $buffer = 'barMaxWidth(';
        $arg0 =\var_export($input->spec->fieldConfig->defaults->custom->barMaxWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

