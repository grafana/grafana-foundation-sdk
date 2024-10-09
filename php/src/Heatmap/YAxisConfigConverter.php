<?php

namespace Grafana\Foundation\Heatmap;

final class YAxisConfigConverter
{
    public static function convert(\Grafana\Foundation\Heatmap\YAxisConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Heatmap\YAxisConfigBuilder())',
        ];
            if ($input->unit !== null && $input->unit !== "") {
    
        
    $buffer = 'unit(';
        $arg0 =\var_export($input->unit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->reverse !== null) {
    
        
    $buffer = 'reverse(';
        $arg0 =\var_export($input->reverse, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->decimals !== null) {
    
        
    $buffer = 'decimals(';
        $arg0 =\var_export($input->decimals, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->min !== null) {
    
        
    $buffer = 'min(';
        $arg0 =\var_export($input->min, true);
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
            if ($input->max !== null) {
    
        
    $buffer = 'max(';
        $arg0 =\var_export($input->max, true);
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

        return \implode("\n\t->", $calls);
    }
}

