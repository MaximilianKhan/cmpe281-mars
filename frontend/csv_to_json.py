import csv
import json

csvfile = open('database.csv', 'r')
jsonfile = open('db.json', 'w')

fieldnames = ("id","product","price","qty","image_path","category")
reader = csv.DictReader( csvfile, fieldnames)
for row in reader:
    json.dump(row, jsonfile)
    jsonfile.write('\n')
