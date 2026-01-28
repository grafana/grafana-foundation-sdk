<?php

namespace Grafana\Foundation\Histogram;

final class FieldConfigConverter
{
    public static function convert(\Grafana\Foundation\Histogram\FieldConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Histogram\FieldConfigBuilder())',
        ];
            if ($input->lineWidth !== null && $input->lineWidth !== 1) {
    
        
    $buffer = 'lineWidth(';
        $arg0 =\var_export($input->lineWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fillOpacity !== null && $input->fillOpacity !== 80) {
    
        
    $buffer = 'fillOpacity(';
        $arg0 =\var_export($input->fillOpacity, true);
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
            if ($input->hideFrom !== null) {
    
        
    $buffer = 'hideFrom(';
        $arg0 = \Grafana\Foundation\Common\HideSeriesConfigConverter::convert($input->hideFrom);
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
            if ($input->gradientMode !== null) {
    
        
    $buffer = 'gradientMode(';
        $arg0 ='\Grafana\Foundation\Common\GraphGradientMode::fromValue("'.$input->gradientMode.'")';
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

        return \implode("\n\t->", $calls);
    }
}

