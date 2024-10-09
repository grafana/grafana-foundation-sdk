<?php

namespace Grafana\Foundation\Elasticsearch;

final class FiltersConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\Filters $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\FiltersBuilder())',
        ];
            if ($input->id !== "") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->settings !== null) {
    
        
    $buffer = 'settings(';
        $arg0 = \Grafana\Foundation\Elasticsearch\ElasticsearchFiltersSettingsConverter::convert($input->settings);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

