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

async function LoginGet() {
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

async function AddUser(username, password) {
    if (username === undefined || username == null || username === '') {
        dealErr('username-为空')
        return null
    }
    if (password === undefined || password == null || password === '') {
        dealErr('password-为空')
        return null
    }

    if (!window.confirm("确定添加？")) {
        return
    }

    return httpPost(instance, '../../api/AddUser', {
        object: [{
            username: username,
            password: password,
        }]
    })
}

async function RemoveUser(id) {
    if (id === undefined || id == null || id === '') {
        dealErr('id-为空')
        return null
    }

    if (!window.confirm("确定删除？")) {
        return
    }

    return httpPost(instance, '../../api/RemoveUser', {
        id: id,
    })
}

async function ListUser() {
    return httpGet(instance, '../../api/ListUser', {})
}

