<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchMovingFunctionSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchMovingFunctionSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchMovingFunctionSettingsBuilder())',
        ];
            if ($input->window !== null && $input->window !== "") {
    
        
    $buffer = 'window(';
        $arg0 =\var_export($input->window, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
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
    
    
            if ($input->shift !== null && $input->shift !== "") {
    
        
    $buffer = 'shift(';
        $arg0 =\var_export($input->shift, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

