<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchInlineScriptConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchInlineScriptBuilder())',
        ];
            if ($input->inline !== null && $input->inline !== "") {
    
        
    $buffer = 'inline(';
        $arg0 =\var_export($input->inline, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

