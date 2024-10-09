<?php

namespace Grafana\Foundation\Elasticsearch;

final class SerialDiffConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\SerialDiff $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\SerialDiffBuilder())',
        ];
            if ($input->pipelineAgg !== null && $input->pipelineAgg !== "") {
    
        
    $buffer = 'pipelineAgg(';
        $arg0 =\var_export($input->pipelineAgg, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->field !== null && $input->field !== "") {
    
        
    $buffer = 'field(';
        $arg0 =\var_export($input->field, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->id !== "") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->settings !== null) {
    
        
    $buffer = 'settings(';
        $arg0 = \Grafana\Foundation\Elasticsearch\ElasticsearchSerialDiffSettingsConverter::convert($input->settings);
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

