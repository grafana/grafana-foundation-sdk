<?php

namespace Grafana\Foundation\Histogram;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Histogram\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Histogram\OptionsBuilder())',
        ];
            if ($input->bucketCount !== null && $input->bucketCount !== 30) {
    
        
    $buffer = 'bucketCount(';
        $arg0 =\var_export($input->bucketCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->bucketSize !== null) {
    
        
    $buffer = 'bucketSize(';
        $arg0 =\var_export($input->bucketSize, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->bucketOffset !== null && $input->bucketOffset !== (float) 0) {
    
        
    $buffer = 'bucketOffset(';
        $arg0 =\var_export($input->bucketOffset, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Common\VizLegendOptionsConverter::convert($input->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'tooltip(';
        $arg0 = \Grafana\Foundation\Common\VizTooltipOptionsConverter::convert($input->tooltip);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->combine !== null) {
    
        
    $buffer = 'combine(';
        $arg0 =\var_export($input->combine, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

