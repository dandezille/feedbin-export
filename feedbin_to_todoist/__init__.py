import sys
from pprint import pprint as pp

from . import feedbin_api, todoist_api


def export(feedbin_user: str, feedbin_password: str, todoist_api_key: str) -> None:
    feedbin = feedbin_api.connect(feedbin_user, feedbin_password)

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
    todoist = todoist_api.connect(todoist_api_key)

    print("Adding urls to inbox")
    todoist.add_tasks(entry_urls)

    print("Removing stars from entries")
    feedbin.remove_starred_entries(starred_ids)
