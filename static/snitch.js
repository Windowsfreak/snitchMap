if (window.timer !== undefined) {
    window.clearTimeout(window.timer)
    window.timer = undefined;
}

const http = function(method, url, params, headers = {}) {
    return new Promise(function (resolve, reject) {
        const xhr = new XMLHttpRequest();
        xhr.open(method, url);
        xhr.onload = function () {
            if (this.status >= 200 && this.status < 300) {
                try {
                    resolve(headers.accept === 'text/html' ? xhr.response : JSON.parse(xhr.response));
                } catch (e) {
                    //xhr.onerror();
                    console.error('error_server_request');
                    reject({
                        method,
                        url,
                        params,
                        headers,
                        status: this.status,
                        statusText: xhr.statusText,
                        message: xhr.response
                    });
                }
            } else {
                xhr.onerror();
            }
        };
        xhr.onerror = function() {
            console.error('error_server_request');
            reject({
                method,
                url,
                params,
                headers,
                status: this.status,
                statusText: xhr.statusText,
                message: xhr.response
            });
        };
        if (headers) {
            Object.keys(headers).forEach(key => xhr.setRequestHeader(key, headers[key]));
        }
        // stringify params if object:
        if (params && typeof params === 'object') {
            xhr.send(Object.keys(params).map(key => encodeURIComponent(key) + '=' + encodeURIComponent(params[key])).join('&'));
        } else {
            xhr.send(params);
        }
    });
};

const get = (url, params, headers) => {
    if (params && typeof params === 'object') {
        return http('GET', url + (url.indexOf('?') > 0 ? '&' : '?') + Object.keys(params).map(key => encodeURIComponent(key) + '=' + encodeURIComponent(params[key])).join('&'), undefined, headers);
    } else if (params) {
        return http('GET', url + (url.indexOf('?') > 0 ? '&' : '?') + params, undefined, headers);
    } else {
        return http('GET', url, params, headers);
    }
};

let time = 1; // -3600;

const updatePlayers = function() {
    get(window.url + '/world/users/after', {
        token: window.token,
        time: time
    }).then(it => {
        time = it.Time - 1;
        if (typeof civMapApi !== 'undefined') {
            it.Users.forEach(user => civMapApi.store.dispatch({
                type: 'UPDATE_FEATURE_IN_COLLECTION',
                collectionId: 'civmap:collection/user',
                feature: {
                    id: 'snitchmap.users.' + user.User,
                    name: user.User,
                    x: user.X, y: user.Y, z: user.Z,
                    image: 'https://mc-heads.net/head/' + user.User + '/32',
                    icon: 'https://mc-heads.net/avatar/' + user.User + '/32',
                    last_seen: new Date(user.Seen * 1000).toLocaleString(undefined, {day: '2-digit', month: '2-digit', year: 'numeric', hour: '2-digit', minute: '2-digit', second: '2-digit'}),
                    snitch_hits: user.Hits
                }
            }));
        } else {
            console.log('no CivMap running!');
        }
    });
}

window.timer = window.setInterval(updatePlayers, 5000);