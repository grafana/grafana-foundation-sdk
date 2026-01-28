<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ActionVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\ActionVariable $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\ActionVariableBuilder())',
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

