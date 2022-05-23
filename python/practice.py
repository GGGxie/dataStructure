#!/usr/bin/python
import json,sys
jsonStr = sys.argv[2]
jsonStr=jsonStr.replace("'",'"')
jsonStr=jsonStr.replace("True","true")
data = json.loads(jsonStr)
idx = int(sys.argv[4]) - 1

count = data["data"]["count"]
if idx+1 <= count:
    host = data["data"]["pair"][idx]["host"]
    operator = data["data"]["pair"][idx]["operator"]
    oTemp="<SOPS_VAR>authman:"+operator+"</SOPS_VAR>"
    hTemp="<SOPS_VAR>authnode:"+host+"</SOPS_VAR>"
    print(oTemp.encode('utf-8'))
    print(hTemp.encode('utf-8'))
count = str(count)
c ="<SOPS_VAR>count:"+count+"</SOPS_VAR>"
print(c.encode('utf-8'))

