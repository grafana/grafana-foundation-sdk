<?php

namespace Grafana\Foundation\Resource;

final class ManifestConverter
{
    public static function convert(\Grafana\Foundation\Resource\Manifest $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Resource\ManifestBuilder())',
        ];
            if ($input->apiVersion !== "") {
    
        
    $buffer = 'apiVersion(';
        $arg0 =\var_export($input->apiVersion, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->kind !== "") {
    
        
    $buffer = 'kind(';
        $arg0 =\var_export($input->kind, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'metadata(';
        $arg0 = \Grafana\Foundation\Resource\MetadataConverter::convert($input->metadata);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec !== null) {
    
        
    $buffer = 'spec(';
        $arg0 =\var_export($input->spec, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

