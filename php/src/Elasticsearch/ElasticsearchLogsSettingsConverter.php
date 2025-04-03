<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchLogsSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchLogsSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchLogsSettingsBuilder())',
        ];
            if ($input->limit !== null && $input->limit !== "") {
    
        
    $buffer = 'limit(';
        $arg0 =\var_export($input->limit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

