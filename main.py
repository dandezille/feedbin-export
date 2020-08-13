import os

from dotenv import load_dotenv

from feedbin_to_todoist import export


def main():
    export(
        os.getenv("FEEDBIN_USER"),
        os.getenv("FEEDBIN_PASSWORD"),
        os.getenv("TODOIST_API_KEY"),
    )


def function(request, msg):
    main()


if __name__ == "__main__":
    print("Loading environment")
    load_dotenv()

    main()
