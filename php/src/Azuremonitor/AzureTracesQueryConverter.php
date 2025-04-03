<?php

namespace Grafana\Foundation\Azuremonitor;

final class AzureTracesQueryConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\AzureTracesQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\AzureTracesQueryBuilder())',
        ];
            if ($input->resultFormat !== null) {
    
        
    $buffer = 'resultFormat(';
        $arg0 ='\Grafana\Foundation\Azuremonitor\ResultFormat::fromValue("'.$input->resultFormat.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resources !== null && count($input->resources) >= 1) {
    
        
    $buffer = 'resources(';
        $tmparg0 = [];
        foreach ($input->resources as $arg1) {
        $tmpresourcesarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpresourcesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->operationId !== null && $input->operationId !== "") {
    
        
    $buffer = 'operationId(';
        $arg0 =\var_export($input->operationId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->traceTypes !== null && count($input->traceTypes) >= 1) {
    
        
    $buffer = 'traceTypes(';
        $tmparg0 = [];
        foreach ($input->traceTypes as $arg1) {
        $tmptraceTypesarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmptraceTypesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->filters !== null && count($input->filters) >= 1) {
    
        
    $buffer = 'filters(';
        $tmparg0 = [];
        foreach ($input->filters as $arg1) {
        $tmpfiltersarg1 = \Grafana\Foundation\Azuremonitor\AzureTracesFilterConverter::convert($arg1);
        $tmparg0[] = $tmpfiltersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->query !== null && $input->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

