import mysql.connector
import json

db = mysql.connector.connect(
    host="localhost",
    user="test",
    password="test123",
    database="sushi",
    port=3307
)

cur = db.cursor()

with open("./sushi.json", "r") as f:
    jsonData = json.loads(f.read())

categories = {
    0: "Все",
    1: "Суши",
    2: "Роллы",
    3: "Сеты",
    4: "Сашими",
    5: "Салаты",
    6: "Десерты",
    7: "Напитки"
}

for x in jsonData:
    nutrition = json.dumps(x['nutrition'])
    query = f"insert into sushi_products (name, image, category, price, weight, nutrition, description) values ('{x['title']}', '{x['image']}', {x['category']}, {x['price']}, {x['weight']}, '{nutrition}', '{x['description']}')"
    cur.execute(query)

for key in categories:
    query = f"insert into sushi_categories (id, name) values ({key}, '{categories[key]}')"
    cur.execute(query)
db.commit()