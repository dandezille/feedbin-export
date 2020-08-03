import json
import os
import sys

import requests
from dotenv import load_dotenv
from todoist.api import TodoistAPI


def fail(msg):
    sys.stderr.write("Fail: {}\n".format(msg))
    exit(1)


def feedbin(path, params=None):
    base_url = "https://api.feedbin.com/v2/"
    credentials = requests.auth.HTTPBasicAuth(
        os.getenv("FEEDBIN_USER"), os.getenv("FEEDBIN_PASSWORD")
    )

    return requests.get(base_url + path, auth=credentials, params=params)


def feedbin_delete(path, data):
    base_url = "https://api.feedbin.com/v2/"
    credentials = requests.auth.HTTPBasicAuth(
        os.getenv("FEEDBIN_USER"), os.getenv("FEEDBIN_PASSWORD")
    )
    headers = {"content-type": "application/json"}

    return requests.delete(
        base_url + path, auth=credentials, headers=headers, data=data
    )


if __name__ == "__main__":
    load_dotenv()

    if not feedbin("authentication.json").status_code == 200:
        fail("Failed to authenticate")

    starred_ids = feedbin("starred_entries.json").json()
    print(starred_ids)

    entries_list = ",".join([str(id) for id in starred_ids])
    entries_params = {"ids": entries_list}

    entry_details = feedbin("entries.json", params=entries_params)
    print(entry_details)
    print(entry_details.request.url)

    urls = [(entry["id"], entry["url"]) for entry in entry_details.json()]
    print(urls)

    todoist = TodoistAPI("cf137683c65c146b2d358fc95c28a73270540e95")
    todoist.sync()

    id = [
        project["id"]
        for project in todoist.state["projects"]
        if project["name"] == "Inbox"
    ]
    if len(id) != 1:
        fail("Expected a single project id")

    inbox = todoist.projects.get_by_id(int(id[0]))

    for (id, url) in urls:
        todoist.items.add(url, project_id=inbox)

    todoist.commit()

    req = feedbin_delete(
        "starred_entries.json", data=json.dumps({"starred_entries": starred_ids})
    )

    print(req.status_code)
    print(req.request.url)
    print(req.request.headers)
    print(req.request.body)
