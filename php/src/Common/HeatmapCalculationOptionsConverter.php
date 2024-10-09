<?php

namespace Grafana\Foundation\Common;

final class HeatmapCalculationOptionsConverter
{
    public static function convert(\Grafana\Foundation\Common\HeatmapCalculationOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\HeatmapCalculationOptionsBuilder())',
        ];
            if ($input->xBuckets !== null) {
    
        
    $buffer = 'xBuckets(';
        $arg0 = \Grafana\Foundation\Common\HeatmapCalculationBucketConfigConverter::convert($input->xBuckets);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->yBuckets !== null) {
    
        
    $buffer = 'yBuckets(';
        $arg0 = \Grafana\Foundation\Common\HeatmapCalculationBucketConfigConverter::convert($input->yBuckets);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

