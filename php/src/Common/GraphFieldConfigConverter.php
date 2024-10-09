<?php

namespace Grafana\Foundation\Common;

final class GraphFieldConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\GraphFieldConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\GraphFieldConfigBuilder())',
        ];
            if ($input->drawStyle !== null) {
    
        
    $buffer = 'drawStyle(';
        $arg0 ='\Grafana\Foundation\Common\GraphDrawStyle::fromValue("'.$input->drawStyle.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->gradientMode !== null) {
    
        
    $buffer = 'gradientMode(';
        $arg0 ='\Grafana\Foundation\Common\GraphGradientMode::fromValue("'.$input->gradientMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->thresholdsStyle !== null) {
    
        
    $buffer = 'thresholdsStyle(';
        $arg0 = \Grafana\Foundation\Common\GraphThresholdsStyleConfigConverter::convert($input->thresholdsStyle);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lineColor !== null && $input->lineColor !== "") {
    
        
    $buffer = 'lineColor(';
        $arg0 =\var_export($input->lineColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lineWidth !== null) {
    
        
    $buffer = 'lineWidth(';
        $arg0 =\var_export($input->lineWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lineInterpolation !== null) {
    
        
    $buffer = 'lineInterpolation(';
        $arg0 ='\Grafana\Foundation\Common\LineInterpolation::fromValue("'.$input->lineInterpolation.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lineStyle !== null) {
    
        
    $buffer = 'lineStyle(';
        $arg0 = \Grafana\Foundation\Common\LineStyleConverter::convert($input->lineStyle);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fillColor !== null && $input->fillColor !== "") {
    
        
    $buffer = 'fillColor(';
        $arg0 =\var_export($input->fillColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fillOpacity !== null) {
    
        
    $buffer = 'fillOpacity(';
        $arg0 =\var_export($input->fillOpacity, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showPoints !== null) {
    
        
    $buffer = 'showPoints(';
        $arg0 ='\Grafana\Foundation\Common\VisibilityMode::fromValue("'.$input->showPoints.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->pointSize !== null) {
    
        
    $buffer = 'pointSize(';
        $arg0 =\var_export($input->pointSize, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->pointColor !== null && $input->pointColor !== "") {
    
        
    $buffer = 'pointColor(';
        $arg0 =\var_export($input->pointColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->axisPlacement !== null) {
    
        
    $buffer = 'axisPlacement(';
        $arg0 ='\Grafana\Foundation\Common\AxisPlacement::fromValue("'.$input->axisPlacement.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->axisColorMode !== null) {
    
        
    $buffer = 'axisColorMode(';
        $arg0 ='\Grafana\Foundation\Common\AxisColorMode::fromValue("'.$input->axisColorMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->axisLabel !== null && $input->axisLabel !== "") {
    
        
    $buffer = 'axisLabel(';
        $arg0 =\var_export($input->axisLabel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->axisWidth !== null) {
    
        
    $buffer = 'axisWidth(';
        $arg0 =\var_export($input->axisWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->axisSoftMin !== null) {
    
        
    $buffer = 'axisSoftMin(';
        $arg0 =\var_export($input->axisSoftMin, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->axisSoftMax !== null) {
    
        
    $buffer = 'axisSoftMax(';
        $arg0 =\var_export($input->axisSoftMax, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->axisGridShow !== null) {
    
        
    $buffer = 'axisGridShow(';
        $arg0 =\var_export($input->axisGridShow, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->scaleDistribution !== null) {
    
        
    $buffer = 'scaleDistribution(';
        $arg0 = \Grafana\Foundation\Common\ScaleDistributionConfigConverter::convert($input->scaleDistribution);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->axisCenteredZero !== null) {
    
        
    $buffer = 'axisCenteredZero(';
        $arg0 =\var_export($input->axisCenteredZero, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->barAlignment !== null) {
    
        
    $buffer = 'barAlignment(';
        $arg0 ='\Grafana\Foundation\Common\BarAlignment::fromValue("'.$input->barAlignment.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->barWidthFactor !== null) {
    
        
    $buffer = 'barWidthFactor(';
        $arg0 =\var_export($input->barWidthFactor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->stacking !== null) {
    
        
    $buffer = 'stacking(';
        $arg0 = \Grafana\Foundation\Common\StackingConfigConverter::convert($input->stacking);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hideFrom !== null) {
    
        
    $buffer = 'hideFrom(';
        $arg0 = \Grafana\Foundation\Common\HideSeriesConfigConverter::convert($input->hideFrom);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->transform !== null) {
    
        
    $buffer = 'transform(';
        $arg0 ='\Grafana\Foundation\Common\GraphTransform::fromValue("'.$input->transform.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spanNulls !== null) {
    
        
    $buffer = 'spanNulls(';
        switch (true) {
            case is_bool($input->spanNulls):
                $disjunctionspanNulls =\var_export($input->spanNulls, true);
                $arg0 = $disjunctionspanNulls;
                break;
            case is_float($input->spanNulls):
                $disjunctionspanNulls =\var_export($input->spanNulls, true);
                $arg0 = $disjunctionspanNulls;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fillBelowTo !== null && $input->fillBelowTo !== "") {
    
        
    $buffer = 'fillBelowTo(';
        $arg0 =\var_export($input->fillBelowTo, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->pointSymbol !== null && $input->pointSymbol !== "") {
    
        
    $buffer = 'pointSymbol(';
        $arg0 =\var_export($input->pointSymbol, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->axisBorderShow !== null) {
    
        
    $buffer = 'axisBorderShow(';
        $arg0 =\var_export($input->axisBorderShow, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->barMaxWidth !== null) {
    
        
    $buffer = 'barMaxWidth(';
        $arg0 =\var_export($input->barMaxWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->insertNulls !== null) {
    
        
    $buffer = 'insertNulls(';
        switch (true) {
            case is_bool($input->insertNulls):
                $disjunctioninsertNulls =\var_export($input->insertNulls, true);
                $arg0 = $disjunctioninsertNulls;
                break;
            case is_int($input->insertNulls):
                $disjunctioninsertNulls =\var_export($input->insertNulls, true);
                $arg0 = $disjunctioninsertNulls;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

