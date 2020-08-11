import sys

from todoist.api import TodoistAPI


class TodoistProject:
    def __init__(self, api, project_id):
        self.__api = api
        self.__project_id = project_id

    def add_tasks(self, urls):
        for url in urls.values():
            self.__api.items.add(url, project_id=self.__project_id)

        self.__api.commit()


def _fail(msg):
    sys.stderr.write("Fail: {}\n".format(msg))
    exit(1)


def get_inbox(api_key):
    todoist = TodoistAPI(api_key)
    todoist.sync()

    id = [
        project["id"]
        for project in todoist.state["projects"]
        if project["name"] == "Inbox"
    ]
    if len(id) != 1:
        _fail("Expected a single project id")

    return TodoistProject(todoist, todoist.projects.get_by_id(int(id[0])))
