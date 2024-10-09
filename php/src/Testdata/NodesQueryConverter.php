<?php

namespace Grafana\Foundation\Testdata;

final class NodesQueryConverter
{
    public static function convert(\Grafana\Foundation\Testdata\NodesQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Testdata\NodesQueryBuilder())',
        ];
            if ($input->type !== null) {
    
        
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Testdata\NodesQueryType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->count !== null) {
    
        
    $buffer = 'count(';
        $arg0 =\var_export($input->count, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

