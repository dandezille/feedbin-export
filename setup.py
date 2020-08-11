import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

setuptools.setup(
    name="feedbin-to-todoist",
    version="0.0.3",
    author="Dan de Zille",
    author_email="dan@ddez.net",
    description="Export starred Feedbin articles to Todoist inbox",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/dandezille/feedbin-to-todoist",
    packages=setuptools.find_packages(exclude=("tests",)),
    install_requires=["requests", "todoist-python"],
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires=">=3.8",
)
