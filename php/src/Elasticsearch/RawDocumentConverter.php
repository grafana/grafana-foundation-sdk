<?php

namespace Grafana\Foundation\Elasticsearch;

final class RawDocumentConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\RawDocument $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\RawDocumentBuilder())',
        ];
            if ($input->id !== "") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->settings !== null) {
    
        
    $buffer = 'settings(';
        $arg0 = \Grafana\Foundation\Elasticsearch\ElasticsearchRawDocumentSettingsConverter::convert($input->settings);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

