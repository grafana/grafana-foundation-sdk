<?php

namespace Grafana\Foundation\Heatmap;

final class FilterValueRangeConverter
{
    public static function convert(\Grafana\Foundation\Heatmap\FilterValueRange $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Heatmap\FilterValueRangeBuilder())',
        ];
            if ($input->le !== null) {
    
        
    $buffer = 'le(';
        $arg0 =\var_export($input->le, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->ge !== null) {
    
        
    $buffer = 'ge(';
        $arg0 =\var_export($input->ge, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

