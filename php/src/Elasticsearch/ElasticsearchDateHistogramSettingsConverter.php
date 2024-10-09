<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchDateHistogramSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchDateHistogramSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchDateHistogramSettingsBuilder())',
        ];
            if ($input->interval !== null && $input->interval !== "") {
    
        
    $buffer = 'interval(';
        $arg0 =\var_export($input->interval, true);
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
            if ($input->trimEdges !== null && $input->trimEdges !== "") {
    
        
    $buffer = 'trimEdges(';
        $arg0 =\var_export($input->trimEdges, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->offset !== null && $input->offset !== "") {
    
        
    $buffer = 'offset(';
        $arg0 =\var_export($input->offset, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeZone !== null && $input->timeZone !== "") {
    
        
    $buffer = 'timeZone(';
        $arg0 =\var_export($input->timeZone, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

