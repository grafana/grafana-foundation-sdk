<?php

namespace Grafana\Foundation\Nodegraph;

final class NodeOptionsConverter
{
    public static function convert(\Grafana\Foundation\Nodegraph\NodeOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Nodegraph\NodeOptionsBuilder())',
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
            if ($input->arcs !== null && count($input->arcs) >= 1) {
    
        
    $buffer = 'arcs(';
        $tmparg0 = [];
        foreach ($input->arcs as $arg1) {
        $tmparcsarg1 = \Grafana\Foundation\Nodegraph\ArcOptionConverter::convert($arg1);
        $tmparg0[] = $tmparcsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

