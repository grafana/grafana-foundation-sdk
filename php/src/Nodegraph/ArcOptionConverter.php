<?php

namespace Grafana\Foundation\Nodegraph;

final class ArcOptionConverter
{
    public static function convert(\Grafana\Foundation\Nodegraph\ArcOption $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Nodegraph\ArcOptionBuilder())',
        ];
            if ($input->field !== null && $input->field !== "") {
    
        
    $buffer = 'field(';
        $arg0 =\var_export($input->field, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->color !== null && $input->color !== "") {
    
        
    $buffer = 'color(';
        $arg0 =\var_export($input->color, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

