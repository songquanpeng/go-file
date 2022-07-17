#!/usr/bin/python3

# Reference: https://support.typora.io/Upload-Image/#custom

import sys
import requests

SERVER_URL = 'http://localhost:3000'
TOKEN = '593ff9f2e91842b497ae79ecae83f412'

for file in sys.argv[1:]:
    r = requests.post(f"{SERVER_URL}/api/image", files={'image': open(file, 'rb')}, headers={
        "Authorization": TOKEN
    })
    if r.status_code != 200:
        sys.exit(1)
    filename = r.json()["data"][0]
    print(f"{SERVER_URL}/image/{filename}")
