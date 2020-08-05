import requests
import responses
import pdb
from nose.tools import assert_equals

import feedbin_api as feedbin


def test_api_url():
    assert_equals(feedbin._api_url("test.json"), "https://api.feedbin.com/v2/test.json")


@responses.activate
def test_check_authenticated_true():
    responses.add(responses.GET, feedbin._api_url("authentication.json"), status=200)
    assert_equals(feedbin.check_authenticated(), True)
    assert_equals(len(responses.calls), 1)
    assert_equals(responses.calls[0].request.headers, {})


@responses.activate
def test_check_authenticated_false():
    responses.add(responses.GET, feedbin._api_url("authentication.json"), status=401)
    assert_equals(feedbin.check_authenticated(), False)
    assert_equals(len(responses.calls), 1)
