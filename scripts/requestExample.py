import requests

def get_all_banners(base_url, token):
    url = f"{base_url}/banner?tag_id=2"
    headers = {'token': token}
    response = requests.get(url, headers=headers)
    print(response.json())

def create_banners_with_same_tag_x_feature(base_url, token):
    url = f"{base_url}/banner?"
    headers = {'token': token}
    args = {
             "tag_ids": [
               1,2
             ],
             "feature_id": 2,
             "content": {
               "title": "some_title",
               "text": "some_text",
               "url": "some_url"
             },
             "is_active": True
           }
    response = requests.post(url, headers=headers, json=args)
    print(response.json())
    url = f"{base_url}/banner?"
    headers = {'token': token}
    args = {
             "tag_ids": [
               2,4
             ],
             "feature_id": 2,
             "content": {
               "title": "some_title",
               "text": "some_text",
               "url": "some_url"
             },
             "is_active": True
           }
    response = requests.post(url, headers=headers, json=args)
    print(response.json())


base_url = 'http://localhost:7070'
admin_token = 'admin_token'
create_banners_with_same_tag_x_feature(base_url, admin_token)
