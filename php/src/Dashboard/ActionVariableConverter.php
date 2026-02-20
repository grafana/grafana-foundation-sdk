<?php

namespace Grafana\Foundation\Dashboard;

final class ActionVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\ActionVariable $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\ActionVariableBuilder())',
        ];
            if ($input->key !== "") {
    
        
    $buffer = 'key(';
        $arg0 =\var_export($input->key, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

