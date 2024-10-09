<?php

namespace Grafana\Foundation\Common;

final class StackingConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\StackingConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\StackingConfigBuilder())',
        ];
            if ($input->mode !== null) {
    
        
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Common\StackingMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->group !== null && $input->group !== "") {
    
        
    $buffer = 'group(';
        $arg0 =\var_export($input->group, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

