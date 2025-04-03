# Goal
- Implement a performant and cost effective api to support stock fundamental data
- The api should be easily extensible
- Rate limiting support
- Api key auth support
- Excellent and fast documentation
- API status page
- Monitoring, tracing and alerting
- Fast (lt; 15 mins) update of new forms from sec
- Data ingested from sec.gov should be reusable without downloading again


# Authentication
API authentication via api key. The key should be present in the authentication header. This is as per standard RFC: https://datatracker.ietf.org/doc/html/rfc7235https://datatracker.ietf.org/doc/html/rfc7235
'Authentication: Bearer token'
Query param should not be used for api key to prevent credential leaking as it is highly unsecure.

# Documentation
- Use swagger to generate api documentation.
- Ensure to apply beautiful themes to documentation for excellent impression.

# API design
Fundamental data: 10-K, 10-K/A, 10-Q, 10-Q/A, 20-F, 40-F.
- /api/income/{company}?period=annual|quarter
- /api/balancesheet/{company}?period=annual|quarter
- /api/cashflow/{company}?period=annual|quarter

Growth
- /api/growth/{company}?period=annual|quarter

Segment information
- /api/segment/{company}/revenue?period=annual|quarter

Insider transaction: Form 3,4,5
- /api/insider/{company} : This should support pagination as the list can be long

Future
- Ratios
- Growth
- Segment
- Insider
- Statement as reported
- Stock ownership: Edgar
- ETF: Edgar
- Mutual funds: Edgar
- Press releases ?
- Quote api: FMP api
- Trade signals: EMA, SMA, ...
- Webhooks: Notifications for trade signals, stock ownership changes, whale moves, press releases etc