<?php

namespace Grafana\Foundation\Common;

final class BarConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\BarConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\BarConfigBuilder())',
        ];
            if ($input->barAlignment !== null) {
    
        
    $buffer = 'barAlignment(';
        $arg0 ='\Grafana\Foundation\Common\BarAlignment::fromValue("'.$input->barAlignment.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->barWidthFactor !== null) {
    
        
    $buffer = 'barWidthFactor(';
        $arg0 =\var_export($input->barWidthFactor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->barMaxWidth !== null) {
    
        
    $buffer = 'barMaxWidth(';
        $arg0 =\var_export($input->barMaxWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

