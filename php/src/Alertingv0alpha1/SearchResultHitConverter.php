<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class SearchResultHitConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\SearchResultHit $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\SearchResultHitBuilder())',
        ];
            
    
        {
    $buffer = 'resource(';
        $arg0 = \Grafana\Foundation\Alertingv0alpha1\SearchResultResourceConverter::convert($input->resource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->score !== null) {
    
        
    $buffer = 'score(';
        $arg0 =\var_export($input->score, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'fields(';
        $arg0 = \Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFieldsConverter::convert($input->fields);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

