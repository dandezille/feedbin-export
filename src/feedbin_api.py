import json

import requests


class FeedbinApi:
    def __init__(self, user, password):
        self.__user = user
        self.__password = password

    def __api_url(self, path):
        return "https://api.feedbin.com/v2/" + path

    def __credentials(self):
        return requests.auth.HTTPBasicAuth(self.__user, self.__password)

    def __get(self, path, params=None):
        return requests.get(self.__api_url(path), auth=self.__credentials(), params=params)

    def __delete(self, path, data={}):
        headers = {"content-type": "application/json"}

        return requests.delete(
            self.__api_url(path), auth=self.__credentials(), headers=headers, data=data
        )

    def check_authenticated(self):
        return self.__get("authentication.json").status_code == 200

    def get_starred_entries(self):
        return self.__get("starred_entries.json").json()

    def get_entry_urls(self, entries):
        entries_list = ",".join([str(id) for id in entries])
        entries_params = {"ids": entries_list}

        response = self.__get("entries.json", params=entries_params)

        return {entry["id"]: entry["url"] for entry in response.json()}

    def remove_starred_entries(self, entries):
        return (
            self.__delete(
                "starred_entries.json", data=json.dumps({"starred_entries": entries})
            ).status_code
            == 200
        )
