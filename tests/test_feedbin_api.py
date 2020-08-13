import base64
import json

import pytest
import requests
import responses

import feedbin_to_todoist.feedbin_api as feedbin_api


@pytest.fixture
def feedbin():
    return feedbin_api.FeedbinApi("user", "password")


def has_auth_header(request):
    return request.headers["Authorization"] == "Basic {}".format(
        base64.b64encode(b"user:password").decode("ascii")
    )


def test_api_url():
    assert feedbin_api._api_url("test.json") == "https://api.feedbin.com/v2/test.json"


@responses.activate
def test_check_authenticated_true(feedbin):
    responses.add(
        responses.GET, feedbin_api._api_url("authentication.json"), status=200
    )

    assert feedbin.check_authenticated() == True
    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)


@responses.activate
def test_check_authenticated_false(feedbin):
    responses.add(
        responses.GET, feedbin_api._api_url("authentication.json"), status=401
    )

    assert feedbin.check_authenticated() == False
    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)


@responses.activate
def test_get_starred_entries(feedbin):
    responses.add(
        responses.GET,
        feedbin_api._api_url("starred_entries.json"),
        status=200,
        body="[42,57]",
    )

    entries = feedbin.get_starred_entries()
    assert entries == [42, 57]

    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)


@responses.activate
def test_get_starred_entries_fail(feedbin):
    responses.add(
        responses.GET,
        feedbin_api._api_url("starred_entries.json"),
        status=401
    )

    with pytest.raises(Exception) as ex:
        feedbin.get_starred_entries()

    assert 'Status code 401' in str(ex.value)

    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)


@responses.activate
def test_get_entry_urls(feedbin):
    data = [
        {"id": 42, "url": "https://test.example.com"},
        {"id": 57, "url": "https://test2.example.com"},
    ]
    responses.add(
        responses.GET,
        feedbin_api._api_url("entries.json"),
        status=200,
        body=json.dumps(data),
    )

    urls = feedbin.get_entry_urls([42, 57])
    assert urls == {42: "https://test.example.com", 57: "https://test2.example.com"}

    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)

@responses.activate
def test_get_entry_urls_fail(feedbin):
    responses.add(
        responses.GET,
        feedbin_api._api_url("entries.json"),
        status=401,
    )

    with pytest.raises(Exception) as ex:
        feedbin.get_entry_urls([42, 57])

    assert 'Status code 401' in str(ex.value)

    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)

@responses.activate
def test_remove_starred_entries(feedbin):
    responses.add(
        responses.DELETE,
        feedbin_api._api_url("starred_entries.json"),
        status=200,
    )

    feedbin.remove_starred_entries([42, 57])
    assert responses.calls[0].request.body == '{"starred_entries": [42, 57]}'

    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)
