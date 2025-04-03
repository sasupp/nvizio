# XBRL to JSON Api
GET JSON representation of XBRL filing

`GET v1/{code}/filing/{filing_id}`

### Parameters

| Parameter   | Description                                  
| ----------- | -------------------------------------------------------------------------------------
| `code`      | Valid values are ```sec``` for USA filings, ```nse``` for India filings.
| `filing_id` | Filing id of the XBRL submission. This is equal to accession number for SEC filings.

### Responses

| Code  | Description                                                                   |
| ----- | ----------------------------------------------------------------------------- |
| `200` | A JSON object containing every information available on a specific filing     |
| `401` | Unauthorized                                                                  |
| `404` | Filing not found                                                              |

### Examples

??? success "200 Successful Response"

    `v1/sec/filing/0001628280-23-027039` returns

    ``` JSON
    {}
    ```

??? fail "404 Not Found"

    `/v1/sec/filing/0000000000-00-000000` returns

    ``` JSON
    {
        "status": 404,
        "error": "FilingNotFound",
    }
    ```