#!/bin/bash

#셀프인증 
openssl x509 -req -days 365 -in localhost.csr -signkey localhost.key -out localhost.crt
