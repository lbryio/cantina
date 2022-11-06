import requests, binascii, simplejson.errors, html, json

as_header = {'Accept': 'application/ld+json; profile="https://www.w3.org/ns/activitystreams"'}
def get_ap(url):
    response = requests.get(url, headers=as_header)
    try:
        return response.json(), None
    except simplejson.errors.JSONDecodeError:
        return None, response.text

def is_link(s):
    return s.lower().startswith("http") and not s.lower().endswith("jpeg") and not s.lower().endswith("jpg") and not s.lower().endswith("png")

def get_links(json_obj):
    if isinstance(json_obj, str):
        if is_link(json_obj):
            return [json_obj]
        else:
            return []

    if isinstance(json_obj, dict):
        return sum(map(get_links, json_obj.values()), [])

    if isinstance(json_obj, list):
        return sum(map(get_links, json_obj), [])

    # Various other expected types
    if isinstance(json_obj, (int, float)) or json_obj is None:
        return []

    raise Exception("unexpected type:", type(json_obj))

def render(json_obj):
    links = get_links(json_obj)

    json_str = json.dumps(json_obj, indent=4)

    out = "<pre>" + html.escape(json_str) + "</pre>"

    for link in links:
        # Assume our json is trustworthy
        # Need to surround by '' because urls are substrings of other urls
        print(html.escape('"%s"' % link))
        out = out.replace(
            html.escape('"%s"' % link),
            "\"<a href='/?l=%s'>%s</a>\"" % (binascii.hexlify(bytes(link, 'utf-8')).decode('utf-8'), link)
        )

    return out

def get(url):
    # here, raw is probably an error
    json_obj, raw = get_ap(url)

    if json_obj:
        return render(json_obj)
    else:
        # probably an error
        return raw
