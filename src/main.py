import os

from dotenv import load_dotenv

from feedbin_to_todoist import feedbin_to_todoist

if __name__ == "__main__":
    print("Loading environment")
    load_dotenv()

    feedbin_to_todoist(
        os.getenv("FEEDBIN_USER"),
        os.getenv("FEEDBIN_PASSWORD"),
        os.getenv("TODOIST_API_KEY"),
    )
