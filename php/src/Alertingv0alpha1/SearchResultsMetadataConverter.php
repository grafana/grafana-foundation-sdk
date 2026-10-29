<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class SearchResultsMetadataConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadataBuilder())',
        ];
            if ($input->continue !== null && $input->continue !== "") {
    
        
    $buffer = 'continue(';
        $arg0 =\var_export($input->continue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->totalHits !== null) {
    
        
    $buffer = 'totalHits(';
        $arg0 =\var_export($input->totalHits, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->totalHitsRelation !== null) {
    
        
    $buffer = 'totalHitsRelation(';
        $arg0 ='\Grafana\Foundation\Alertingv0alpha1\TotalHitsRelation::fromValue("'.$input->totalHitsRelation.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

