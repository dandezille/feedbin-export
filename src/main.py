import json
import sys
from pprint import pp

from dotenv import load_dotenv
from todoist.api import TodoistAPI

import feedbin_api as feedbin


def fail(msg):
    sys.stderr.write("Fail: {}\n".format(msg))
    exit(1)


if __name__ == "__main__":
    print("Loading environment")
    load_dotenv()

    print("Authenticating")
    if not feedbin.check_authenticated():
        fail("Failed to authenticate")

    print("Fetching starred entries")
    starred_ids = feedbin.get_starred_entries()
    print("Received starred entries:")
    pp(starred_ids)

    print("Fetching entry urls")
    entry_urls = feedbin.get_entry_urls(starred_ids)
    print("Received urls:")
    pp(entry_urls)

    exit(0)

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
