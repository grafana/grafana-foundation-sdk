<?php

namespace Grafana\Foundation\Rolebinding;

final class CustomRoleRefConverter
{
    public static function convert(\Grafana\Foundation\Rolebinding\CustomRoleRef $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Rolebinding\CustomRoleRefBuilder())',
        ];
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

