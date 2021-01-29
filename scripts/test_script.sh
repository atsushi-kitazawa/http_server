#!/bin/sh

echo "GET [http://localhost:9999/]"
curl http://localhost:9999/
echo

echo "GET [http://localhost:9999/pages/hello.html]"
curl http://localhost:9999/pages/hello.html
echo

echo "POST [http://localhost:9999/pages/hello.html]"
curl -X POST -d key1=valu1 -d key2=value2 http://localhost:9999/pages/hello.html
echo

echo "GET [http://localhost:9999/pages/auth_dir/hello.html]"
curl -u test:password -i -s http://localhost:9999/pages/auth_dir/hello.html
echo

echo "GET [http://localhost:9999/pages/auth_dir/hello.html]"
curl -u test:invalidpassword -i -s http://localhost:9999/pages/auth_dir/hello.html
echo
