<?php

namespace Grafana\Foundation\Nodegraph;

final class EdgeOptionsConverter
{
    public static function convert(\Grafana\Foundation\Nodegraph\EdgeOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Nodegraph\EdgeOptionsBuilder())',
        ];
            if ($input->mainStatUnit !== null && $input->mainStatUnit !== "") {
    
        
    $buffer = 'mainStatUnit(';
        $arg0 =\var_export($input->mainStatUnit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->secondaryStatUnit !== null && $input->secondaryStatUnit !== "") {
    
        
    $buffer = 'secondaryStatUnit(';
        $arg0 =\var_export($input->secondaryStatUnit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

