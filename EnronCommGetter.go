//Package to acquire or load the demonstration enron datasets from Azure file storage
package main

import (
"github.com/Azure/azure-storage-file-go/2017-07-29/azfile"
)

func main() {
    var basicHeaders = azfile.FileHTTPHeaders{ContentType: "my_type", ContentDisposition: "my_disposition",
	   CacheControl: "control", ContentMD5: nil, ContentLanguage: "my_language", ContentEncoding: "my_encoding"}
    var basicMetadata = azfile.Metadata{"foo": "bar"}
    var name = 
    var key = 
    credential, err := azfile.NewSharedKeyCredential(name, key)
    var shareName
    var directoryName
    var shareUrl
    var directoryUrl
    
    
}
