import json
import os
import sys
from pprint import pp

from dotenv import load_dotenv

import todoist_api as todoist
from feedbin_api import FeedbinApi


def fail(msg):
    sys.stderr.write("Fail: {}\n".format(msg))
    exit(1)


if __name__ == "__main__":
    print("Loading environment")
    load_dotenv()

    feedbin = FeedbinApi(os.getenv("FEEDBIN_USER"), os.getenv("FEEDBIN_PASSWORD"))

    print("Authenticating")
    if not feedbin.check_authenticated():
        fail("Failed to authenticate")

    print("Fetching starred entries")
    starred_ids = feedbin.get_starred_entries()

    if not starred_ids:
        print("No starred entries found")
        exit(0)

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
