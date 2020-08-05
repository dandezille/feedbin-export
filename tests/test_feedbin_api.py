import base64
import json

import pytest
import requests
import responses

import src.feedbin_api as feedbin


@pytest.fixture(autouse=True)
def run_around_tests(monkeypatch):
    monkeypatch.setenv("FEEDBIN_USER", "user")
    monkeypatch.setenv("FEEDBIN_PASSWORD", "password")
    yield


def has_auth_header(request):
    return request.headers["Authorization"] == "Basic {}".format(
        base64.b64encode(b"user:password").decode("ascii")
    )


def test_api_url():
    assert feedbin._api_url("test.json") == "https://api.feedbin.com/v2/test.json"


@responses.activate
def test_check_authenticated_true():
    responses.add(responses.GET, feedbin._api_url("authentication.json"), status=200)

    assert feedbin.check_authenticated() == True
    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)


@responses.activate
def test_check_authenticated_false():
    responses.add(responses.GET, feedbin._api_url("authentication.json"), status=401)

    assert feedbin.check_authenticated() == False
    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)


@responses.activate
def test_get_starred_entries():
    responses.add(
        responses.GET,
        feedbin._api_url("starred_entries.json"),
        status=200,
        body="[42,57]",
    )

    entries = feedbin.get_starred_entries()
    assert entries == [42, 57]

    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)


@responses.activate
def test_get_entry_urls():
    data = [
        {"id": 42, "url": "https://test.example.com"},
        {"id": 57, "url": "https://test2.example.com"},
    ]
    responses.add(
        responses.GET,
        feedbin._api_url("entries.json"),
        status=200,
        body=json.dumps(data),
    )

    urls = feedbin.get_entry_urls([42, 57])
    assert urls == {42: "https://test.example.com", 57: "https://test2.example.com"}

    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)
