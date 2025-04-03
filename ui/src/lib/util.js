const companyUrl = (ticker, api) => {
    return `/nse/company/${ticker}/${api}`
}

const to_lakhs = (n, decimals) => {
    return (n / 100000).toFixed(decimals);
}

const to_crores = (n, decimals) => {
    return (n / 10000000).toFixed(decimals);
}

export { companyUrl, to_lakhs, to_crores } ;