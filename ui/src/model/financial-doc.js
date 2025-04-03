import m from 'mithril'
import { to_lakhs, to_crores } from '../lib/util';

const isEmpty = value => {
    if (value === 0) {
        return true
    }else if (typeof value === 'string' || Array.isArray(value)) {
        return value.length === 0;
    } else if (typeof value === 'object' && value !== null) {
        return Object.keys(value).length === 0;
    } else {
        return value === null || value === undefined;
    }
};

const allValuesEmpty = arr => arr.every(isEmpty);

function doc_header(doc, unit_function) {
    let headers = []
    let unit_string = unit_function === to_lakhs ? 'Lakhs' : 'Crores';
    headers.push(`(INR in ${unit_string})`)
    headers.push(...doc.data.contexts.map(ctx => {
        if (ctx.type === 'duration') {
            return `${ctx.duration} months<br/>${ctx.end_date}`
        }
        return ctx.end_date
    }))
    return headers
}

function doc_row(item, unit_function) {
    if ((item.concept.dimensions === null || item.concept.dimensions.length === 0) && item.facts && item.facts.length > 0) {
        if (item.concept.name === 'DisclosureOfNotesOnFinancialResultsExplanatoryTextBlock') {
            return []
        }
        let is_monetary = item.concept.type === 'monetaryItemType';
        let is_percent = item.concept.type === 'percentItemType';
        const facts = item.facts.map(f => {
            if (!f.is_nil) {
                if (f.is_text) {
                    return f.text;
                } else {
                    if (f.value == 0) {
                        return 0
                    }
                    if (is_monetary) {
                        return unit_function(f.value, 0)
                    } if (is_percent) {
                        let d = ((f.value) * (100.0)).toFixed(2);
                        return `${d} %`
                    } else {
                        return f.value;
                    }
                }
            } else {
                return []
            }
        });
        if (!allValuesEmpty(facts)) {
            return [item.concept.label, ...facts] // Use spread operator to flatten the array
        }
    }
    return []
}


var FinancialDoc = {
    unit_function: to_lakhs,
    headers: [],
    rows: [],
    result: [],
    loadSampleFinancialDocs: function (ticker, filing_id, doc_id) {
        FinancialDoc.headers = []
        FinancialDoc.rows = []
        FinancialDoc.result = []
    },
    load: function (ticker, filing_id, doc_id) {
        return m.request({
            method: "GET",
            url: `/nse/company/${ticker}/financials/${filing_id}/docs/${doc_id}`
        }).then(function (result) {
            FinancialDoc.result = result
            if (result.length > 0) {
                FinancialDoc.headers = doc_header(result[0], FinancialDoc.unit_function)
                FinancialDoc.rows = result[0].data.items.map(item => ({ 
                    is_total: item.concept.is_total || item.concept.label_type === 'net',
                    is_abstract: item.concept.is_abstract,
                    elements: doc_row(item, FinancialDoc.unit_function)
                }))
            }
        }).catch(function (error) {
            console.error(error)
        })
    },
    toggleUnit: function(isChecked) {
        if (isChecked) {
            FinancialDoc.unit_function = to_crores;
        } else {
            FinancialDoc.unit_function = to_lakhs;
        }
        FinancialDoc.headers = doc_header(FinancialDoc.result[0], FinancialDoc.unit_function)
        FinancialDoc.rows = FinancialDoc.result[0].data.items.map(item => ({ 
            is_total: item.concept.is_total,
            is_abstract: item.concept.is_abstract,
            elements: doc_row(item, FinancialDoc.unit_function)
        }))
    }
}

export default FinancialDoc;