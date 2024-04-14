import requests

def get_all_banners(base_url, token):
    url = f"{base_url}/banner?tag_id=2"
    headers = {'token': token}
    response = requests.get(url, headers=headers)
    print(response.json())


base_url = 'http://localhost:7070'
admin_token = 'admin_token'
get_all_banners(base_url, admin_token)
