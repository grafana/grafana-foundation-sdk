<?php

namespace Grafana\Foundation\Accesspolicy;

final class RoleRefConverter
{
    public static function convert(\Grafana\Foundation\Accesspolicy\RoleRef $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Accesspolicy\RoleRefBuilder())',
        ];
            
    
        {
    $buffer = 'kind(';
        $arg0 ='\Grafana\Foundation\Accesspolicy\RoleRefKind::fromValue("'.$input->kind.'")';
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
            if ($input->xname !== "") {
    
        
    $buffer = 'xname(';
        $arg0 =\var_export($input->xname, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

