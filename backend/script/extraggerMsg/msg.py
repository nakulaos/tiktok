import re
import json
import sys

def extract_msg_tags(file_path):
    with open(file_path, 'r') as file:
        content = file.read()

    pattern = re.compile(r'msg:"(.*?)"')
    matches = pattern.findall(content)

    result = {}
    added = set()

    for match in matches:
        if match not in added:
            added.add(match)
            result[match] = ""

    with open('result.json', 'w') as outfile:
        json.dump(result, outfile, indent=4)

if len(sys.argv) > 1:
    extract_msg_tags(sys.argv[1])
else:
    print("Please provide a file name as a command line argument.")