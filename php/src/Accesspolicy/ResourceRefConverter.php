<?php

namespace Grafana\Foundation\Accesspolicy;

final class ResourceRefConverter
{
    public static function convert(\Grafana\Foundation\Accesspolicy\ResourceRef $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Accesspolicy\ResourceRefBuilder())',
        ];
            if ($input->kind !== "") {
    
        
    $buffer = 'kind(';
        $arg0 =\var_export($input->kind, true);
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

