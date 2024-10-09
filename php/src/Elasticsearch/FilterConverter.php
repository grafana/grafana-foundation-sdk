<?php

namespace Grafana\Foundation\Elasticsearch;

final class FilterConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\Filter $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\FilterBuilder())',
        ];
            if ($input->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->label !== "") {
    
        
    $buffer = 'label(';
        $arg0 =\var_export($input->label, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

