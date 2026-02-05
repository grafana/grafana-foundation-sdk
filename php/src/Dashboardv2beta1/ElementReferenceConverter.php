<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ElementReferenceConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\ElementReference $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\ElementReferenceBuilder())',
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

