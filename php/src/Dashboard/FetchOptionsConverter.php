<?php

namespace Grafana\Foundation\Dashboard;

final class FetchOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\FetchOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\FetchOptionsBuilder())',
        ];
            
    
        {
    $buffer = 'method(';
        $arg0 ='\Grafana\Foundation\Dashboard\HttpRequestMethod::fromValue("'.$input->method.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->url !== "") {
    
        
    $buffer = 'url(';
        $arg0 =\var_export($input->url, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->body !== null && $input->body !== "") {
    
        
    $buffer = 'body(';
        $arg0 =\var_export($input->body, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryParams !== null && count($input->queryParams) >= 1) {
    
        
    $buffer = 'queryParams(';
        $tmparg0 = [];
        foreach ($input->queryParams as $arg1) {
        $tmptmpqueryParamsarg1 = [];
        foreach ($arg1 as $arg1Value) {
        $tmparg1arg1Value =\var_export($arg1Value, true);
        $tmptmpqueryParamsarg1[] = $tmparg1arg1Value;
        }
        $tmpqueryParamsarg1 = "[" . implode(", \n", $tmptmpqueryParamsarg1) . "]";
        $tmparg0[] = $tmpqueryParamsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->headers !== null && count($input->headers) >= 1) {
    
        
    $buffer = 'headers(';
        $tmparg0 = [];
        foreach ($input->headers as $arg1) {
        $tmptmpheadersarg1 = [];
        foreach ($arg1 as $arg1Value) {
        $tmparg1arg1Value =\var_export($arg1Value, true);
        $tmptmpheadersarg1[] = $tmparg1arg1Value;
        }
        $tmpheadersarg1 = "[" . implode(", \n", $tmptmpheadersarg1) . "]";
        $tmparg0[] = $tmpheadersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

