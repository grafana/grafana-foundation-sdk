<?php

namespace Grafana\Foundation\Elasticsearch;

final class FiltersSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\FiltersSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\FiltersSettingsBuilder())',
        ];
            if ($input->filters !== null && count($input->filters) >= 1) {
    
        
    $buffer = 'filters(';
        $tmparg0 = [];
        foreach ($input->filters as $arg1) {
        $tmpfiltersarg1 = \Grafana\Foundation\Elasticsearch\FilterConverter::convert($arg1);
        $tmparg0[] = $tmpfiltersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

