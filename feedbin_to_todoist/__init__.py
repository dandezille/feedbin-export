import sys
from pprint import pp

from . import todoist_api as todoist
from .feedbin_api import FeedbinApi


def export(feedbin_user, feedbin_password, todoist_api_key):
    feedbin = FeedbinApi(feedbin_user, feedbin_password)

    print("Authenticating")
    if not feedbin.check_authenticated():
        print("Failed to authenticate")
        return

    print("Fetching starred entries")
    starred_ids = feedbin.get_starred_entries()

    if not starred_ids:
        print("No starred entries found")
        return

    print("Received starred entries:")
    pp(starred_ids)

    print("Fetching entry urls")
    entry_urls = feedbin.get_entry_urls(starred_ids)
    print("Received urls:")
    pp(entry_urls)

    print("Fetching inbox project")
    inbox = todoist.get_inbox(todoist_api_key)

    print("Adding urls to inbox")
    inbox.add_tasks(entry_urls)

    print("Removing stars from entries")
    feedbin.remove_starred_entries(starred_ids)
