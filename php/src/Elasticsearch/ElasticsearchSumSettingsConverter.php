<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchSumSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchSumSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchSumSettingsBuilder())',
        ];
            
    
        {
    $buffer = 'script(';
        switch (true) {
            case is_string($input->script):
                $disjunctionscript =\var_export($input->script, true);
                $arg0 = $disjunctionscript;
                break;
            case $input->script instanceof \Grafana\Foundation\Elasticsearch\ElasticsearchInlineScript:
                $disjunctionscript = \Grafana\Foundation\Elasticsearch\ElasticsearchInlineScriptConverter::convert($input->script);
                $arg0 = $disjunctionscript;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->missing !== null && $input->missing !== "") {
    
        
    $buffer = 'missing(';
        $arg0 =\var_export($input->missing, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}
