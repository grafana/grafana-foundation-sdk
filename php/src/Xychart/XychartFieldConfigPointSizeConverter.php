<?php

namespace Grafana\Foundation\Xychart;

final class XychartFieldConfigPointSizeConverter
{
    public static function convert(\Grafana\Foundation\Xychart\XychartFieldConfigPointSize $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Xychart\XychartFieldConfigPointSizeBuilder())',
        ];
            if ($input->fixed !== null) {
    
        
    $buffer = 'fixed(';
        $arg0 =\var_export($input->fixed, true);
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
            if ($input->max !== null) {
    
        
    $buffer = 'max(';
        $arg0 =\var_export($input->max, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

