import os

from dotenv import load_dotenv

from feedbin_to_todoist import export


def getenv(name: str) -> str:
    value = os.getenv(name)
    if value:
        return value

    raise KeyError("Environment variable {} not found".format(name))


def main() -> None:
    export(
        getenv("FEEDBIN_USER"), getenv("FEEDBIN_PASSWORD"), getenv("TODOIST_API_KEY"),
    )


def function(request, msg) -> None:
    main()


if __name__ == "__main__":
    print("Loading environment")
    load_dotenv()

    main()
