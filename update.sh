#!/bin/bash
mkdir /tmp/sla || true
d=$(date +%Y-%m-%d).txt
if ! test -s $d
then
	curl 'http://www.landapplications.gov.sg/SPIOWeb/SPIO/Public/ResidentialMashupMap.aspx/GetSPIOArray' -X POST -H 'Pragma: no-cache' -H 'Origin: http://www.landapplications.gov.sg' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-GB,en;q=0.8' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.10 Safari/537.36' -H 'Content-Type: application/json; charset=UTF-8' -H 'Accept: */*' -H 'Cache-Control: no-cache' -H 'Referer: http://www.landapplications.gov.sg/SPIOWeb/SPIO/Public/ResidentialMashupMap.aspx' -H 'Cookie: ASP.NET_SessionId=oyurzo55drvcos55l2ajrl55' -H 'Connection: keep-alive' -H 'Content-Length: 0' -H 'DNT: 1' --compressed > $d
fi
./sg-accomodation < $d
mv *.html /tmp/sla
