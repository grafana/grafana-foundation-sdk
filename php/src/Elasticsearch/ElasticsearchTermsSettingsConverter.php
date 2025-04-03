<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchTermsSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchTermsSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchTermsSettingsBuilder())',
        ];
            if ($input->order !== null) {
    
        
    $buffer = 'order(';
        $arg0 ='\Grafana\Foundation\Elasticsearch\TermsOrder::fromValue("'.$input->order.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->size !== null && $input->size !== "") {
    
        
    $buffer = 'size(';
        $arg0 =\var_export($input->size, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->minDocCount !== null && $input->minDocCount !== "") {
    
        
    $buffer = 'minDocCount(';
        $arg0 =\var_export($input->minDocCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->orderBy !== null && $input->orderBy !== "") {
    
        
    $buffer = 'orderBy(';
        $arg0 =\var_export($input->orderBy, true);
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

