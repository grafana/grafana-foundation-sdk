<?php

namespace Grafana\Foundation\Common;

final class VizTextDisplayOptionsConverter
{
    public static function convert(\Grafana\Foundation\Common\VizTextDisplayOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\VizTextDisplayOptionsBuilder())',
        ];
            if ($input->titleSize !== null) {
    
        
    $buffer = 'titleSize(';
        $arg0 =\var_export($input->titleSize, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->valueSize !== null) {
    
        
    $buffer = 'valueSize(';
        $arg0 =\var_export($input->valueSize, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

