const instance = axios.create({timeout: 60 * 1000})
instance.interceptors.request.use(
    config => {
        config.headers['Authorization'] = 'Bearer ' + enJwt()
        return config
    },
    error => Promise.reject(error))

async function ping() {
    return httpPost(instance, '../../api/ping', {})
}

async function GetConfig() {
    let url = '../../api/GetConfig'
    if (document.domain === 'localhost') {
        url += '.json'
    }
    try {
        let resp = await instance.get(url, {
            params: {},
            paramsSerializer: params => {
                return Qs.stringify(params, {indices: false})
            }
        })
        return dealResp(resp)
    } catch (error) {
        // dealErr(error)
    }
    return null
}

async function Register(expire_date, activation_code, baidu_ocr_api_key, baidu_ocr_secret_key) {
    if (expire_date === undefined || expire_date == null || expire_date === '') {
        dealErr('激活到期日为空')
        return null
    }
    expire_date = getTimestamp(new Date(expire_date)) + (60 * 60 * 24)
    if (expire_date <= 0) {
        dealErr('激活到期日为空')
        return null
    }
    if (activation_code === undefined || activation_code == null || activation_code === '') {
        dealErr('激活码为空')
        return null
    }
    if (baidu_ocr_api_key === undefined || baidu_ocr_api_key == null || baidu_ocr_api_key === '') {
        dealErr('百度OCR ApiKey为空')
        return null
    }
    if (baidu_ocr_secret_key === undefined || baidu_ocr_secret_key == null || baidu_ocr_secret_key === '') {
        dealErr('百度OCR SecretKey为空')
        return null
    }

    return httpPost(instance, '../../api/Register', {
        expire_date: expire_date,
        activation_code: activation_code,
        baidu_ocr_api_key: baidu_ocr_api_key,
        baidu_ocr_secret_key: baidu_ocr_secret_key,
    })
}

async function AddJob(sites, word) {
    if (sites === undefined || sites == null || sites === '' || sites.length === 0) {
        dealErr('所选爬取网址为空')
        return null
    }
    if (word === undefined || word == null || word === '') {
        dealErr('所填爬取关键字为空')
        return null
    }

    if (!window.confirm("确定创建？")) {
        return
    }

    return httpPost(instance, '../../api/AddJob', {
        sites: sites,
        word: word,
    })
}

async function ListJob() {
    let url = '../../api/ListJob'
    if (document.domain === 'localhost') {
        url += '.json'
    }
    try {
        let resp = await instance.get(url, {
            params: {},
            paramsSerializer: params => {
                return Qs.stringify(params, {indices: false})
            }
        })
        return dealResp(resp)
    } catch (error) {
        // dealErr(error)
    }
    return null
}

