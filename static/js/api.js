const instance = axios.create({timeout: 60 * 1000})
// instance.interceptors.request.use(
//     config => {
//         config.headers['Authorization'] = 'Bearer ' + enJwt()
//         return config
//     },
//     error => Promise.reject(error))

function dealResponse(response) {
    let result = response.data
    if (result.code !== 1) {
        dealErr(result.msg)
        return null
    }
    return result.data
}

function dealErr(error) {
    let msg = JSON.stringify(error)
    if (msg === undefined || msg == null || msg === '' || msg === '{}' || msg === '[]') {
        msg = error
    }
    alert("error: " + msg)
    log(msg)
}

async function ping() {
    let url = '../../api/ping'
    if (document.domain === 'localhost') {
        url += '.json'
    }
    try {
        let response = await instance.post(url, {
            params: {},
            paramsSerializer: params => {
                return Qs.stringify(params, {indices: false})
            }
        })
        return dealResponse(response)
    } catch (error) {
        dealErr(error)
    }
    return null
}

async function LoginPost(username, password) {
    if (username === undefined || username == null || username === '') {
        dealErr('username为空')
        return null
    }
    if (password === undefined || password == null || password === '') {
        dealErr('password为空')
        return null
    }

    if (!window.confirm("确定？")) {
        return
    }

    let url = '../../api/LoginPost'
    if (document.domain === 'localhost') {
        url += '.json'
    }
    try {
        let response = await instance.post(url, {
            username: username,
            password: password,
        })
        return dealResponse(response)
    } catch (error) {
        dealErr(error)
    }
    return null
}

async function LoginGet(username, password) {
    if (username === undefined || username == null || username === '') {
        dealErr('username为空')
        return null
    }
    if (password === undefined || password == null || password === '') {
        dealErr('password为空')
        return null
    }

    if (!window.confirm("确定？")) {
        return
    }

    let url = '../../api/LoginGet'
    if (document.domain === 'localhost') {
        url += '.json'
    }
    try {
        let response = await instance.get(url, {
            params: {
                username: username,
                password: password,
            },
            paramsSerializer: params => {
                return Qs.stringify(params, {indices: false})
            }
        })
        return dealResponse(response)
    } catch (error) {
        dealErr(error)
    }
    return null
}

