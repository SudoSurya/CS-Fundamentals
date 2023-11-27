
import json

def getJSONDiff(json1, json2):
    obj1 = json.loads(json1)
    obj2 = json.loads(json2)

    keys1 = set(obj1.keys())
    keys2 = set(obj2.keys())

    common_keys = keys1.intersection(keys2)

    different_keys = [key for key in common_keys if obj1[key] != obj2[key]]

    sorted_result = sorted(different_keys)

    return sorted_result

# Example usage:
json1 = '{"hacker":"rank","input":"wrong"}'
json2= '{}'

result = getJSONDiff(json1, json2)
print(result)  # Output: ["hi"]
