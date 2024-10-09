<?php

namespace Grafana\Foundation\Rolebinding;

final class BuiltinRoleRefConverter
{
    public static function convert(\Grafana\Foundation\Rolebinding\BuiltinRoleRef $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Rolebinding\BuiltinRoleRefBuilder())',
        ];
            
    
        {
    $buffer = 'name(';
        $arg0 ='\Grafana\Foundation\Rolebinding\BuiltinRoleRefName::fromValue("'.$input->name.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

