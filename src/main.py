import os
import sys

import requests
from dotenv import load_dotenv


def fail(msg):
    sys.stderr.write("Fail: {}\n".format(msg))
    exit(1)


def feedbin_request(path, params=None):
    base_url = "https://api.feedbin.com/v2/"
    credentials = requests.auth.HTTPBasicAuth(
        os.getenv("FEEDBIN_USER"), os.getenv("FEEDBIN_PASSWORD")
    )

    return requests.get(base_url + path, auth=credentials, params=params)


def authenticated():
    req = feedbin_request("authentication.json")
    return req.status_code == 200


def get_starred_ids():
    return feedbin_request("starred_entries.json").json()


if __name__ == "__main__":
    load_dotenv()

    if not authenticated():
        fail("Failed to authenticate")

    starred_ids = get_starred_ids()
    print(starred_ids)

    entries_params = { 'ids': ','.join([str(id) for id in starred_ids]) }

    entry_details = feedbin_request("entries.json", params=entries_params)
    print(entry_details)
    print(entry_details.request.url)

    urls = [(entry["id"], entry["url"]) for entry in entry_details.json()]
    print(urls)
