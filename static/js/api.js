const instance = axios.create({timeout: 60 * 1000})
instance.interceptors.request.use(
    config => {
        config.headers['Authorization'] = 'Bearer ' + enJwt()
        return config
    },
    error => Promise.reject(error))

async function ping() {
    return httpGet(instance, '../../api/ping', {})
}

async function pong() {
    return httpPost(instance, '../../api/ping', {})
}

async function LoginGet(username, password) {
    return httpGet(instance, '../../api/LoginGet', {})
}

async function LoginPost(username, password) {
    if (username === undefined || username == null || username === '') {
        dealErr('username-为空')
        return null
    }
    if (password === undefined || password == null || password === '') {
        dealErr('password-为空')
        return null
    }

    return httpPost(instance, '../../api/LoginPost', {
        username: username,
        password: password,
    })
}

async function ListUser() {
    return httpGet(instance, '../../api/ListUser', {})
}

