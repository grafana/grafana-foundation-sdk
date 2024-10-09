<?php

namespace Grafana\Foundation\Common;

final class ColorDimensionConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\ColorDimensionConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\ColorDimensionConfigBuilder())',
        ];
            if ($input->fixed !== null && $input->fixed !== "") {
    
        
    $buffer = 'fixed(';
        $arg0 =\var_export($input->fixed, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->field !== null && $input->field !== "") {
    
        
    $buffer = 'field(';
        $arg0 =\var_export($input->field, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

