import json
from typing import Dict, List

import requests


def _api_url(path: str) -> str:
    return "https://api.feedbin.com/v2/" + path


class FeedbinApi:
    def __init__(self, user: str, password: str) -> None:
        self.__user = user
        self.__password = password

    def __credentials(self) -> requests.auth.HTTPBasicAuth:
        return requests.auth.HTTPBasicAuth(self.__user, self.__password)

    def __get(self, path: str, params=None) -> requests.Response:
        return requests.get(_api_url(path), auth=self.__credentials(), params=params)

    def __delete(self, path: str, data={}) -> requests.Response:
        headers = {"content-type": "application/json"}

        return requests.delete(
            _api_url(path), auth=self.__credentials(), headers=headers, data=data
        )

    def check_authenticated(self) -> bool:
        return self.__get("authentication.json").status_code == 200

    def get_starred_entries(self) -> List[int]:
        return self.__get("starred_entries.json").json()

    def get_entry_urls(self, entries: List[int]) -> Dict[int, str]:
        entries_list = ",".join([str(id) for id in entries])
        entries_params = {"ids": entries_list}

        response = self.__get("entries.json", params=entries_params)

        return {entry["id"]: entry["url"] for entry in response.json()}

    def remove_starred_entries(self, entries: List[int]) -> bool:
        return (
            self.__delete(
                "starred_entries.json", data=json.dumps({"starred_entries": entries})
            ).status_code
            == 200
        )
