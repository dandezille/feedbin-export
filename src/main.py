import json
import sys
from pprint import pp

from dotenv import load_dotenv

import feedbin_api as feedbin
import todoist_api as todoist


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

    print("Fetching inbox project")
    inbox = todoist.get_inbox()

    print("Adding urls to inbox")
    inbox.add_tasks(entry_urls)

    print("Removing stars from entries")
    feedbin.remove_starred_entries(starred_ids)
