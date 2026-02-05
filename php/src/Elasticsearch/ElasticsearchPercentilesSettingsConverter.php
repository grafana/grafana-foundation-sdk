<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchPercentilesSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchPercentilesSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchPercentilesSettingsBuilder())',
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
            if ($input->percents !== null && count($input->percents) >= 1) {
    
        
    $buffer = 'percents(';
        $tmparg0 = [];
        foreach ($input->percents as $arg1) {
        $tmppercentsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmppercentsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

