<?php

namespace Grafana\Foundation\Common;

final class HeatmapCalculationBucketConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\HeatmapCalculationBucketConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\HeatmapCalculationBucketConfigBuilder())',
        ];
            if ($input->mode !== null) {
    
        
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Common\HeatmapCalculationMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->value !== null && $input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->scale !== null) {
    
        
    $buffer = 'scale(';
        $arg0 = \Grafana\Foundation\Common\ScaleDistributionConfigConverter::convert($input->scale);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

