<?php

namespace Grafana\Foundation\Statushistory;

final class FieldConfigConverter
{
    public static function convert(\Grafana\Foundation\Statushistory\FieldConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Statushistory\FieldConfigBuilder())',
        ];
            if ($input->lineWidth !== null && $input->lineWidth !== 1) {
    
        
    $buffer = 'lineWidth(';
        $arg0 =\var_export($input->lineWidth, true);
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
            if ($input->fillOpacity !== null && $input->fillOpacity !== 70) {
    
        
    $buffer = 'fillOpacity(';
        $arg0 =\var_export($input->fillOpacity, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

