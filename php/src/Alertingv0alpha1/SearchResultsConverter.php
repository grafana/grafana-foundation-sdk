<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class SearchResultsConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\SearchResults $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\SearchResultsBuilder())',
        ];
            
    
        {
    $buffer = 'metadata(';
        $arg0 = \Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadataConverter::convert($input->metadata);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->items) >= 1) {
    
        
    $buffer = 'items(';
        $tmparg0 = [];
        foreach ($input->items as $arg1) {
        $tmpitemsarg1 = \Grafana\Foundation\Alertingv0alpha1\SearchResultHitConverter::convert($arg1);
        $tmparg0[] = $tmpitemsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->facets !== null) {
    
        
    $buffer = 'facets(';
        $arg0 = "[";
        foreach ($input->facets as $key => $arg1) {
            $tmptmpfacetsarg1 = [];
        foreach ($arg1 as $arg1Value) {
        $tmparg1arg1Value = \Grafana\Foundation\Alertingv0alpha1\FacetValueConverter::convert($arg1Value);
        $tmptmpfacetsarg1[] = $tmparg1arg1Value;
        }
        $tmpfacetsarg1 = "[" . implode(", \n", $tmptmpfacetsarg1) . "]";
            $arg0 .= "\t".var_export($key, true)." => $tmpfacetsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

