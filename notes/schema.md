# API
## Get filings of a company
GET /api/v1/symbol/:symbol/filings?limit=limit?format=html
[{filing_id, filing_date, is_amended, filing_type, filing_link}, ...]

## Get single filing
GET /api/v1/symbol/:symbol/filings/:filing_id?format=html
[{filing_id, filing_date, is_amended, filing_type, filing_link}] only instance

## Get documents within a filing
GET /api/v1/filings/:filing_id/docs?type=statements,format=html
[{id, name, type, title}, ... ]

## Get single document within a filing
GET /api/v1/filings/:filing_id/docs/:doc_id?format=html
[{id, name, type, title, data}]


# Schema

## Symbols
symbol text PK
entity int64 CK

## Filings
filing_id text     PK
filing_date date   CK
entity int Global secondary index
symbols set<text>
is_amended boolean,
company_name text,
filing_link text

## Documents
filing_id PK
type      CK
doc_id    CK  Local secondary index on filing_id, doc_id
name
title
data blob


# Filing -> Dei, [Document]
# Document -> seq, Definition, [Context], [Item]
# Definition -> _id, type, tag, title, sort_code
# Context -> _xml_id, type, end_date, start_date, _period_type, duration
# Item -> seq, level, Concept, Fact, [Dimension]
# Concept -> id, name, type, balance, label
# Fact -> value, decimals, unit, unit_short, text, is_text
# Dimension -> axis, axis_label, member, member_label

# Archive
## Statements
entity: int # cik, scrip code
statement_type: text # IS, BS, CF
period_type: int # qtr, semi, annual, tri (9)
filing_date: date # form filing date
end_date: date # end date of reporting period
seq: int # dictates order of the facts

#
Axis                    Context
    Member
        Concept