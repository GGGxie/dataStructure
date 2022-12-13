# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.

import json
import ast

data = "{'message': '', 'code': 200, 'data': {'is_done': True}, 'result': True}"
data2=ast.literal_eval(data)
print(data2) 
temp = "<SOPS_VAR>done:"+str(data2['data']['is_done'])+"</SOPS_VAR>"
print(temp)
