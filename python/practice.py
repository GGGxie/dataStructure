#!/usr/bin/python
import sys
str = sys.argv[1]
str = str.encode("utf-8").decode("unicode_escape")
str = str.strip("/")
temp = "<SOPS_VAR>text:"+str+"</SOPS_VAR>"
print(temp.encode("utf-8"))

str = sys.argv[2]
str = str.strip("/")
temp = "<SOPS_VAR>person:"+str+"</SOPS_VAR>"
print(temp.encode("utf-8"))