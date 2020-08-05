import base64

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
def test_check_authenticated_true(monkeypatch):
    responses.add(responses.GET, feedbin._api_url("authentication.json"), status=200)

    assert feedbin.check_authenticated() == True
    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)


@responses.activate
def test_check_authenticated_false(monkeypatch):
    responses.add(responses.GET, feedbin._api_url("authentication.json"), status=401)

    assert feedbin.check_authenticated() == False
    assert len(responses.calls) == 1
    assert has_auth_header(responses.calls[0].request)
