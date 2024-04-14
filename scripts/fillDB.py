import requests
import json
import random

base_url = 'http://localhost:7070'

token = 'admin_token'

count = 100

tag_ids_array = [['1', '2'], ['2', '3'], ['1', '3'], ['3', '4'], ['4', '5'],
['11', '12'], ['12', '31'], ['11', '13'], ['13', '14'], ['14', '15']]
feature_ids = [100, 101, 102, 103, 104, 105, 106, 107]
contents = [
    {"title": "Sale", "text": "Big sale this weekend!", "url": "http://sale.com"},
    {"title": "New Product", "text": "Check out our new product!", "url": "http://product.com"},
    {"title": "Event", "text": "Join our event tomorrow!", "url": "http://event.com"}
]

for _ in range(count):
    tag_ids = random.choice(tag_ids_array)
    feature_id = random.choice(feature_ids)
    content = random.choice(contents)
    is_active = random.choice([True, False])

    json_data = json.dumps({
        "tag_ids": list(map(int, tag_ids)),
        "feature_id": feature_id,
        "content": content,
        "is_active": is_active
    })

    headers = {
        'Content-Type': 'application/json',
        'token': token
    }

    response = requests.post(f"{base_url}/banner", headers=headers, data=json_data)
    print(f"Status Code: {response.status_code}, Response: {response.json()}")


print("Создание баннеров завершено.")
